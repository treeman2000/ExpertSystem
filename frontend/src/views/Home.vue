<template>
  <div class="full">
    <h1 id="header">依恋类型测试</h1>
    <br/>
    <div class="full" id="body">
      <el-button type='primary' @click='startQuestionnaire'>开始测试</el-button>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Home',
  methods:{
    startQuestionnaire(){
      this.$router.push('Questionnaire')
    }
  }
}
</script>

<style scoped>
#header{
  padding:5% 0
}
.full{
  text-align: center;
}
</style>
