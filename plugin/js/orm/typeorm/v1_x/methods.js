
var common= require("../../../common");
function detect(input, callback) {
  let path = input.request.root;
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

function isUsed(input, callback) {
  let path = input.request.root;
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
  let path = input.request.root;
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
  isUsed: isUsed,
  percentOfUsed: percentOfUsed
};