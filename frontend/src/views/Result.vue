<template>
<div>
    <h1> 您的依恋类型测试结果报告</h1>
    <h3>{{$route.query.attachmentStyle}}</h3>
    <h3>{{$route.query.description}}</h3>

</div>
</template>

<script>
export default {
    name:"Result"
}
</script>

<style scoped>
h1{
    text-align: center;
}
h3{
    padding: 5%;
}
</style>