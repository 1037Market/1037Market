<template>
  <div class="goods">
      <div class="goods-item" v-for="productDetail in products" @click="itemClick(productDetail.productId)">
          <van-image :src="Array.isArray(productDetail.imageURIs) && productDetail.imageURIs.length > 0 ? 'http://franky.pro:7301/api/image?imageURI=' + productDetail.imageURIs[0] : ''"/>
          <div class="goods-info">
              <p style="font-size: 14px"><strong>{{ productDetail.title }}</strong> </p>
              <span class="price" style="font-size: 13px"><small>￥</small>{{ productDetail.price }}</span>
          </div>
      </div>
  </div>
</template>

<script setup>
    import {getDetail} from "@/network/detail";
    import {onMounted, ref, watch} from "vue";
    import {useRouter} from "vue-router";

    const router = useRouter()

    const props = defineProps({
        productIDs: Array
    })

    watch(() => props.productIDs, () => {
        updatePage()
    })

    const products = ref([])
    const updatePage = () => {
        products.value.length = 0
        for(const productId of props.productIDs)
            getDetail(productId).then((resp) => {
                products.value.push(resp)
            })
    }

    onMounted(() => {
        updatePage()
    });

    const itemClick = (productId) => {
        router.push({ path: `/detail/${productId}`});
    };

</script>

<style scoped>
.goods {
  flex-wrap: wrap;
  justify-content: space-around;
  padding: 5px;
    columns: 2;
    column-gap: 16px;
    width: 100%;
}

.goods-item {
    break-inside: avoid-column;
    margin-bottom: 16px;
    display: inline-block;
    width: 100%;
}
.goods-item van-image{
    width: 100%;
    height: auto;
    border-radius: 5px;
}
.goods-item .goods-info {
    font-size: 12px;
    position: relative;
    margin-top: 10px;
    bottom: 5px;
    left: 0;
    right: 0;
    overflow: hidden;
    text-align: center;
}
.goods-item .goods-info p {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    margin-bottom: 3px;
}
.goods-item .goods-info .price {
    color: red;
    margin-right: 20px;
}
</style>