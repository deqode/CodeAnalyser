# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from protos.plugin import common_pb2 as protos_dot_plugin_dot_common__pb2
from protos.plugin import db_pb2 as protos_dot_plugin_dot_db__pb2


class DbServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Detect = channel.unary_unary(
                '/proto.DbService/Detect',
                request_serializer=protos_dot_plugin_dot_common__pb2.ServiceInput.SerializeToString,
                response_deserializer=protos_dot_plugin_dot_db__pb2.ServiceOutputBoolInt.FromString,
                )
        self.IsDbUsed = channel.unary_unary(
                '/proto.DbService/IsDbUsed',
                request_serializer=protos_dot_plugin_dot_common__pb2.ServiceInput.SerializeToString,
                response_deserializer=protos_dot_plugin_dot_common__pb2.ServiceOutputBool.FromString,
                )
        self.PercentOfDbUsed = channel.unary_unary(
                '/proto.DbService/PercentOfDbUsed',
                request_serializer=protos_dot_plugin_dot_common__pb2.ServiceInput.SerializeToString,
                response_deserializer=protos_dot_plugin_dot_common__pb2.ServiceOutputFloat.FromString,
                )


class DbServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Detect(self, request, context):
        """boolint stand for bool-detect , int-port number of db
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def IsDbUsed(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PercentOfDbUsed(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DbServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Detect': grpc.unary_unary_rpc_method_handler(
                    servicer.Detect,
                    request_deserializer=protos_dot_plugin_dot_common__pb2.ServiceInput.FromString,
                    response_serializer=protos_dot_plugin_dot_db__pb2.ServiceOutputBoolInt.SerializeToString,
            ),
            'IsDbUsed': grpc.unary_unary_rpc_method_handler(
                    servicer.IsDbUsed,
                    request_deserializer=protos_dot_plugin_dot_common__pb2.ServiceInput.FromString,
                    response_serializer=protos_dot_plugin_dot_common__pb2.ServiceOutputBool.SerializeToString,
            ),
            'PercentOfDbUsed': grpc.unary_unary_rpc_method_handler(
                    servicer.PercentOfDbUsed,
                    request_deserializer=protos_dot_plugin_dot_common__pb2.ServiceInput.FromString,
                    response_serializer=protos_dot_plugin_dot_common__pb2.ServiceOutputFloat.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'proto.DbService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class DbService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Detect(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.DbService/Detect',
            protos_dot_plugin_dot_common__pb2.ServiceInput.SerializeToString,
            protos_dot_plugin_dot_db__pb2.ServiceOutputBoolInt.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def IsDbUsed(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.DbService/IsDbUsed',
            protos_dot_plugin_dot_common__pb2.ServiceInput.SerializeToString,
            protos_dot_plugin_dot_common__pb2.ServiceOutputBool.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PercentOfDbUsed(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.DbService/PercentOfDbUsed',
            protos_dot_plugin_dot_common__pb2.ServiceInput.SerializeToString,
            protos_dot_plugin_dot_common__pb2.ServiceOutputFloat.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
