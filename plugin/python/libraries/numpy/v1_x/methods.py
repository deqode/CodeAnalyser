from protos.plugin import library_pb2_grpc, common_pb2


class Methods(library_pb2_grpc.LibraryServiceServicer):
    def Detect(self, request, context):
        return common_pb2.ServiceOutputBool(value=True, error=None)

    def IsUsed(self, request, context):
        return common_pb2.ServiceOutputBool(value=True, error=None)

    def PercentOfUsed(self, request, context):
        return common_pb2.ServiceOutputFloat(value=78, error=None)
