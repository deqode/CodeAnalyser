# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from protos.plugin import commands_pb2 as protos_dot_plugin_dot_commands__pb2


class CommandsServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.DetectBuildCommands = channel.unary_unary(
                '/proto.CommandsService/DetectBuildCommands',
                request_serializer=protos_dot_plugin_dot_commands__pb2.ServiceCommandsInput.SerializeToString,
                response_deserializer=protos_dot_plugin_dot_commands__pb2.ServiceOutputDetectBuildCommands.FromString,
                )
        self.DetectStartUpCommands = channel.unary_unary(
                '/proto.CommandsService/DetectStartUpCommands',
                request_serializer=protos_dot_plugin_dot_commands__pb2.ServiceCommandsInput.SerializeToString,
                response_deserializer=protos_dot_plugin_dot_commands__pb2.ServiceOutputStartUpCommands.FromString,
                )
        self.DetectSeedCommands = channel.unary_unary(
                '/proto.CommandsService/DetectSeedCommands',
                request_serializer=protos_dot_plugin_dot_commands__pb2.ServiceCommandsInput.SerializeToString,
                response_deserializer=protos_dot_plugin_dot_commands__pb2.ServiceOutputSeedCommands.FromString,
                )
        self.DetectMigrationCommands = channel.unary_unary(
                '/proto.CommandsService/DetectMigrationCommands',
                request_serializer=protos_dot_plugin_dot_commands__pb2.ServiceCommandsInput.SerializeToString,
                response_deserializer=protos_dot_plugin_dot_commands__pb2.ServiceMigrationCommands.FromString,
                )
        self.DetectAdHocScripts = channel.unary_unary(
                '/proto.CommandsService/DetectAdHocScripts',
                request_serializer=protos_dot_plugin_dot_commands__pb2.ServiceCommandsInput.SerializeToString,
                response_deserializer=protos_dot_plugin_dot_commands__pb2.ServiceAdHocScripts.FromString,
                )


class CommandsServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def DetectBuildCommands(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DetectStartUpCommands(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DetectSeedCommands(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DetectMigrationCommands(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DetectAdHocScripts(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_CommandsServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'DetectBuildCommands': grpc.unary_unary_rpc_method_handler(
                    servicer.DetectBuildCommands,
                    request_deserializer=protos_dot_plugin_dot_commands__pb2.ServiceCommandsInput.FromString,
                    response_serializer=protos_dot_plugin_dot_commands__pb2.ServiceOutputDetectBuildCommands.SerializeToString,
            ),
            'DetectStartUpCommands': grpc.unary_unary_rpc_method_handler(
                    servicer.DetectStartUpCommands,
                    request_deserializer=protos_dot_plugin_dot_commands__pb2.ServiceCommandsInput.FromString,
                    response_serializer=protos_dot_plugin_dot_commands__pb2.ServiceOutputStartUpCommands.SerializeToString,
            ),
            'DetectSeedCommands': grpc.unary_unary_rpc_method_handler(
                    servicer.DetectSeedCommands,
                    request_deserializer=protos_dot_plugin_dot_commands__pb2.ServiceCommandsInput.FromString,
                    response_serializer=protos_dot_plugin_dot_commands__pb2.ServiceOutputSeedCommands.SerializeToString,
            ),
            'DetectMigrationCommands': grpc.unary_unary_rpc_method_handler(
                    servicer.DetectMigrationCommands,
                    request_deserializer=protos_dot_plugin_dot_commands__pb2.ServiceCommandsInput.FromString,
                    response_serializer=protos_dot_plugin_dot_commands__pb2.ServiceMigrationCommands.SerializeToString,
            ),
            'DetectAdHocScripts': grpc.unary_unary_rpc_method_handler(
                    servicer.DetectAdHocScripts,
                    request_deserializer=protos_dot_plugin_dot_commands__pb2.ServiceCommandsInput.FromString,
                    response_serializer=protos_dot_plugin_dot_commands__pb2.ServiceAdHocScripts.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'proto.CommandsService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class CommandsService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def DetectBuildCommands(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.CommandsService/DetectBuildCommands',
            protos_dot_plugin_dot_commands__pb2.ServiceCommandsInput.SerializeToString,
            protos_dot_plugin_dot_commands__pb2.ServiceOutputDetectBuildCommands.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DetectStartUpCommands(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.CommandsService/DetectStartUpCommands',
            protos_dot_plugin_dot_commands__pb2.ServiceCommandsInput.SerializeToString,
            protos_dot_plugin_dot_commands__pb2.ServiceOutputStartUpCommands.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DetectSeedCommands(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.CommandsService/DetectSeedCommands',
            protos_dot_plugin_dot_commands__pb2.ServiceCommandsInput.SerializeToString,
            protos_dot_plugin_dot_commands__pb2.ServiceOutputSeedCommands.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DetectMigrationCommands(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.CommandsService/DetectMigrationCommands',
            protos_dot_plugin_dot_commands__pb2.ServiceCommandsInput.SerializeToString,
            protos_dot_plugin_dot_commands__pb2.ServiceMigrationCommands.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DetectAdHocScripts(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.CommandsService/DetectAdHocScripts',
            protos_dot_plugin_dot_commands__pb2.ServiceCommandsInput.SerializeToString,
            protos_dot_plugin_dot_commands__pb2.ServiceAdHocScripts.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)