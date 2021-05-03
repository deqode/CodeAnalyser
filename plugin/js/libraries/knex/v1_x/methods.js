
function detect(input, callback) {
  let path = input.request.root;
  callback(null, {
    value: true,
    error: null,
  });
}

function isUsed(input, callback) {
  let path = input.request.root;
  callback(null, {
    value: true,
    error: null,
  });
}

function percentOfUsed(input, callback) {
  let path = input.request.root;
  callback(null, {
    error: null,
  });
}
module.exports = {
  detect: detect,
  isUsed: isUsed,
  percentOfUsed: percentOfUsed
};