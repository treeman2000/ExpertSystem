<template>
<div>
    <el-form>
        <el-form-item v-for="(q,i) in questions" :key=i>
            {{i}}. {{q}}
            <br/>
            <el-radio v-model="ans[i]" label=1>完全不符合</el-radio>
            <el-radio v-model="ans[i]" label=2>较不符合</el-radio>
            <el-radio v-model="ans[i]" label=3>不确定</el-radio>
            <el-radio v-model="ans[i]" label=4>较符合</el-radio>
            <el-radio v-model="ans[i]" label=5>完全符合</el-radio>
            ans[i]={{ans[i]}}
            <br/>
        </el-form-item>
        <el-form-item id="submitButton">
            <el-button type='primary' @click='submit'>提交</el-button>
        </el-form-item>
    </el-form>
</div>
</template>

<script>
import axios from 'axios'
const length=18//number of questions
export default {
    name:"Questionnaire",
    data(){
        return {
            questions:["我发现与人亲近比较容易。",
            "我发现要我去依赖别人很困难。",
            "我时常担心情侣并不真心爱我。",
            "我发现别人并不愿像我希望的那样亲近我。",
            "能依赖别人让我感到很舒服。",
            "我不在乎别人太亲近我。",
            "我发现当我需要别人帮助时，没人会帮我。",
            "和别人亲近使我感到有些不舒服。",
            "我时常担心情侣不想和我在一起。",
            "当我对别人表达我的情感时，我害怕他们与我的感觉会不一样。",
            "我时常怀疑情侣是否真正关心我。",
            "我对别人建立亲密的关系感到很舒服。",
            "当有人在情感上太亲近我时，我感到不舒服。",
            "我知道当我需要别人帮助时，总有人会帮我。",
            "我想与人亲近，但担心自己会受到伤害。",
            "我发现我很难完全信赖别人。",
            "情侣想要我在情感上更亲近一些，这常使我感到不舒服。",
            "我不能肯定，在我需要时，总找得到可以依赖的人。"
            ],
            ans: new Array(length).fill(0)
        }
    },
    methods:{
        submit(){
            for(let i=0;i<length;i++){
                if(this.ans[i]==0){
                    this.$alert("您尚未选择问题"+i+"的答案，请全部填写完毕后再提交","提示")
                    return
                }
            }
            axios.post('/submitQuestionnaire',{
                questionnaire:this.ans
            }).then(res=>{
                console.log(res)
            }).catch(err=>{
                console.log('error!')
                console.log(err)
            })
        }
    }
}
</script>

<style>
#submitButton{
    text-align: center;
}
</style>