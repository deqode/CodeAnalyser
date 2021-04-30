let health = require("grpc-health-check");

const statusMap = {
  ServiceFoo: proto.grpc.health.v1.HealthCheckResponse.ServingStatus.SERVING,
  ServiceBar:
    proto.grpc.health.v1.HealthCheckResponse.ServingStatus.NOT_SERVING,
  "": proto.grpc.health.v1.HealthCheckResponse.ServingStatus.NOT_SERVING,
};
let healthImpl = new health.Implementation(statusMap);

module.exports = { healthe: health, healthImpl };
