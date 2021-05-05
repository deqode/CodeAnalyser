
function detect(input, callback) {
  let path = input.request.root;
  callback(null, {
    used: true,
    error: null,
  });
}

function isOrmUsed(input, callback) {
  let path = input.request.root;
  callback(null, {
    value: true,
    error: null,
  });
}

function percentOfORMUsed(input, callback) {
  let path = input.request.root;
  callback(null, {
    error: null,
  });
}
module.exports = {
  detect: detect,
  isOrmUsed: isOrmUsed,
  percentOfORMUsed: percentOfORMUsed
};