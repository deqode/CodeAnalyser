const fs = require("fs");

/**
 * check path exist or not
 * @param {string} path - path of any file
 * @param {Function} callback - callback function of grpc
 * @param {string} [message= path does'nt exist] - error message for callback 
 * @returns {JSON | null} -return jsonFile object if path exist,
 *  otherwise returns null and call callback with error
 */
function requirePathCheck(path, callback, message = "path doesn't exist ") {
  if (!fs.existsSync(path)) {
    callback(null, {
      value: null,
      error: {
        message: message,
        cause: `${path}`,
      },
    });
  }
  else
    return true
}

module.exports = {
  requirePathCheck: requirePathCheck,
  requireJsonFile: requireJsonFile,
};

/**
 * check path exist or not
 * @param {string} path - path of any file
 * @param {Function} callback - callback function of grpc
 * @param {string} [message= package.json file not found] - error message for callback 
 * @returns {JSON | null} -return jsonFile object if path exist,
 *  otherwise returns null and call callback with error
 */
function requireJsonFile(path, callback, message = "package.json file not found") {
  try {
    let jsonFile = require(path);
    return jsonFile;
  } catch (error) {
    callback(null, {
      value: null,
      error: {
        message: message,
        cause: `${path}`,
      },
    });
    return null;
  }
}