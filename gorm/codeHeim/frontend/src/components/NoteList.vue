<template>
  <div class="note-list">
    <h2>All Notes</h2>
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
const loading = ref(true)
const error = ref(null)

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

onMounted(fetchNotes)
</script>

<style scoped>
.note-list {
  padding: 20px;
}

h2 {
  margin-bottom: 20px;
  color: #333;
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
