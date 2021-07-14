const common = require("../common");

function detect(input, callback) {
  let path = input.request.value;
  let parsedFile = common.requireJsonFile(
    path + "/package.json",
    callback,
    "package.json file not found"
  );
  if (parsedFile) {
    if (parsedFile.engines && parsedFile.engines.node)
      callback(null, { value: parsedFile.engines.node, error: null });
    else callback(null, { value: process.version.slice(1), error: null });
  }
}

module.exports = {
  detect: detect,
};
