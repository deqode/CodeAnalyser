from protos.plugin import detectRuntime_pb2_grpc, common_pb2


class Methods(detectRuntime_pb2_grpc.DetectRuntimeServiceServicer):
    def DetectRuntime(self, request, context):
        return common_pb2.ServiceOutputString(value="b55", error=None)
