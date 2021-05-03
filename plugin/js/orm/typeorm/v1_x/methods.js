
function detect(input, callback) {
  let path = input.request.root;
  callback(null, {
    used: true,
    error: null,
  });
}

function isORMUsed(input, callback) {
  let path = input.request.root;
  callback(null, {
    value: true,
    error: null,
  });
}

function PercentOfORMUsed(input, callback) {
  let path = input.request.root;
  callback(null, {
    error: null,
  });
}
module.exports = {
  detect: detect,
  isUsed: isORMUsed,
  percentOfUsed: isORMUsed
};