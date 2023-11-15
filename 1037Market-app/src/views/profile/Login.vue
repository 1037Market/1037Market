<template>
    <div>
        <van-nav-bar title="用户登录" fixed style="--van-nav-bar-background: linear-gradient(rgba(66, 185, 131, 0.9),rgba(66,185,131,0.45));--van-nav-bar-title-text-color: rgba(255,255,255,1);"
        />
        <div style="text-align: center; padding-top: 100px">
            <img style="display: block;
                width: 250px;
                height: 125px;
                margin: 0 auto;
                "
                src="@/assets/images/logo.png"

            />
        </div>
        <div style="margin-top: 50px" class="display">

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
                    >提交
                    </van-button
                    >
                </div>
            </van-form>
        </div>
    </div>
</template>

<script>
import {ref, reactive, toRefs} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {login, hashPassword} from "@/network/user";
import {showSuccessToast, showFailToast} from 'vant';

export default {
    components: {},
    setup() {
        const router = useRouter();
        const store = useStore();
        const userinfo = reactive({
            studentId: "",
            hashedPassword: "",
        });

        const onSubmit = async () => {
            let hashedUserInfo = {
                studentId: userinfo.studentId,
                hashedPassword: ''
            }
            hashedUserInfo.hashedPassword = await hashPassword(userinfo.hashedPassword)
            console.log(userinfo)
            console.log(hashedUserInfo)
            login(hashedUserInfo).then((res) => {
                if(res === undefined){
                    showFailToast("登录失败")
                    return
                }
                window.localStorage.setItem('token', res)
                store.commit("setIsLogin", true);
                window.localStorage.setItem('studentId', userinfo.studentId);

                showSuccessToast("登录成功");
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
.display {
    font-family: 'Poppins', sans-serif;
    margin: 10px;
    padding: 15px;
    border-radius: 15px;
    box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
    color: #333;
    line-break: anywhere;
    text-align: left;
}

.link-login {
    font-size: 14px;
    margin-bottom: 20px;
    color: #42b983;
    display: inline-block;
    text-align: left;
}
</style>