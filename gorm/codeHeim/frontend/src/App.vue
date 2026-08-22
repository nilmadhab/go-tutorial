<script setup>
import { ref } from 'vue'
import UserList from './components/UserList.vue'
import NoteList from './components/NoteList.vue'

const activeTab = ref('users')
</script>

<template>
  <div class="app">
    <header>
      <h1>CodeHeim GORM Demo</h1>
      <nav>
        <button
          :class="{ active: activeTab === 'users' }"
          @click="activeTab = 'users'"
        >
          Users
        </button>
        <button
          :class="{ active: activeTab === 'notes' }"
          @click="activeTab = 'notes'"
        >
          Notes
        </button>
      </nav>
    </header>
    <main>
      <UserList v-if="activeTab === 'users'" />
      <NoteList v-if="activeTab === 'notes'" />
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

nav button {
  padding: 10px 20px;
  border: none;
  background: #34495e;
  color: white;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1em;
  transition: background 0.2s;
}

nav button:hover {
  background: #4a6278;
}

nav button.active {
  background: #3498db;
}

main {
  background: #fff;
  border-radius: 8px;
  min-height: 400px;
}
</style>
