const fs = require("fs");
const common = require("../common");

function detectBuildCommands(input, callback) {
  let path = input.request.root;
  let jsonFile = common.requireJsonFile(path + "/package.json", callback);
  if (jsonFile) {
    let command = [];
    if (jsonFile.scripts) {
      Object.keys(jsonFile.scripts).some((element) => {
        if (element.search("build") != -1) command.push(element);
      });
    }
    callback(null, {
      buildCommands: {
        used: command.length ? true : false,
        buildCommands: command.length ? command.map(element => ({
          command: "npm run ",
          args: element.split(' '),
        })
        ) : [{
          command: "",
          args: [],
        }]
      },
      error: null,
    });
  }
}

function detectStartUpCommands(input, callback) {
  let path = input.request.root;
  let jsonFile = common.requireJsonFile(path + "/package.json", callback);
  if (jsonFile) {
    let command = jsonFile.scripts
      ? jsonFile.scripts.start
        ? jsonFile.scripts.start
        : false
      : false;
    callback(null, {
      startUpCommands: {
        used: command,
        startUpCommands: [
          {
            command: command ? "npm" : "",
            args: command ? [" start"] : [],
          },
        ],
      },
      error: null,
    });
  }
}

function detectSeedCommands(input, callback) {
  let path = input.request.root;
  let jsonFile = common.requireJsonFile(path + "/package.json", callback);
  if (jsonFile) {
    let command = [];
    if (jsonFile.scripts) {
      Object.keys(jsonFile.scripts).some((element) => {
        if (element.search("seed") != -1) command.push(element);
      });
    }
    callback(null, {
      seedCommands: {
        used: command.length ? true : false,
        seedCommands: command.length ? command.map(element => ({
          command: "npm run ",
          args: element.split(' '),
        })
        ) : [{
          command: "",
          args: [],
        }]
      },
      error: null,
    });
  }
}

//TODO implement logic in this method
function detectMigrationCommands(input, callback) {
  let path = input.request.root;
  let jsonFile = common.requireJsonFile(path + "/package.json", callback);
  if (jsonFile) {
    let command = [];
    if (jsonFile.scripts) {
      Object.keys(jsonFile.scripts).some((element) => {
        if (element.search("migrate") != -1) command.push(element);
      });
    }
    callback(null, {
      migrationCommands: {
        used: command.length ? true : false,
        migrationCommand: command.length ? command.map(element => ({
          command: "npm run ",
          args: element.split(' '),
        })
        ) : [{
          command: "",
          args: [],
        }]
      },
      error: null,
    });
  }
}

//TODO implement logic in this method
function detectAdHocScripts(input, callback) {
  let path = input.request.root;
  let jsonFile = common.requireJsonFile(path + "/package.json", callback);
  if (jsonFile) {
    let command = [];
    if (jsonFile.scripts) {
      Object.keys(jsonFile.scripts).some((element) => {
        if (element.search("adhoc") != -1) command.push(element);
      });
    }
    callback(null, {
      adHocScripts: {
        used: command.length ? true : false,
        adHocScripts: command.length ? command.map(element => ({
          command: "npm run ",
          args: element.split(' '),
        })
        ) : [{
          command: "",
          args: [],
        }]
      },
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
