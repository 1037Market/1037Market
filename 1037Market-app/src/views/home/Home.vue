<template>
    <div id="home">

        <van-nav-bar
            title="1037集市"
            fixed
            placeholder
        />


        <form action="/">
            <van-search
                v-model="searchInfo"
                placeholder="请输入搜索关键词"
                @search="onSearch"
            />

        </form>

        <div>
            <van-tabs v-model:active="currentType">
                <van-tab title="搜索结果" name="search" v-if="searched"/>
                <van-tab title="推荐" name="recommend"/>
                <van-tab title="二手书" name="books"/>
                <van-tab title="闲置物品" name="items"/>
                <van-tab title="更多分类" name="more"/>
            </van-tabs>
        </div>

        <div class="wrapper" @touchend="handleTouchEnd">
            <div>
                <p style="line-height: 30px; margin-top: -30px;text-align: center;color: #646566">{{pullingDownHint}}</p>
                <p v-if="currentType==='search' && searchFail">没有找到相关商品</p>
                <goods-list :showGoods="showGoods" v-if="typeof (showGoods) !== 'undefined'"></goods-list>
            </div>
        </div>

        <back-top @goback="goback" v-show="isShowBackTop"></back-top>
    </div>
</template>

<script>
import {onMounted, ref, reactive, computed, watchEffect, nextTick, watch} from "vue";
import {useRouter} from 'vue-router'
import {getHomeAllData, getHomeGoodsData, getSearchData, refresh} from "@/network/home";
import GoodsList from "@/components/content/goods/GoodsList.vue";
import BackTop from "@/components/common/backtop/BackTop.vue";
import BetterScroll from "@better-scroll/core";
import PullUp from '@better-scroll/pull-up'
import PullDown from "@better-scroll/pull-down";

BetterScroll.use(PullUp).use(PullDown)

export default {
    name: "Home",
    components: {
        GoodsList,
        BackTop,
    },
    setup() {
        const recommends = ref([]);
        const router = useRouter()

        const searchInfo = ref('')
        const searchFail = computed(() => {
            return goods['search'].value.length === 0
        })

        //复制TabControl
        const isTabFixed = ref(false);

        const isShowBackTop = ref(false);

        //商品列表对象模型,里面三个选项卡的页码和列表
        const goods = {
            recommend: ref([]),
            books: ref([]),
            items: ref([]),
            search: ref([])
        };

        const currentType = ref("recommend");
        const showGoods = computed(() => {
            return goods[currentType.value].value
        })

        watch(currentType, (newValue, oldValue) => {
            if (newValue === "more") {
                nextTick(() => {
                    currentType.value = oldValue
                })
                router.push('/category')
            } else if (newValue === "search") {
                // if(goods.search.length === 0)
                //   searchFail.value = true
            } else if (goods[newValue].value.length === 0)
                getShowGoods()
        })

        const getShowGoods = () => {
            getHomeGoodsData(currentType.value).then((res) => {
                goods[currentType.value].value.push(...res);
            });
        }

        let bscroll = reactive({});

        const searched = ref(false)

        const onSearch = () => {
            // console.log("on search")
            getSearchData(searchInfo.value).then((res) => {
                searched.value = true
                goods.search.value.length = 0
                goods.search.value.push(...res);
                nextTick(() => {
                    currentType.value = 'search'
                    console.log('search success', currentType.value)
                })

            });
        }

        const onCancel = () => {
            console.log("on cancel")
            searchInfo.value = ''
        }

        //监听，任何一个变量有变化
        watchEffect(() => {
            //nextTick：Dom加载完成后
            nextTick(() => {
                //只要页面有变化就要调用refresh
                bscroll && bscroll.refresh();
            });
        });

        //回到顶部
        const goback = () => {
            bscroll.scrollTo(0, 0);
        };

        const pullingDown = ref(false)
        const pullingUp = ref(false)

        function sleep(ms) {
            return new Promise(resolve => setTimeout(resolve, ms));
        }

        const pullingDownHandler = () => {
            console.log('pull down')
            pullingDownHint.value = '加载中'
            refresh()
            getHomeGoodsData(currentType.value).then((res) => {
                goods[currentType.value].value = []
                goods[currentType.value].value.push(...res);
                // console.log('pulling down')
                bscroll.finishPullDown();
                bscroll.refresh();
                pullingDown.value = false;
            }).catch((error) => {
                console.log('refresh fail')
                bscroll.finishPullDown();
                bscroll.refresh();
                pullingDown.value = false;
            })
        };

        const handleTouchEnd = () => {
            // console.log('touch end')
            const BSRefreshTimer = setTimeout(() => {
                bscroll.finishPullUp();
                bscroll.finishPullDown();
                bscroll.refresh();
                pullingUp.value = false;
                clearTimeout(BSRefreshTimer)
            }, 5)
        };

        const pullingUpHandler = () => {
            if (pullingUp.value === true)
                return
            console.log('pulling up')
            pullingUp.value = true;
            getHomeGoodsData(currentType.value, goods[currentType.value].value.length).then((res) => {
                goods[currentType.value].value.push(...res);
            }).catch((error) => {
                console.log('get more fail')
            })
        };
        const pullingDownHint = ref('继续下拉刷新页面')

        onMounted(() => {
            currentType.value = 'recommend'
            getShowGoods()

            // 创建BS对象
            bscroll = new BetterScroll(document.querySelector(".wrapper"), {
                probeType: 3, // 0, 1, 2, 3, 3 只要在运动就触发scroll事件
                click: true, // 是否允许点击
                pullUpLoad: {
                    threshold: -20
                },
                pullDownRefresh: {
                    threshold: 50,
                    stop: 20
                }
            });

            bscroll.on('pullingDown', pullingDownHandler);
            bscroll.on('pullingUp', pullingUpHandler);
            bscroll.on('scroll', (position) => {
                if(pullingDownHint.value === '加载中'){
                    if(position.y <= 0)
                        pullingDownHint.value = '继续下拉刷新页面'
                }else if(position.y > 30)
                    pullingDownHint.value ='松手刷新页面'
                else pullingDownHint.value = '继续下拉刷新页面'
            })
        });

        const debug = () => {
            console.log("type", currentType.value)
        }

        return {
            recommends,
            goods,
            currentType,
            showGoods,
            bscroll,
            isTabFixed,
            goback,
            isShowBackTop,
            searchInfo,
            onSearch,
            onCancel,
            debug,
            searched,
            searchFail,
            pullingDown,
            handleTouchEnd,
            pullingDownHint
        };
    },
};
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
    overflow: hidden;
    height: calc(100vh - 200px);
}
</style>