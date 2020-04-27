import Vue from 'vue'
import Vuex from 'vuex'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import '@/lib/filter'
import 'element-ui/lib/theme-chalk/index.css'
import '@/assets/styles/theme-default.scss'
import '@/assets/styles/main.scss'

Vue.config.productionTip = false

Vue.use(ElementUI)
Vue.use(Vuex)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
