<template>
    <div>
        <van-nav-bar title="商品详情" fixed placeholder
                     left-arrow @click-left="router.go(-1)"
        />

        <van-swipe :autoplay="3000" indicator-color="#44b883" style="height: 300px; margin-top: 1px;" lazy-render>
            <van-swipe-item v-for="uri in productDetail.imageURIs" :key="uri">
                <van-image :src="'http://franky.pro:7301/api/image?imageURI=' + uri" fit="contain"/>
            </van-swipe-item>
        </van-swipe>


        <div>

            <price-display v-if="typeof (productDetail.title) !== 'undefined' " :name="productDetail.title"
                           :price="productDetail.price" :categories="productDetail.categories"/>
            <product-description v-if="typeof (productDetail.content) != 'undefined'"
                                 :description="productDetail.content"
                                 style="line-height: 25px;font-family: Arial,sans-serif;font-weight: 600;letter-spacing:0.03em;"/>

        </div>

        <van-action-bar>
            <van-icon name="manager-o" style="margin-left: 20px"/>
            <van-action-bar-icon text="卖家信息" @click="userInfo"/>
            <van-action-bar-button type="warning" text="联系卖家" style="margin-left: 20px"/>
            <van-action-bar-button type="danger" text="收藏商品" style="margin-right: 20px" @click="handleCart"/>
        </van-action-bar>

    </div>
</template>

<script>
import GoodsList from "@/components/content/goods/GoodsList.vue";

import {onBeforeRouteLeave, useRoute} from "vue-router";
import {useRouter} from "vue-router";
import {useStore} from "vuex";

import {ref, onMounted, reactive, toRefs} from "vue";
import {getDetail} from "@/network/detail";
import {addCart} from "@/network/cart";

import PriceDisplay from "@/components/content/goods/PriceDisplay.vue"
import ProductDescription from "@/components/content/goods/ProductDescription.vue";
import {showSuccessToast, showFailToast} from 'vant';

export default {
    name: "Detail",
    components: {
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
            imageURIs: [],
            categories: []
        })
        let active = ref(1);

        const handleCart = () => {
            addCart(id.value).then((res) => {
                console.log(res)
                if (res === 'ok')
                    showSuccessToast('收藏成功')
                else showFailToast('收藏失败')
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


        onMounted(() => {
            getDetail(id.value).then((res) => {
                productDetail.value = res;
            }).catch((err) => {
                console.log(err);
            });
        });

        const userInfo = () => {
            console.log(productDetail.value.publisher);
            router.push({path: `/seller/${productDetail.value.publisher}`});
        }

        return {
            id,
            active,
            handleCart,
            goToCart,
            productDetail,
            userInfo
        };
    },
};
</script>

<style scoped lang="scss">
#con1 {
    padding: 10px;
}
</style>

<style scoped>
.display {
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
