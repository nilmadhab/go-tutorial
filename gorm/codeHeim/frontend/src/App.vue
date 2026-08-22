<script setup>
import { useRoute } from 'vue-router'
import { computed } from 'vue'

const route = useRoute()
const isDetailPage = computed(() => route.name === 'user-detail')
</script>

<template>
  <div class="app">
    <header v-if="!isDetailPage">
      <h1>CodeHeim GORM Demo</h1>
      <nav>
        <router-link to="/" :class="{ active: route.name === 'users' }">
          Users
        </router-link>
        <router-link to="/notes" :class="{ active: route.name === 'notes' }">
          Notes
        </router-link>
      </nav>
    </header>
    <main>
      <router-view />
    </main>
  </div>
</template>

<style>
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, sans-serif;
  background: #f5f5f5;
  min-height: 100vh;
}

.app {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

header {
  background: #2c3e50;
  color: white;
  padding: 20px;
  border-radius: 8px;
  margin-bottom: 20px;
}

header h1 {
  margin-bottom: 15px;
  font-size: 1.5em;
}

nav {
  display: flex;
  gap: 10px;
}

nav a {
  padding: 10px 20px;
  background: #34495e;
  color: white;
  border-radius: 4px;
  text-decoration: none;
  font-size: 1em;
  transition: background 0.2s;
}

nav a:hover {
  background: #4a6278;
}

nav a.active {
  background: #3498db;
}

main {
  background: #fff;
  border-radius: 8px;
  min-height: 400px;
}
</style>
