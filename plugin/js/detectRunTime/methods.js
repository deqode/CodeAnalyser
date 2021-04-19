const fs = require("fs");

function detectRuntime(input, callback) {
  let path = input.request.value;
  let promise = new Promise((resolve, reject) => {
    fs.readdir(path, (err, data) => {
      if (err) reject(err);
      if (data.find((element) => element === "package.json")) {
        let rawFile = fs.readFileSync(path + "/package.json");
        let parsedFile = JSON.parse(rawFile);
        if (parsedFile.engines && parsedFile.engines.node)
          resolve(parsedFile.engines.node);
        else resolve(process.version.slice(1));
      }
    });
  });

  promise.then(
    (resolve) => callback(null, { value: resolve, error: null }),
    (reject) => callback(reject, { value: null, error: reject })
  );
}

module.exports = {
  detectRuntime: detectRuntime,
};
