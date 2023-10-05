// rollup.config.mjs
import resolve from '@rollup/plugin-node-resolve';
import commonjs from '@rollup/plugin-commonjs';
import terser from '@rollup/plugin-terser';
export default {
    input: 'src/index.js',
    output: {
        file: 'main.js',
        format: 'cjs'
    },
    plugins: [
        resolve({
            // 将自定义选项传递给解析插件
            moduleDirectories: ['node_modules']
        }),
        commonjs(),
        terser({
            compress:{
                drop_console: false
            }
        })
    ]
};