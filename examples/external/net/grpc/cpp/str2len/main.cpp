#include <iostream>
#include <memory>
#include <string>
#include <optional>
#include <functional>

#include <grpcpp/grpcpp.h>

#include "sl.grpc.pb.h"
#include "sl.pb.h"

using google::protobuf::Timestamp;

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;

using str2len::v1::StringLengthService;
using str2len::v1::StringLengthRequest;
using str2len::v1::StringLengthResponse;

using str2len::v1::SendEventService;
using str2len::v1::Proxy_SendEventRequest;
using str2len::v1::Proxy_SendEventResponse;

using str2len::v1::Proxy_RequestId;
using str2len::v1::Proxy_ReplyId;

using str2len::v1::GetCmdService;
using str2len::v1::Proxy_GetCmdRequest;
using str2len::v1::Proxy_GetCmdResponse;

using str2len::v1::Proxy_Target;
using str2len::v1::Proxy_Uuid;

using str2len::v1::Proxy_StringLengthCommand;
using str2len::v1::Proxy_StringLengthResult;

using namespace std;

class Client {
	public: Client(shared_ptr<Channel> chan)
	: s(StringLengthService::NewStub(chan))
	{}

	private: unique_ptr<StringLengthService::Stub> s;

	public: const optional<uint64_t> StringLength(const string& src){
		StringLengthRequest req;
		req.set_source(src);

		StringLengthResponse res;
		ClientContext ctx;

		Status status = this->s->StringLength(&ctx, req, &res);
		if(status.ok()) return res.length();
		cerr << "Unable to get length." << endl;
		cerr << "Error code:    " << status.error_code() << endl;
		cerr << "Error message: " << status.error_message() << endl;
		return {};
	}
};

class Sender {
	public: Sender(shared_ptr<Channel> chan)
	: s(SendEventService::NewStub(chan))
	{}

	private: unique_ptr<SendEventService::Stub> s;

	public: const optional<Proxy_SendEventResponse> Send(const Proxy_SendEventRequest& req){
		Proxy_SendEventResponse res;
		ClientContext ctx;
		Status status = this->s->SendEvent(&ctx, req, &res);
		if(status.ok()) return res;
		cerr << "Unable to send." << endl;
		cerr << "Error code:    " << status.error_code() << endl;
		cerr << "Error message: " << status.error_message() << endl;
		return {};
	}

	/*
	private: const optional<Proxy_Uuid> reqid2uuid(const Proxy_RequestId& reqid) const {
		const bool has_id = reqid.has_id();
		if(! has_id) return {};
		const Proxy_Uuid& id = reqid.id();
		return id;
	}

	private: const optional<Proxy_Uuid> request2uuid(const Proxy_SendEventRequest& request) const {
		const bool has_sendid = request.has_send_id();
		if(! has_sendid) return {};
		const Proxy_RequestId& id = request.send_id();
		return reqid2uuid(id);
	}
	*/

	/*
	private: const optional<string> send(
		const optional<Proxy_RequestId> request,
		Proxy_ReplyId reply,
		optional<StringLengthResponse> response
	){
		Proxy_StringLengthResult result;
		Proxy_SendEventRequest event_request;
		if(request.has_value()){
			const Proxy_RequestId send_id = request.value();
			Proxy_RequestId* id = event_request.mutable_send_id();
			Proxy_Uuid* uuid = id->mutable_id();
		}
		return {};
	}
	*/
};

class Converter {
	public: virtual Status Convert(
		const StringLengthRequest& request,
		StringLengthResponse* out
	) = 0;
};

typedef function<Status(const StringLengthRequest&, StringLengthResponse*)> ConvFunc;

class ConverterFn: public Converter {
	private: ConvFunc internal;

	public: virtual Status Convert(
		const StringLengthRequest& request,
		StringLengthResponse* out
	){ return this->internal(request, out); }

	public: ConverterFn(ConvFunc f):internal(f){}
};

class Getter {
	public: Getter(shared_ptr<Channel> chan)
	: s(GetCmdService::NewStub(chan))
	{}

	private: unique_ptr<GetCmdService::Stub> s;

	/*
	private: const optional<StringLengthRequest> cmd2req(const Proxy_StringLengthCommand& c) const {
		const bool has_request = c.has_request();
		if(! has_request) return {};
		const StringLengthRequest& request = c.request();
		return request;
	}
	*/

	public: const optional<Proxy_StringLengthCommand> getcmd_res2cmd(
		const Proxy_GetCmdResponse& res
	) const {
		const bool has_cmd = res.has_strlen_cmd();
		if(! has_cmd) return {};
		const Proxy_StringLengthCommand& cmd = res.strlen_cmd();
		return cmd;
	}

	public: const optional<Proxy_GetCmdResponse> get_command(const Proxy_GetCmdRequest& request){
		Proxy_GetCmdResponse response;
		ClientContext ctx;
		Status status = this->s->GetCmd(&ctx, request, &response);
		if(status.ok()) return response;
		cerr << "Unable to get command." << endl;
		cerr << "Error code:    " << status.error_code() << endl;
		cerr << "Error message: " << status.error_message() << endl;
		return {};
	}
};

