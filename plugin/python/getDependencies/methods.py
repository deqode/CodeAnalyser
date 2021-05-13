from protos.plugin import dependencies_pb2_grpc, common_pb2


class Methods(dependencies_pb2_grpc.DependenciesServiceServicer):
    def GetDependencies(self, request, context):
        return common_pb2.ServiceOutputStringMap(value={"grpcio": "5.6"}, error=None)
