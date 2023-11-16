<template>
    <div>
        <van-nav-bar title="购物车" fixed placeholder style="--van-nav-bar-background: rgba(66, 185, 131, 0.9);--van-nav-bar-title-text-color: rgba(255,255,255,1);"
                     left-arrow @click-left="router.go(-1)"
        />
        <div class="cart-box">

            <div class="cardList">
                <van-swipe-cell v-for="product in products" class="display">
                    <van-config-provider :theme-vars="cardTheme">
                        <van-card style="--van-card-price-color: #FF4C0A;width: 100%;--van-card-price-integer-font-size: 24px;--van-card-price-font-size: 16px;"
                            :key="product.productId"
                            :price="product.price"
                            :title="product.title"
                            :thumb="'https://franky.pro:7301/api/image?imageURI=' + product.imageURIs[0]"
                            @click="navigateToProduct(product.productId)"
                        >
                            <template #tags>
                                <van-tag type="primary" size="medium"
                                    v-for="tag in product.categories"
                                    :key="category"
                                    style="margin: 10px 3px"
                                >{{ tag }}
                                </van-tag>
                            </template>
                        </van-card>
                    </van-config-provider>

                    <template #right>
                        <van-config-provider :theme-vars="buttonTheme" style="margin-left: 15px;">
                            <van-button 
                                        type="danger"
                                        size="small"
                                        @click="deleteFavorites(product.productId)"
                            >删除
                            </van-button>
                        </van-config-provider>
                    </template>
                </van-swipe-cell>
            </div>
            <div class="empty" v-if="!products.length">
                <img
                    class="empty-cart"
                    src="@/assets/images/empty-car.png"
                    alt="空购物车"
                />
                <div class="title" style="text-align: center;padding: 15px">购物车空空如也</div>
                <van-button round color="#1baeae" type="primary" block @click="goTo"
                >前往选购
                </van-button
                >
            </div>
        </div>
    </div>
</template>

<script>
import {ref, reactive, toRefs, onMounted, computed, nextTick, watch} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {getCart, deleteCartItem} from "@/network/cart";
import {getDetail} from "@/network/detail"
import {showSuccessToast, showFailToast, showLoadingToast, closeToast} from 'vant';

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

        watch(productIDs, (newIDs, oldIDs) => {
            products.value = products.value.filter(product => {
                return newIDs.includes(product.productId)
            })
            for (const productId of newIDs) {
                if (!oldIDs.includes(productId)) {
                    getDetail(productId).then((resp) => {
                        products.value.push(resp)
                        // console.log(products)
                    })
                }
            }
        })

        // 前往购物
        const goTo = () => {
            router.push({path: "/"});
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
                showSuccessToast("删除成功")
                getCart().then((resp) => {
                    productIDs.value = resp;
                });
            });
        };

        const buttonTheme = {
            buttonSmallFontSize: '16px',
            buttonSmallHeight: '104px'
        }

        const cardTheme = {
            cardFontSize: '16px',
            cardTitleLineHeight: '20px',
            tagPadding: '3px',
            cardBackground: 'white'
        }

        return {
            goTo,
            deleteFavorites,
            debug,
            products,
            navigateToProduct,
            buttonTheme,
            cardTheme,
            router
        };
    },

    components: {},
};
</script>

<style scoped lang="scss">
.display {
    justify-content: space-between;
    align-items: center;
    padding: 15px;
    margin: 10px; /* Added horizontal margin */
    color: #333;
    border-radius: 15px;
    box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
    font-family: 'Poppins', sans-serif;
    
}
.cardList {
    margin-top: 10px;
}

.delete-button {
    height: 100%;
}
</style>