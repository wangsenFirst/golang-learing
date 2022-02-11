var identifyCodes='23456789abcdefghjkmnpqrstuvwxyz';
var identifyCode='';
function yzm(str) {
    this.identifyCode=str;
    this.fontSizeMin=34;
    this.fontSizeMax=40;
    this.backgroundColorMin=180;
    this.backgroundColorMax=240;
    this.colorMin=50;
    this.colorMax=160;
    this.lineColorMin=40;
    this.lineColorMax=180;
    this.contentWidth=90;
    this.contentHeight=40;

    // 生成一个随机数
    this.randomNum=function(min, max) {
        return Math.floor(Math.random() * (max - min) + min)
    }
// 生成一个随机的颜色
    this.randomColor=function(min, max) {
        let r = this.randomNum(min, max)
        let g = this.randomNum(min, max)
        let b = this.randomNum(min, max)
        return 'rgb(' + r + ',' + g + ',' + b + ')'
    }
    this.drawPic=function() {
        let canvas = getDom('s-canvas')
        let ctx = canvas.getContext('2d')
        ctx.textBaseline = 'bottom'
        // 绘制背景
        ctx.fillStyle = this.randomColor(this.backgroundColorMin, this.backgroundColorMax)
        ctx.fillRect(0, 0, this.contentWidth, this.contentHeight)
        // 绘制文字
        for (let i = 0; i < this.identifyCode.length; i++) {
            this.drawText(ctx, this.identifyCode[i], i)
        }
        this.drawLine(ctx)
        this.drawDot(ctx)
    }
    this.drawText=function(ctx, txt, i) {
        ctx.fillStyle = this.randomColor(this.colorMin, this.colorMax)
        ctx.font = this.randomNum(this.fontSizeMin, this.fontSizeMax) + 'px SimHei'
        let x = (i + 1) * (this.contentWidth / (this.identifyCode.length + 2))
        let y = this.randomNum(this.fontSizeMax, this.contentHeight - 5)
        var deg = this.randomNum(-10, 10)
        // 修改坐标原点和旋转角度
        ctx.translate(x, y)
        ctx.rotate(deg * Math.PI / 180)
        ctx.fillText(txt, 0, 0)
        // 恢复坐标原点和旋转角度
        ctx.rotate(-deg * Math.PI / 180)
        ctx.translate(-x, -y)
    }
    this.drawLine=function(ctx) {
        // 绘制干扰线
        for (let i = 0; i < 4; i++) {
            ctx.strokeStyle = this.randomColor(this.lineColorMin, this.lineColorMax)
            ctx.beginPath()
            ctx.moveTo(this.randomNum(0, this.contentWidth), this.randomNum(0, this.contentHeight))
            ctx.lineTo(this.randomNum(0, this.contentWidth), this.randomNum(0, this.contentHeight))
            ctx.stroke()
        }
    }
    this.drawDot=function(ctx) {
        // 绘制干扰点
        for (let i = 0; i < 60; i++) {
            ctx.fillStyle = this.randomColor(0, 255)
            ctx.beginPath()
            ctx.arc(this.randomNum(0, this.contentWidth), this.randomNum(0, this.contentHeight), 1, 0, 2 * Math.PI)
            ctx.fill()
        }
    }
    //this.drawPic();
}

function randomNum(min, max) {
    return Math.floor(Math.random() * (max - min) + min);
}
function refreshCode() {
    identifyCode = "";
    makeCode(identifyCodes, 4);
    var abc=new yzm(identifyCode);
}
function makeCode(o, l) {
    for (let i = 0; i < l; i++) {
        identifyCode += identifyCodes[
            randomNum(0, identifyCodes.length)
            ];
    }

    //console.log(this.identifyCode);
}
function SavePassword() {
    if(getDom("savePassword").checked) {
        var username = getDom("username").value;
        var password = getDom("password").value;
        window.localStorage.username =  username;
        window.localStorage.password = password;
        localStorage.rmbPassword = true;
    } else {
        localStorage.rmbPassword = false;
        localStorage.clear()
    }
}
function getDom(s) {
    return document.getElementById(s)
}
function checkForm(){
    SavePassword()
    var that = getDom('password');
    if(getDom("username").value==""){
        getDom("error").innerHTML="请输入用户名";
        getDom("error").style.display="";
        return false;
    }else if(getDom("password").value==""){
        getDom("error").innerHTML="请输入密码";
        getDom("error").style.display="";
        return false;
    }/*else if(getDom("yzm").value!=identifyCode){
        getDom("error").innerHTML="请输入正确的验证码";
        getDom("error").style.display="";
        return false;
    }*/else {
        ajax({
            method : 'post',
            url : '/auth/getPasswordKey',//获取公钥
            data : {},
            success : function (res) {
                var publicKey = res;
                //利用公钥对密码进行加密
                var encrypt = new JSEncrypt();
                encrypt.setPublicKey(publicKey);
                var password = that.value;
                that.value=encrypt.encrypt(password);
                getDom('myform').submit();
            },
            async : true
        })

    }
}
function getQueryString(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return unescape(r[2]);
    return null;
}
window.onload=function () {
    var str=getQueryString('error');
    if(!!str){
        getDom("error").innerHTML="用户名或密码错误";
        getDom("error").style.display="";
    }
    if(localStorage.rmbPassword){
        getDom("savePassword").checked=true;
        getDom("username").value=localStorage.username;
        getDom("password").value=localStorage.password
    }else {
        getDom("savePassword").checked=false;
    }
    refreshCode();
};
