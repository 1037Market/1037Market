<template>
  <div class="price-display">
    <div class="price-container">
      <span class="currency">¥</span>
      <span class="price">{{ price }}</span>
    </div>
    <div class="name-and-tags">
      <div class="name">{{ name }}</div>
      <div class="tags">
        <van-tag type="primary" size="medium"
                 v-for="category in categories"
                 :key="category"
                 style="margin: 5px"
        >
          {{ category }}
        </van-tag>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PriceDisplay',
  props: {
    name: {
      type: String,
      required: true,
    },
    price: {
      type: Number,
      required: true,
    },
    categories: {
      type: Array,
      required: true
    }
  },
};
</script>

<style scoped>
.price-display {
  display: flex;
  flex-direction: column; /* Change to column layout */
  align-items: flex-start; /* Align to the start of the flex container */
  padding: 15px;
  margin: 10px;
  color: #333;
  border-radius: 15px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  font-family: 'Poppins', sans-serif;
}

.price-container {
  display: flex;
  align-items: baseline;
}

.currency {
  font-size: 1.6em;
  margin-right: 5px;
  color: #333;
}

.price {
  font-size: 2.8em;
  font-weight: bold;
  letter-spacing: 1px;
}

.name-and-tags {
  align-self: stretch; /* Stretch to fill the width */
}

.name {
  font-size: 1.2em;
  text-align: right;
  opacity: 0.85;
  text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.1);
}

.tags {
  text-align: right; /* Align tags to the right */
  margin-top: 5px; /* Spacing between name and tags */
}
</style>
