// rollup.config.prod.mjs
import pkg from '../package.json' assert {type: 'json'};
import serve from 'rollup-plugin-serve';
const input = './src/index.js';
import baseConfig from './rollup.config.mjs';
import resolve from "@rollup/plugin-node-resolve";
export default [
    ...baseConfig,
    {
        input,
        output: {
            file: pkg.exports.require,
            format: 'cjs',
            generatedCode: 'es2015',
            sourcemap: false,
        },
        plugins: [
            resolve({
                // 将自定义选项传递给解析插件
                moduleDirectories: ['node_modules']
            }),
            serve({
                open: true,
                contentBase:['example'],
                port: 3000
            })
        ]
    }
]