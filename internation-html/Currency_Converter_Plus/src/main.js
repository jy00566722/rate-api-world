import { createApp } from 'vue'
import { createPinia } from 'pinia'
import './assets/ghost.css'
import App from './App.vue'
import router from './router'

createApp(App).use(router).mount('#app')
