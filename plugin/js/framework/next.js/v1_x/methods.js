function detect(input, callback) {
  let path = input.request.root;
  callback(null, { value: true, error: null });
  // callback(reject, { value: null, error: reject })
}

function isFrameworkUsed(input, callback) {
  let path = input.request.root;
  callback(null, { value: true, error: null });
  // callback(reject, { value: null, error: reject })
}

function percentOfFrameworkUsed(input, callback) {
  let path = input.request.root;
  callback(null, { error: null });
  // callback(reject, { value: null, error: reject })
}

module.exports = {
  detect: detect,
  isFrameworkUsed: isFrameworkUsed,
  percentOfFrameworkUsed: percentOfFrameworkUsed,
};
