<template>
  <div class="recommend">
    <div
      class="reommend-item"
      v-for="(item, index) in recommends.splice(0, 4)"
      :key="index"
    >
      <a href="" @click.prevent="godetail(item.id)">
        <!-- 跳转到详情页需要携带id -->
        <img v-lazy="item.cover_url" alt="" />
        <div>{{ item.title }}</div>
      </a>
    </div>
  </div>
</template>

<script>
import { useRouter } from "vue-router";
export default {
  props: {
    recommends: {
      type: Array,
      default() {
        return [];
      },
    },
  },
  setup() {
    const router = useRouter();
    const godetail = (id) => {
      router.push({
        path: "/detail",
        //需要传递参数
        query: {
          id
        },
      });
    };
    return { godetail };
  },
};
</script>

<style scoped lang="scss">
.recommend {
  display: flex;
  width: 100%;
  text-align: center;
  padding: 15px 0 30px;
  border-bottom: 8px solid #eee;
  font-size: 12px;
}
.recommend-item {
  flex: 1;
  img {
    width: 70px;
    height: 70px;
    margin-bottom: 10px;
  }
}
</style>