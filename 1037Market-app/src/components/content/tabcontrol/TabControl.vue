<template>
  <div class="tab-control">
    <!-- 点击要传入index -->
    <div
      class="tab-control-item"
      v-for="(item, index) in titles"
      :key="index"
      @click="itemClick(index)"
      :class="{ active: index == currentIndex }"
    >
      <span>{{ item }}</span>
    </div>
  </div>
</template>

<script>
import { ref } from "vue";
export default {
  props: {
    titles: {
      type: Array,
      default() {
        return [];
      },
    },
  },
  setup(props, ctx) {
    let currentIndex = ref(0);
    //点击切换选项卡
    let itemClick = (index) => {
      currentIndex.value = index;
      //切换选项卡后下面内容要对应变化，因此携带index发送请求到父组件
      ctx.emit('tabClick',index)
    };

    return {
      currentIndex,
      itemClick,
    };
  },
};
</script>
<style scoped lang="scss">
.tab-control {
  display: flex;
  height: 40px;
  line-height: 40px;
  text-align: center;
  font-size: 14px;
  background-color: #ffffff;
  width: 100%;
  z-index: 30;

  position: sticky;
  top: 44px;

  .tab-control-item {
    flex: 1;
    span {
      padding: 6px;
    }
  }
  .active {
    color: var(--color-tint);
    span {
      border-bottom: 3px solid var(--color-tint);
    }
  }
}
</style>