
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

function isOrmUsed(input, callback) {
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

function percentOfORMUsed(input, callback) {
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
  isOrmUsed: isOrmUsed,
  percentOfORMUsed: percentOfORMUsed
};