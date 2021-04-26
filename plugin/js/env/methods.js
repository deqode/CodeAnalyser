const fs = require("fs");
const path = require("path");
const dotenv = require("dotenv");

function detect(input, callback) {
  let dirPath = input.request.root;
  let envKeyValues = {};

  walkDir(dirPath, function (name, filePath) {
    if (name === ".env") {
      envKeyValues = {
        ...envKeyValues,
        ...dotenv.config({ path: filePath }).parsed,
      };
    }
  });
  callback(null, { envKeyValues: envKeyValues, error: null, keys: [] });
}

module.exports = {
  detect: detect,
};

/**
 * search all files recursively exclude node_modules folder
 * @param {string} dir path of directory
 * @param {Function} callback function which will call on every file found
 */
function walkDir(dir, callback) {
  fs.readdirSync(dir).forEach((f) => {
    let dirPath = path.join(dir, f);
    let isDirectory = fs.statSync(dirPath).isDirectory();
    isDirectory
      ? f !== "node_modules"
        ? walkDir(dirPath, callback)
        : ""
      : callback(f, dirPath);
  });
}
