<template>
    <div class="goods" ref="goodsContainer">
        <!--        <button @click="debug">debug</button>-->
        <div class="goods-item" 
             v-for="productDetail in products"
             :key="productDetail.productId"
             @click="itemClick(productDetail.productId)"
             :style="productDetail.style"
             :data-key="productDetail.productId"
             ref="productsRef"
        >
            <img
                :src="Array.isArray(productDetail.imageURIs) && productDetail.imageURIs.length > 0 ? 'http://franky.pro:7301/api/image?imageURI=' + productDetail.imageURIs[0] : ''"
                @load="calculatePosition"
                alt="商品"
            />
            <div class="goods-info">
                <p style="font-size: 14px;font-weight: 600;margin: 5px auto;">{{ productDetail.title }}</p>
                <div style="display: inline;margin-left: -12px;">￥</div><span class="price" style="font-size: 16px;font-weight: 600;margin: 0px auto;">{{ productDetail.price }}</span>
            </div>
        </div>
    </div>
</template>

<script setup>
import {getDetail} from "@/network/detail";
import {nextTick, ref, watch, watchEffect} from "vue";
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

const goodsContainer = ref(null)

watchEffect(() => {
    productIDs.value = [...props.showGoods]
    // debouncedHandle(productIDs.value)
})

const debouncedHandle = debounce((newIDs) => {
    products.value = products.value.filter(item => newIDs.includes(item.productId))
    calculatePosition()
    renderIDs = new Set(newIDs.filter(item => renderIDs.has(item)))
    newIDs.forEach(function (productId) {
        if(renderIDs.has(productId) === false){
            renderIDs.add(productId)
            getDetail(productId).then((resp) => {
                products.value.push(resp)
            }).catch((error) => {
                renderIDs.delete(productId)
                console.log("Load error:",error)
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

let calculating = false
let columnHeights = [0,0]
const productsRef = ref([])
function getProductHeight(productId){
    for (const productRef of productsRef.value) {
        if(productId == productRef.dataset.key){
            return productRef.clientHeight;
        }
    }
    return 0;
}

const calculatePosition = () => {
    // console.log('start cal')
    if(calculating) return
    calculating = true
    columnHeights[0] = 0
    columnHeights[1] = 0
    products.value.forEach((product) => {
        const column = columnHeights[0] <= columnHeights[1] ? 0 : 1
        const height = getProductHeight(product.productId)
        // console.log(`${product.title} at ${column ? 'left' : 'right'}`)
        product.style = ref({
            position: 'absolute',
            left: `${column*50}%`,
            top: `${columnHeights[column]}px`,
            width: '50%',
            height: 'auto'
        })
        columnHeights[column] += height
        goodsContainer.value.style.height = `${Math.max(...columnHeights)+30}px`
    })
    calculating = false
}

</script>

<style scoped>
.goods {
    flex-wrap: wrap;
    justify-content: space-around;
    padding: 5px;
    position: relative;
    //columns: 2;
    //column-gap: 16px;
    //display: grid;
    //grid-template-columns: repeat(2, 1fr);
    width: 100%;
}

.goods-item {
    position: absolute;
    left: 0;
    top: 0;
    width: 50%;
    height: auto;
    padding: 10px;
}

.goods-item van-image {
    width: 100%;
    height: auto;
    border-radius: 5px;
}

.goods-item .goods-info {
    font-size: 12px;
    position: relative;
    margin-top: 5px;
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
.display {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 15px;
    margin: 10px; /* Added horizontal margin */
    color: #333;
    border-radius: 15px;
    box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
    font-family: 'Poppins', sans-serif;
}
.goods-item .goods-info .price {
    color: #FF4C0A;
}
</style>