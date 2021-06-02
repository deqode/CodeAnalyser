const common = require("../common");

function detect(input, callback) {
  let path = input.request.root;
  let jsonFile = common.requireJsonFile(path + "/package.json", callback);
  if (jsonFile) {
    let command = [];
    if (jsonFile.scripts) {
      Object.keys(jsonFile.scripts).some((element) => {
        if (element.search("test") != -1) command.push(element);
      });
    }
    callback(null, {
      commands: command.length ? command.map(element => ({
        command: "npm",
        args: [element],
      }))
        : [],
      error: null,
    });
  }
}

module.exports = {
  detect: detect,
};