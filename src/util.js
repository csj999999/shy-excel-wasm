
function common_extend(obj1, obj2) {
    let ret = {};

    for (let attr in obj1) {
        ret[attr] = obj1[attr];
    }
    if (obj2 === undefined) {
        return ret;
    }
    for (let attr in obj2) {
        if (obj2[attr] === undefined) {
            continue;
        }
        ret[attr] = obj2[attr];
    }

    return ret;
}

function deepMerge(target, source) {
    if (typeof target !== 'object' || target === null) {
        target = {};
    }
    if (typeof source !== 'object' || source === null) {
        source = {};
    }

    for (const key in source) {
        if (source.hasOwnProperty(key)) {
            const value = source[key];
            if (typeof value !== 'object' || value === null) {
                target[key] = value;
            } else {
                target[key] = deepMerge(target[key], value);
            }
        }
    }

    return target;
}

function random() {
    // 获取当前日期和时间并转换为字符串格式
    var dateTime = new Date().toISOString().replace(/[-T:.Z]/g, '');

    // 生成一个随机数，长度为 30 减去日期时间字符串的长度
    var randomPartLength = 30 - dateTime.length;
    var randomPart = '';
    for (var i = 0; i < randomPartLength; i++) {
        randomPart += Math.floor(Math.random() * 10);
    }

    // 将两者连接起来
    return dateTime + randomPart;
}

export {
    common_extend,
    deepMerge,
    random
}