import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Questionnaire from '../views/Questionnaire.vue'
import Result from '../views/Result.vue'
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/questionnaire',
    component: Questionnaire
  },
  {
    path: '/result',
    component: Result
  }
]

const router = new VueRouter({
  routes
})

export default router
