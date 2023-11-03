<template>
  <div>
    <nav-bar>
      <!-- vue3插槽写法 -->
      <template v-slot:center>个人中心</template>
    </nav-bar>



    <div style="margin: 60px">
      <van-image
          width="10rem"
          height="10rem"
          radius="5rem"
          fit="contain"
          :src="'http://franky.pro:7301/api/image?imageURI=' + userInfo['avatar']"
      />
      <van-cell-group>
        <van-field v-model="userInfo.nickName" label="昵称"
                   :right-icon="infoEditing.nickName?'sign':'edit'"
                   @update:model-value="updateNickName"
                   @click-right-icon="saveUserInfo"
        />
        <van-field v-model="userInfo.studentId" label="学号" readonly/>
        <van-field v-model="userInfo.contact" label="联系方式"
                   :right-icon="infoEditing.contact?'sign':'edit'"
                   @update:model-value="updateContact"
                   @click-right-icon="saveUserInfo"
        />
      </van-cell-group>
    </div>


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
import { updateUser } from "@/network/user";
import { Toast } from "vant";

import {reactive} from "vue";
import {getUser} from "@/network/user";
export default {
  name: "Profile",
  components: {
    NavBar,
  },
  setup() {

    const router = useRouter();
    const store = useStore();
    const tologout = () => {
        window.localStorage.removeItem('token')

        Cookies.remove('user')

        setTimeout(() => {
            router.push({ path: "/login" });
        }, 500);
    };
    const userInfo = reactive({
      studentId: "",
      avatar: "",
      contact: "",
      nickName: ""
    })

    const infoEditing = reactive(
        {
          nickName: false,
          contact: false
        }
    )
    getUser().then((data) => {
      userInfo.studentId = data['userId'];
      userInfo.avatar = data['avatar'];
      userInfo.contact = data['contact'];
      userInfo.nickName = data['nickName'];
    })

    const updateNickName = () => {
      infoEditing.nickName = true;
    }

    const updateContact = () => {
      infoEditing.contact = true;
    }

    const saveUserInfo = () => {
      console.log(userInfo)
      updateUser(userInfo).then((response) => {
        if(response !== 'OK') {
          Toast.fail('保存失败');
        } else {
          infoEditing.contact = false;
          infoEditing.nickName = false;
          Toast.success('保存成功');

        }
      })

    }

    return {
      tologout,
      userInfo,
      infoEditing,
      updateContact,
      updateNickName,
      saveUserInfo
    };
  },
};
</script>

<style>
</style>