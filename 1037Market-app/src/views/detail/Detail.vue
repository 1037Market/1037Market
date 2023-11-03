<template>
  <div>
    <nav-bar>
      <template v-slot:center>商品详情:{{ id }}</template>
    </nav-bar>

    <van-image
      style="margin-top: 50px"
      width="100%"
      lazy-load
      src="https://picgo-xqaqyn.oss-cn-shanghai.aliyuncs.com/img/b_256781c21be0a7ed01ebd36a6d7c72c0.jpg"
    />

    <van-card
      style="text-align: left"
      :price="productDetail.price + '.00'"
      :desc="productDetail.content"
      :title="productDetail.title"
    >
<!--      <template #tags>-->
<!--        <van-tag plain type="danger">新书</van-tag>-->
<!--        <van-tag plain type="danger">推荐</van-tag>-->
<!--      </template>-->
      <template #footer>
        <van-button type="warning" @click="handleCart">加入购物车</van-button>
        <van-button type="danger" @click="goToCart">立即购买</van-button>
      </template>
    </van-card>

    <van-tabs v-model="active">
      <van-tab title="概述">
        <div id="con1" v-html="productDetail.content"></div>
      </van-tab>
      <van-tab title="热评"> </van-tab>
<!--      <van-tab title="相关图书">-->
<!--        <goods-list :goods="like_goods"></goods-list>-->
<!--      </van-tab>-->
    </van-tabs>
  </div>
</template>

<script>
import NavBar from "components/common/navbar/NavBar";
import GoodsList from "components/content/goods/GoodsList";

import { useRoute } from "vue-router";
import { useRouter } from "vue-router";
import { useStore } from "vuex";

import { ref, onMounted, reactive, toRefs } from "vue";
import { getDetail } from "network/detail";
import { Toast } from "vant";
import { addCart } from "network/cart";
export default {
  name: "Detail",
  components: {
    NavBar,
    GoodsList,
  },
  setup() {
    const route = useRoute();
    const router = useRouter();
    const store = useStore();

    let id = ref(route.params.id);
    const productDetail = ref({})
    let active = ref(1);

    const handleCart = () => {
        console.log('购物车功能未实现')
      // addCart({ goods_id: book.detail.id, num: 1 }).then((res) => {
      //   if (res.status == "204" || res.status == "201") {
      //     Toast.success("添加成功");
      //     调用actions
          // store.dispatch("updateCart");
        // }
      // });
    };

    const goToCart = () => {
        console.log('购物车功能未实现')
      // addCart({ goods_id: book.detail.id, num: 1 }).then((res) => {
      //   if (res.status == "204" || res.status == "201") {
      //     Toast.success("添加成功,跳转到购物车");
      //     router.push({ path: "/shopcart" });
      //     store.dispatch("updateCart");
      //   }
      // });
    };

    onMounted(() => {

      getDetail(id.value).then((res) => {
          productDetail.value = res
      })
    });

    return {
        id,
        active,
        handleCart,
        goToCart,
        productDetail
    };
  },
};
</script>

<style scoped lang="scss">
#con1 {
  padding: 10px;
}
</style>