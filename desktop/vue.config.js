const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  configureWebpack: {
    devtool: 'cheap-module-source-map',
  },
  devServer: {
    client: {
      webSocketURL: {
        hostname: '0.0.0.0',
        pathname: '/ws',
        password: 'dev-server',
        port: 8080,
        protocol: 'ws',
        username: 'webpack',
      },
    },
    proxy: {
      '^/api': {
        target: 'http://localhost',
        changeOrigin: true
      }
    }
  },
})
