<template>
  <router-view />
  <div id="nav" style="z-index: 20">
      <router-link to="/" class="tab-bar-item">
          <div class="icon">
              <i class="iconfont icon-shouye"></i>
          </div>
          <div>首页</div>
      </router-link>

      <router-link to="/chat/0" class="tab-bar-item">
          <div class="icon">
              <van-icon name="chat-o"/>
          </div>
          <div>聊天</div>
      </router-link>

      <router-link to="/publish" class="tab-bar-item">
          <div class="icon">
              <i class="iconfont icon-fabu"></i>
          </div>
          <div>发布</div>
      </router-link>

      <router-link to="/shopcart" class="tab-bar-item">
          <div class="icon">
              <i class="iconfont icon-gouwuche"></i>
          </div>
          <div>购物车</div>
      </router-link>

      <router-link to="/profile" class="tab-bar-item">
          <div class="icon">
              <i class="iconfont icon-yonghu-xianxing"></i>
          </div>
          <div>我的</div>
      </router-link>
  </div>
</template>

<script>
import { useStore } from "vuex";
export default {
  setup() {
    const store = useStore();

    return {};
  },
};
</script>

<style lang="scss">
@import "assets/css/base.css";
@import "assets/css/iconfont.css";
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  //text-align: center;
  color: #2c3e50;
}

#nav {
  background-color: #f6f6f6;
  display: flex;
  position: fixed;
  left: 0;
  right: 0;
  bottom: 0;
  box-shadow: 0 -3px 1px rgba(100, 100, 100, 0.1);

  a {
    color: var(--color-text);

    &.router-link-exact-active {
      color: var(--color-high-text);
    }
  }

  .tab-bar-item {
    flex: 1;
    text-align: center;
    height: 50px;
    font-size: var(--font-size);
  }
  .tab-bar-item .icon {
    width: 24px;
    height: 24px;
    margin-top: 5px;
    vertical-align: middle;
    display: inline-block;
  }
}
</style>
