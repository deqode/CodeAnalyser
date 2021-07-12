const depcheck = require("depcheck");
const fs = require("fs");
const common = require("../common");

function runPreDetect(input, callback) {
  let path = input.request.rootPath;
  var packageJsonFile = common.requireJsonFile(
    path + "/package.json",
    callback,
    "package.json file not found"
  );
  if (packageJsonFile) {
    const options = {
      skipMissing: false,
      ignoreDirs: ["node_modules"],
      detectors: Object.values(depcheck.detector),
    };
    depcheck(path, options, async (unused) => {
      removeUnusedFromPackageJson(unused, packageJsonFile);
      await addMissingToPackageJson(
        Object.keys(unused.missing),
        packageJsonFile
      );
      writePackageJsonFile(packageJsonFile, path).then(
        () => callback(null, { error: null }),
        (err) =>
          callback(null, {
            error: {
              message: err,
            },
          })
      );
    });
  }
}

/**
 * remove unused dependencies from package.json file object for example
 * express in npm dependency but never used
 * @param {depcheck.Results} unused - result of depcheck method
 * @param {JSON} packageJsonFile - package.json file object
 */
function removeUnusedFromPackageJson(unused, packageJsonFile) {
  if (unused.dependencies.length || unused.devDependencies.length) {
    unused.dependencies.forEach((element) => {
      delete packageJsonFile.dependencies[element];
    });
    unused.devDependencies.forEach((element) => {
      delete packageJsonFile.devDependencies[element];
    });
  }
}

/**
 * write json object to package.json file
 * @param {JSON} packageJsonFile  package.json file object which we have to write
 * @param {string} path path of repository
 * @returns {Promise<string>}  resolve on no error otherwise reject
 */
function writePackageJsonFile(packageJsonFile, path) {
  var promise = new Promise((resolve, reject) => {
    fs.writeFile(
      path + "/package.json",
      JSON.stringify(packageJsonFile, null, 4),
      (err) => {
        if (err) reject(err);
        resolve("");
      }
    );
  });
  return promise;
}

/**
 * add missing dependencies to package.json file object
 * @param {Array<string>} missing - array of name of dependencies which are missing
 * @param {JSON} packageJsonFile - package.json file object where we will add dependencies
 */
async function addMissingToPackageJson(missing, packageJsonFile) {
  if (missing.length) {
    const latestVersion = require("latest-version");
    for (let index = 0; index < missing.length; index++) {
      try {
        packageJsonFile.dependencies[missing[index]] = await latestVersion(missing[index]);
      } catch (error) {
        // handle this
      }
    }
  }
}

module.exports = {
  runPreDetect: runPreDetect,
};
