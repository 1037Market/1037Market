<template>
  <div id="home">
    <nav-bar>
      <!-- vue3插槽写法 -->
      <template v-slot:center>1037集市</template>
    </nav-bar>

    <!-- 复制一份tab-control选择性显示 -->
    <tab-control
      :titles="['推荐', '二手书', '闲置物品']"
      @tabClick="tabClick"
    ></tab-control>

    <!-- 使用better-scroll -->
    <div class="wrapper">
      <div class="content">
        <div ref="banref">
<!--          <home-swiper :banners="banners"></home-swiper>-->

<!--          <recommend-view :recommends="recommends"></recommend-view>-->
        </div>

<!--        <tab-control-->
<!--          :titles="['推荐', '二手书', '闲置物品']"-->
<!--          @tabClick="tabClick"-->
<!--        ></tab-control>-->

        <!-- 因为是切换选项卡，所以只显示一个，只传一个类型的数据，需要知道当前是哪个选项卡，使用计算属性 -->
        <goods-list :goods="showGoods"></goods-list>
      </div>
    </div>
    <back-top @goback="goback" v-show="isShowBackTop"></back-top>
  </div>
</template>

<script>
import { onMounted, ref, reactive, computed, watchEffect, nextTick } from "vue";
import { getHomeAllData, getHomeGoodsData } from "network/home";
import HomeSwiper from "views/home/childComps/HomeSwiper";
import NavBar from "components/common/navbar/NavBar";
import recommendView from "./childComps/RecommendView";
import TabControl from "components/content/tabcontrol/TabControl";
import GoodsList from "components/content/goods/GoodsList";
import BackTop from "components/common/backtop/BackTop";

import BScroll from "better-scroll";

export default {
  name: "Home",
  components: {
    NavBar,
    recommendView,
    TabControl,
    GoodsList,
    BackTop,
    HomeSwiper,
  },
  setup() {
    const recommends = ref([]);

    //复制TabControl
    const isTabFixed = ref(false);

    const isShowBackTop = ref(false);

    let banref = ref(null);

    //商品列表对象模型,里面三个选项卡的页码和列表
    const goods = reactive({
      recommend: [],
      books: [],
      items: []
    });

    const currentType = ref("recommend");
    const showGoods = computed(() => {
      return goods[currentType.value];
    });

    let bscroll = reactive({});

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

    onMounted(() => {
      //分别传参获取三个选项卡的数据

        getHomeGoodsData("recommend").then((res) => {
            goods.recommend.push(...res);
        });

        getHomeGoodsData("books").then((res) => {
            goods.books.push(...res);
        });

      getHomeGoodsData("items").then((res) => {
        goods.items.push(...res);
      });

      // 创建BS对象
      bscroll = new BScroll(document.querySelector(".wrapper"), {
        probeType: 3, // 0, 1, 2, 3, 3 只要在运动就触发scroll事件
        click: true, // 是否允许点击
        pullUpLoad: true, //上拉加载更多， 默认是false
      });

      // 触发滚动事件
      bscroll.on("scroll", (position) => {
        // console.log(-position.y);
      //   悬停条件(返回顶部和Tabcontrol)
        isShowBackTop.value = isTabFixed.value =
          -position.y > banref.value.offsetHeight;
      });

      //上拉加载数据,触发pullingUp
      bscroll.on("pullingUp", () => {
        // console.log("上拉加载更多");
        bscroll.refresh();
        // const page = goods[currentType.value].page + 1;

        // 获取下拉数据
        // getHomeGoodsData(currentType.value, page).then((res) => {
        //   goods[currentType.value].push(...res.goods.data);
        //   goods[currentType.value].page += 1;
        // });

        // 完成上拉， 等数据请求完成， 要将新数据展示出来
        bscroll.finishPullUp();

        //重新计算高度
        bscroll.refresh();
      });
    });

    //子组件TabControl传过来的
    const tabClick = (index) => {
      let types = ["sales", "new", "recommend"];
      //选项卡切换后下面内容才会变
      currentType.value = types[index];
      // console.log(currentType.value);

      //监听，任何一个变量有变化
      watchEffect(() => {
        //nextTick：Dom加载完成后
        nextTick(() => {
          //只要页面有变化就要调用refresh
          bscroll && bscroll.refresh();
        });
      });
    };

    return {
      recommends,
      tabClick,
      goods,
      currentType,
      showGoods,
      bscroll,
      isTabFixed,
      banref,
      goback,
      isShowBackTop,
    };
  },
};
</script>

<style>
.banners img {
  width: 100%;
  height: auto;
}

#home {
  height: 100vh;
  position: relative;
}

.wrapper {
  position: absolute;
  top: 45px;
  bottom: 50px;
  right: 0;
  left: 0;
  overflow: hidden;
}
</style>