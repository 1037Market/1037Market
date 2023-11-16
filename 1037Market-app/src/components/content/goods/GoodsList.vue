<template>
    <div class="goods" ref="goodsContainer" >
        <!--        <button @click="debug">debug</button>-->
        <div class="goods-item" 
             v-for="productDetail in filteredProducts"
             :key="productDetail.productId"
             :style="productDetail.style"
             :data-key="productDetail.productId"
             ref="productsRef"
        >
        <div @click="itemClick(productDetail.productId)">
            <img
                :src="Array.isArray(productDetail.imageURIs) && productDetail.imageURIs.length > 0 ? 'https://franky.pro:7301/api/image?imageURI=' + productDetail.imageURIs[0] : ''"
                @load="calculatePosition"
                alt="商品"
                style="border-radius: 10%"
                

            />
            <div class="goods-info">
                <p style="font-size: 20px;font-weight: 600;margin: 5px auto;">{{ productDetail.title }}</p>
                <div style="display: inline;margin-left: 0px;">￥</div><span class="price" style="font-size: 25px;font-weight: 500;margin: 0px auto;">{{ productDetail.price }}</span>
                
            </div>
        </div>
        <div @click="clickAvatar(productDetail.publisher)">
            <div>
              <van-image style="display: inline-block; vertical-align: middle"
                      width="20px"
                      height="20px"
                      radius="10px"
                      fit="cover"
                      :src="'https://franky.pro:7301/api/image?imageURI=' + productDetail.avatar"
                      
              />
              <span style="vertical-align: middle; font-weight: 400; font-size: 15px; margin-left: 10px">{{ productDetail.nickName }}</span>

            </div>
        </div>
        </div>
    </div>
</template>

<script setup>
import {getDetail} from "@/network/detail";
import {nextTick, ref, watch, watchEffect, computed} from "vue";
import {useRouter} from "vue-router";
import {debounce} from "lodash";

const router = useRouter()
const debug = () => {
    console.log(productIDs.value)
}

const props = defineProps({
    showGoods: Array,
    showPositive: Number
})

const productIDs = ref(props.showGoods)

const products = ref([])
const filteredProducts = computed(() => {
    return products.value.filter((item) => {
        return (item.price > 0) === (props.showPositive === 0)
    })
})
let renderIDs = new Set()

const goodsContainer = ref(null)

watchEffect(() => {
    productIDs.value = [...props.showGoods]
    // debouncedHandle(productIDs.value)
})

function shuffleArray(array) {
    for (let i = array.length - 1; i > 0; i--) {
        // 生成一个随机索引
        const j = Math.floor(Math.random() * (i + 1));

        // 交换元素 array[i] 和 array[j]
        [array[i], array[j]] = [array[j], array[i]];
    }
    return array;
}

const debouncedHandleRefresh = debounce((newIDs) => {
    products.value = products.value.filter(item => newIDs.includes(item.productId))
    products.value = shuffleArray(products.value)
    calculatePosition()
    renderIDs = new Set(newIDs.filter(item => renderIDs.has(item)))
    newIDs.forEach(function (productId) {
        if(renderIDs.has(productId) === false){
            renderIDs.add(productId)
            getDetail(productId).then((resp) => {
                products.value.unshift(resp)
            }).catch((error) => {
                renderIDs.delete(productId)
                console.log("Load error:",error)
            })
        }
    })
}, 100)

const debouncedHandleLoadMore = debounce((newIDs) => {
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
    if(products.value.length !== 0 && newIDs[0] === products.value[0].productId){
        debouncedHandleLoadMore(newIDs)
        // console.log('load more')
    }
    else{
        debouncedHandleRefresh(newIDs)
        // console.log('refresh')
    }
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

const clickAvatar = (studentId) => {
    router.push({path: `/seller/${studentId}`});
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
    word-wrap: anywhere;
    word-break: keep-all;
}

.goods-item .goods-info p {
    
    word-wrap: anywhere;
    word-break: break-all;
    margin-bottom: 3px;
}

.goods-item .goods-info .price {
    color: #FF4C0A;
}
</style>