<template>
  <div class="product-card">
    <div class="product-image">
      <img :src="`https://via.placeholder.com/200?text=${product.name.split(' ')[0]}`" :alt="product.name" />
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
</template>

<script setup>
defineProps({
  product: {
    type: Object,
    required: true
  }
})
</script>

<style scoped>
.product-card {
  background: white;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  transition: transform 0.2s, box-shadow 0.2s;
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
</style>
