import { createApp } from 'vue'
import App from './App.vue'

import './assets/main.css'

// import Router from 'vue-router'
import router from './router'

import ViewUIPlus from 'view-ui-plus'
import 'view-ui-plus/dist/styles/viewuiplus.css'

import {saveAs} from 'file-saver';

const app = createApp(App)

app.use(ViewUIPlus)
    .use(router)
    .mount('#app')

// export default new Router({
//     routes: [
//         {
//         path: '/',
//         name: 'home',
//         component: Home
//         },
//         {
//         path: './components/database.vue',
//         name: 'about',
//         component: About
//         }
//     ]
//     })