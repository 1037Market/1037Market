<template>
  <!-- Published Products -->
  <div class="display seller-products">
    <h3>{{ title }}</h3>
    <div class="products-grid">
      <div class="product-card" v-for="product in seller.products" :key="product.id" @click="clickProduct(product.id)">
        <van-image :src="product.image" alt="Product Image" class="product-image" />
        <div class="product-info">
          <h4 class="product-name">{{ product.name }}</h4>
          <p class="product-description">{{ product.description }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {useRouter} from "vue-router";


export default {
  name: "ProfileProducts",
  props: {
    title: '',
    seller: {
      avatar: '',
      nickname: '',
      comments: [
      ],
      products: [
      ]
    }
  },

  setup() {
    const router = useRouter()
    const clickProduct = (productId) => {
      router.push({path: `/detail/${productId}`})
    }
    return {
      clickProduct
    }
  }
}
</script>

<style scoped>

.products-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
  gap: 10px;
  margin-bottom: 45px;
}

.product-card {
  border: 1px solid #e1e1e1;
  border-radius: 6px;
  overflow: hidden;
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
  height: 200px;
}

.product-image {
  width: 100%;
  height: 70%;
  overflow: hidden;
  object-fit: contain;
}

.product-info {
  padding: 8px;
}

.product-name {
  font-size: 16px;
  font-weight: bold;
  color: #333;
  margin: 0 0 5px 0;
}

.product-description {
  font-size: 12px;
  color: #666;
}

.seller-products h3 {
  margin-top: 20px;
  margin-bottom: 20px;
  text-align: center;
}
</style>