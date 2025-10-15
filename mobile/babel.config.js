module.exports = function (api) {
  api.cache(true);
  return {
    presets: ["babel-preset-expo"],
    plugins: [
      [
        "module-resolver",
        {
          alias: {
            "@shared/api-client": "../shared/fe/api-client/dist/index",
          },
        },
      ],
    ],
  };
};
