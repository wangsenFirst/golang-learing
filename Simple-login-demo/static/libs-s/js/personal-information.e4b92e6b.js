(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["personal-information"],{"80d3":function(e,t,o){"use strict";o.r(t);var r=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{staticStyle:{background:"#f1f1f1",width:"100%",height:"100%","overflow-y":"auto"}},[o("div",{class:"0"===e.full?"nofull":"",staticStyle:{height:"100%"}},[o("div",{staticStyle:{"box-shadow":"0 3px 8px 0 rgba(21,52,91,0.1)",background:"#fff"},style:"0"===e.full?"":"height:100%"},[o("div",{staticStyle:{width:"100%",height:"50px",background:"#08213E","text-align":"center","line-height":"50px",color:"#fff","font-size":"14px"}},[e._v("个人信息")]),o("el-form",{ref:"form",staticClass:"my-form",attrs:{rules:e.rules,model:e.form,"label-width":"80px"}},[o("el-form-item",{attrs:{prop:"",label:"姓名"}},[o("el-input",{attrs:{disabled:""},model:{value:e.form.nickname,callback:function(t){e.$set(e.form,"nickname",t)},expression:"form.nickname"}})],1),o("el-form-item",{attrs:{prop:"",label:"账号"}},[o("el-input",{attrs:{disabled:""},model:{value:e.form.username,callback:function(t){e.$set(e.form,"username",t)},expression:"form.username"}})],1),o("el-form-item",{attrs:{prop:"",label:"部门"}},[o("el-input",{attrs:{disabled:""},model:{value:e.form.orgName,callback:function(t){e.$set(e.form,"orgName",t)},expression:"form.orgName"}})],1),o("el-form-item",{attrs:{prop:"",label:"手机号"}},[o("el-input",{model:{value:e.form.phoneNumber,callback:function(t){e.$set(e.form,"phoneNumber",t)},expression:"form.phoneNumber"}})],1),o("el-form-item",{attrs:{prop:"",label:"邮箱"}},[o("el-input",{model:{value:e.form.email,callback:function(t){e.$set(e.form,"email",t)},expression:"form.email"}})],1),o("div",{staticClass:"dialog-footer"},[o("el-button",{attrs:{size:"default",type:"primary",loading:e.loading},on:{click:e.submit}},[e._v("确 定")]),o("el-button",{attrs:{size:"default",plain:""},on:{click:e.cancel}},[e._v("取 消")])],1)],1)],1)])])},a=[],n=(o("96cf"),o("1da1")),s=o("5530"),i=o("2f62"),l=o("c24f"),u={nickname:"",username:"",orgName:"",phoneNumber:"",email:""},c={name:"personal-information",components:{},props:{},data:function(){return{form:Object(s["a"])({},u),loading:!1,full:"0",passwordType:"password",passwordType2:"password",passwordType3:"password",rules:{phoneNumber:[{required:!1,trigger:"blur",message:"请输入正确的手机号!"}],email:[{required:!1,trigger:"blur",message:"请输入正确的邮箱!"}]}}},computed:Object(s["a"])({},Object(i["b"])({user:function(e){return e.user.userInfo}})),mounted:function(){document.title="个人信息",this.full=this.$route.query.full||"0",this.form=Object(s["a"])({},this.user),console.log(this.user)},watch:{},methods:{showPwd:function(e){"password"===this["passwordType".concat(e)]?this["passwordType".concat(e)]="":this["passwordType".concat(e)]="password"},cancel:function(){window.close()},submit:function(){var e=this,t=this;this.$refs["form"].validate(function(){var o=Object(n["a"])(regeneratorRuntime.mark((function o(r){return regeneratorRuntime.wrap((function(o){while(1)switch(o.prev=o.next){case 0:r&&(e.loading=!0,Object(l["g"])({id:t.form.username,phoneNumber:t.form.phoneNumber,email:t.form.email}).then((function(){e.$message({message:"修改成功!",type:"success"}),e.$store.dispatch("user/getInfo").then((function(){setTimeout((function(){t.loading=!1,window.close()}),1e3)}))})).catch((function(){t.loading=!1})));case 1:case"end":return o.stop()}}),o)})));return function(e){return o.apply(this,arguments)}}())}}},f=c,m=(o("aff3"),o("2877")),p=Object(m["a"])(f,r,a,!1,null,"1a8df3e2",null);t["default"]=p.exports},aff3:function(e,t,o){"use strict";o("dd29")},dd29:function(e,t,o){}}]);