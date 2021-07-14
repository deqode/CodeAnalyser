const fs = require("fs");
const common = require("../common");

function detectBuildCommands(input, callback) {
  let path = input.request.value;
  let jsonFile = common.requireJsonFile(path + "/package.json", callback);
  if (jsonFile) {
    let command = [];
    if (jsonFile.scripts) {
      Object.keys(jsonFile.scripts).some((element) => {
        if (element.search("build") != -1) command.push(element);
      });
    }
    callback(null, {
      commands: command.length ? command.map(element => ({
        command: "npm run ",
        args: element.split(' '),
      })
      ) : [],
      used: command.length ? true : false,
      error: null,
    });
  }
}

function detectStartUpCommands(input, callback) {
  let path = input.request.value;
  let jsonFile = common.requireJsonFile(path + "/package.json", callback);
  if (jsonFile) {
    let command = jsonFile.scripts
      ? jsonFile.scripts.start
        ? jsonFile.scripts.start
        : false
      : false;
    callback(null, {
      commands: command ?
        [
          {
            command: "npm ",
            args: ["start"],
          },
        ] : [],
      used: command ? true : false,
      error: null,
    });
  }
}

function detectSeedCommands(input, callback) {
  let path = input.request.value;
  let jsonFile = common.requireJsonFile(path + "/package.json", callback);
  if (jsonFile) {
    let command = [];
    if (jsonFile.scripts) {
      Object.keys(jsonFile.scripts).some((element) => {
        if (element.search("seed") != -1) command.push(element);
      });
    }
    callback(null, {
      commands: command.length ? command.map(element => ({
        command: "npm run ",
        args: element.split(' '),
      })
      ) : [],
      used: command.length ? true : false,
      error: null,
    });
  }
}

//TODO implement logic in this method
function detectMigrationCommands(input, callback) {
  let path = input.request.value;
  let jsonFile = common.requireJsonFile(path + "/package.json", callback);
  if (jsonFile) {
    let command = [];
    if (jsonFile.scripts) {
      Object.keys(jsonFile.scripts).some((element) => {
        if (element.search("migrat") != -1) command.push(element);
      });
    }
    callback(null, {
      commands: command.length ? command.map(element => ({
        command: "npm run ",
        args: element.split(' '),
      })
      ) : [],
      used: command.length ? true : false,
      error: null,
    });
  }
}

//TODO implement logic in this method
function detectAdHocScripts(input, callback) {
  let path = input.request.value;
  let jsonFile = common.requireJsonFile(path + "/package.json", callback);
  if (jsonFile) {
    let command = [];
    if (jsonFile.scripts) {
      Object.keys(jsonFile.scripts).some((element) => {
        if (element.search("adhoc") != -1) command.push(element);
      });
    }
    callback(null, {
      commands: command.length ? command.map(element => ({
        command: "npm run ",
        args: element.split(' '),
      })
      ) : [],
      used: command.length ? true : false,
      error: null,
    });
  }
  //callback(reject, { value: null, error: reject });
}

module.exports = {
  detectBuildCommands: detectBuildCommands,
  detectStartUpCommands: detectStartUpCommands,
  detectSeedCommands: detectSeedCommands,
  detectMigrationCommands: detectMigrationCommands,
  detectAdHocScripts: detectAdHocScripts,
};
