from protos.plugin import framework_pb2_grpc, common_pb2


class Methods(framework_pb2_grpc.FrameworkServiceServicer):
    def Detect(self, request, context):
        return common_pb2.ServiceOutputBool(value=True, error=None)

    def IsFrameworkUsed(self, request, context):
        return common_pb2.ServiceOutputBool(value=True, error=None)

    def PercentOfFrameworkUsed(self, request, context):
        return common_pb2.ServiceOutputFloat(value=78, error=None)
