module.exports = {
  // For multi-page mode
  pages: {
    index: {
      // entry for the page
      entry: 'src/main.js',
      // the source template
      template: 'public/index.html',
      // output as dist/index.html
      filename: 'index.html',
      // when using title option,
      // template title tag needs to be <title><%= htmlWebpackPlugin.options.title %></title>
      title: 'Home',
      // chunks to include on this page, by default includes
      // extracted common chunks and vendor chunks.
      chunks: ['chunk-vendors', 'chunk-common', 'index']
    }
  },
  devServer: {
    proxy: {
      '/api': {
        target: 'https://localhost:3003',
        changeOrigin: true,
        cookieDomainRewrite: '',
        headers: {
          'X-COOKIE-AUTH-SECURE': false
        },
        ws: true,
        secure: false
      }
    },
    historyApiFallback: {
      rewrites: [
      ]
    }
  },

  // For i18n support
  // refs:
  // http://kazupon.github.io/vue-i18n/guide/sfc.html#vue-cli-3-0-beta
  // https://github.com/kazupon/vue-i18n-loader/issues/17#issuecomment-411750335
  chainWebpack: config => {
    config.module
      .rule('i18n')
      .resourceQuery(/blockType=i18n/)
      .use('i18n')
        .loader('@kazupon/vue-i18n-loader')
        .end()
      .use('yaml')
        .loader('yaml-loader')
        .end()
  }
}
