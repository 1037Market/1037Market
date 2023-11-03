<template>
  <div>
    <nav-bar>
      <template v-slot:center>用户注册</template>
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
              name="学号"
              label="学号"
              placeholder="学号"
              :rules="[{ required: true, message: '请填写学号' }]"
          />
          <van-button :disabled="waitCaptcha" @click="clickCaptcha">{{captchaHint}}</van-button>

        <van-field
          v-model="password"
          type="password"
          name="密码"
          label="密码"
          placeholder="密码"
          :rules="[{ required: true, message: '请填写密码' }]"
        />

        <van-field
          v-model="password_confirmation"
          type="password"
          name="确认密码"
          label="确认密码"
          placeholder="确认密码"
          :rules="[{ required: true, message: '请填写一致密码' }]"
        />

        <van-field
          v-model="captcha"
          name="验证码"
          label="验证码"
          placeholder="请输入HUST邮箱中的验证码,如果没有请在垃圾箱寻找"
          :rules="[{ required: true, message: '请填写正确的验证码' }]"
        />
        <div style="margin: 16px">
          <div class="link-login" @click="$router.push({ path: '/login' })">
            已有账号，立即登录
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
import { register, getCaptcha } from "network/user";
import { Notify } from "vant";
import { Toast } from "vant";
import NavBar from "components/common/navbar/NavBar";
export default {
  components: {
    NavBar,
  },
  setup() {
    const router = useRouter();
    const userinfo = reactive({
        studentId: "",
        password: "",
        password_confirmation: "",
        captcha: "",
    });
    const waitCaptcha = ref(false);
    const captchaHint = ref('点击获取验证码')

    const disableButton = (disabled) => {
        waitCaptcha.value = true
        captchaHint.value = '60s后再次获取验证码'
        let remaining = 60
        const interval = setInterval(() => {
            remaining--;
            captchaHint.value = remaining.toString() + 's后再次获取验证码'
            if(remaining <= 0){
                clearInterval(interval)
            }
        }, 1000)
        captchaHint.value = '点击获取验证码'
    }

    const clickCaptcha = () => {
        getCaptcha({studentId: userinfo.studentId})
            .then((res) => {
                if(res == 'OK')
                    disableButton(waitCaptcha)
                else{
                    captchaHint.value = '获取失败，请重新点击'
                }
            })
    }

    const onSubmit = () => {
      //先验证
      if (userinfo.password != userinfo.password_confirmation) {
        Notify("两次密码不一致");
      } else {
        register({
            studentId: userinfo.studentId,
            hashedPassword: userinfo.password,
            emailCaptcha: userinfo.captcha
        }).then((res) => {
          console.log(res);
          if (res == 'OK') {
            Toast.success("注册成功");

            setTimeout(() => {
              router.push({ path: "/login" });
            }, 1000);
          }

          userinfo.password = "";
          userinfo.password_confirmation = "";
        });
      }
    };
    return {
        ...toRefs(userinfo),
        onSubmit,
        waitCaptcha,
        captchaHint,
        clickCaptcha
    };
  },
};
</script>

<style scoped>
.link-login {
    font-size:14px;
    margin-bottom: 20px;
    color:#42b983;
    display: inline-block;
    text-align: left;
   
}
</style>