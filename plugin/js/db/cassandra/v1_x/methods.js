function detect(input, callback) {
  let path = input.request.root;
  callback(null, { value: true, intValue: null, error: null });
  // callback(reject, { value: null, error: reject })
}

function isDbUsed(input, callback) {
  let path = input.request.root;
  callback(null, { value: true, error: null });
  // callback(reject, { value: null, error: reject })
}

function percentOfDbUsed(input, callback) {
  let path = input.request.root;
  callback(null, {  error: null });
  // callback(reject, { value: null, error: reject })
}

module.exports = {
  detect: detect,
  isDbUsed: isDbUsed,
  percentOfDbUsed: percentOfDbUsed,
};
