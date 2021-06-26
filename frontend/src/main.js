import Vue from 'vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import App from './App.vue'
import router from './router'
Vue.use(ElementUI);


Vue.config.productionTip = false

new Vue({
  router,//增强写法，实例和参数名相同时，只写一次就行
  render: h => h(App),
}).$mount('#app')
