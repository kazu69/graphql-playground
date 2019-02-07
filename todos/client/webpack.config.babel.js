import * as path from 'path';
import * as webpack from 'webpack';
import HtmlWebpackPlugin from 'html-webpack-plugin';
import CleanWebpackPlugin from 'clean-webpack-plugin';


export default {
  mode: 'development',
  cache: true,
  entry: {
    app: './src/index.tsx',
  },
  output: {
    path: path.resolve(__dirname, 'dist'),
    publicPath: '/',
    filename: '[name].[hash].js',
    chunkFilename: '[name].[chunkhash].[id].js',
  },
  resolve: {
    extensions: ['.ts', '.tsx', '.js', '.json', '.mjs'],
    alias: { '~': path.resolve(__dirname, 'src') },
  },
  plugins: [
    new webpack.EnvironmentPlugin({
      NODE_ENV: JSON.stringify('development'),
    }),
    new webpack.HotModuleReplacementPlugin(),
    new webpack.NamedModulesPlugin(),
    new CleanWebpackPlugin(['dist']),
    new HtmlWebpackPlugin({
      title: 'Todo App',
      minify: true,
      cache: true,
      template: 'src/index.html',
    })
  ],
  module: {
    rules: [
      {
        test: /\.mjs$/,
        include: /node_modules/,
        type: 'javascript/auto'
      },
      {
        test: /\.ts(x?)$/,
        exclude: /node_modules/,
        use: 'ts-loader',
      },
      {
        test: /\.(js|jsx)$/,
        use: 'babel-loader'
      },
      {
        enforce: 'pre',
        test: /\.js$/,
        loader: 'source-map-loader'
      },
      {
        test: /\.html$/,
        loader: 'html-loader'
      },
    ]
  },
  devServer: {
    before: app => app.get('/favicon.ico', (_, res) => res.status(200).send()),
    historyApiFallback: true,
    contentBase: './dist',
    hot: true,
    publicPath: '/',
    host: '0.0.0.0',
    port: 9000,
    proxy: { '/graphql': 'http://localhost:10001' },
    clientLogLevel: 'info',
  }
};
