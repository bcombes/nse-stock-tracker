import Vue from 'vue'
import App from './App.vue'
import VuePaginate from 'vue-paginate'


Vue.config.devtools = true
Vue.use(VuePaginate)
new Vue({
  el: '#app',
  render: h => h(App)
})
