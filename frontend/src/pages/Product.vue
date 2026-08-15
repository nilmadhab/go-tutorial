<template>
  <div class="product-page">
    <div v-if="loading" class="loading">Loading product...</div>
    <div v-else-if="error" class="error">{{ error }}</div>
    <div v-else-if="product" class="product-detail">
      <div class="product-image">
        <img :src="product.image" :alt="product.name" />
      </div>
      <div class="product-info">
        <span class="category">{{ product.category }}</span>
        <h1>{{ product.name }}</h1>
        <p class="description">{{ product.description }}</p>

        <div class="price-section">
          <span class="price">${{ product.price.toFixed(2) }}</span>
          <span class="stock" :class="{ 'low-stock': product.stock < 50, 'out-of-stock': product.stock === 0 }">
            <template v-if="product.stock === 0">Out of Stock</template>
            <template v-else-if="product.stock < 50">Only {{ product.stock }} left!</template>
            <template v-else>{{ product.stock }} in stock</template>
          </span>
        </div>

        <div class="quantity-section">
          <label>Quantity:</label>
          <div class="quantity-controls">
            <button @click="quantity > 1 && quantity--" :disabled="quantity <= 1">-</button>
            <span>{{ quantity }}</span>
            <button @click="quantity < product.stock && quantity++" :disabled="quantity >= product.stock">+</button>
          </div>
        </div>

        <div class="action-buttons">
          <button class="btn-cart" :disabled="product.stock === 0">
            Add to Cart - ${{ (product.price * quantity).toFixed(2) }}
          </button>
          <button class="btn-buy" :disabled="product.stock === 0">
            Buy Now
          </button>
        </div>

        <div class="product-meta">
          <div class="meta-item">
            <span class="label">Product ID:</span>
            <span>{{ product.id }}</span>
          </div>
          <div class="meta-item">
            <span class="label">Category:</span>
            <span>{{ product.category }}</span>
          </div>
        </div>

        <router-link to="/" class="back-link">&larr; Back to Products</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const product = ref(null)
const loading = ref(true)
const error = ref(null)
const quantity = ref(1)

const fetchProduct = async () => {
  try {
    const response = await fetch(`http://localhost:8080/products/${route.params.id}`)
    if (!response.ok) throw new Error('Product not found')
    product.value = await response.json()
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchProduct()
})
</script>

<style scoped>
.product-page {
  padding: 40px 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.product-detail {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 60px;
  background: white;
  border-radius: 16px;
  padding: 40px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

@media (max-width: 768px) {
  .product-detail {
    grid-template-columns: 1fr;
    gap: 30px;
  }
}

.product-image {
  border-radius: 12px;
  overflow: hidden;
  background: #f5f5f5;
}

.product-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  min-height: 400px;
}

.product-info {
  display: flex;
  flex-direction: column;
}

.category {
  font-size: 14px;
  color: #666;
  text-transform: uppercase;
  letter-spacing: 2px;
  margin-bottom: 8px;
}

h1 {
  font-size: 32px;
  color: #2c3e50;
  margin-bottom: 16px;
  line-height: 1.2;
}

.description {
  font-size: 16px;
  color: #666;
  line-height: 1.6;
  margin-bottom: 24px;
}

.price-section {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 24px;
  padding: 16px 0;
  border-top: 1px solid #eee;
  border-bottom: 1px solid #eee;
}

.price {
  font-size: 36px;
  font-weight: bold;
  color: #2c3e50;
}

.stock {
  font-size: 14px;
  color: #27ae60;
  padding: 4px 12px;
  background: #e8f8f0;
  border-radius: 20px;
}

.stock.low-stock {
  color: #e67e22;
  background: #fef5e7;
}

.stock.out-of-stock {
  color: #e74c3c;
  background: #fdedec;
}

.quantity-section {
  display: flex;
  align-items: center;
  gap: 16px;
  margin-bottom: 24px;
}

.quantity-section label {
  font-weight: 500;
  color: #333;
}

.quantity-controls {
  display: flex;
  align-items: center;
  gap: 12px;
  background: #f5f5f5;
  border-radius: 8px;
  padding: 8px 16px;
}

.quantity-controls button {
  width: 32px;
  height: 32px;
  border: none;
  background: white;
  border-radius: 6px;
  font-size: 18px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.quantity-controls button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.quantity-controls span {
  min-width: 40px;
  text-align: center;
  font-size: 18px;
  font-weight: 600;
}

.action-buttons {
  display: flex;
  gap: 16px;
  margin-bottom: 32px;
}

.btn-cart, .btn-buy {
  flex: 1;
  padding: 16px 24px;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-cart {
  background: #3498db;
  color: white;
}

.btn-cart:hover:not(:disabled) {
  background: #2980b9;
}

.btn-buy {
  background: #27ae60;
  color: white;
}

.btn-buy:hover:not(:disabled) {
  background: #219a52;
}

.btn-cart:disabled, .btn-buy:disabled {
  background: #bdc3c7;
  cursor: not-allowed;
}

.product-meta {
  background: #f8f9fa;
  padding: 16px;
  border-radius: 8px;
  margin-bottom: 24px;
}

.meta-item {
  display: flex;
  gap: 8px;
  padding: 8px 0;
}

.meta-item .label {
  color: #666;
  font-weight: 500;
}

.back-link {
  color: #3498db;
  text-decoration: none;
  font-weight: 500;
}

.back-link:hover {
  text-decoration: underline;
}

.loading, .error {
  text-align: center;
  padding: 60px;
  font-size: 18px;
}

.error {
  color: #e74c3c;
}
</style>
