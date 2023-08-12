function upload_data() {
    jsonObj = getData()
    for(let item in jsonObj){
        jsonObj[item] = jsonObj[item].replaceAll(",","，")
    }
    axios.post("/getall",jsonObj).catch(function(error){
        console.log(error)
    })
    alert("上传成功")
}


function button_click() {
    jsonObj = getData()
    for(let item in jsonObj){
        jsonObj[item] = jsonObj[item].replaceAll(",","，")
    }
    console.log(jsonObj)
    mytoCSV(jsonObj)
}


function mytoCSV(jsonData){
    //列标题
    var str = '姓名,手机,单位名称,职务/职称,科研发展思路,科研发展方向,科研管理机制优化改革,科研团队组建,重大科技项目攻关,区域科技合作,其它\n';
    //具体数值 遍历
    for(let item in jsonData){

        //增加\t为了不让表格显示科学计数法或者其他格式
        //此处用`取代'，具体用法搜索模板字符串 ES6特性
        str+=`${jsonData[item] + '\t,'}`;
    }
    str+='\n';
    let uri = 'data:text/csv;charset=utf-8,\ufeff' + encodeURIComponent(str);
    var link = document.createElement("a");
    link.href = uri;
    link.download = "导出.csv";
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
}

function getData(){
    name = document.getElementById("name").value;
    phone = document.getElementById("phone").value;
    dwmc = document.getElementById("dwmc").value;
    zwzc = document.getElementById("zwzc").value;
    kyfzsl = document.getElementById("exampleFormControlTextarea1").value;
    kyfzfx = document.getElementById("exampleFormControlTextarea2").value;
    kygljzyhgg = document.getElementById("exampleFormControlTextarea3").value;
    kytdzj = document.getElementById("exampleFormControlTextarea4").value;
    zdkjxmgg = document.getElementById("exampleFormControlTextarea5").value;
    qykjhz = document.getElementById("exampleFormControlTextarea6").value;
    qt = document.getElementById("exampleFormControlTextarea7").value;
    let jsonObj = {};
    jsonObj.name = name
    jsonObj.phone = phone
    jsonObj.dwmc = dwmc
    jsonObj.zwzc = zwzc
    jsonObj.kyfzsl = kyfzsl
    jsonObj.kyfzfx = kyfzfx
    jsonObj.kygljzyhgg = kygljzyhgg
    jsonObj.kytdzj = kytdzj
    jsonObj.zdkjxmgg = zdkjxmgg
    jsonObj.qykjhz = qykjhz
    jsonObj.qt = qt
    return jsonObj

}