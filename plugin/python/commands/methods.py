from protos.plugin import commands_pb2_grpc, common_pb2, commands_pb2, buildCommands_pb2, command_pb2


class Methods(commands_pb2_grpc.CommandsServiceServicer):
    def DetectBuildCommands(self, request, context):
        return commands_pb2.ServiceOutputDetectBuildCommands(error=None,
                                                             buildCommands=buildCommands_pb2.BuildCommandsOutput(
                                                                 used=True, buildCommands=[
                                                                     command_pb2.Command(command="bolo",
                                                                                         args=["kjjljl"])]
                                                             ))

    def DetectStartUpCommands(self, request, context):
        pass

    def DetectSeedCommands(self, request, context):
        pass

    def DetectMigrationCommands(self, request, context):
        pass

    def DetectAdHocScripts(self, request, context):
        pass
