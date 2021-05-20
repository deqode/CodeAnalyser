const common = require("../common");

function getDependencies(input, callback) {
  let path = input.request.root;
  let parsedFile = common.requireJsonFile(
    path + "/package.json",
    callback,
    "package.json file not found"
  );
  if (parsedFile) {
    parsedFile.dependencies
      ? callback(null, { value: parsedFile.dependencies, error: null })
      : callback(null, { value: {}, error: null });
  }
}

module.exports = {
  getDependencies: getDependencies,
};
