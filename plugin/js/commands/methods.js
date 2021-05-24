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
        ) : []
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
      startUpCommands:
        command ? {
          used: true,
          startUpCommands: [
            {
              command: "npm ",
              args: ["start"],
            },
          ],
        } : {
          used: false,
          startUpCommands: []
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
      seedCommands: command.length ? {
        used: true,
        seedCommands: command.map(
          element => ({
            command: "npm run ",
            args: element.split(' '),
          })
        )
      } : {
        used: false,
        seedCommands: []
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
        if (element.search("migrat") != -1) command.push(element);
      });
    }
    callback(null, {
      migrationCommands: command.length ? {
        used: true,
        migrationCommands: command.map(
          element => ({
            command: "npm run ",
            args: element.split(' '),
          })
        )
      } : {
        used: false,
        migrationCommands: []
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
        ) : []
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
