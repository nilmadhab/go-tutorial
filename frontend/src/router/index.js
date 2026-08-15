import { createRouter, createWebHistory } from 'vue-router'
import Home from '../pages/Home.vue'
import Admin from '../pages/Admin.vue'
import Product from '../pages/Product.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/admin',
    name: 'Admin',
    component: Admin
  },
  {
    path: '/product/:id',
    name: 'Product',
    component: Product
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
