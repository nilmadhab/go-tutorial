<template>
  <div class="admin">
    <header class="admin-header">
      <h1>Admin Dashboard</h1>
      <p>Manage your products</p>
    </header>

    <div class="stats">
      <div class="stat-card">
        <h3>{{ products.length }}</h3>
        <p>Total Products</p>
      </div>
      <div class="stat-card">
        <h3>{{ totalStock }}</h3>
        <p>Total Stock</p>
      </div>
      <div class="stat-card">
        <h3>${{ totalValue.toFixed(2) }}</h3>
        <p>Inventory Value</p>
      </div>
      <div class="stat-card">
        <h3>{{ lowStockCount }}</h3>
        <p>Low Stock Items</p>
      </div>
    </div>

    <div class="admin-actions">
      <h2>Products</h2>
      <button class="btn-add" @click="openForm()">+ Add Product</button>
    </div>

    <div v-if="loading" class="loading">Loading products...</div>
    <div v-else-if="error" class="error">{{ error }}</div>
    <table v-else class="products-table">
      <thead>
        <tr>
          <th>ID</th>
          <th>Image</th>
          <th>Name</th>
          <th>Category</th>
          <th>Price</th>
          <th>Stock</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="product in products" :key="product.id">
          <td>{{ product.id }}</td>
          <td>
            <img :src="product.image" :alt="product.name" class="table-image" />
          </td>
          <td>{{ product.name }}</td>
          <td>{{ product.category }}</td>
          <td>${{ product.price.toFixed(2) }}</td>
          <td :class="{ 'low-stock': product.stock < 50 }">{{ product.stock }}</td>
          <td class="actions">
            <button class="btn-edit" @click="openForm(product)">Edit</button>
            <button class="btn-delete" @click="deleteProduct(product.id)">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>

    <!-- Product Form Modal -->
    <div v-if="showForm" class="modal-overlay" @click.self="closeForm">
      <div class="modal">
        <h2>{{ editingProduct ? 'Edit Product' : 'Add Product' }}</h2>
        <form @submit.prevent="saveProduct">
          <div class="form-group">
            <label>Name</label>
            <input v-model="form.name" type="text" required />
          </div>
          <div class="form-group">
            <label>Description</label>
            <textarea v-model="form.description" rows="3"></textarea>
          </div>
          <div class="form-row">
            <div class="form-group">
              <label>Price ($)</label>
              <input v-model.number="form.price" type="number" step="0.01" min="0" required />
            </div>
            <div class="form-group">
              <label>Stock</label>
              <input v-model.number="form.stock" type="number" min="0" required />
            </div>
          </div>
          <div class="form-group">
            <label>Category</label>
            <select v-model="form.category" required>
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
            <label>Image URL</label>
            <input v-model="form.image" type="url" placeholder="https://..." />
          </div>
          <div class="form-actions">
            <button type="button" class="btn-cancel" @click="closeForm">Cancel</button>
            <button type="submit" class="btn-save" :disabled="saving">
              {{ saving ? 'Saving...' : 'Save' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'

const products = ref([])
const loading = ref(true)
const error = ref(null)
const showForm = ref(false)
const saving = ref(false)
const editingProduct = ref(null)

const form = ref({
  name: '',
  description: '',
  price: 0,
  stock: 0,
  category: '',
  image: ''
})

const totalStock = computed(() => products.value.reduce((sum, p) => sum + p.stock, 0))
const totalValue = computed(() => products.value.reduce((sum, p) => sum + (p.price * p.stock), 0))
const lowStockCount = computed(() => products.value.filter(p => p.stock < 50).length)

const fetchProducts = async () => {
  try {
    const response = await fetch('http://localhost:8080/products')
    if (!response.ok) throw new Error('Failed to fetch products')
    const data = await response.json()
    products.value = data.products || []
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

const openForm = (product = null) => {
  editingProduct.value = product
  if (product) {
    form.value = { ...product }
  } else {
    form.value = { name: '', description: '', price: 0, stock: 0, category: '', image: '' }
  }
  showForm.value = true
}

const closeForm = () => {
  showForm.value = false
  editingProduct.value = null
}

const saveProduct = async () => {
  saving.value = true
  try {
    const url = editingProduct.value
      ? `http://localhost:8080/products/${editingProduct.value.id}`
      : 'http://localhost:8080/products/create'
    const method = editingProduct.value ? 'PUT' : 'POST'

    const response = await fetch(url, {
      method,
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(form.value)
    })

    if (!response.ok) throw new Error('Failed to save product')

    await fetchProducts()
    closeForm()
  } catch (err) {
    alert(err.message)
  } finally {
    saving.value = false
  }
}

const deleteProduct = async (id) => {
  if (!confirm('Are you sure you want to delete this product?')) return

  try {
    const response = await fetch(`http://localhost:8080/products/${id}`, {
      method: 'DELETE'
    })
    if (!response.ok) throw new Error('Failed to delete product')
    products.value = products.value.filter(p => p.id !== id)
  } catch (err) {
    alert(err.message)
  }
}

onMounted(() => {
  fetchProducts()
})
</script>

<style scoped>
.admin {
  padding: 20px;
}

.admin-header {
  margin-bottom: 30px;
}

.admin-header h1 {
  font-size: 32px;
  color: #2c3e50;
  margin-bottom: 8px;
}

.admin-header p {
  color: #666;
}

.stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
  margin-bottom: 30px;
}

.stat-card {
  background: white;
  padding: 24px;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.stat-card h3 {
  font-size: 28px;
  color: #3498db;
  margin-bottom: 8px;
}

.stat-card p {
  color: #666;
  font-size: 14px;
}

.admin-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.admin-actions h2 {
  font-size: 24px;
  color: #2c3e50;
}

.btn-add {
  padding: 10px 20px;
  background: #27ae60;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
}

.btn-add:hover {
  background: #219a52;
}

.products-table {
  width: 100%;
  background: white;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  border-collapse: collapse;
  overflow: hidden;
}

.products-table th,
.products-table td {
  padding: 16px;
  text-align: left;
  border-bottom: 1px solid #eee;
}

.products-table th {
  background: #f8f9fa;
  font-weight: 600;
  color: #333;
}

.products-table tr:hover {
  background: #f8f9fa;
}

.table-image {
  width: 50px;
  height: 50px;
  object-fit: cover;
  border-radius: 6px;
}

.low-stock {
  color: #e74c3c;
  font-weight: 600;
}

.actions {
  display: flex;
  gap: 8px;
}

.btn-edit, .btn-delete {
  padding: 6px 12px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 12px;
}

.btn-edit {
  background: #3498db;
  color: white;
}

.btn-edit:hover {
  background: #2980b9;
}

.btn-delete {
  background: #e74c3c;
  color: white;
}

.btn-delete:hover {
  background: #c0392b;
}

.modal-overlay {
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

.modal {
  background: white;
  padding: 32px;
  border-radius: 12px;
  width: 100%;
  max-width: 500px;
}

.modal h2 {
  margin-bottom: 24px;
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

.form-group label {
  display: block;
  margin-bottom: 6px;
  font-weight: 500;
  color: #333;
}

.form-group input,
.form-group textarea,
.form-group select {
  width: 100%;
  padding: 10px 12px;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 14px;
}

.form-group input:focus,
.form-group textarea:focus,
.form-group select:focus {
  outline: none;
  border-color: #3498db;
}

.form-actions {
  display: flex;
  gap: 12px;
  margin-top: 24px;
}

.btn-cancel, .btn-save {
  flex: 1;
  padding: 12px;
  border: none;
  border-radius: 6px;
  font-size: 16px;
  cursor: pointer;
}

.btn-cancel {
  background: #e0e0e0;
  color: #333;
}

.btn-save {
  background: #3498db;
  color: white;
}

.btn-save:disabled {
  background: #bdc3c7;
}

.loading, .error {
  text-align: center;
  padding: 40px;
  font-size: 18px;
}

.error {
  color: #e74c3c;
}
</style>
