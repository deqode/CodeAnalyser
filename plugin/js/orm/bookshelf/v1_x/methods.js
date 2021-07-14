
var common = require("../../../common");

function detect(input, callback) {
  let path = input.request.rootPath;
  common.requirePathCheck(
    path,
    callback,
    "root path not found"
  );
  callback(null, {
    used: true,
    error: null,
  });
}

function IsUsed(input, callback) {
  let path = input.request.rootPath;
  common.requirePathCheck(
    path,
    callback,
    "root path not found"
  );
  callback(null, {
    value: true,
    error: null,
  });
}

function percentOfUsed(input, callback) {
  let path = input.request.rootPath;
  common.requirePathCheck(
    path,
    callback,
    "root path not found"
  );
  callback(null, {
    error: null,
  });
}
module.exports = {
  detect: detect,
  IsUsed: IsUsed,
  percentOfUsed: percentOfUsed
};