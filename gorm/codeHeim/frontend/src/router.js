import { createRouter, createWebHistory } from 'vue-router'
import UserList from './components/UserList.vue'
import UserDetail from './components/UserDetail.vue'
import NoteList from './components/NoteList.vue'

const routes = [
  { path: '/', name: 'users', component: UserList },
  { path: '/users/:id', name: 'user-detail', component: UserDetail },
  { path: '/notes', name: 'notes', component: NoteList }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
