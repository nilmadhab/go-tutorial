<template>
  <div class="user-detail">
    <button class="back-btn" @click="$router.push('/')">
      &larr; Back to Users
    </button>

    <div v-if="loading" class="loading">Loading...</div>
    <div v-else-if="error" class="error">{{ error }}</div>
    <div v-else class="user-content">
      <div class="user-header">
        <h1>{{ user.username }}</h1>
        <span class="user-id">User #{{ user.id }}</span>
      </div>

      <!-- Credit Card Section -->
      <section class="section">
        <h2>Credit Card</h2>
        <div v-if="user.credit_card" class="credit-card-display">
          <div class="card-visual">
            <div class="card-chip"></div>
            <div class="card-number">{{ user.credit_card.number }}</div>
            <div class="card-holder">{{ user.username }}</div>
          </div>
        </div>
        <div v-else class="empty-state">
          No credit card on file
        </div>
      </section>

      <!-- Notes Section -->
      <section class="section">
        <div class="section-header">
          <h2>Notes ({{ user.notes?.length || 0 }})</h2>
          <button class="add-btn" @click="showNoteForm = !showNoteForm">
            {{ showNoteForm ? 'Cancel' : '+ Add Note' }}
          </button>
        </div>

        <!-- Add Note Form -->
        <div v-if="showNoteForm" class="add-form">
          <form @submit.prevent="createNote">
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

        <div v-if="user.notes && user.notes.length" class="notes-list">
          <div v-for="note in user.notes" :key="note.id" class="note-item">
            <div class="note-header">
              <h3>{{ note.name }}</h3>
              <span class="note-id">#{{ note.id }}</span>
            </div>
            <p>{{ note.content }}</p>
          </div>
        </div>
        <div v-else class="empty-state">
          No notes yet
        </div>
      </section>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const user = ref({})
const loading = ref(true)
const error = ref(null)
const showNoteForm = ref(false)
const creating = ref(false)
const formError = ref(null)
const newNote = ref({ name: '', content: '' })

const fetchUser = async () => {
  try {
    const response = await fetch(`http://localhost:8080/api/users/${route.params.id}`)
    if (!response.ok) throw new Error('User not found')
    user.value = await response.json()
  } catch (e) {
    error.value = e.message
  } finally {
    loading.value = false
  }
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
        user_id: Number(route.params.id)
      })
    })

    if (!response.ok) {
      const data = await response.json()
      throw new Error(data.error || 'Failed to create note')
    }

    newNote.value = { name: '', content: '' }
    showNoteForm.value = false
    await fetchUser()
  } catch (e) {
    formError.value = e.message
  } finally {
    creating.value = false
  }
}

onMounted(fetchUser)
</script>

<style scoped>
.user-detail {
  padding: 20px;
}

.back-btn {
  padding: 10px 20px;
  background: #ecf0f1;
  color: #2c3e50;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.95em;
  margin-bottom: 20px;
}

.back-btn:hover {
  background: #bdc3c7;
}

.user-content {
  max-width: 800px;
}

.user-header {
  display: flex;
  align-items: center;
  gap: 15px;
  margin-bottom: 30px;
  padding-bottom: 20px;
  border-bottom: 2px solid #eee;
}

.user-header h1 {
  margin: 0;
  color: #2c3e50;
}

.user-id {
  background: #3498db;
  color: white;
  padding: 5px 12px;
  border-radius: 20px;
  font-size: 0.85em;
}

.section {
  margin-bottom: 30px;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

.section h2 {
  margin: 0 0 15px 0;
  color: #2c3e50;
  font-size: 1.3em;
}

.section-header h2 {
  margin: 0;
}

/* Credit Card Styles */
.credit-card-display {
  max-width: 380px;
}

.card-visual {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 16px;
  padding: 25px;
  color: white;
  box-shadow: 0 10px 30px rgba(102, 126, 234, 0.4);
}

.card-chip {
  width: 50px;
  height: 40px;
  background: linear-gradient(135deg, #ffd700 0%, #ffaa00 100%);
  border-radius: 8px;
  margin-bottom: 30px;
}

.card-number {
  font-family: 'Courier New', monospace;
  font-size: 1.4em;
  letter-spacing: 3px;
  margin-bottom: 20px;
}

.card-holder {
  text-transform: uppercase;
  font-size: 0.9em;
  letter-spacing: 2px;
  opacity: 0.9;
}

/* Notes Styles */
.notes-list {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.note-item {
  background: #f8f9fa;
  border-radius: 8px;
  padding: 20px;
  border-left: 4px solid #3498db;
}

.note-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.note-header h3 {
  margin: 0;
  color: #2c3e50;
}

.note-id {
  color: #888;
  font-size: 0.85em;
}

.note-item p {
  margin: 0;
  color: #555;
  line-height: 1.6;
}

.empty-state {
  color: #888;
  font-style: italic;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 8px;
  text-align: center;
}

.add-btn {
  padding: 8px 16px;
  background: #27ae60;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 0.9em;
}

.add-btn:hover {
  background: #219a52;
}

.add-form {
  background: #fff;
  padding: 20px;
  border-radius: 8px;
  margin-bottom: 20px;
  border: 1px solid #e0e0e0;
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
.form-group textarea {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1em;
  font-family: inherit;
}

.form-group input:focus,
.form-group textarea:focus {
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
}

.form-actions button:disabled {
  background: #95a5a6;
  cursor: not-allowed;
}

.form-error {
  color: #e74c3c;
  margin-top: 10px;
}

.loading, .error {
  text-align: center;
  padding: 40px;
}

.error {
  color: #e74c3c;
}
</style>
