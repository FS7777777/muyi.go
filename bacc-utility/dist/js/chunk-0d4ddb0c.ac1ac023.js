(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-0d4ddb0c"],{"1e68":function(t,e,a){"use strict";var s=a("af96"),r=a.n(s);r.a},"2d51":function(t,e,a){"use strict";a.r(e);var s=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("a-card",{attrs:{bordered:!1}},[a("a-row",[a("a-col",{attrs:{sm:8,xs:24}},[a("head-info",{attrs:{title:"我的待办",content:"8个任务",bordered:!0}})],1),a("a-col",{attrs:{sm:8,xs:24}},[a("head-info",{attrs:{title:"本周任务平均处理时间",content:"32分钟",bordered:!0}})],1),a("a-col",{attrs:{sm:8,xs:24}},[a("head-info",{attrs:{title:"本周完成任务数",content:"24个"}})],1)],1)],1),a("a-card",{staticStyle:{"margin-top":"24px"},attrs:{bordered:!1,title:"标准列表"}},[a("div",{attrs:{slot:"extra"},slot:"extra"},[a("a-radio-group",[a("a-radio-button",[t._v("全部")]),a("a-radio-button",[t._v("进行中")]),a("a-radio-button",[t._v("等待中")])],1),a("a-input-search",{staticStyle:{"margin-left":"16px",width:"272px"}})],1),a("div",{staticClass:"operate"},[a("a-button",{staticStyle:{width:"100%"},attrs:{type:"dashed",icon:"plus"},on:{click:function(e){return t.$refs.taskForm.add()}}},[t._v("添加")])],1),a("a-list",{attrs:{size:"large",pagination:{showSizeChanger:!0,showQuickJumper:!0,pageSize:5,total:50}}},t._l(t.data,function(e,s){return a("a-list-item",{key:s},[a("a-list-item-meta",{attrs:{description:e.description}},[a("a-avatar",{attrs:{slot:"avatar",size:"large",shape:"square",src:e.avatar},slot:"avatar"}),a("a",{attrs:{slot:"title"},slot:"title"},[t._v(t._s(e.title))])],1),a("div",{attrs:{slot:"actions"},slot:"actions"},[a("a",[t._v("编辑")])]),a("div",{attrs:{slot:"actions"},slot:"actions"},[a("a-dropdown",[a("a-menu",{attrs:{slot:"overlay"},slot:"overlay"},[a("a-menu-item",[a("a",[t._v("编辑")])]),a("a-menu-item",[a("a",[t._v("删除")])])],1),a("a",[t._v("更多"),a("a-icon",{attrs:{type:"down"}})],1)],1)],1),a("div",{staticClass:"list-content"},[a("div",{staticClass:"list-content-item"},[a("span",[t._v("Owner")]),a("p",[t._v(t._s(e.owner))])]),a("div",{staticClass:"list-content-item"},[a("span",[t._v("开始时间")]),a("p",[t._v(t._s(e.startAt))])]),a("div",{staticClass:"list-content-item"},[a("a-progress",{staticStyle:{width:"180px"},attrs:{percent:e.progress.value,status:e.progress.status?e.progress.status:null}})],1)])],1)}),1),a("task-form",{ref:"taskForm"})],1)],1)},r=[],o=a("81d1"),i=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("a-modal",{attrs:{width:640,visible:t.visible,title:"任务添加"},on:{ok:t.handleSubmit,cancel:function(e){t.visible=!1}}},[a("a-form",{attrs:{form:t.form},on:{submit:t.handleSubmit}},[a("a-form-item",{attrs:{label:"任务名称",labelCol:t.labelCol,wrapperCol:t.wrapperCol}},[a("a-input",{directives:[{name:"decorator",rawName:"v-decorator",value:["taskName",{rules:[{required:!0,message:"请输入任务名称"}]}],expression:"['taskName', {rules:[{required: true, message: '请输入任务名称'}]}]"}]})],1),a("a-form-item",{attrs:{label:"开始时间",labelCol:t.labelCol,wrapperCol:t.wrapperCol}},[a("a-date-picker",{directives:[{name:"decorator",rawName:"v-decorator",value:["startTime",{rules:[{required:!0,message:"请选择开始时间"}]}],expression:"['startTime', {rules:[{required: true, message: '请选择开始时间'}]}]"}],staticStyle:{width:"100%"}})],1),a("a-form-item",{attrs:{label:"任务负责人",labelCol:t.labelCol,wrapperCol:t.wrapperCol}},[a("a-select",{directives:[{name:"decorator",rawName:"v-decorator",value:["owner",{rules:[{required:!0,message:"请选择开始时间"}]}],expression:"['owner', {rules:[{required: true, message: '请选择开始时间'}]}]"}]},[a("a-select-option",{attrs:{value:0}},[t._v("付晓晓")]),a("a-select-option",{attrs:{value:1}},[t._v("周毛毛")])],1)],1),a("a-form-item",{attrs:{label:"产品描述",labelCol:t.labelCol,wrapperCol:t.wrapperCol}},[a("a-textarea",{directives:[{name:"decorator",rawName:"v-decorator",value:["desc"],expression:"['desc']"}]})],1)],1)],1)},l=[],n={name:"TaskForm",data:function(){return{labelCol:{xs:{span:24},sm:{span:7}},wrapperCol:{xs:{span:24},sm:{span:13}},visible:!1,form:this.$form.createForm(this)}},methods:{add:function(){this.visible=!0},edit:function(t){var e=this.form.setFieldsValue;this.visible=!0,this.$nextTick(function(){e({taskName:"test"})})},handleSubmit:function(){var t=this.form.validateFields;this.visible=!0,t(function(t,e){t||console.log("values",e)})}}},c=n,p=a("6691"),d=Object(p["a"])(c,i,l,!1,null,null,null),u=d.exports,m=[];m.push({title:"Alipay",avatar:"https://gw.alipayobjects.com/zos/rmsportal/WdGqmHpayyMjiEhcKoVE.png",description:"那是一种内在的东西， 他们到达不了，也无法触及的",owner:"付晓晓",startAt:"2018-07-26 22:44",progress:{value:90}}),m.push({title:"Angular",avatar:"https://gw.alipayobjects.com/zos/rmsportal/zOsKZmFRdUtvpqCImOVY.png",description:"希望是一个好东西，也许是最好的，好东西是不会消亡的",owner:"曲丽丽",startAt:"2018-07-26 22:44",progress:{value:54}}),m.push({title:"Ant Design",avatar:"https://gw.alipayobjects.com/zos/rmsportal/dURIMkkrRFpPgTuzkwnB.png",description:"生命就像一盒巧克力，结果往往出人意料",owner:"林东东",startAt:"2018-07-26 22:44",progress:{value:66}}),m.push({title:"Ant Design Pro",avatar:"https://gw.alipayobjects.com/zos/rmsportal/sfjbOqnsXXJgNCjCzDBL.png",description:"城镇中有那么多的酒馆，她却偏偏走进了我的酒馆",owner:"周星星",startAt:"2018-07-26 22:44",progress:{value:30}}),m.push({title:"Bootstrap",avatar:"https://gw.alipayobjects.com/zos/rmsportal/siCrBXXhmvTQGWPNLBow.png",description:"那时候我只会想自己想要什么，从不想自己拥有什么",owner:"吴加好",startAt:"2018-07-26 22:44",progress:{status:"exception",value:100}});var v={name:"StandardList",components:{HeadInfo:o["a"],TaskForm:u},data:function(){return{data:m}}},f=v,b=(a("1e68"),Object(p["a"])(f,s,r,!1,null,"395fcc17",null));e["default"]=b.exports},"81d1":function(t,e,a){"use strict";var s=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"head-info",class:t.center&&"center"},[a("span",[t._v(t._s(t.title))]),a("p",[t._v(t._s(t.content))]),t.bordered?a("em"):t._e()])},r=[],o={name:"HeadInfo",props:{title:{type:String,default:""},content:{type:String,default:""},bordered:{type:Boolean,default:!1},center:{type:Boolean,default:!0}}},i=o,l=(a("9d9c"),a("6691")),n=Object(l["a"])(i,s,r,!1,null,"7002d88c",null);e["a"]=n.exports},"9d9c":function(t,e,a){"use strict";var s=a("d98a"),r=a.n(s);r.a},af96:function(t,e,a){},d98a:function(t,e,a){}}]);