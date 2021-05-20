const common = require("../common");

function detect(input, callback) {
  let path = input.request.root;
  let jsonFile = common.requireJsonFile(path + "/package.json", callback);
  if (jsonFile) {
    let command = jsonFile.scripts ?
      jsonFile.scripts.test ?
        jsonFile.scripts.test :
        false :
      false;
    callback(null, {
      commands: [{
        command: command ? "npm" : "",
        args: command ? [" test"] : [],
      }],
      error: null,
    });
  }
}

module.exports = {
  detect: detect,
};