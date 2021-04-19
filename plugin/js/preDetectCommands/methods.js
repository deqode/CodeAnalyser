const depcheck = require("depcheck");
const fs = require("fs");
let childProcess = require("child_process");

function runPreDetect(input, callback) {
  let path = input.request.root;
  console.log(path);
  var packageJsonFile = require(path + "/package.json");
  const options = {
    skipMissing: false,
    ignoreDirs: ["node_modules"],
    detectors: Object.values(depcheck.detector),
  };
  depcheck(path, options, (unused) => {
    removeUnusedFromPackageJson(unused, packageJsonFile, path).then(
      (resolve) => {
        installMissingDependencies(Object.keys(unused.missing), path);
        callback(null, { error: null });
      },
      (err) => callback(err, { error: err })
    );
  });
}

/**
 *
 * @param {depcheck.Results} unused - result of depcheck method
 * @param {JSON} packageJsonFile - package.json file object
 * @param {string} path - path of project
 * @returns {Promise<string>} 
 */
function removeUnusedFromPackageJson(unused, packageJsonFile, path) {
  var promise = new Promise((resolve, reject) => {
    if (unused.dependencies.length || unused.devDependencies.length) {
      unused.dependencies.forEach((element) => {
        delete packageJsonFile.dependencies[element];
      });
      unused.devDependencies.forEach((element) => {
        delete packageJsonFile.devDependencies[element];
      });

      fs.writeFile(
        path + "/package.json",
        JSON.stringify(packageJsonFile, null, 4),
        (err) => {
          if (err) reject(err);
          resolve("");
        }
      );
    } else resolve("");
  });
  return promise;
}

/**
 * @param {Array<string>} missing - array of name of dependencies which are missing
 * @param {string} path - path where we have to install dependencies
 */
function installMissingDependencies(missing, path) {
  if (missing.length)
    childProcess.exec(
      `cd ${path} && npm i ${missing.toString().replace(",", " ")} `
    );
}

module.exports = {
  runPreDetect: runPreDetect,
};