class CheckedRequest {
	private: optional<Proxy_RequestId> req_id;
	private: Proxy_Uuid reply_uuid;
	private: optional<StringLengthRequest> request;

	public: CheckedRequest(
		const optional<Proxy_RequestId> id,
		const Proxy_Uuid reply,
		const optional<StringLengthRequest> req
	){
		this->req_id = id;
		this->reply_uuid = reply;
		this->request = req;
	}

	public: const optional<StringLengthRequest> Request() const { return this->request; }
	public: const optional<Proxy_RequestId> RequestId() const { return this->req_id; }
	public: const Proxy_Uuid ReplyId() const { return this->reply_uuid; }
};

class Svc {
	private: Getter& getter;
	private: Converter& converter;
	private: Sender& sender;

	public: Svc(
		Getter& g,
		Converter& c,
		Sender& s
	)
	: getter(g)
	, converter(c)
	, sender(s)
	{}

	private: const optional<StringLengthResponse> convert(
		const StringLengthRequest& request
	) const {
		StringLengthResponse response;
		Status status = this->converter.Convert(request, &response);
		if(status.ok()) return response;
		cerr << "Unable to convert." << endl;
		cerr << "Error code:    " << status.error_code() << endl;
		cerr << "Error message: " << status.error_message() << endl;
		return {};
	}

	private: const optional<Proxy_StringLengthCommand> get_command(
		const Proxy_GetCmdRequest& request
	){
		const optional<Proxy_GetCmdResponse> o = this->getter.get_command(request);
		if(! o.has_value()) return {};
		const Proxy_GetCmdResponse res = o.value();
		return this->getter.getcmd_res2cmd(res);
	}

	private: const optional<Proxy_RequestId> cmd2reqid(
		const Proxy_StringLengthCommand& cmd
	) const {
		const bool has_reqid = cmd.has_strlen_req_id();
		if(! has_reqid) return {};
		const Proxy_RequestId& reqid = cmd.strlen_req_id();
		return reqid;
	}

	private: const optional<Proxy_ReplyId> cmd2reply(
		const Proxy_StringLengthCommand& cmd
	) const {
		const bool has_reply = cmd.has_reply();
		if(! has_reply) return {};
		const Proxy_ReplyId& reply = cmd.reply();
		return reply;
	}

	private: const optional<Proxy_Uuid> reply2uuid(const Proxy_ReplyId& reply) const {
		const bool has_id = reply.has_id();
		if(! has_id) return {};
		const Proxy_Uuid& id = reply.id();
		return id;
	}

	private: const optional<Proxy_Uuid> cmd2uuid(const Proxy_StringLengthCommand& cmd) const {
		const optional<Proxy_ReplyId> oreply = this->cmd2reply(cmd);
		if(! oreply.has_value()) return {};
		const Proxy_ReplyId reply = oreply.value();
		return this->reply2uuid(reply);
	}

	private: const optional<StringLengthRequest> cmd2req(
		const Proxy_StringLengthCommand& cmd
	) const {
		const bool has_request = cmd.has_request();
		if(! has_request) return {};
		const StringLengthRequest& request = cmd.request();
		return request;
	}

	private: const optional<CheckedRequest> cmd2checked(
		const Proxy_StringLengthCommand& cmd
	) const {
		const optional<Proxy_Uuid> ouuid = this->cmd2uuid(cmd);
		if(! ouuid.has_value()) return {};
		const Proxy_Uuid uuid = ouuid.value();
		const optional<Proxy_RequestId> reqid = this->cmd2reqid(cmd);
		const optional<StringLengthRequest> req = this->cmd2req(cmd);
		return CheckedRequest(reqid, uuid, req);
	}

	private: const optional<StringLengthResponse> checked2response(const CheckedRequest& checked){
		const optional<StringLengthRequest> oreq = checked.Request();
		if(! oreq.has_value()) return {};
		const StringLengthRequest req = oreq.value();
		return this->convert(req);
	}

	private: void uuid_copy(const Proxy_Uuid& src, Proxy_Uuid* dst) const {
		dst->set_hi(src.hi());
		dst->set_lo(src.lo());
	}

	private: void reqid_copy(const Proxy_RequestId& src, Proxy_RequestId* dst) const {
		const Proxy_Uuid& s = src.id();
		Proxy_Uuid* d = dst->mutable_id();
		this->uuid_copy(s, d);
	}

	private: void response_copy(const StringLengthResponse& src, StringLengthResponse* dst) const {
		dst->set_length(src.length());
	}

	private: Proxy_ReplyId uuid2reply(const Proxy_Uuid& uuid) const {
		Proxy_ReplyId reply;
		Proxy_Uuid* out = reply.mutable_id();
		this->uuid_copy(uuid, out);
		return reply;
	}

