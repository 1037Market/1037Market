<template>
  <div class="display seller-products">
    <h3>{{ title }}</h3>
    <div class="products-grid">
      <div class="product-card" v-for="product in displayedProducts" :key="product.id" @click="clickProduct(product.id)">
        <div v-if="product.soldout" class="soldout-overlay">已售出</div>
        <van-image :src="product.image" alt="Product Image" class="product-image" />
        <div class="product-info">
          <h4 class="product-name">{{ product.name }}</h4>
          <p class="product-description">{{ product.description }}</p>
        </div>
      </div>
    </div>
    <button v-if="seller.products.length > 2 && !showAllProducts" @click="toggleDisplay" class="view-more-btn">
      查看更多
    </button>
    <button v-if="showAllProducts" @click="toggleDisplay" class="collapse-btn">
      收起
    </button>
  </div>
</template>


<script setup>
import { useRouter } from 'vue-router';
import { ref, computed } from 'vue';

const props = defineProps({
  title: String,
  seller: {
    type: Object,
    default: () => ({
      avatar: '',
      nickname: '',
      comments: [],
      products: []
    }),
  },
});

const router = useRouter();
const showAllProducts = ref(false);

const displayedProducts = computed(() => {
  return showAllProducts.value ? props.seller.products : props.seller.products.slice(0, 2);
});

const toggleDisplay = () => {
  showAllProducts.value = !showAllProducts.value;
};

const clickProduct = (productId) => {
  router.push({ path: `/detail/${productId}` });
};
</script>

<style scoped>
.products-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
  gap: 10px;
  margin-bottom: 10px;
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

.view-more-btn,  .collapse-btn {
  background-color: transparent;
  border: none;
  color: #1f8efa;
  padding: 10px;
  font-size: 14px;
  cursor: pointer;
  display: block;
  margin: 10px auto; /* Center the button */
}

.view-more-btn:hover, .collapse-btn:hover {
  text-decoration: underline;
}

.soldout-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 70%; /* 与.product-image高度相同 */
  background-color: rgba(0, 0, 0, 0.5); /* 半透明背景 */
  color: white;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 20px;
  font-weight: bold;
  z-index: 10;
}

.product-card {
  /* ...其他样式... */
  position: relative;
}


</style>
