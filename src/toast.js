function start(val, num, buffer, fileName, error,shyexcel) {
    var divs = document.getElementsByTagName('span');
    for (var i = 0; i < divs.length; i++) {
        // 设置每个div的display属性为none
        // divs[i].style.display = 'none';
        divs[i].remove()
    }
    var parentDiv = document.createElement('span');

    var img = document.createElement('img');
    img.src = '../example/static/del.png';
    img.style.width = '14px';
    img.style.height = '14px';
    img.style.position = "absolute"; // 设置为绝对定位
    img.style.right = 10 + "px"; // 设置 div2 的左边界为 div1 的右边界
    img.style.top = 10 + "px"; // 设置 div2 的顶部与 div1 的顶部对齐
    // img.style.display = 'flex';
    // img.style.justifyContent = 'flex-end';
    // img.style.top = '5px';
    parentDiv.appendChild(img);
    img.addEventListener('click', function() {
        var divs = document.getElementsByTagName('span');
            for (var i = 0; i < divs.length; i++) {
                // 设置每个div的display属性为none
                divs[i].style.display = 'none';
                shyexcel._status = 0
            }
    })

    var div = document.createElement('span');
    div.innerHTML = val;
    parentDiv.style.display = 'block';
    div.style.display = 'block';
    parentDiv.style.padding = '10px';
    parentDiv.style.width = '10%';
    parentDiv.style.position = 'fixed';
    parentDiv.style.bottom = '30px';
    parentDiv.style.right = '20px';
    parentDiv.style.borderRadius = '4px';
    parentDiv.style.backgroundColor = 'white';
    parentDiv.style.fontSize = '16px';
    parentDiv.style.boxShadow = '0 1px 6px rgba(0,0,0,.2)';
    parentDiv.appendChild(div);

    if(num == 1) {
        var loadimg = document.createElement('img');
        loadimg.src = '../example/static/loadingProgress.gif';
        loadimg.style.marginTop = '10px';
        loadimg.style.width = '80%';
        loadimg.style.height = '10px';
        parentDiv.appendChild(loadimg);
    }

    if(num == 2) {
        var div1 = document.createElement('span');
        div1.innerHTML = '下载';
        div1.style.display = 'block';
        div1.style.color = '#2d8cf0';
        div1.style.marginTop = '20px';
        div1.style.cursor = 'pointer';
        parentDiv.appendChild(div1);
        div1.addEventListener('click', function() {
            const link = document.createElement('a');
            link.download = fileName === null || fileName === '' ? random()+".xlsx" : fileName;
            link.href = URL.createObjectURL(
            new Blob([buffer], {
                type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
            })
            );
            link.click();
            
            var divs = document.getElementsByTagName('span');
            for (var i = 0; i < divs.length; i++) {
                // 设置每个div的display属性为none
                divs[i].style.display = 'none';
                shyexcel._status = 0
            }
            // div.style.display = 'none';
            // div1.style.display = 'none';
        });
    }



    if(num == 3) {
        var div1 = document.createElement('span');
        div1.innerHTML = '查看原因';
        div1.style.display = 'block';
        div1.style.color = 'red';
        div1.style.marginTop = '20px';
        div1.style.cursor = 'pointer';
        parentDiv.appendChild(div1);
        div1.addEventListener('click', function() {
            var div2 = document.createElement('span');
            div2.innerHTML = error;
            div2.style.display = 'block';
            div2.style.marginTop = '20px';
            parentDiv.appendChild(div2);
        });
    }
    document.body.appendChild(parentDiv);
}
export {
    start
}