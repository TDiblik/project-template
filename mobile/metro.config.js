const path = require("path");
const {getDefaultConfig} = require("expo/metro-config");
const {withNativewind} = require("nativewind/metro");

const projectRoot = __dirname;
const workspaceRoot = path.resolve(projectRoot, "..");

/** @type {import('expo/metro-config').MetroConfig} */
const config = getDefaultConfig(projectRoot);
config.watchFolders = [...config.watchFolders, path.resolve(workspaceRoot, "shared/fe/api-client")];
config.resolver.extraNodeModules = {
  ...config.resolver.extraNodeModules,
  "@shared/api-client": path.resolve(workspaceRoot, "shared/fe/api-client/dist"),
};

module.exports = withNativewind(config);
