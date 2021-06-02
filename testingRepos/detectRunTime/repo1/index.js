let a = {
    a: 8,
    b: 6
};

console.log(a);

function detect(input, callback) {
    let dirPath = input.request.root;
    common.requirePathCheck(
        dirPath,
        callback,
        "directory path not found" + dirPath.toString()
    );
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