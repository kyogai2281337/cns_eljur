const { defineConfig } = require('@vue/cli-service')
const fs = require('fs');
const path = require('path');

module.exports = defineConfig({
  transpileDependencies: true,
  configureWebpack: {
    devtool: 'cheap-module-source-map',
  },
  devServer: {
    https: {
      key: fs.readFileSync(path.resolve(__dirname, '../ssl/cert.key')),
      cert: fs.readFileSync(path.resolve(__dirname, '../ssl/cert.crt')),
    },
    client: {
      webSocketURL: {
        hostname: '0.0.0.0',
        pathname: '/ws',
        password: 'dev-server',
        port: 8080,
        protocol: 'wss',
        username: 'webpack',
      },
    },
    proxy: {
      '^/api': {
        target: 'https://localhost',
        changeOrigin: true
      }
    }
  },
  pluginOptions: {
    electronBuilder: {
      nodeIntegration: true,
      mainProcessFile: 'src/background.ts',
      builderOptions: {
        files: [
          "**/*",
          "!**/node_modules/*/{CHANGELOG.md,README.md,README,readme.md,readme}",
          "!**/node_modules/*/{test,__tests__,tests,powered-test,example,examples}",
          "!**/node_modules/*.d.ts",
          "!**/node_modules/.bin",
          "!**/*.{iml,o,hprof,orig,pyc,pyo,rbc,swp,csproj,sln,xproj}",
          "!.editorconfig",
          "!**/._*",
          "!**/{.DS_Store,.git,.hg,.svn,CVS,RCS,SCCS,.gitignore,.gitattributes}",
          "!**/{__pycache__,thumbs.db,.flowconfig,.idea,.vs,.nyc_output}",
          "!**/{appveyor.yml,.travis.yml,circle.yml}",
          "!**/{npm-debug.log,yarn.lock,.yarn-integrity,.yarn-metadata.json}"
        ]
      }
    }
  }
})