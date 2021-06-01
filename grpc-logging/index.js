var PROTO_PATH =[ __dirname + '/repo.proto',__dirname + '/graphql.proto'];

var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });
var protoDescriptor = grpc.loadPackageDefinition(packageDefinition);
var helloworld = protoDescriptor.alfred;

function CloneRepository(call, callback) {
  callback(null, {
    message: 'Hello! ' + call.request.name
  });
}

function getServer() {
  var server = new grpc.Server();
  server.addService(helloworld.GitService.service, {
    CloneRepository: CloneRepository,
  });
  return server;
}

console.log(helloworld.GitService)


(async () => {
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