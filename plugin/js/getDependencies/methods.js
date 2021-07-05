const common = require("../common");

function getDependencies(input, callback) {
  let path = input.request.rootPath;
  let parsedFile = common.requireJsonFile(
    path + "/package.json",
    callback,
    "package.json file not found"
  );
  if (parsedFile) {
    devDependencies = {};
    parsedFile.devDependencies ? devDependencies = { ...parsedFile.devDependencies } : {}
    parsedFile.dependencies
      ? callback(null, { value: {...parsedFile.dependencies,...devDependencies}, error: null })
      : callback(null, { value: {}, error: null });
  }
}

module.exports = {
  getDependencies: getDependencies,
};
