const depcheck = require("depcheck");
const fs = require("fs");

function runPreDetect(input, callback) {
  let path = input.request.root;
  var packageJsonFile = require(path + "/package.json");
  const options = {
    skipMissing: false,
    ignoreDirs: ["node_modules"],
    detectors: Object.values(depcheck.detector),
  };
  depcheck(path, options, async (unused) => {
    removeUnusedFromPackageJson(unused, packageJsonFile);
    await addMissingToPackageJson(Object.keys(unused.missing), packageJsonFile);
    writePackageJsonFile(packageJsonFile,path).then(
      () => callback(null, { error: null }),
      (err) => callback(err, { error: err })
    );
  });
}

/**
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
 *
 * @param {JSON} packageJsonFile - package.json file object which we have to write
 * @param {string} path - path of repository
 * @returns {Promise<string>} - resolve on no error otherwise reject
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
 * @param {Array<string>} missing - array of name of dependencies which are missing
 * @param {JSON} packageJsonFile - package.json file object where we will add dependencies
 */
async function addMissingToPackageJson(missing, packageJsonFile) {
  if (missing.length) {
    const latestVersion = require("latest-version");
    for (let index = 0; index < missing.length; index++) {
      packageJsonFile.dependencies[missing[index]] = await latestVersion(
        missing[index]
      );
    }
  }
}

module.exports = {
  runPreDetect: runPreDetect,
};
