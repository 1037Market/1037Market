<template>
    <div class="goods">
        <!--        <button @click="debug">debug</button>-->
        <div class="goods-item" v-for="productDetail in products" :key="productDetail.productId" @click="itemClick(productDetail.productId)">
            <van-image
                :src="Array.isArray(productDetail.imageURIs) && productDetail.imageURIs.length > 0 ? 'http://franky.pro:7301/api/image?imageURI=' + productDetail.imageURIs[0] : ''"/>
            <div class="goods-info">
                <p style="font-size: 14px"><strong>{{ productDetail.title }}</strong></p>
                <span class="price" style="font-size: 13px"><small>￥</small>{{ productDetail.price }}</span>
            </div>
        </div>
    </div>
</template>

<script setup>
import {getDetail} from "@/network/detail";
import {onMounted, ref, watch, watchEffect} from "vue";
import {useRouter} from "vue-router";
import {debounce} from "lodash";

const router = useRouter()
const debug = () => {
    console.log(productIDs.value)
}

const props = defineProps({
    showGoods: Array
})

const productIDs = ref(props.showGoods)

const products = ref([])
let renderIDs = new Set()

watchEffect(() => {
    productIDs.value = [...props.showGoods]
    // debouncedHandle(productIDs.value)
})

const debouncedHandle = debounce((newIDs) => {
    products.value = products.value.filter(item => newIDs.includes(item.productId))
    renderIDs = new Set(newIDs.filter(item => renderIDs.has(item)))
    newIDs.forEach(function (productId, index) {
        if(renderIDs.has(productId) === false){
            renderIDs.add(productId)
            getDetail(productId).then((resp) => {
                products.value.push(resp)
            }).catch((error) => {
                renderIDs.delete(productId)
                console.log("Load error")
            })
        }
    })
}, 100)

watch(productIDs, (newIDs) => {
    debouncedHandle(newIDs)
})

const itemClick = (productId) => {
    router.push({path: `/detail/${productId}`});
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

.goods-item van-image {
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