const fs = require("fs");

function getDependencies(input, callback) {
  let path = input.request.root;
  let promise = new Promise((resolve, reject) => {
    fs.readdir(path, (err, data) => {
      if (err) reject(err);
      if (data.find((element) => element === "package.json")) {
        let rawFile = fs.readFileSync(path + "/package.json");
        let parsedFile = JSON.parse(rawFile);
        parsedFile.dependencies
          ? resolve(parsedFile.dependencies)
          : resolve({});
      } else reject("package.json file not found ");
    });
  });

  promise.then(
    (resolve) => callback(null, { value: resolve, error: null }),
    (reject) => callback(reject, { value: null, error: reject })
  );
}

module.exports = {
  getDependencies: getDependencies,
};
