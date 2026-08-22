<template>
  <div class="user-list">
    <h2>Users</h2>
    <div v-if="loading" class="loading">Loading...</div>
    <div v-else-if="error" class="error">{{ error }}</div>
    <div v-else class="users-grid">
      <div v-for="user in users" :key="user.id" class="user-card">
        <div class="user-header">
          <h3>{{ user.username }}</h3>
          <span class="user-id">#{{ user.id }}</span>
        </div>

        <div v-if="user.credit_card" class="credit-card">
          <strong>Credit Card:</strong> {{ maskCardNumber(user.credit_card.number) }}
        </div>

        <div v-if="user.notes && user.notes.length" class="user-notes">
          <strong>Notes ({{ user.notes.length }}):</strong>
          <ul>
            <li v-for="note in user.notes" :key="note.id">
              <span class="note-name">{{ note.name }}</span>
              <span class="note-content">{{ note.content }}</span>
            </li>
          </ul>
        </div>
        <div v-else class="no-notes">No notes</div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const users = ref([])
const loading = ref(true)
const error = ref(null)

const maskCardNumber = (number) => {
  if (!number) return ''
  const parts = number.split('-')
  return `****-****-****-${parts[parts.length - 1]}`
}

const fetchUsers = async () => {
  try {
    const response = await fetch('http://localhost:8080/api/users')
    if (!response.ok) throw new Error('Failed to fetch users')
    users.value = await response.json()
  } catch (e) {
    error.value = e.message
  } finally {
    loading.value = false
  }
}

onMounted(fetchUsers)
</script>

<style scoped>
.user-list {
  padding: 20px;
}

h2 {
  margin-bottom: 20px;
  color: #333;
}

.users-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
}

.user-card {
  background: #fff;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.user-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
  padding-bottom: 10px;
  border-bottom: 1px solid #eee;
}

.user-header h3 {
  margin: 0;
  color: #2c3e50;
}

.user-id {
  color: #888;
  font-size: 0.9em;
}

.credit-card {
  background: #f8f9fa;
  padding: 10px;
  border-radius: 4px;
  margin-bottom: 15px;
  font-family: monospace;
}

.user-notes ul {
  list-style: none;
  padding: 0;
  margin: 10px 0 0 0;
}

.user-notes li {
  background: #f0f7ff;
  padding: 10px;
  margin: 8px 0;
  border-radius: 4px;
  border-left: 3px solid #3498db;
}

.note-name {
  font-weight: bold;
  display: block;
  color: #2c3e50;
}

.note-content {
  color: #666;
  font-size: 0.9em;
}

.no-notes {
  color: #999;
  font-style: italic;
}

.loading, .error {
  text-align: center;
  padding: 40px;
}

.error {
  color: #e74c3c;
}
</style>
