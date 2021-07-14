from protos.plugin import db_pb2_grpc, common_pb2, db_pb2


class Methods(db_pb2_grpc.DbServiceServicer):
    def Detect(self, request, context):
        return db_pb2.ServiceOutputBoolInt(value=True, error=None)

    def IsDbUsed(self, request, context):
        return common_pb2.ServiceOutputBool(value=True, error=None)

    def PercentOfDbUsed(self, request, context):
        return common_pb2.ServiceOutputFloat(value=55.6, error=None)
