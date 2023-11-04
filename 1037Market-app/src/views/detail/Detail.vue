<template>
  <div>
    <nav-bar>
      <template v-slot:center>商品详情:{{ id }}</template>
    </nav-bar>

    <van-swipe :autoplay="3000" indicator-color="#44b883" style="height: 300px; margin-top: 45px;" lazy-render>
      <van-swipe-item v-for="uri in productDetail.imageURIs" :key="uri">
        <van-image :src="'http://franky.pro:7301/api/image?imageURI=' + uri" fit="contain"/>
      </van-swipe-item>
    </van-swipe>

    <div>
      <price-display :name="productDetail.title" :price="productDetail.price" />
      <product-description :description="productDetail.content" />
    </div>

    <van-action-bar>
      <van-action-bar-icon icon="chat-o" text="联系卖家"/>
      <van-action-bar-icon icon="friends-o" text="卖家信息"/>
      <van-action-bar-button type="danger" text="立即购买"/>
    </van-action-bar>

  </div>
</template>

<script>
import NavBar from "components/common/navbar/NavBar";
import GoodsList from "components/content/goods/GoodsList";

import {onBeforeRouteLeave, useRoute} from "vue-router";
import { useRouter } from "vue-router";
import { useStore } from "vuex";

import { ref, onMounted, reactive, toRefs } from "vue";
import { getDetail } from "network/detail";
import { Toast } from "vant";
import { addCart } from "network/cart";

import PriceDisplay from "components/content/goods/PriceDisplay"
import ProductDescription from "components/content/goods/ProductDescription";
export default {
  name: "Detail",
  components: {
    NavBar,
    GoodsList,
    PriceDisplay,
    ProductDescription
  },
  setup() {
    const route = useRoute();
    const router = useRouter();
    const store = useStore();

    let id = ref(route.params.id);
    const productDetail = ref({
      imageURIs: []
    })
    let active = ref(1);

    const handleCart = () => {
      addCart(id.value).then((res) => {
        if (res.status == "204" || res.status == "201") {
          Toast.success("添加成功");
      //     调用actions
      //     store.dispatch("updateCart");
        }
      });
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

    onBeforeRouteLeave((to, from, next) => { // 离开前将底部tabbar恢复
      const nav = document.getElementById('nav');
      if(nav) {
        nav.style.visibility = 'visible';
      }
      next();
    })

    onMounted(() => {
      // 进入时隐藏底部tabbar
      const nav = document.getElementById('nav');
      if(nav) {
        nav.style.visibility = 'hidden';
      }

      getDetail(id.value).then((res) => {
          productDetail.value = res;
          console.log(res);
      }).catch((err) => {
        console.log(err);
      });

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