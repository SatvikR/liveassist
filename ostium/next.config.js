const withPlugins = require("next-compose-plugins");
const withTM = require("next-transpile-modules")(["@liveassist/liber"]);

module.exports = withPlugins([withTM]);
