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
  try {
    let jsonFile = require(path);
    return jsonFile;
  } catch (error) {
    console.log(error);
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

module.exports = {
  requirePathCheck: requirePathCheck,
};
