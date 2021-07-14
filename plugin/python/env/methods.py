from protos.plugin import env_pb2_grpc, env_pb2


class Methods(env_pb2_grpc.EnvServiceServicer):
    def Detect(self, request, context):
        return env_pb2.ServiceOutputEnv(envKeyValues={
            "name": "tera baap"
        })
