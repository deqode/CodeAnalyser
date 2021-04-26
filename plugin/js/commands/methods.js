const fs = require("fs");
const common = require("../common");

function detectBuildCommands(input, callback) {
  let path = input.request.root;
  let jsonFile = common.requirePathCheck(path + "/package.json", callback);
  if (jsonFile) {
    let command;
    if (jsonFile.scripts) {
      Object.keys(jsonFile.scripts).some((element) => {
        if (element.search("build") != -1) command = element;
      });
    }
    callback(null, {
      buildCommands: {
        used: command ? true : false,
        buildCommands: [
          {
            command: command ? "npm run " : "",
            args: command ? [command] : [],
          },
        ],
      },
      error: null,
    });
  }
}

function detectStartUpCommands(input, callback) {
  let path = input.request.root;
  let jsonFile = common.requirePathCheck(path + "/package.json", callback);
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
  let jsonFile = common.requirePathCheck(path + "/package.json", callback);
  if (jsonFile) {
    let command;
    if (jsonFile.scripts) {
      Object.keys(jsonFile.scripts).some((element) => {
        if (element.search("seed") != -1) command = element;
      });
    }
    callback(null, {
      seedCommands: {
        used: command ? true : false,
        seedCommands: [
          {
            command: command ? "npm run " : "",
            args: command ? [command] : [],
          },
        ],
      },
      error: null,
    });
  }
}

function detectMigrationCommands(input, callback) {
  let path = input.request.root;
  let jsonFile = common.requirePathCheck(path + "/package.json", callback);
  if (jsonFile) {
    let command;
    if (jsonFile.scripts) {
      Object.keys(jsonFile.scripts).some((element) => {
        if (element.search("build") != -1) command = element;
      });
    }
    callback(null, {
      migrationCommands: {
        used: command ? true : false,
        migrationCommand: [
          {
            command: command ? "npm run " : "",
            args: command ? [command] : [],
          },
        ],
      },
      error: null,
    });
  }
}

function detectAdHocScripts(input, callback) {
  let path = input.request.root;
  let jsonFile = common.requirePathCheck(path + "/package.json", callback);
  if (jsonFile) {
    let command;
    if (jsonFile.scripts) {
      Object.keys(jsonFile.scripts).some((element) => {
        if (element.search("build") != -1) command = element;
      });
    }
    callback(null, {
      adHocScripts: {
        used: command ? true : false,
        adHocScripts: [
          {
            command: command ? "npm run " : "",
            args: command ? [command] : [],
          },
        ],
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
