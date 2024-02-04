import { createRouter, createWebHistory } from 'vue-router'
import Home from './views/Home.vue'
import Edit from './views/Edit.vue'
import Query from './views/Query.vue'


const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/edit',
        name: 'Edit',
        component: Edit
    },
    {
        path: '/query',
        name: 'Query',
        component: Query
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router