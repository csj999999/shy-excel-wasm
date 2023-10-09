// rollup.config.prod.mjs
import resolve from '@rollup/plugin-node-resolve';
import nodePolyfills from 'rollup-plugin-polyfill-node';
import commonjs from '@rollup/plugin-commonjs';
import terser from '@rollup/plugin-terser';
import pkg from '../package.json' assert {type: 'json'};
const input = './src/index.js';
export default [
    {
        input,
        output: {
            file: pkg.exports.script,
            format: 'iife',
            generatedCode: 'es2015',
            name: 'shyexcel',
            sourcemap: false,
        },
        plugins: [
            resolve({
                // 将自定义选项传递给解析插件
                moduleDirectories: ['node_modules']
            }),
            nodePolyfills(),
            commonjs(),
            terser({
                compress:{
                    drop_console: false
                }
            })
        ]
    },
    {
        // ES6 module and <script type="module">
        input,
        output: {
            file: pkg.exports.default,
            format: 'esm',
            generatedCode: 'es2015',
            sourcemap: false,
        },
        plugins: [
            commonjs(),
            nodePolyfills(),
            resolve(),
            terser(),

        ]
    },
    {
        // CommonJS Node module
        input,
        output: {
            file: pkg.exports.require,
            format: 'cjs',
            generatedCode: 'es2015',
            sourcemap: false,
        },
        external: ['path', 'fs'],
        plugins: [
            commonjs(),
            resolve(),
            terser()
        ]
    }
]