(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-4a9950e0"],{"23f1":function(t,a,e){},"33e6":function(t,a,e){"use strict";e.r(a);var o=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("div",{staticClass:"account-settings-info-view"},[e("a-row",{attrs:{gutter:16}},[e("a-col",{attrs:{md:24,lg:16}},[e("a-form",{attrs:{layout:"vertical"}},[e("a-form-item",{attrs:{label:"昵称"}},[e("a-input",{attrs:{placeholder:"给自己起个名字"}})],1),e("a-form-item",{attrs:{label:"Bio"}},[e("a-textarea",{attrs:{rows:"4",placeholder:"You are not alone."}})],1),e("a-form-item",{attrs:{label:"电子邮件",required:!1}},[e("a-input",{attrs:{placeholder:"exp@admin.com"}})],1),e("a-form-item",{attrs:{label:"加密方式",required:!1}},[e("a-select",{attrs:{defaultValue:"aes-256-cfb"}},[e("a-select-option",{attrs:{value:"aes-256-cfb"}},[t._v("aes-256-cfb")]),e("a-select-option",{attrs:{value:"aes-128-cfb"}},[t._v("aes-128-cfb")]),e("a-select-option",{attrs:{value:"chacha20"}},[t._v("chacha20")])],1)],1),e("a-form-item",{attrs:{label:"连接密码",required:!1}},[e("a-input",{attrs:{placeholder:"h3gSbecd"}})],1),e("a-form-item",{attrs:{label:"登录密码",required:!1}},[e("a-input",{attrs:{placeholder:"密码"}})],1),e("a-form-item",[e("a-button",{attrs:{type:"primary"}},[t._v("提交")]),e("a-button",{staticStyle:{"margin-left":"8px"}},[t._v("保存")])],1)],1)],1),e("a-col",{style:{minHeight:"180px"},attrs:{md:24,lg:8}},[e("div",{staticClass:"ant-upload-preview",on:{click:function(a){return t.$refs.modal.edit(1)}}},[e("a-icon",{staticClass:"upload-icon",attrs:{type:"cloud-upload-o"}}),e("div",{staticClass:"mask"},[e("a-icon",{attrs:{type:"plus"}})],1),e("img",{attrs:{src:t.option.img}})],1)])],1),e("avatar-modal",{ref:"modal"})],1)},i=[],r=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("a-modal",{attrs:{title:"修改头像",visible:t.visible,maskClosable:!1,confirmLoading:t.confirmLoading,width:800},on:{cancel:t.cancelHandel}},[e("a-row",[e("a-col",{style:{height:"350px"},attrs:{xs:24,md:12}},[e("vue-cropper",{ref:"cropper",attrs:{img:t.options.img,info:!0,autoCrop:t.options.autoCrop,autoCropWidth:t.options.autoCropWidth,autoCropHeight:t.options.autoCropHeight,fixedBox:t.options.fixedBox},on:{realTime:t.realTime}})],1),e("a-col",{style:{height:"350px"},attrs:{xs:24,md:12}},[e("div",{staticClass:"avatar-upload-preview"},[e("img",{style:t.previews.img,attrs:{src:t.previews.url}})])])],1),e("template",{slot:"footer"},[e("a-button",{key:"back",on:{click:t.cancelHandel}},[t._v("取消")]),e("a-button",{key:"submit",attrs:{type:"primary",loading:t.confirmLoading},on:{click:t.okHandel}},[t._v("保存")])],1)],2)},s=[],n={data:function(){return{visible:!1,id:null,confirmLoading:!1,options:{img:"/avatar2.jpg",autoCrop:!0,autoCropWidth:200,autoCropHeight:200,fixedBox:!0},previews:{}}},methods:{edit:function(t){this.visible=!0,this.id=t},close:function(){this.id=null,this.visible=!1},cancelHandel:function(){this.close()},okHandel:function(){var t=this;t.confirmLoading=!0,setTimeout(function(){t.confirmLoading=!1,t.close(),t.$message.success("上传头像成功")},2e3)},realTime:function(t){this.previews=t}}},l=n,c=(e("46df"),e("6691")),u=Object(c["a"])(l,r,s,!1,null,"52ec7f14",null),p=u.exports,d={components:{AvatarModal:p},data:function(){return{preview:{},option:{img:"/avatar2.jpg",info:!0,size:1,outputType:"jpeg",canScale:!1,autoCrop:!0,autoCropWidth:180,autoCropHeight:180,fixedBox:!0,fixed:!0,fixedNumber:[1,1]}}},methods:{}},f=d,m=(e("66a1"),Object(c["a"])(f,o,i,!1,null,"24d3fbf6",null));a["default"]=m.exports},"46df":function(t,a,e){"use strict";var o=e("23f1"),i=e.n(o);i.a},"4e19":function(t,a,e){},"66a1":function(t,a,e){"use strict";var o=e("4e19"),i=e.n(o);i.a}}]);