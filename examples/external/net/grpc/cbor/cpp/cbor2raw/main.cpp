#include <iostream>
#include <string>
#include <optional>

#include <immintrin.h>

#include <cbor.h>

#include "httplib.h"

using namespace std;

class CborUser {
	public: virtual bool use_cbor(const cbor_item_t* item) = 0;
};

class CborReader {
	private: cbor_item_t* internal = nullptr;

	public: CborReader(const uint8_t* data, const size_t size){
		struct cbor_load_result result = {0};
		cbor_item_t* item = cbor_load(data, size, &result);
		if(CBOR_ERR_NONE != result.error.code) return;
		this->internal = item;
	}

	public: bool has_value() const { return nullptr != this->internal; }

	private: const cbor_item_t* value() const { return this->internal; }

	public: bool to_cbor_user(CborUser& user) const {
		if(! this->has_value()) return false;
		const cbor_item_t* val = this->value();
		return user.use_cbor(val);
	}

	public: ~CborReader(){
		if(! this->has_value()) return;
		cbor_decref(&this->internal);
	}
};

enum class SeverityType {
	UNSPECIFIED = 0,
	TRACE = 1,
	DEBUG = 5,
	INFO = 9,
	WARN = 13,
	ERROR = 17,
	FATAL = 21,
};

class Severity {
	private: SeverityType internal = SeverityType::UNSPECIFIED;

	private: SeverityType convert(const uint8_t raw){
		switch(raw){
			case 1: return SeverityType::TRACE;
			case 5: return SeverityType::DEBUG;
			case 9: return SeverityType::INFO;
			case 13: return SeverityType::WARN;
			case 17: return SeverityType::ERROR;
			case 21: return SeverityType::FATAL;
			default: return SeverityType::UNSPECIFIED;
		}
	}

	public: virtual bool use_cbor(const cbor_item_t* item){
		if(CBOR_TYPE_UINT != cbor_typeof(item)) return false;
		const uint8_t raw = cbor_get_uint8(item);
		this->internal = this->convert(raw);
		return true;
	}
};

class Uuid: public CborUser {
	private: uint64_t hi = 0;
	private: uint64_t lo = 0;

	private: bool cbor2u(const cbor_item_t* i, uint64_t& out){
		if(CBOR_TYPE_UINT != cbor_typeof(i)) return false;
		out = cbor_get_uint64(i);
		return true;
	}

	public: virtual bool use_cbor(const cbor_item_t* item){
		if(CBOR_TYPE_ARRAY != cbor_typeof(item)) return false;
		if(2 != cbor_array_size(item)) return false;
		cbor_item_t* h = cbor_array_get(item, 0);
		cbor_item_t* l = cbor_array_get(item, 0);
		const bool bh = this->cbor2u(h, this->hi);
		const bool bl = this->cbor2u(l, this->lo);
		cbor_decref(&h);
		cbor_decref(&l);
		return bh && bl;
	}
};

class Timestamp {
	private: uint64_t unixtime_nanos = 0;

	public: virtual bool use_cbor(const cbor_item_t* item){
		if(CBOR_TYPE_UINT != cbor_typeof(item)) return false;
		this->unixtime_nanos = cbor_get_uint64(item);
		return true;
	}
};

class CommonString: public CborUser {
	private: string internal = "";

	public: virtual bool use_cbor(const cbor_item_t* item){
		if(CBOR_TYPE_STRING != cbor_typeof(item)) return false;
		const auto s = cbor_string_handle(item);
		const size_t sz = cbor_string_length(item);
		this->internal = string((const char*)s, sz);
		return true;
	}

	public: const string as_string() const { return this->internal; }
};

class Body {
	private: CommonString s;
	public: virtual bool use_cbor(const cbor_item_t* item){ return this->s.use_cbor(item); }
};

class Name {
	private: CommonString s;
	public: virtual bool use_cbor(const cbor_item_t* item){ return this->s.use_cbor(item); }

	public: const string as_string() const { return this->s.as_string(); }
};

class NameSpace {
	private: CommonString s;
	public: virtual bool use_cbor(const cbor_item_t* item){ return this->s.use_cbor(item); }

	public: const string as_string() const { return this->s.as_string(); }
};

class Version {
	private: CommonString s;
	public: virtual bool use_cbor(const cbor_item_t* item){ return this->s.use_cbor(item); }
};

class ConvertRequest: public CborUser {
	// ti: timestamp(0x7469)
	private: Timestamp timestamp;

	// tr: trace_id(0x7472)
	private: Uuid trace_id;

	// se: severity(0x7365)
	private: Severity severity;

	// bo: body(0x626f)
	private: Body body;

	// na: name(0x6e61)
	private: Name name;

	// ns: namespace(0x6e73)
	private: NameSpace ns;

	// in: instance_id(0x696e)
	private: Uuid instance_id;

	// ve: version(0x7665)
	private: Version ver;

	private: bool use(const uint16_t typ, const cbor_item_t* item){
		switch(typ){
			case 0x7469: return this->timestamp.use_cbor(item);
			case 0x7472: return this->trace_id.use_cbor(item);
			case 0x7365: return this->severity.use_cbor(item);
			case 0x626f: return this->body.use_cbor(item);
			case 0x6e61: return this->name.use_cbor(item);
			case 0x6e73: return this->ns.use_cbor(item);
			case 0x696e: return this->instance_id.use_cbor(item);
			case 0x7665: return this->ver.use_cbor(item);
			default: return false;
		}
	}

	public: virtual bool use_cbor(const cbor_item_t* item){
		if(CBOR_TYPE_MAP != cbor_typeof(item)) return false;
		const int sz = cbor_map_size(item);
		if(8 != sz) return false;
		for(int i=0; i<sz; i++){
			const cbor_item_t* key = cbor_map_handle(item)[i].key;
			const cbor_item_t* val = cbor_map_handle(item)[i].value;
			if(CBOR_TYPE_STRING != cbor_typeof(key)) return false;
			if(cbor_string_length(key) < 2) return false;
			uint16_t buf;
			memcpy(&buf, cbor_string_handle(key), 2);

			const uint16_t ho = be16toh(buf);
			const bool ok = this->use(ho, val);
			if(!ok) return false;
		}
		return true;
	}

	public: const string to_string() const {
		return this->ns.as_string() + "." + this->name.as_string();
	}
};

class ConvertResponse {
	private: string raw = "";

	public: ConvertResponse(const ConvertRequest& req){ this->raw = req.to_string(); }

	public: const string as_string() const { return this->raw; }
};

int main(int argc, char** argv){
	httplib::Server srv;
	srv.Post("/", [](const httplib::Request& req, httplib::Response& res){
		const string body = req.body;
		CborReader reader((const uint8_t*)body.c_str(), body.size());
		ConvertRequest creq;
		const bool ok = reader.to_cbor_user(creq);
		if(! ok){
			res.set_content("invalid request", "text/plain");
			res.status = 400;
			return;
		}
		ConvertResponse cres(creq);
		res.set_content(cres.as_string(), "text/plain");
	});
	srv.listen("127.0.0.1", 8148);
	return 0;
}
