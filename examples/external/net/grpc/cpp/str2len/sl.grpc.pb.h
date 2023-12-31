// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: sl.proto
#ifndef GRPC_sl_2eproto__INCLUDED
#define GRPC_sl_2eproto__INCLUDED

#include "sl.pb.h"

#include <functional>
#include <grpcpp/generic/async_generic_service.h>
#include <grpcpp/support/async_stream.h>
#include <grpcpp/support/async_unary_call.h>
#include <grpcpp/support/client_callback.h>
#include <grpcpp/client_context.h>
#include <grpcpp/completion_queue.h>
#include <grpcpp/support/message_allocator.h>
#include <grpcpp/support/method_handler.h>
#include <grpcpp/impl/proto_utils.h>
#include <grpcpp/impl/rpc_method.h>
#include <grpcpp/support/server_callback.h>
#include <grpcpp/impl/server_callback_handlers.h>
#include <grpcpp/server_context.h>
#include <grpcpp/impl/service_type.h>
#include <grpcpp/support/status.h>
#include <grpcpp/support/stub_options.h>
#include <grpcpp/support/sync_stream.h>

namespace str2len {
namespace v1 {

class StringLengthService final {
 public:
  static constexpr char const* service_full_name() {
    return "str2len.v1.StringLengthService";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status StringLength(::grpc::ClientContext* context, const ::str2len::v1::StringLengthRequest& request, ::str2len::v1::StringLengthResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::StringLengthResponse>> AsyncStringLength(::grpc::ClientContext* context, const ::str2len::v1::StringLengthRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::StringLengthResponse>>(AsyncStringLengthRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::StringLengthResponse>> PrepareAsyncStringLength(::grpc::ClientContext* context, const ::str2len::v1::StringLengthRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::StringLengthResponse>>(PrepareAsyncStringLengthRaw(context, request, cq));
    }
    class async_interface {
     public:
      virtual ~async_interface() {}
      virtual void StringLength(::grpc::ClientContext* context, const ::str2len::v1::StringLengthRequest* request, ::str2len::v1::StringLengthResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void StringLength(::grpc::ClientContext* context, const ::str2len::v1::StringLengthRequest* request, ::str2len::v1::StringLengthResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
    };
    typedef class async_interface experimental_async_interface;
    virtual class async_interface* async() { return nullptr; }
    class async_interface* experimental_async() { return async(); }
   private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::StringLengthResponse>* AsyncStringLengthRaw(::grpc::ClientContext* context, const ::str2len::v1::StringLengthRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::StringLengthResponse>* PrepareAsyncStringLengthRaw(::grpc::ClientContext* context, const ::str2len::v1::StringLengthRequest& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
    ::grpc::Status StringLength(::grpc::ClientContext* context, const ::str2len::v1::StringLengthRequest& request, ::str2len::v1::StringLengthResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::str2len::v1::StringLengthResponse>> AsyncStringLength(::grpc::ClientContext* context, const ::str2len::v1::StringLengthRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::str2len::v1::StringLengthResponse>>(AsyncStringLengthRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::str2len::v1::StringLengthResponse>> PrepareAsyncStringLength(::grpc::ClientContext* context, const ::str2len::v1::StringLengthRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::str2len::v1::StringLengthResponse>>(PrepareAsyncStringLengthRaw(context, request, cq));
    }
    class async final :
      public StubInterface::async_interface {
     public:
      void StringLength(::grpc::ClientContext* context, const ::str2len::v1::StringLengthRequest* request, ::str2len::v1::StringLengthResponse* response, std::function<void(::grpc::Status)>) override;
      void StringLength(::grpc::ClientContext* context, const ::str2len::v1::StringLengthRequest* request, ::str2len::v1::StringLengthResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
     private:
      friend class Stub;
      explicit async(Stub* stub): stub_(stub) { }
      Stub* stub() { return stub_; }
      Stub* stub_;
    };
    class async* async() override { return &async_stub_; }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    class async async_stub_{this};
    ::grpc::ClientAsyncResponseReader< ::str2len::v1::StringLengthResponse>* AsyncStringLengthRaw(::grpc::ClientContext* context, const ::str2len::v1::StringLengthRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::str2len::v1::StringLengthResponse>* PrepareAsyncStringLengthRaw(::grpc::ClientContext* context, const ::str2len::v1::StringLengthRequest& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_StringLength_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status StringLength(::grpc::ServerContext* context, const ::str2len::v1::StringLengthRequest* request, ::str2len::v1::StringLengthResponse* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_StringLength : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_StringLength() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_StringLength() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status StringLength(::grpc::ServerContext* /*context*/, const ::str2len::v1::StringLengthRequest* /*request*/, ::str2len::v1::StringLengthResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestStringLength(::grpc::ServerContext* context, ::str2len::v1::StringLengthRequest* request, ::grpc::ServerAsyncResponseWriter< ::str2len::v1::StringLengthResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_StringLength<Service > AsyncService;
  template <class BaseClass>
  class WithCallbackMethod_StringLength : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_StringLength() {
      ::grpc::Service::MarkMethodCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::str2len::v1::StringLengthRequest, ::str2len::v1::StringLengthResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::str2len::v1::StringLengthRequest* request, ::str2len::v1::StringLengthResponse* response) { return this->StringLength(context, request, response); }));}
    void SetMessageAllocatorFor_StringLength(
        ::grpc::MessageAllocator< ::str2len::v1::StringLengthRequest, ::str2len::v1::StringLengthResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(0);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::str2len::v1::StringLengthRequest, ::str2len::v1::StringLengthResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_StringLength() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status StringLength(::grpc::ServerContext* /*context*/, const ::str2len::v1::StringLengthRequest* /*request*/, ::str2len::v1::StringLengthResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* StringLength(
      ::grpc::CallbackServerContext* /*context*/, const ::str2len::v1::StringLengthRequest* /*request*/, ::str2len::v1::StringLengthResponse* /*response*/)  { return nullptr; }
  };
  typedef WithCallbackMethod_StringLength<Service > CallbackService;
  typedef CallbackService ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_StringLength : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_StringLength() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_StringLength() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status StringLength(::grpc::ServerContext* /*context*/, const ::str2len::v1::StringLengthRequest* /*request*/, ::str2len::v1::StringLengthResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_StringLength : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_StringLength() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_StringLength() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status StringLength(::grpc::ServerContext* /*context*/, const ::str2len::v1::StringLengthRequest* /*request*/, ::str2len::v1::StringLengthResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestStringLength(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_StringLength : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_StringLength() {
      ::grpc::Service::MarkMethodRawCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->StringLength(context, request, response); }));
    }
    ~WithRawCallbackMethod_StringLength() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status StringLength(::grpc::ServerContext* /*context*/, const ::str2len::v1::StringLengthRequest* /*request*/, ::str2len::v1::StringLengthResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* StringLength(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_StringLength : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_StringLength() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler<
          ::str2len::v1::StringLengthRequest, ::str2len::v1::StringLengthResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::str2len::v1::StringLengthRequest, ::str2len::v1::StringLengthResponse>* streamer) {
                       return this->StreamedStringLength(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_StringLength() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status StringLength(::grpc::ServerContext* /*context*/, const ::str2len::v1::StringLengthRequest* /*request*/, ::str2len::v1::StringLengthResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedStringLength(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::str2len::v1::StringLengthRequest,::str2len::v1::StringLengthResponse>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_StringLength<Service > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_StringLength<Service > StreamedService;
};

class GetCmdService final {
 public:
  static constexpr char const* service_full_name() {
    return "str2len.v1.GetCmdService";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status GetCmd(::grpc::ClientContext* context, const ::str2len::v1::Proxy_GetCmdRequest& request, ::str2len::v1::Proxy_GetCmdResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::Proxy_GetCmdResponse>> AsyncGetCmd(::grpc::ClientContext* context, const ::str2len::v1::Proxy_GetCmdRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::Proxy_GetCmdResponse>>(AsyncGetCmdRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::Proxy_GetCmdResponse>> PrepareAsyncGetCmd(::grpc::ClientContext* context, const ::str2len::v1::Proxy_GetCmdRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::Proxy_GetCmdResponse>>(PrepareAsyncGetCmdRaw(context, request, cq));
    }
    class async_interface {
     public:
      virtual ~async_interface() {}
      virtual void GetCmd(::grpc::ClientContext* context, const ::str2len::v1::Proxy_GetCmdRequest* request, ::str2len::v1::Proxy_GetCmdResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void GetCmd(::grpc::ClientContext* context, const ::str2len::v1::Proxy_GetCmdRequest* request, ::str2len::v1::Proxy_GetCmdResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
    };
    typedef class async_interface experimental_async_interface;
    virtual class async_interface* async() { return nullptr; }
    class async_interface* experimental_async() { return async(); }
   private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::Proxy_GetCmdResponse>* AsyncGetCmdRaw(::grpc::ClientContext* context, const ::str2len::v1::Proxy_GetCmdRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::Proxy_GetCmdResponse>* PrepareAsyncGetCmdRaw(::grpc::ClientContext* context, const ::str2len::v1::Proxy_GetCmdRequest& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
    ::grpc::Status GetCmd(::grpc::ClientContext* context, const ::str2len::v1::Proxy_GetCmdRequest& request, ::str2len::v1::Proxy_GetCmdResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::str2len::v1::Proxy_GetCmdResponse>> AsyncGetCmd(::grpc::ClientContext* context, const ::str2len::v1::Proxy_GetCmdRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::str2len::v1::Proxy_GetCmdResponse>>(AsyncGetCmdRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::str2len::v1::Proxy_GetCmdResponse>> PrepareAsyncGetCmd(::grpc::ClientContext* context, const ::str2len::v1::Proxy_GetCmdRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::str2len::v1::Proxy_GetCmdResponse>>(PrepareAsyncGetCmdRaw(context, request, cq));
    }
    class async final :
      public StubInterface::async_interface {
     public:
      void GetCmd(::grpc::ClientContext* context, const ::str2len::v1::Proxy_GetCmdRequest* request, ::str2len::v1::Proxy_GetCmdResponse* response, std::function<void(::grpc::Status)>) override;
      void GetCmd(::grpc::ClientContext* context, const ::str2len::v1::Proxy_GetCmdRequest* request, ::str2len::v1::Proxy_GetCmdResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
     private:
      friend class Stub;
      explicit async(Stub* stub): stub_(stub) { }
      Stub* stub() { return stub_; }
      Stub* stub_;
    };
    class async* async() override { return &async_stub_; }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    class async async_stub_{this};
    ::grpc::ClientAsyncResponseReader< ::str2len::v1::Proxy_GetCmdResponse>* AsyncGetCmdRaw(::grpc::ClientContext* context, const ::str2len::v1::Proxy_GetCmdRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::str2len::v1::Proxy_GetCmdResponse>* PrepareAsyncGetCmdRaw(::grpc::ClientContext* context, const ::str2len::v1::Proxy_GetCmdRequest& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_GetCmd_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status GetCmd(::grpc::ServerContext* context, const ::str2len::v1::Proxy_GetCmdRequest* request, ::str2len::v1::Proxy_GetCmdResponse* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_GetCmd : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_GetCmd() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_GetCmd() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetCmd(::grpc::ServerContext* /*context*/, const ::str2len::v1::Proxy_GetCmdRequest* /*request*/, ::str2len::v1::Proxy_GetCmdResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGetCmd(::grpc::ServerContext* context, ::str2len::v1::Proxy_GetCmdRequest* request, ::grpc::ServerAsyncResponseWriter< ::str2len::v1::Proxy_GetCmdResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_GetCmd<Service > AsyncService;
  template <class BaseClass>
  class WithCallbackMethod_GetCmd : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_GetCmd() {
      ::grpc::Service::MarkMethodCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::str2len::v1::Proxy_GetCmdRequest, ::str2len::v1::Proxy_GetCmdResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::str2len::v1::Proxy_GetCmdRequest* request, ::str2len::v1::Proxy_GetCmdResponse* response) { return this->GetCmd(context, request, response); }));}
    void SetMessageAllocatorFor_GetCmd(
        ::grpc::MessageAllocator< ::str2len::v1::Proxy_GetCmdRequest, ::str2len::v1::Proxy_GetCmdResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(0);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::str2len::v1::Proxy_GetCmdRequest, ::str2len::v1::Proxy_GetCmdResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_GetCmd() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetCmd(::grpc::ServerContext* /*context*/, const ::str2len::v1::Proxy_GetCmdRequest* /*request*/, ::str2len::v1::Proxy_GetCmdResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* GetCmd(
      ::grpc::CallbackServerContext* /*context*/, const ::str2len::v1::Proxy_GetCmdRequest* /*request*/, ::str2len::v1::Proxy_GetCmdResponse* /*response*/)  { return nullptr; }
  };
  typedef WithCallbackMethod_GetCmd<Service > CallbackService;
  typedef CallbackService ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_GetCmd : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_GetCmd() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_GetCmd() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetCmd(::grpc::ServerContext* /*context*/, const ::str2len::v1::Proxy_GetCmdRequest* /*request*/, ::str2len::v1::Proxy_GetCmdResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_GetCmd : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_GetCmd() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_GetCmd() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetCmd(::grpc::ServerContext* /*context*/, const ::str2len::v1::Proxy_GetCmdRequest* /*request*/, ::str2len::v1::Proxy_GetCmdResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestGetCmd(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_GetCmd : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_GetCmd() {
      ::grpc::Service::MarkMethodRawCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->GetCmd(context, request, response); }));
    }
    ~WithRawCallbackMethod_GetCmd() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status GetCmd(::grpc::ServerContext* /*context*/, const ::str2len::v1::Proxy_GetCmdRequest* /*request*/, ::str2len::v1::Proxy_GetCmdResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* GetCmd(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_GetCmd : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_GetCmd() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler<
          ::str2len::v1::Proxy_GetCmdRequest, ::str2len::v1::Proxy_GetCmdResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::str2len::v1::Proxy_GetCmdRequest, ::str2len::v1::Proxy_GetCmdResponse>* streamer) {
                       return this->StreamedGetCmd(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_GetCmd() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status GetCmd(::grpc::ServerContext* /*context*/, const ::str2len::v1::Proxy_GetCmdRequest* /*request*/, ::str2len::v1::Proxy_GetCmdResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedGetCmd(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::str2len::v1::Proxy_GetCmdRequest,::str2len::v1::Proxy_GetCmdResponse>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_GetCmd<Service > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_GetCmd<Service > StreamedService;
};

class SendEventService final {
 public:
  static constexpr char const* service_full_name() {
    return "str2len.v1.SendEventService";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    virtual ::grpc::Status SendEvent(::grpc::ClientContext* context, const ::str2len::v1::Proxy_SendEventRequest& request, ::str2len::v1::Proxy_SendEventResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::Proxy_SendEventResponse>> AsyncSendEvent(::grpc::ClientContext* context, const ::str2len::v1::Proxy_SendEventRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::Proxy_SendEventResponse>>(AsyncSendEventRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::Proxy_SendEventResponse>> PrepareAsyncSendEvent(::grpc::ClientContext* context, const ::str2len::v1::Proxy_SendEventRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::Proxy_SendEventResponse>>(PrepareAsyncSendEventRaw(context, request, cq));
    }
    class async_interface {
     public:
      virtual ~async_interface() {}
      virtual void SendEvent(::grpc::ClientContext* context, const ::str2len::v1::Proxy_SendEventRequest* request, ::str2len::v1::Proxy_SendEventResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void SendEvent(::grpc::ClientContext* context, const ::str2len::v1::Proxy_SendEventRequest* request, ::str2len::v1::Proxy_SendEventResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
    };
    typedef class async_interface experimental_async_interface;
    virtual class async_interface* async() { return nullptr; }
    class async_interface* experimental_async() { return async(); }
   private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::Proxy_SendEventResponse>* AsyncSendEventRaw(::grpc::ClientContext* context, const ::str2len::v1::Proxy_SendEventRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::str2len::v1::Proxy_SendEventResponse>* PrepareAsyncSendEventRaw(::grpc::ClientContext* context, const ::str2len::v1::Proxy_SendEventRequest& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
    ::grpc::Status SendEvent(::grpc::ClientContext* context, const ::str2len::v1::Proxy_SendEventRequest& request, ::str2len::v1::Proxy_SendEventResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::str2len::v1::Proxy_SendEventResponse>> AsyncSendEvent(::grpc::ClientContext* context, const ::str2len::v1::Proxy_SendEventRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::str2len::v1::Proxy_SendEventResponse>>(AsyncSendEventRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::str2len::v1::Proxy_SendEventResponse>> PrepareAsyncSendEvent(::grpc::ClientContext* context, const ::str2len::v1::Proxy_SendEventRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::str2len::v1::Proxy_SendEventResponse>>(PrepareAsyncSendEventRaw(context, request, cq));
    }
    class async final :
      public StubInterface::async_interface {
     public:
      void SendEvent(::grpc::ClientContext* context, const ::str2len::v1::Proxy_SendEventRequest* request, ::str2len::v1::Proxy_SendEventResponse* response, std::function<void(::grpc::Status)>) override;
      void SendEvent(::grpc::ClientContext* context, const ::str2len::v1::Proxy_SendEventRequest* request, ::str2len::v1::Proxy_SendEventResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
     private:
      friend class Stub;
      explicit async(Stub* stub): stub_(stub) { }
      Stub* stub() { return stub_; }
      Stub* stub_;
    };
    class async* async() override { return &async_stub_; }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    class async async_stub_{this};
    ::grpc::ClientAsyncResponseReader< ::str2len::v1::Proxy_SendEventResponse>* AsyncSendEventRaw(::grpc::ClientContext* context, const ::str2len::v1::Proxy_SendEventRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::str2len::v1::Proxy_SendEventResponse>* PrepareAsyncSendEventRaw(::grpc::ClientContext* context, const ::str2len::v1::Proxy_SendEventRequest& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_SendEvent_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    virtual ::grpc::Status SendEvent(::grpc::ServerContext* context, const ::str2len::v1::Proxy_SendEventRequest* request, ::str2len::v1::Proxy_SendEventResponse* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_SendEvent : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_SendEvent() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_SendEvent() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendEvent(::grpc::ServerContext* /*context*/, const ::str2len::v1::Proxy_SendEventRequest* /*request*/, ::str2len::v1::Proxy_SendEventResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestSendEvent(::grpc::ServerContext* context, ::str2len::v1::Proxy_SendEventRequest* request, ::grpc::ServerAsyncResponseWriter< ::str2len::v1::Proxy_SendEventResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_SendEvent<Service > AsyncService;
  template <class BaseClass>
  class WithCallbackMethod_SendEvent : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_SendEvent() {
      ::grpc::Service::MarkMethodCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::str2len::v1::Proxy_SendEventRequest, ::str2len::v1::Proxy_SendEventResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::str2len::v1::Proxy_SendEventRequest* request, ::str2len::v1::Proxy_SendEventResponse* response) { return this->SendEvent(context, request, response); }));}
    void SetMessageAllocatorFor_SendEvent(
        ::grpc::MessageAllocator< ::str2len::v1::Proxy_SendEventRequest, ::str2len::v1::Proxy_SendEventResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(0);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::str2len::v1::Proxy_SendEventRequest, ::str2len::v1::Proxy_SendEventResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_SendEvent() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendEvent(::grpc::ServerContext* /*context*/, const ::str2len::v1::Proxy_SendEventRequest* /*request*/, ::str2len::v1::Proxy_SendEventResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* SendEvent(
      ::grpc::CallbackServerContext* /*context*/, const ::str2len::v1::Proxy_SendEventRequest* /*request*/, ::str2len::v1::Proxy_SendEventResponse* /*response*/)  { return nullptr; }
  };
  typedef WithCallbackMethod_SendEvent<Service > CallbackService;
  typedef CallbackService ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_SendEvent : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_SendEvent() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_SendEvent() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendEvent(::grpc::ServerContext* /*context*/, const ::str2len::v1::Proxy_SendEventRequest* /*request*/, ::str2len::v1::Proxy_SendEventResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_SendEvent : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_SendEvent() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_SendEvent() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendEvent(::grpc::ServerContext* /*context*/, const ::str2len::v1::Proxy_SendEventRequest* /*request*/, ::str2len::v1::Proxy_SendEventResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestSendEvent(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_SendEvent : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_SendEvent() {
      ::grpc::Service::MarkMethodRawCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->SendEvent(context, request, response); }));
    }
    ~WithRawCallbackMethod_SendEvent() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status SendEvent(::grpc::ServerContext* /*context*/, const ::str2len::v1::Proxy_SendEventRequest* /*request*/, ::str2len::v1::Proxy_SendEventResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* SendEvent(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_SendEvent : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_SendEvent() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler<
          ::str2len::v1::Proxy_SendEventRequest, ::str2len::v1::Proxy_SendEventResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::str2len::v1::Proxy_SendEventRequest, ::str2len::v1::Proxy_SendEventResponse>* streamer) {
                       return this->StreamedSendEvent(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_SendEvent() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status SendEvent(::grpc::ServerContext* /*context*/, const ::str2len::v1::Proxy_SendEventRequest* /*request*/, ::str2len::v1::Proxy_SendEventResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedSendEvent(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::str2len::v1::Proxy_SendEventRequest,::str2len::v1::Proxy_SendEventResponse>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_SendEvent<Service > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_SendEvent<Service > StreamedService;
};

}  // namespace v1
}  // namespace str2len


#endif  // GRPC_sl_2eproto__INCLUDED
