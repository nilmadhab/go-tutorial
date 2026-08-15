<template>
  <div class="form-overlay" @click.self="$emit('close')">
    <div class="form-container">
      <h2>Add New Product</h2>
      <form @submit.prevent="submitForm">
        <div class="form-group">
          <label for="name">Product Name</label>
          <input type="text" id="name" v-model="form.name" required />
        </div>

        <div class="form-group">
          <label for="description">Description</label>
          <textarea id="description" v-model="form.description" rows="3"></textarea>
        </div>

        <div class="form-row">
          <div class="form-group">
            <label for="price">Price ($)</label>
            <input type="number" id="price" v-model.number="form.price" step="0.01" min="0" required />
          </div>

          <div class="form-group">
            <label for="stock">Stock</label>
            <input type="number" id="stock" v-model.number="form.stock" min="0" required />
          </div>
        </div>

        <div class="form-group">
          <label for="category">Category</label>
          <select id="category" v-model="form.category" required>
            <option value="">Select category</option>
            <option value="Electronics">Electronics</option>
            <option value="Clothing">Clothing</option>
            <option value="Home & Kitchen">Home & Kitchen</option>
            <option value="Sports">Sports</option>
            <option value="Books">Books</option>
            <option value="Other">Other</option>
          </select>
        </div>

        <div class="form-group">
          <label for="image">Image URL</label>
          <input type="url" id="image" v-model="form.image" placeholder="https://..." />
        </div>

        <div class="form-actions">
          <button type="button" class="btn-cancel" @click="$emit('close')">Cancel</button>
          <button type="submit" class="btn-submit" :disabled="loading">
            {{ loading ? 'Adding...' : 'Add Product' }}
          </button>
        </div>

        <p v-if="error" class="error-message">{{ error }}</p>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const emit = defineEmits(['close', 'product-added'])

const loading = ref(false)
const error = ref(null)

const form = ref({
  name: '',
  description: '',
  price: 0,
  stock: 0,
  category: '',
  image: ''
})

const submitForm = async () => {
  loading.value = true
  error.value = null

  try {
    const response = await fetch('http://localhost:8080/products/create', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(form.value)
    })

    if (!response.ok) {
      throw new Error('Failed to create product')
    }

    const newProduct = await response.json()
    emit('product-added', newProduct)
    emit('close')
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.form-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.form-container {
  background: white;
  padding: 32px;
  border-radius: 12px;
  width: 100%;
  max-width: 500px;
  max-height: 90vh;
  overflow-y: auto;
}

h2 {
  margin: 0 0 24px;
  color: #2c3e50;
}

.form-group {
  margin-bottom: 16px;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}

label {
  display: block;
  margin-bottom: 6px;
  font-weight: 500;
  color: #333;
}

input, textarea, select {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 14px;
  transition: border-color 0.2s;
}

input:focus, textarea:focus, select:focus {
  outline: none;
  border-color: #3498db;
}

.form-actions {
  display: flex;
  gap: 12px;
  margin-top: 24px;
}

.btn-cancel, .btn-submit {
  flex: 1;
  padding: 12px;
  border: none;
  border-radius: 6px;
  font-size: 16px;
  cursor: pointer;
  transition: background 0.2s;
}

.btn-cancel {
  background: #e0e0e0;
  color: #333;
}

.btn-cancel:hover {
  background: #d0d0d0;
}

.btn-submit {
  background: #3498db;
  color: white;
}

.btn-submit:hover {
  background: #2980b9;
}

.btn-submit:disabled {
  background: #bdc3c7;
  cursor: not-allowed;
}

.error-message {
  color: #e74c3c;
  margin-top: 16px;
  text-align: center;
}
</style>
