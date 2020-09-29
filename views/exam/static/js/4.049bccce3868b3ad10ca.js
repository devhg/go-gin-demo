webpackJsonp([4],{CtEG:function(t,e,s){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var i=s("aCbg"),n={name:"Exam",components:{VueCountdown:s.n(i).a},data:function(){return{judgeProblems:[],multiSelects:[],singleSelects:[],judgeAnswer:[],multiAnswer:[],singleAnswer:[],time:0,counting:!0,startTime:0,endTime:0,startTimer:"",progressNum:0,success:null,score:[],percentage:0,color:""}},methods:{countDone:function(t){for(var e=0,s=0;s<this.judgeAnswer.length;s++)void 0!==this.judgeAnswer[s]&&e++;for(s=0;s<this.singleAnswer.length;s++)void 0!==this.singleAnswer[s]&&e++;for(s=0;s<this.multiSelects.length;s++)0!=this.multiAnswer[s]&&e++;this.percentage=Math.floor(e/(this.judgeProblems.length+this.singleSelects.length+this.multiSelects.length)*100),this.color=100===this.percentage?"#67C23A":"#409EFF"},questionTable:function(){var t=this;this.$axios.get("contest/cid/"+this.$route.params.cid).then(function(e){t.startTime=e.data.data[2].startTime,t.endTime=e.data.data[2].endTime,t.time=e.data.data[2].endTime-(new Date).getTime(),t.score=e.data.data[1],t.judgeProblems=void 0!=e.data.data[0].judgeProblems?e.data.data[0].judgeProblems:[],t.multiSelects=void 0!=e.data.data[0].multiSelects?e.data.data[0].multiSelects:[],t.singleSelects=void 0!=e.data.data[0].singleSelects?e.data.data[0].singleSelects:[];for(var s=0;s<t.multiSelects.length;s++)t.multiAnswer[s]=[]}).catch(function(e){t.$notify.error({title:"网络错误",message:"网络错误, 请稍后重试",duration:1e3})})},submitForm:function(){var t=this;if(100==this.percentage)if(Date.now()>this.endTime)this.$notify.error({title:"测试过期",message:"无法提交",duration:1e3});else{for(var e=[],s=0;s<this.multiAnswer.length;s++)e[s]=this.multiAnswer[s].sort().toString();this.$axios.post("contest/submit/cid/"+this.$route.params.cid+"/"+localStorage.getItem("ms_username")+"/?time="+Date.parse(new Date),{judgeProblems:this.judgeAnswer,multiSelects:e,singleSelects:this.singleAnswer}).then(function(e){if(200==e.data.code)t.$notify({title:"成功",message:"提交成功",type:"success",duration:1e3}),t.$router.push("/contests");else if(-200==e.data.code){t.$notify.warning({title:"此次测试您已提交过, 请勿重复提交！",message:"2秒后自动退出。",duration:2e3});var s=t;setTimeout(function(){s.$router.push("/contests")},2e3)}}).catch(function(e){t.$notify.error({title:"提交失败",message:"稍后重试",duration:2e3})})}else this.$notify.error({title:"试题未完成",message:"请完成所有试题后提交",duration:1e3})},startProgress:function(){var t=this;this.startTimer=setInterval(function(){var e=(Date.now()-t.startTime)/(t.endTime-t.startTime);t.progressNum=Math.floor(100*e),100==t.progressNum&&(t.counting=!1,clearInterval(t.startTimer))},2e3)},format:function(t){100===t&&(this.success="success")}},created:function(){this.questionTable(),this.startProgress()}},o={render:function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"exam"},[s("div",[s("el-progress",{attrs:{percentage:t.progressNum,status:t.success,format:t.format}}),t._v(" "),t.counting?s("VueCountdown",{attrs:{time:t.time},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v("距离测试结束："+t._s(e.days)+" 天, "+t._s(e.hours)+" 时, "+t._s(e.minutes)+" 分, "+t._s(e.seconds)+" 秒.")]}}],null,!1,3292030193)}):t._e()],1),t._v(" "),t.counting?t._e():s("span",[t._v("测试已结束")]),t._v(" "),s("el-table",{staticStyle:{"margin-top":"10px","box-shadow":"0 2px 12px 0 rgba(0, 0, 0, 0.2)"},attrs:{border:"",data:t.singleSelects}},[s("el-table-column",{attrs:{resizable:!1,"header-align":"center",align:"left",label:"单选题"},scopedSlots:t._u([{key:"default",fn:function(e){return[s("el-form",{ref:"form",attrs:{"label-width":"20px",size:"mini"}},[s("div",{staticClass:"question"},[t._v("\n            "+t._s(e.$index+1)+".\n            "),s("span",{staticStyle:{"font-weight":"bold"}},[t._v("[ "+t._s(t.score.singleScore)+" 分 ]")]),t._v("\n            "+t._s(e.row.title)+"\n          ")]),t._v(" "),s("el-form-item",[s("el-radio-group",{attrs:{size:"medium"},model:{value:t.singleAnswer[e.$index],callback:function(s){t.$set(t.singleAnswer,e.$index,s)},expression:"singleAnswer[scope.$index]"}},[s("el-radio",{attrs:{label:"A"},on:{change:t.countDone}},[t._v("A."+t._s(e.row.optionA))]),t._v(" "),s("el-radio",{attrs:{label:"B"},on:{change:t.countDone}},[t._v("B. "+t._s(e.row.optionB))]),t._v(" "),s("el-radio",{attrs:{label:"C"},on:{change:t.countDone}},[t._v("C. "+t._s(e.row.optionC))]),t._v(" "),s("el-radio",{attrs:{label:"D"},on:{change:t.countDone}},[t._v("D. "+t._s(e.row.optionD))]),t._v(" "),s("el-radio",{directives:[{name:"show",rawName:"v-show",value:""!=e.row.optionE,expression:"scope.row.optionE != ''"}],attrs:{label:"E"},on:{change:t.countDone}},[t._v("E. "+t._s(e.row.optionE))])],1)],1)],1)]}}])},[s("el-divider")],1)],1),t._v(" "),s("el-table",{staticStyle:{"box-shadow":"0 2px 12px 0 rgba(0, 0, 0, 0.2)"},attrs:{border:"",data:t.judgeProblems}},[s("el-table-column",{attrs:{resizable:!1,"header-align":"center",align:"left",label:"判断题"},scopedSlots:t._u([{key:"default",fn:function(e){return[s("el-form",{ref:"form",attrs:{"label-width":"20px",size:"mini"}},[s("div",{staticClass:"question"},[t._v("\n            "+t._s(e.$index+1+t.singleSelects.length)+".\n            "),s("span",{staticStyle:{"font-weight":"bold"}},[t._v("[ "+t._s(t.score.judgeScore)+" 分 ]")]),t._v("\n            "+t._s(e.row.title)+"\n          ")]),t._v(" "),s("el-form-item",[s("el-radio-group",{attrs:{size:"mini"},model:{value:t.judgeAnswer[e.$index],callback:function(s){t.$set(t.judgeAnswer,e.$index,s)},expression:"judgeAnswer[scope.$index]"}},[s("el-radio",{attrs:{label:"T"},on:{change:t.countDone}},[t._v(t._s(e.row.optionT))]),t._v(" "),s("el-radio",{attrs:{label:"F"},on:{change:t.countDone}},[t._v(t._s(e.row.optionF))])],1)],1)],1)]}}])})],1),t._v(" "),s("el-table",{staticStyle:{"box-shadow":"0 2px 12px 0 rgba(0, 0, 0, 0.2)"},attrs:{border:"",data:t.multiSelects}},[s("el-table-column",{attrs:{resizable:!1,"header-align":"center",align:"left",label:"多选题"},scopedSlots:t._u([{key:"default",fn:function(e){return[s("el-form",{ref:"form",attrs:{"label-width":"20px",size:"mini"}},[s("div",{staticClass:"question"},[t._v("\n            "+t._s(e.$index+1+t.singleSelects.length+t.judgeProblems.length)+".\n            "),s("span",{staticStyle:{"font-weight":"bold"}},[t._v("[ "+t._s(t.score.multiScore)+" 分 ]")]),t._v("\n            "+t._s(e.row.title)+"\n          ")]),t._v(" "),s("el-form-item",[s("el-checkbox-group",{attrs:{size:"mini"},model:{value:t.multiAnswer[e.$index],callback:function(s){t.$set(t.multiAnswer,e.$index,s)},expression:"multiAnswer[scope.$index]"}},[s("el-checkbox",{attrs:{label:"A"},on:{change:t.countDone}},[t._v("A. "+t._s(e.row.optionA))]),t._v(" "),s("el-checkbox",{attrs:{label:"B"},on:{change:t.countDone}},[t._v("B. "+t._s(e.row.optionB))]),t._v(" "),s("el-checkbox",{attrs:{label:"C"},on:{change:t.countDone}},[t._v("C. "+t._s(e.row.optionC))]),t._v(" "),s("el-checkbox",{staticClass:"aaa",attrs:{label:"D"},on:{change:t.countDone}},[t._v("D. "+t._s(e.row.optionD))]),t._v(" "),s("el-checkbox",{directives:[{name:"show",rawName:"v-show",value:""!=e.row.optionE,expression:"scope.row.optionE != ''"}],attrs:{label:"E"},on:{change:t.countDone}},[t._v("E. "+t._s(e.row.optionE))])],1)],1)],1)]}}])})],1),t._v(" "),s("div",{staticClass:"progressBar"},[s("el-progress",{attrs:{percentage:t.percentage,color:t.color,"stroke-width":20,"text-inside":""}})],1),t._v(" "),s("el-button",{staticClass:"btn",attrs:{type:"primary",plain:"","native-type":"submint"},on:{click:function(e){return t.submitForm()}}},[t._v("提交")])],1)},staticRenderFns:[]};var a=s("VU/8")(n,o,!1,function(t){s("lMAl")},"data-v-d78cc034",null);e.default=a.exports},aCbg:function(t,e,s){
/*!
 * vue-countdown v1.1.5
 * https://fengyuanchen.github.io/vue-countdown
 *
 * Copyright 2018-present Chen Fengyuan
 * Released under the MIT license
 *
 * Date: 2020-02-25T01:19:32.769Z
 */var i;i=function(){"use strict";return{name:"countdown",data:function(){return{counting:!1,endTime:0,totalMilliseconds:0}},props:{autoStart:{type:Boolean,default:!0},emitEvents:{type:Boolean,default:!0},interval:{type:Number,default:1e3,validator:function(t){return t>=0}},now:{type:Function,default:function(){return Date.now()}},tag:{type:String,default:"span"},time:{type:Number,default:0,validator:function(t){return t>=0}},transform:{type:Function,default:function(t){return t}}},computed:{days:function(){return Math.floor(this.totalMilliseconds/864e5)},hours:function(){return Math.floor(this.totalMilliseconds%864e5/36e5)},minutes:function(){return Math.floor(this.totalMilliseconds%36e5/6e4)},seconds:function(){return Math.floor(this.totalMilliseconds%6e4/1e3)},milliseconds:function(){return Math.floor(this.totalMilliseconds%1e3)},totalDays:function(){return this.days},totalHours:function(){return Math.floor(this.totalMilliseconds/36e5)},totalMinutes:function(){return Math.floor(this.totalMilliseconds/6e4)},totalSeconds:function(){return Math.floor(this.totalMilliseconds/1e3)}},render:function(t){return t(this.tag,this.$scopedSlots.default?[this.$scopedSlots.default(this.transform({days:this.days,hours:this.hours,minutes:this.minutes,seconds:this.seconds,milliseconds:this.milliseconds,totalDays:this.totalDays,totalHours:this.totalHours,totalMinutes:this.totalMinutes,totalSeconds:this.totalSeconds,totalMilliseconds:this.totalMilliseconds}))]:this.$slots.default)},watch:{$props:{deep:!0,immediate:!0,handler:function(){this.totalMilliseconds=this.time,this.endTime=this.now()+this.time,this.autoStart&&this.start()}}},methods:{start:function(){this.counting||(this.counting=!0,this.emitEvents&&this.$emit("start"),"visible"===document.visibilityState&&this.continue())},continue:function(){var t=this;if(this.counting){var e=Math.min(this.totalMilliseconds,this.interval);if(e>0)if(window.requestAnimationFrame){var s,i;this.requestId=requestAnimationFrame(function n(o){s||(s=o),i||(i=o);var a=o-s;a>=e||a+(o-i)/2>=e?t.progress():t.requestId=requestAnimationFrame(n),i=o})}else this.timeoutId=setTimeout(function(){t.progress()},e);else this.end()}},pause:function(){window.requestAnimationFrame?cancelAnimationFrame(this.requestId):clearTimeout(this.timeoutId)},progress:function(){this.counting&&(this.totalMilliseconds-=this.interval,this.emitEvents&&this.totalMilliseconds>0&&this.$emit("progress",{days:this.days,hours:this.hours,minutes:this.minutes,seconds:this.seconds,milliseconds:this.milliseconds,totalDays:this.totalDays,totalHours:this.totalHours,totalMinutes:this.totalMinutes,totalSeconds:this.totalSeconds,totalMilliseconds:this.totalMilliseconds}),this.continue())},abort:function(){this.counting&&(this.pause(),this.counting=!1,this.emitEvents&&this.$emit("abort"))},end:function(){this.counting&&(this.pause(),this.totalMilliseconds=0,this.counting=!1,this.emitEvents&&this.$emit("end"))},update:function(){this.counting&&(this.totalMilliseconds=Math.max(0,this.endTime-this.now()))},handleVisibilityChange:function(){switch(document.visibilityState){case"visible":this.update(),this.continue();break;case"hidden":this.pause()}}},mounted:function(){document.addEventListener("visibilitychange",this.handleVisibilityChange)},beforeDestroy:function(){document.removeEventListener("visibilitychange",this.handleVisibilityChange),this.pause()}}},t.exports=i()},lMAl:function(t,e){}});
//# sourceMappingURL=4.049bccce3868b3ad10ca.js.map