	private: void reply_copy(const Proxy_ReplyId& src, Proxy_ReplyId* dst) const {
		const optional<Proxy_Uuid> ou = this->reply2uuid(src);
		if(! ou.has_value()) return;
		const Proxy_Uuid u = ou.value();
		Proxy_Uuid* out = dst->mutable_id();
		this->uuid_copy(u, out);
	}

	private: const Proxy_StringLengthResult checked2res(const CheckedRequest& checked){
		const optional<Proxy_RequestId> oreq = checked.RequestId();
		const Proxy_Uuid uuid = checked.ReplyId();
		const optional<StringLengthResponse> ores = this->checked2response(checked);

		Proxy_StringLengthResult result;
		if(oreq.has_value()){
			const Proxy_RequestId reqid = oreq.value();
			Proxy_RequestId* out = result.mutable_strlen_req_id();
			this->reqid_copy(reqid, out);
		}
		if(ores.has_value()){
			const StringLengthResponse res = ores.value();
			StringLengthResponse* out = result.mutable_response();
			this->response_copy(res, out);
		}
		const Proxy_ReplyId repl = this->uuid2reply(uuid);
		Proxy_ReplyId* out = result.mutable_reply();
		this->reply_copy(repl, out);
		return result;
	}

	private: void slr_copy(
		const Proxy_StringLengthResult& src,
		Proxy_StringLengthResult* dst
	) const {
		*dst = src;
	}

	private: const Proxy_SendEventRequest checked2evt(const CheckedRequest& checked){
		const Proxy_StringLengthResult res = this->checked2res(checked);
		Proxy_SendEventRequest req;
		Proxy_StringLengthResult* out = req.mutable_reply();
		this->slr_copy(res, out);
		const optional<Proxy_RequestId> oreq = checked.RequestId();
		if(oreq.has_value()){
			const Proxy_RequestId reqid = oreq.value();
			Proxy_RequestId* oid = req.mutable_send_id();
			*oid = reqid;
		}
		return req;
	}

	private: const optional<Proxy_SendEventRequest> cmd2evt(const Proxy_StringLengthCommand& c){
		const optional<CheckedRequest> oc = this->cmd2checked(c);
		if(! oc.has_value()) return {};
		const CheckedRequest checked = oc.value();
		const Proxy_SendEventRequest req = this->checked2evt(checked);
		return req;
	}

	private: const optional<Proxy_SendEventResponse> cmd2res(const Proxy_StringLengthCommand& c){
		const optional<Proxy_SendEventRequest> oreq = this->cmd2evt(c);
		if(! oreq.has_value()) return {};
		const Proxy_SendEventRequest req = oreq.value();
		return this->sender.Send(req);
	}

	private: const optional<Proxy_SendEventResponse> get2res(const Proxy_GetCmdResponse& c){
		const optional<Proxy_StringLengthCommand> ostr = this->getter.getcmd_res2cmd(c);
		if(! ostr.has_value()) return {};
		const Proxy_StringLengthCommand scmd = ostr.value();
		return this->cmd2res(scmd);
	}

	public: const optional<Proxy_SendEventResponse> greq2res(const Proxy_GetCmdRequest& c){
		const optional<Proxy_GetCmdResponse> ogres = this->getter.get_command(c);
		if(! ogres.has_value()) return {};
		const Proxy_GetCmdResponse gres = ogres.value();
		return this->get2res(gres);
	}
};

class App {
	private: Svc& svc;
	public: App(Svc& service) : svc(service){}

	public: const optional<Timestamp> GetAndSend(const Proxy_Uuid& target){
		Proxy_GetCmdRequest req;
		Proxy_Target* tgt = req.mutable_target();
		Proxy_Uuid* tgt_id = tgt->mutable_id();
		*tgt_id = target;
		const optional<Proxy_SendEventResponse> ores = this->svc.greq2res(req);
		if(! ores.has_value()) return {};
		const Proxy_SendEventResponse res = ores.value();
		if(! res.has_sent()) return {};
		const Timestamp sent = res.sent();
		return sent;
	}
};

int main(int argc, char** argv){
	auto fconv = ConverterFn([](
		const StringLengthRequest& req,
		StringLengthResponse* out
	){
		const string& raw = req.source();
		out->set_length(raw.size());
		return Status::OK;
	});
	Converter& converter = fconv;

	const char* addr = "localhost:50051";

	shared_ptr<Channel> chan = grpc::CreateChannel(
		addr,
		grpc::InsecureChannelCredentials()
	);

	Getter getter(chan);
	Sender sender(chan);

	Svc svc(getter, converter, sender);
	App app(svc);

	Proxy_Uuid uuid;
	uuid.set_hi(0xcafef00ddeadbeaf);
	uuid.set_lo(0xface864299792458);
	const optional<Timestamp> sent = app.GetAndSend(uuid);

	if(! sent.has_value()) return 1;
	const Timestamp ts = sent.value();

	cout << "timestamp seconds: " << ts.seconds() << endl;

	return 0;
}
