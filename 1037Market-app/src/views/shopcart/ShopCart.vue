<template>
  <div>
    <nav-bar class="nav-bar">
      <template v-slot:center>
        购物车
      </template>
    </nav-bar>
    <div class="cart-box">
        
        <div class="cardList">
            <van-card
                v-for="product in products"
                :key="product.productId"
                :price="product.price"
                :title="product.title"
                :thumb="'http://franky.pro:7301/api/image?imageURI=' + product.imageURIs[0]"
            >
                <template #tags>
                    <van-tag v-for="tag in product.categories" plain type="danger">{{tag}}</van-tag>
                </template>
                <template #footer>
                    <van-button type="danger" size="mini" @click="deleteFavorites(product.productId)">删除收藏</van-button>
                </template>
            </van-card>
        </div>
        <div class="empty" v-if="!products.length">
        <img
          class="empty-cart"
          src="@/assets/images/empty-car.png"
          alt="空购物车"
        />
        <div class="title">购物车空空如也</div>
        <van-button round color="#1baeae" type="primary" block @click="goTo"
          >前往选购</van-button
        >
      </div>
    </div>
  </div>
</template>

<script>
import {ref, reactive, toRefs, onMounted, computed, nextTick, watch} from "vue";
import { useRouter } from "vue-router";
import { useStore } from "vuex";
import { getCart, deleteCartItem} from "@/network/cart";
import { getDetail } from "@/network/detail"
import { showLoadingToast, closeToast } from 'vant';
import NavBar from "@/components/common/navbar/NavBar.vue";
export default {
  name: "ShopCart",
  setup() {
    const router = useRouter();
    const productIDs = ref([])
    const products = ref([])

    onMounted(() => {
        showLoadingToast({
            message: '加载中...',
            forbidClick: true,
        });

        getCart().then((res) => {
            productIDs.value = res;
            closeToast(true)
        });
    });

    watch(productIDs, (newIDs) => {
        for(const productId of newIDs){
            getDetail(productId).then((resp) => {
                products.value.push(resp)
                // console.log(products)
            })
        }
    })

    // 前往购物
    const goTo = () => {
      router.push({ path: "/" });
    };

    const navigateToProduct = (productId) => {
        router.push({path: `/detail/${productId}`})
    }

    const debug = () => {
        console.log(products)
    }

    // 删除商品
    const deleteFavorites = (id) => {
      deleteCartItem(id).then((res) => {
          Toast.success("删除成功")
          getCart().then((resp) => {
              products.value.length = 0
              productIDs.value = resp;
          });
      });
    };

    const cardTheme = {
        cardTitleLineHeight: '20px'
    }

    return {
        goTo,
        deleteFavorites,
        debug,
        products,
        navigateToProduct,
        cardTheme
    };
  },

  components: {
    NavBar,
  },
};
</script>

<style scoped lang="scss">

.cardList{
    margin-top: 55px;
}

.delete-button {
    margin-left: auto; /* Push the button to the right */
}
</style>