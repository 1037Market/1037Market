<template>
  <div>
    <nav-bar>
      <template v-slot:center>用户登录</template>
    </nav-bar>

    <div style="margin-top: 50px">
      <div style="text-align: center; padding-top: 50px">
        <van-image
          width="10rem"
          height="5rem"
          fit="contain"
          src="https://cdn2.lmonkey.com/94f152aaa94d937ccf5de78f3fcac59f/3b6e32e7bc8145a283431f260c3c1d1a.png"
        />
      </div>
      <van-form @submit="onSubmit">
        <van-field
          v-model="studentId"
          name="电子邮箱"
          label="电子邮箱"
          placeholder="请输入正确电子邮箱格式"
          :rules="[{ required: true, message: '请填写用户名' }]"
        />
        <van-field
          v-model="hashedPassword"
          type="password"
          name="密码"
          label="密码"
          placeholder="密码"
          :rules="[{ required: true, message: '请填写密码' }]"
        />

        <div style="margin: 16px">
          <div class="link-login" @click="$router.push({ path: '/register' })">
            没有账号，立即注册
          </div>
          <van-button
            round
            block
            type="info"
            color="#44b883"
            native-type="submit"
            >提交</van-button
          >
        </div>
      </van-form>
    </div>
  </div>
</template>

<script>
import { ref, reactive, toRefs } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "vuex";
import { login } from "network/user";
import { Notify } from "vant";
import { Toast } from "vant";
import NavBar from "components/common/navbar/NavBar";
export default {
  components: {
    NavBar,
  },
  setup() {
    const router = useRouter();
    const store = useStore();
    const userinfo = reactive({
        studentId: "",
        hashedPassword: "",
    });

    const onSubmit = () => {
      login(userinfo).then((res) => {
        //在Vuex isLogin
        store.commit("setIsLogin", true);

        Toast.success("登录成功");
        userinfo.studentId = "";
        userinfo.hashedPassword = "";

        setTimeout(() => {
          router.push({path: '/profile'});
        }, 500);
      });
    };
    return {
      ...toRefs(userinfo),
      onSubmit,
    };
  },
};
</script>

<style scoped>
.link-login {
  font-size: 14px;
  margin-bottom: 20px;
  color: #42b983;
  display: inline-block;
  text-align: left;
}
</style>