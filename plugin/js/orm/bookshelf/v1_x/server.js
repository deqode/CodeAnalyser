const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const healthCheck = require("./healthCheck");
const getPort = require("get-port");
const methods = require("./methods");

const protoPath = [
  __dirname + "../../../../../../protos/plugin/orm.proto",
  __dirname + "../../../../../../protos/plugin/helpers.proto",
];

//load proto file
const packageDefinition = protoLoader.loadSync(protoPath, {
  keepCase: true,
  defaults: true,
  objects: true,
  oneofs: true,
});
const protoFile = grpc.loadPackageDefinition(packageDefinition);

const server = new grpc.Server();

server.addService(healthCheck.healthe.service, healthCheck.healthImpl);
server.addService(protoFile.proto.OrmService.service, methods);

// server creation
(async () => {
  const PORT = await getPort();
  server.bindAsync(
    `127.0.0.1:${PORT}`,
    grpc.ServerCredentials.createInsecure(),
    (err, port) => {
      if (err) {
        console.error(err);
        return;
      }
      server.start();
      console.log(`1|1|tcp|127.0.0.1:${port}|grpc`);
      setTimeout(() => {
        server.forceShutdown();
      }, 1000 * 60 * 2);
    }
  );
})();
