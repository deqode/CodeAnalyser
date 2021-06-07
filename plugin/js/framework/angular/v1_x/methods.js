const fs = require('fs');
const common= require("../../../common");

function detect(input, callback) {
  let path = input.request.root;
  common.requirePathCheck(
    path,
    callback,
    "root path not found"
  );
  callback(null, { value: true, error: null });
  // callback(reject, { value: null, error: reject })
}

function isUsed(input, callback) {
  let path = input.request.root;
  common.requirePathCheck(
    path,
    callback,
    "root path not found"
  );
  if (fs.existsSync(`${path}/angular.json`))
    callback(null, { value: true, error: null });
  else
    callback(null, { value: false, error: null });
}

function percentOfUsed(input, callback) {
  let path = input.request.root;
  common.requirePathCheck(
    path,
    callback,
    "root path not found"
  );
  callback(null, { error: null });
  // callback(reject, { value: null, error: reject })
}

module.exports = {
  detect: detect,
  isUsed: isUsed,
  percentOfUsed: percentOfUsed,
};
