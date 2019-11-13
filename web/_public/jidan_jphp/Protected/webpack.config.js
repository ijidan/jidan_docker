const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const {CleanWebpackPlugin} = require('clean-webpack-plugin');

module.exports = {
	entry: {
		print: './Src/print.js',
		app: './Src/index.js'
	},
	plugins:[
		new CleanWebpackPlugin({options:['dist']}),
		new HtmlWebpackPlugin({
			title: 'Output Manager'
		})
	],
	output: {
		filename: '[name].bundle.js',
		path: path.resolve(__dirname, 'Dist')
	},
	module: {
		rules: [
			{
				test: /\.css$/,
				use: [
					'style-loader',
					'css-loader'
				]
			},
			{
				test: /\.(gif|jpg|jpeg|png)$/,
				use: [
					'file-loader'
				]
			},
			{
				test: /\.(woff|woff2|eot|ttf|otf)$/,
				use: [
					'file-loader'
				]
			}
		]
	},
	"watch": true
};