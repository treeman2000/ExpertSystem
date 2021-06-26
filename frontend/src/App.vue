<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  methods:{
    startQuestionnaire(){
      console.log("clicked")
      this.$router.push('Questionnaire')
    }
  }
}
</script>

<style>
</style>
