import '@babel/polyfill'
import Vue from 'vue'
import VueMeta from 'vue-meta'
import VueI18n from 'vue-i18n'
import './plugins/vuetify'
import App from './App.vue'
import router from './router'
import store from './store'

Vue.config.productionTip = false

Vue.use(VueMeta)
Vue.use(VueI18n)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
