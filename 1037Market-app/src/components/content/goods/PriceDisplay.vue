<template>
  <div class="price-display">
    <div class="price-container">
      <span class="currency">¥</span>
      <span class="price">{{ price }}</span>
    </div>
    <div class="name">{{ name }}</div>
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
  },
};
</script>

<style scoped>
.price-display {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px;
  margin: 10px; /* Added horizontal margin */
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
  font-size: 1.6em; /* Slightly larger to emphasize currency */
  margin-right: 5px;
  color: #333; /* Dark color for contrast */
}

.price {
  font-size: 2.8em; /* Slightly larger for more emphasis */
  font-weight: bold;
  letter-spacing: 1px;
}

.name {
  font-size: 1.2em;
  text-align: right;
  flex: 1;
  opacity: 0.85; /* Slightly transparent to blend with the background */
  text-shadow: 1px 1px 1px rgba(0, 0, 0, 0.1);
}
</style>
