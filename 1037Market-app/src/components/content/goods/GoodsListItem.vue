<template>
    <div class="goods-item" @click="itemClick()">
<!--    <img v-lazy="product.cover_url" alt="" />-->
        <img v-lazy="tempImg" alt=""/>
      <div class="goods-info">
          <p>{{ productDetail.title }}</p>
          <span class="price"><small>￥</small>{{ productDetail.price }}</span>
<!--          <span class="collect">{{ product.collects_count }}</span>-->
      </div>
    </div>
</template>

<script setup>
import { useRouter } from "vue-router";
import { onMounted, ref } from "vue"
import { getDetail } from "network/detail"
const props = defineProps(["productId"])

const router = useRouter();
const productDetail = ref({})
const tempImg = ref('https://picgo-xqaqyn.oss-cn-shanghai.aliyuncs.com/img/b_256781c21be0a7ed01ebd36a6d7c72c0.jpg')
const itemClick = () => {
    router.push({ path: `/detail/${props.productId}`});
};

onMounted(() => {
    // console.log('item id:',props.productId)
    getDetail(props.productId).then((res) => {
        // console.log("response:",res)
        productDetail.value = res
        // console.log("detail:",productDetail)
    })
})
</script>

<style scoped lang="scss">
.goods-item {
  width: 46%;
  padding-bottom: 40px;
  position: relative;

  img {
    width: 100%;
    border-radius: 5px;
  }

  .goods-info {
    font-size: 12px;
    position: absolute;
    bottom: 5px;
    left: 0;
    right: 0;
    overflow: hidden;
    text-align: center;
    p {
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      margin-bottom: 3px;
    }

    .price {
      color: red;
      margin-right: 20px;
    }

    .collect {
      position: relative;
    }
    .collect::before {
      content: "";
      position: absolute;
      left: -15px;
      width: 14px;
      height: 14px;
      top: -1px;
      //background: url("~assets/images/collect.png") 0 0/14px 14px;
    }
  }
}
</style>