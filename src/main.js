// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import FastClick from 'fastclick'
import VueRouter from 'vue-router'
import App from './App'
import Home from './components/HelloFromVux'
import Form from "./components/Form"
import Search from "./components/Search"
import Info from "./components/Info"
import End from "./components/End"
import axios from 'axios'
Vue.prototype.$axios = axios



Vue.use(VueRouter)

const routes = [{
  path: '/',
  component: Form
},{
	path: '/search',
  component: Search
},
{
	name: "info",
	path: '/info',
  component: Info
},
{
	name:"end",
	path: '/end',
	component: End
}
]

const router = new VueRouter({
  routes
})

FastClick.attach(document.body)

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  router,
  render: h => h(App)
}).$mount('#app-box')
