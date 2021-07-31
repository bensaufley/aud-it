/* eslint-disable import/no-extraneous-dependencies */
import type {} from 'browserslist';
import HtmlWebpackPlugin from 'html-webpack-plugin';
import { resolve } from 'path';
import { Configuration, DefinePlugin, RuleSetRule } from 'webpack';

const extensions = ['.ts', '.tsx', '.js', '.jsx', '.json', '.mjs'];

const mode = process.env.NODE_ENV === 'development' ? 'development' : 'production';

const babelLoaderRule: RuleSetRule = {
  test: /\.[tj]sx?$/,
  loader: 'babel-loader',
  exclude: /node_modules/,
  resolve: {
    extensions,
  },
};

const postcssRule: RuleSetRule = {
  test: /\.css$/,
  use: [
    'style-loader',
    {
      loader: 'css-loader',
      options: {
        modules: {
          auto: true,
        },
      },
    },
    {
      loader: 'postcss-loader',
      options: {
        postcssOptions: {
          plugins: [
            [
              'postcss-preset-env',
              {
                stage: 0,
                features: {
                  'nesting-rules': true,
                },
              },
            ],
          ],
        },
      },
    },
  ],
};

const alias = {
  '~components': resolve(__dirname, 'src/components/'),
  '~contexts': resolve(__dirname, 'src/contexts/'),
  '~graphql': resolve(__dirname, 'src/graphql/'),
  '~hooks': resolve(__dirname, 'src/hooks/'),
  '~lib': resolve(__dirname, 'src/lib/'),
  '~mocks': resolve(__dirname, '__mocks__/'),
  '~spec': resolve(__dirname, 'spec/'),
};

const optimization: Configuration['optimization'] = {
  minimize: mode === 'production',
};

const config: Configuration = {
  entry: () => ['./src/index.tsx'], // https://github.com/webpack-contrib/webpack-hot-client/issues/11
  output: {
    clean: {
      keep: '.keep',
    },
    filename: 'bundle-[contenthash].js',
    path: resolve(__dirname, '.build'),
    publicPath: '/static/',
  },

  devtool: mode === 'production' ? 'hidden-source-map' : 'eval-source-map',
  mode,
  module: {
    rules: [babelLoaderRule, postcssRule],
  },
  optimization,
  plugins: [
    new HtmlWebpackPlugin({
      minify: false,
      template: resolve(__dirname, 'src/index.html'),
    }),
    new DefinePlugin({
      'process.env.NODE_ENV': JSON.stringify(process.env.NODE_ENV),
    }),
  ],
  resolve: { alias },
  target: 'web',
};

export default config;
