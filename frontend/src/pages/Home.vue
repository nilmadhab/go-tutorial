<template>
  <div class="home">
    <header>
      <h1>E-Commerce Store</h1>
      <p>Browse our amazing products</p>
    </header>

    <main>
      <div v-if="loading" class="loading">Loading products...</div>
      <div v-else-if="error" class="error">{{ error }}</div>
      <div v-else class="products-grid">
        <div
          v-for="product in products"
          :key="product.id"
          class="product-card"
          @click="$router.push(`/product/${product.id}`)"
        >
          <div class="product-image">
            <img :src="product.image" :alt="product.name" />
          </div>
          <div class="product-info">
            <span class="category">{{ product.category }}</span>
            <h3>{{ product.name }}</h3>
            <p class="description">{{ product.description }}</p>
            <div class="product-footer">
              <span class="price">${{ product.price.toFixed(2) }}</span>
              <span class="stock" :class="{ 'low-stock': product.stock < 50 }">
                {{ product.stock }} in stock
              </span>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const products = ref([])
const loading = ref(true)
const error = ref(null)

const fetchProducts = async () => {
  try {
    const response = await fetch('http://localhost:8080/products')
    if (!response.ok) {
      throw new Error('Failed to fetch products')
    }
    const data = await response.json()
    products.value = data.products || []
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchProducts()
})
</script>

<style scoped>
header {
  text-align: center;
  padding: 40px 0;
}

header h1 {
  font-size: 36px;
  color: #2c3e50;
  margin-bottom: 8px;
}

header p {
  color: #666;
  font-size: 18px;
}

.products-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 24px;
  padding: 20px 0;
}

.product-card {
  background: white;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  transition: transform 0.2s, box-shadow 0.2s;
  cursor: pointer;
}

.product-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
}

.product-image {
  width: 100%;
  height: 200px;
  background: #f5f5f5;
}

.product-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.product-info {
  padding: 16px;
}

.category {
  font-size: 12px;
  color: #666;
  text-transform: uppercase;
  letter-spacing: 1px;
}

h3 {
  margin: 8px 0;
  font-size: 18px;
  color: #333;
}

.description {
  font-size: 14px;
  color: #666;
  line-height: 1.4;
  margin-bottom: 16px;
}

.product-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.price {
  font-size: 20px;
  font-weight: bold;
  color: #2c3e50;
}

.stock {
  font-size: 12px;
  color: #27ae60;
}

.stock.low-stock {
  color: #e74c3c;
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
