<template>
  <div class="user-list">
    <div class="header-row">
      <h2>Users</h2>
      <button class="add-btn" @click="showForm = !showForm">
        {{ showForm ? 'Cancel' : '+ Add User' }}
      </button>
    </div>

    <!-- Add User Form -->
    <div v-if="showForm" class="add-form">
      <h3>Add New User</h3>
      <form @submit.prevent="createUser">
        <div class="form-group">
          <label for="username">Username</label>
          <input
            id="username"
            v-model="newUser.username"
            type="text"
            placeholder="Enter username"
            required
          />
        </div>
        <div class="form-group">
          <label for="password">Password</label>
          <input
            id="password"
            v-model="newUser.password"
            type="password"
            placeholder="Enter password"
            required
          />
        </div>
        <div class="form-actions">
          <button type="submit" :disabled="creating">
            {{ creating ? 'Creating...' : 'Create User' }}
          </button>
        </div>
        <p v-if="formError" class="form-error">{{ formError }}</p>
      </form>
    </div>

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

        <div class="user-summary">
          <span>{{ user.notes?.length || 0 }} notes</span>
        </div>

        <router-link :to="`/users/${user.id}`" class="view-details-btn">
          View Details &rarr;
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const users = ref([])
const loading = ref(true)
const error = ref(null)
const showForm = ref(false)
const creating = ref(false)
const formError = ref(null)
const newUser = ref({ username: '', password: '' })

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

const createUser = async () => {
  creating.value = true
  formError.value = null

  try {
    const response = await fetch('http://localhost:8080/api/users', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newUser.value)
    })

    if (!response.ok) {
      const data = await response.json()
      throw new Error(data.error || 'Failed to create user')
    }

    // Reset form and refresh list
    newUser.value = { username: '', password: '' }
    showForm.value = false
    await fetchUsers()
  } catch (e) {
    formError.value = e.message
  } finally {
    creating.value = false
  }
}

onMounted(fetchUsers)
</script>

<style scoped>
.user-list {
  padding: 20px;
}

.header-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

h2 {
  margin: 0;
  color: #333;
}

.add-btn {
  padding: 10px 20px;
  background: #27ae60;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.95em;
}

.add-btn:hover {
  background: #219a52;
}

.add-form {
  background: #f8f9fa;
  padding: 20px;
  border-radius: 8px;
  margin-bottom: 20px;
  border: 1px solid #e0e0e0;
}

.add-form h3 {
  margin: 0 0 15px 0;
  color: #2c3e50;
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
  color: #555;
  font-weight: 500;
}

.form-group input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1em;
}

.form-group input:focus {
  outline: none;
  border-color: #3498db;
}

.form-actions {
  margin-top: 15px;
}

.form-actions button {
  padding: 10px 25px;
  background: #3498db;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1em;
}

.form-actions button:hover:not(:disabled) {
  background: #2980b9;
}

.form-actions button:disabled {
  background: #95a5a6;
  cursor: not-allowed;
}

.form-error {
  color: #e74c3c;
  margin-top: 10px;
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

.user-summary {
  color: #666;
  font-size: 0.9em;
  margin-bottom: 15px;
}

.view-details-btn {
  display: block;
  text-align: center;
  padding: 10px 15px;
  background: #3498db;
  color: white;
  text-decoration: none;
  border-radius: 4px;
  font-size: 0.95em;
  transition: background 0.2s;
}

.view-details-btn:hover {
  background: #2980b9;
}

.loading, .error {
  text-align: center;
  padding: 40px;
}

.error {
  color: #e74c3c;
}
</style>
