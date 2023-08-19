# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import conv_pb2 as conv__pb2


class ColorConvServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.FromHsv = channel.unary_unary(
                '/color.hsv.v1.ColorConvService/FromHsv',
                request_serializer=conv__pb2.FromHsvRequest.SerializeToString,
                response_deserializer=conv__pb2.FromHsvResponse.FromString,
                )


class ColorConvServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def FromHsv(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ColorConvServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'FromHsv': grpc.unary_unary_rpc_method_handler(
                    servicer.FromHsv,
                    request_deserializer=conv__pb2.FromHsvRequest.FromString,
                    response_serializer=conv__pb2.FromHsvResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'color.hsv.v1.ColorConvService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ColorConvService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def FromHsv(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/color.hsv.v1.ColorConvService/FromHsv',
            conv__pb2.FromHsvRequest.SerializeToString,
            conv__pb2.FromHsvResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
