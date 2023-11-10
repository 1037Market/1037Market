<template>
    <div style="background-color: rgba(255, 255, 255, 1);">
    <div id="home">

        <van-nav-bar title="1037集市" fixed placeholder style="--van-nav-bar-background: linear-gradient(rgba(66, 185, 131, 0.9),rgba(66,185,131,0.45));--van-nav-bar-title-text-color: rgba(255,255,255,1);"
                     left-arrow @click-left="router.go(-1)"
        />


        <form action="/" style="--van-search-background: linear-gradient(rgba(66, 185, 131, 0.45),rgba(66,185,131,0));margin-top: -1px;">
            <van-search style="--van-search-content-background: linear-gradient(rgba(66, 185, 131, 0),rgba(66,185,131,0));"
                v-model="searchInfo"
                placeholder="请输入搜索关键词"
                @search="onSearch"
                @cancel="onCancel"
            />

        </form>

        <div>
            <van-tabs v-model:active="currentType" scrollspy>
                <van-tab title="搜索结果" name="搜索结果" v-if="searched"/>
                <van-tab title="推荐" name="推荐"/>
                <van-tab v-for="category in categories" :title="category" :name="category" :key="category"/>
            </van-tabs>
        </div>

        <div class="wrapper">
            <van-pull-refresh
                v-model="pullingDown"
                success-text="加载成功"
                @refresh="pullingDownHandler"
            >
                <van-list
                    v-model="pullingUp"
                    :finished="loadFinished"
                    finished-text="没有更多了"
                    @load="pullingUpHandler"
                >
<!--                    <p style="line-height: 30px; margin-top: -30px;text-align: center;color: #646566">{{ pullingDownHint }}</p>-->
                    <p v-if="currentType==='搜索结果' && searchFail">没有找到相关商品</p>
                    <goods-list :showGoods="showGoods" v-if="typeof (showGoods) !== 'undefined'"></goods-list>
                </van-list>
            </van-pull-refresh>

        </div>

        <back-top @goback="goback" v-show="isShowBackTop"></back-top>
    </div>
    </div>
</template>

<script setup>
import {onMounted, ref, computed, watchEffect, nextTick, watch} from "vue";
import {useRouter} from 'vue-router'
import {getHomeGoodsData, getSearchData, refresh} from "@/network/home";
import { debounce } from "lodash";
import {getCategoryData} from "@/network/category";
import GoodsList from "@/components/content/goods/GoodsList.vue";
import BackTop from "@/components/common/backtop/BackTop.vue";

const router = useRouter()

const searchInfo = ref('')
const searchFail = computed(() => {
    return goods['搜索结果'].value.length === 0
})

const isShowBackTop = ref(false);

//商品列表对象模型,里面三个选项卡的页码和列表
const goods = {
    '推荐': ref([]),
    '搜索结果': ref([])
};

const currentType = ref("推荐");
const showGoods = computed(() => {
    if(typeof (goods[currentType.value].value) === "undefined")
        goods[currentType.value] = ref([])
    return goods[currentType.value].value
})

watch(currentType, (newValue, oldValue) => {
    loadFinished.value = true
    if (newValue === "more") {
        nextTick(() => {
            currentType.value = oldValue
        })
        router.push('/category')
    } else if (newValue === "搜索结果") {
        // if(goods.search.length === 0)
        //   searchFail.value = true
    } else if (typeof (goods[newValue]) === "undefined" || goods[newValue].value.length === 0)
        getShowGoods()
})

const getShowGoods = async () => {
    if(typeof (goods[currentType.value]) === "undefined")
        goods[currentType.value] = ref([])
    getHomeGoodsData(currentType.value).then((res) => {
        goods[currentType.value].value.push(...res);
    });
}

const getCategories = async () => {
    getCategoryData().then((res) => {
        categories.value.length = 0
        categories.value = res
    })
}

const searched = ref(false)

const onSearch = () => {
    // console.log("on search")
    getSearchData(searchInfo.value).then((res) => {
        searched.value = true
        goods['搜索结果'].value.length = 0
        goods['搜索结果'].value.push(...res);
        nextTick(() => {
            currentType.value = '搜索结果'
            console.log('search success', currentType.value)
        })

    });
}

const onCancel = () => {
    console.log("on cancel")
    searchInfo.value = ''
}

//回到顶部
const goback = () => {
    // TODO: 实现滑至页面顶端
};

const pullingDown = ref(false)
const pullingUp = ref(false)
const loadFinished = ref(false)

const debouncedPullingDownHandler = debounce(async () => {
    refresh()
    loadFinished.value = false
    getHomeGoodsData(currentType.value).then((res) => {
        goods[currentType.value].value = []
        goods[currentType.value].value.push(...res);
        pullingDown.value = false
    }).catch((error) => {
        console.log('refresh fail',error)
        pullingDown.value = false
    })
}, 300)
const pullingDownHandler = () => {
    debouncedPullingDownHandler()
};

const debouncedPullingUpHandler = debounce(async () => {
    if (pullingUp.value === true)
        return
    getHomeGoodsData(currentType.value, goods[currentType.value].value.length).then((res) => {
        goods[currentType.value].value.push(...res);
        if(res.length === 0)
            loadFinished.value = true
        pullingUp.value = false
    }).catch((error) => {
        console.log('get more fail',error)
        pullingUp.value = false
    })
},300)
const pullingUpHandler = () => {
    debouncedPullingUpHandler()
};
const pullingDownHint = ref('继续下拉刷新页面')

const categories = ref([])
onMounted(() => {
    currentType.value = "推荐"
    getCategories()
    getShowGoods()
});
</script>

<style scoped>
#home {
    text-align: left;
}

.banners img {
    width: 100%;
    height: auto;
}

#home {
    height: 100vh;
    position: relative;
}

.wrapper {
    position: relative;
    top: 0;
    right: 0;
    left: 0;
    overflow: auto;
    height: calc(100vh - 200px);
}
</style>