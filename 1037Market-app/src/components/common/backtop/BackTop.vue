<template>
  <div class="backtop" @click="goback"></div>
</template>

<script>
export default {
  setup(props, ctx) {
    const goback = () => {
      ctx.emit("goback");
    };
    return {
      goback,
    };
  },
};
</script>

<style scoped>
.backtop {
  width: 40px;
  height: 40px;
  position: fixed;
  bottom: 80px;
  right: 30px;
  z-index: 90;
  box-shadow: 5px 5px 5px #888888;
  border-radius: 50%;
  background: #fcfcfc url("~assets/images/up.png") no-repeat center center;
}
</style>