<template>
  <div class="note-list">
    <div class="header-row">
      <h2>All Notes</h2>
      <button class="add-btn" @click="toggleForm">
        {{ showForm ? 'Cancel' : '+ Add Note' }}
      </button>
    </div>

    <!-- Add Note Form -->
    <div v-if="showForm" class="add-form">
      <h3>Add New Note</h3>
      <form @submit.prevent="createNote">
        <div class="form-group">
          <label for="user">User</label>
          <select id="user" v-model="newNote.user_id" required>
            <option value="" disabled>Select a user</option>
            <option v-for="user in users" :key="user.id" :value="user.id">
              {{ user.username }} (#{{ user.id }})
            </option>
          </select>
        </div>
        <div class="form-group">
          <label for="name">Note Title</label>
          <input
            id="name"
            v-model="newNote.name"
            type="text"
            placeholder="Enter note title"
            required
          />
        </div>
        <div class="form-group">
          <label for="content">Content</label>
          <textarea
            id="content"
            v-model="newNote.content"
            placeholder="Enter note content"
            rows="3"
            required
          ></textarea>
        </div>
        <div class="form-actions">
          <button type="submit" :disabled="creating">
            {{ creating ? 'Creating...' : 'Create Note' }}
          </button>
        </div>
        <p v-if="formError" class="form-error">{{ formError }}</p>
      </form>
    </div>

    <div v-if="loading" class="loading">Loading...</div>
    <div v-else-if="error" class="error">{{ error }}</div>
    <div v-else class="notes-grid">
      <div v-for="note in notes" :key="note.id" class="note-card">
        <div class="note-header">
          <h3>{{ note.name }}</h3>
          <span class="note-id">#{{ note.id }}</span>
        </div>
        <p class="note-content">{{ note.content }}</p>
        <div class="note-meta">
          User ID: {{ note.user_id }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const notes = ref([])
const users = ref([])
const loading = ref(true)
const error = ref(null)
const showForm = ref(false)
const creating = ref(false)
const formError = ref(null)
const newNote = ref({ name: '', content: '', user_id: '' })

const fetchNotes = async () => {
  try {
    const response = await fetch('http://localhost:8080/api/notes')
    if (!response.ok) throw new Error('Failed to fetch notes')
    notes.value = await response.json()
  } catch (e) {
    error.value = e.message
  } finally {
    loading.value = false
  }
}

const fetchUsers = async () => {
  try {
    const response = await fetch('http://localhost:8080/api/users')
    if (!response.ok) throw new Error('Failed to fetch users')
    users.value = await response.json()
  } catch (e) {
    console.error('Failed to fetch users:', e)
  }
}

const toggleForm = async () => {
  if (!showForm.value) {
    await fetchUsers()
  }
  showForm.value = !showForm.value
}

const createNote = async () => {
  creating.value = true
  formError.value = null

  try {
    const response = await fetch('http://localhost:8080/api/notes', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        ...newNote.value,
        user_id: Number(newNote.value.user_id)
      })
    })

    if (!response.ok) {
      const data = await response.json()
      throw new Error(data.error || 'Failed to create note')
    }

    // Reset form and refresh list
    newNote.value = { name: '', content: '', user_id: '' }
    showForm.value = false
    await fetchNotes()
  } catch (e) {
    formError.value = e.message
  } finally {
    creating.value = false
  }
}

onMounted(fetchNotes)
</script>

<style scoped>
.note-list {
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

.form-group input,
.form-group select,
.form-group textarea {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1em;
  font-family: inherit;
}

.form-group input:focus,
.form-group select:focus,
.form-group textarea:focus {
  outline: none;
  border-color: #3498db;
}

.form-group textarea {
  resize: vertical;
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

.notes-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 16px;
}

.note-card {
  background: #fff;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  padding: 16px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s, box-shadow 0.2s;
}

.note-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
}

.note-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}

.note-header h3 {
  margin: 0;
  color: #2c3e50;
  font-size: 1.1em;
}

.note-id {
  color: #888;
  font-size: 0.85em;
}

.note-content {
  color: #555;
  line-height: 1.5;
  margin: 0 0 12px 0;
}

.note-meta {
  font-size: 0.85em;
  color: #888;
  padding-top: 10px;
  border-top: 1px solid #eee;
}

.loading, .error {
  text-align: center;
  padding: 40px;
}

.error {
  color: #e74c3c;
}
</style>
