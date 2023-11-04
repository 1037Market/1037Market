<template>
  <div class="product-description">
    <p class="description-content">
      {{ description }}
    </p>
  </div>
</template>

<script>
export default {
  name: 'ProductDescription',
  props: {
    description: {
      type: String,
      required: true,
    },
  },
};
</script>

<style scoped>
.product-description {
  font-family: 'Poppins', sans-serif;
  margin: 10px 10px 50px;
  padding: 15px;
  border-radius: 15px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  color: #333;
  line-break: anywhere;
  text-align: left;

}
</style>
