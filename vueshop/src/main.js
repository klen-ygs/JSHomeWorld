import Vue from 'vue'
import App from './App.vue'
import vueRouter from 'vue-router'
import router from './router'
import Store from './store'

Vue.config.productionTip = false
Vue.use(vueRouter)

new Vue({
  render: h => h(App),
  router: router,
  store: Store,
}).$mount('#app')
