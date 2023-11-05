<template>
  <div id="profile">
    <van-nav-bar title="个人中心"
                 fixed
                 placeholder

    />
      <van-image
            width="10rem"
            height="10rem"
            radius="5rem"
            fit="contain"
            :src="'http://franky.pro:7301/api/image?imageURI=' + userInfo['avatar']"
            @click="clickAvatar"
      />
      <div style="margin-bottom: 10px">
        <van-uploader v-model="fileList" max-count="1" :after-read="afterReadAvatar">
          <van-button block color="#42b983" @click="tologout">上传头像</van-button>
        </van-uploader>
      </div>

      <div class="display">
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

    </div>


    <div style="margin: 60px">

      <van-button round block color="#42b983" @click="tologout"
        >退出登录</van-button
      >

    </div>
</template>

<script>
import NavBar from "@/components/common/navbar/NavBar.vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import Cookies from 'js-cookie'
import { updateUser } from "@/network/user";
import { Toast } from "vant";

import {reactive, ref} from "vue";
import {getUser} from "@/network/user";
import {uploadImage} from "@/network/image";
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
          contact: false,
          avatar: false
        }
    )

    const fetchUserInfo = () => {
      getUser().then((data) => {
        userInfo.studentId = data['userId'];
        userInfo.avatar = data['avatar'];
        userInfo.contact = data['contact'];
        userInfo.nickName = data['nickName'];
      })
    }

    fetchUserInfo();

    const updateNickName = () => {
      infoEditing.nickName = true;
    }

    const updateContact = () => {
      infoEditing.contact = true;
    }

    const clickAvatar = () => {
      infoEditing.avatar = true;
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
          fetchUserInfo();
        }
      })
    }

    const fileList = ref([

    ]);



    const afterReadAvatar = (file, detail) => {
      let formData = new FormData()
      formData.append('file', file.file)
      console.log(formData)
      uploadImage(formData).then((response) => {
        userInfo.avatar = response;
        updateUser(userInfo).then(() => {
          Toast.success('上传成功');
          fetchUserInfo();
          fileList.value = [];
        }).catch((err) => {
          Toast.fail('上传失败');
          console.log(err)
        })
      }).catch((err) => {
        Toast.fail('上传失败');
        console.log(err);
      });
    }

    return {
      tologout,
      userInfo,
      infoEditing,
      updateContact,
      updateNickName,
      saveUserInfo,
      fileList,
      clickAvatar,
      afterReadAvatar
    };
  },
};
</script>

<style scoped>
#profile {
    text-align: center;
}

.display {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px;
  margin: 10px; /* Added horizontal margin */
  color: #333;
  border-radius: 15px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  font-family: 'Poppins', sans-serif;
}
</style>