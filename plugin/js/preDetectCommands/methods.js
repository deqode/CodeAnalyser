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
    removeUnusedFromPackageJson(unused, packageJsonFile,path).then(
      (resolve) => {
        installMissingDependencies(Object.keys(unused.missing), path);
        callback(null, { error: null });
      },
      (err) => callback(err, { error: err })
    );
  });
}

function removeUnusedFromPackageJson(unused, packageJsonFile,path) {
  return new Promise((resolve, reject) => {
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
}

function installMissingDependencies(missing, path) {
  if (missing.length)
    childProcess.exec(
      `cd ${path} && npm i ${missing.toString().replace(",", " ")} `
    );
}

module.exports = {
  runPreDetect: runPreDetect,
};
