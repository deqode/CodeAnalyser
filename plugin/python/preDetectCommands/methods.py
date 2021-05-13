from protos.plugin import predetectcommand_pb2_grpc, common_pb2


class Methods(predetectcommand_pb2_grpc.PreDetectCommandServicer):
    def RunPreDetect(self, request, context):
        return common_pb2.ServiceEmptyOutput(error=None)
