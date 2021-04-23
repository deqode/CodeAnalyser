const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const healthCheck = require("./healthCheck");
const getPort = require("get-port");
const methods = require("./methods");

let absPath = "";
const protoPath = [
  absPath + "protos/plugin/commands.proto",
  absPath + "protos/plugin/common.proto",
  absPath + "protos/output/global/startUpCommands.proto",
  absPath + "protos/output/global/buildCommands.proto",
  absPath + "protos/output/global/adHocScripts.proto",
  absPath + "protos/output/global/migrationCommands.proto",
  absPath + "protos/output/global/seedCommands.proto",
  absPath + "protos/output/global/command.proto",
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
server.addService(protoFile.proto.CommandsService.service, methods);

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
      }, 1000 * 60);
    }
  );
})();
