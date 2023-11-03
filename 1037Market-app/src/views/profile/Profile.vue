<template>
  <div>
    <nav-bar>
      <!-- vue3插槽写法 -->
      <template v-slot:center>个人中心</template>
    </nav-bar>

    <div style="margin: 60px">
      <van-button round block color="#42b983" @click="tologout"
        >退出登录</van-button
      >
    </div>
  </div>
</template>

<script>
import NavBar from "components/common/navbar/NavBar";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import Cookies from 'js-cookie'
import { logout } from "network/user";
import { Toast } from "vant";
export default {
  name: "Profile",
  components: {
    NavBar,
  },
  setup() {
    const router = useRouter();
    const store = useStore();
    const tologout = () => {
        store.commit("setIsLogin", false);

        Cookies.remove('user')

        setTimeout(() => {
            router.push({ path: "/login" });
        }, 500);
    };

    return {
      tologout,
    };
  },
};
</script>

<style>
</style>