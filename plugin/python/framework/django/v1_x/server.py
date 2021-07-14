import sys

from methods import Methods
from protos.plugin import framework_pb2_grpc
from random_open_port import random_port

from grpc_health.v1 import health_pb2, health_pb2_grpc, health as healthServicer

import grpc
from concurrent import futures


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    framework_pb2_grpc.add_FrameworkServiceServicer_to_server(Methods(), server)

    health = healthServicer.HealthServicer()
    health.set("plugin", health_pb2.HealthCheckResponse.ServingStatus.Value('SERVING'))
    health_pb2_grpc.add_HealthServicer_to_server(health, server)
    port = random_port()
    serving_address = f'127.0.0.1:{port}'
    server.add_insecure_port(serving_address)
    server.start()
    print(f'1|1|tcp|{serving_address}|grpc')
    sys.stdout.flush()
    server.wait_for_termination(timeout=60 * 4)


if __name__ == '__main__':
    serve()
