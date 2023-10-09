// rollup.config.prod.mjs
import pkg from '../package.json' assert {type: 'json'};
import serve from 'rollup-plugin-serve';
const input = './src/index.js';
import baseConfig from './rollup.config.mjs';
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
            serve({
                open: true,
                contentBase:['example'],
                port: 3000
            })
        ]
    }
]