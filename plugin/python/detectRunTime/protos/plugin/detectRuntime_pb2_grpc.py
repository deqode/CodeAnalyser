# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from protos.plugin import common_pb2 as protos_dot_plugin_dot_common__pb2


class DetectRuntimeServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.DetectRuntime = channel.unary_unary(
                '/proto.DetectRuntimeService/DetectRuntime',
                request_serializer=protos_dot_plugin_dot_common__pb2.ServiceInputString.SerializeToString,
                response_deserializer=protos_dot_plugin_dot_common__pb2.ServiceOutputString.FromString,
                )


class DetectRuntimeServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def DetectRuntime(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DetectRuntimeServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'DetectRuntime': grpc.unary_unary_rpc_method_handler(
                    servicer.DetectRuntime,
                    request_deserializer=protos_dot_plugin_dot_common__pb2.ServiceInputString.FromString,
                    response_serializer=protos_dot_plugin_dot_common__pb2.ServiceOutputString.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'proto.DetectRuntimeService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class DetectRuntimeService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def DetectRuntime(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.DetectRuntimeService/DetectRuntime',
            protos_dot_plugin_dot_common__pb2.ServiceInputString.SerializeToString,
            protos_dot_plugin_dot_common__pb2.ServiceOutputString.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
