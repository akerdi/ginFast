module.exports = {
  lintOnSave: true,
  productionSourceMap: false,
  devServer: {
    open: true,
    host: '127.0.0.1',
    port: 8001,
    https: false,
    hotOnly: false,
    proxy:{
      '/api': {
        target: 'http://127.0.0.1:8000/',
        secure: false,
        changeOrigin: true,
        pathRewrite: {
          '^/api': ''
        }
      }
    }
  }
};