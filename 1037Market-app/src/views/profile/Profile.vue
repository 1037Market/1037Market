<template>
    <div id="profile">
        <van-nav-bar title="个人中心" fixed placeholder
                     left-arrow @click-left="router.go(-1)"
        />
        <van-uploader v-model="fileList" max-count="1" :after-read="afterReadAvatar">
            <van-image
                width="10rem"
                height="10rem"
                radius="5rem"
                fit="contain"
                :src="'http://franky.pro:7301/api/image?imageURI=' + userInfo['avatar']"
                @click="clickAvatar"
            />
            <van-icon name="photograph" id="photograph"
                      size="large"
            />
        </van-uploader>


        <!--TODO:调整文字输入位置,还有地址的model-value-->
        <div class="display">
            <van-cell-group>
                <van-field left-icon="contact" v-model="userInfo.nickName" label="昵称"
                           :right-icon="infoEditing.nickName?'sign':'edit'"
                           @update:model-value="updateNickName"
                           @click-right-icon="saveUserInfo"
                />
                <van-field left-icon="bookmark" v-model="userInfo.studentId" label="学号" readonly/>
                <van-field left-icon="phone" v-model="userInfo.contact" label="联系方式"
                           :right-icon="infoEditing.contact?'sign':'edit'"
                           @update:model-value="updateContact"
                           @click-right-icon="saveUserInfo"
                />
                <van-field left-icon="map-marked" v-model="userInfo.address" label="地址"
                           :right-icon="infoEditing.address?'sign':'edit'"
                           @update:model-value="updateAddress"
                           @click-right-icon="saveUserInfo"
                />
            </van-cell-group>
        </div>

    </div>

    <profile-comments :seller="seller" title="我收到的评论"/>
    <profile-products :seller="seller" title="我发布的商品"/>

    <div style="margin: 60px">
        <van-button round block color="#42b983" @click="tologout"
        >退出登录
        </van-button
        >
    </div>
</template>

<script>
import {useStore} from "vuex";
import {useRouter} from "vue-router";
import Cookies from 'js-cookie'
import {updateUser} from "@/network/user";
import {showSuccessToast, showFailToast} from 'vant';

import {reactive, ref} from "vue";
import {getUser} from "@/network/user";
import {uploadImage} from "@/network/image";
import {getCommentDetail, getUserCommentIds} from "../../network/comment";
import {getDetail} from "../../network/detail";
import {getUserPublishedProductIds} from "../../network/user";
import {route} from "vant/es/composables/use-route";
import ProfileComments from "@/components/content/profile/Comments.vue";
import ProfileProducts from "@/components/content/profile/Products.vue";

export default {
    name: "Profile",
    components: {
        ProfileProducts,
        ProfileComments,
    },
    setup() {

        const router = useRouter();
        const store = useStore();
        const tologout = () => {
            window.localStorage.removeItem('token')

            Cookies.remove('user')

            setTimeout(() => {
                router.push({path: "/login"});
            }, 500);
        };
        const userInfo = reactive({
            studentId: "",
            avatar: "",
            contact: "",
            nickName: "",
            address: ""
        })

        const infoEditing = reactive(
            {
                nickName: false,
                contact: false,
                avatar: false,
                address: false
            }
        )

        const fetchUserInfo = () => {
            getUser().then((data) => {
                userInfo.studentId = data['userId'];
                userInfo.avatar = data['avatar'];
                userInfo.contact = data['contact'];
                userInfo.nickName = data['nickName'];
                userInfo.address = data['address'];
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

        const updateAddress = () => {
            infoEditing.address = true;
        }

        const saveUserInfo = () => {
            console.log(userInfo)
            updateUser(userInfo).then((response) => {
                if (response !== 'OK') {
                    showFailToast('保存失败');
                } else {
                    infoEditing.contact = false;
                    infoEditing.nickName = false;
                    infoEditing.address = false;
                    showSuccessToast('保存成功');
                    fetchUserInfo();
                }
            })
        }

        const fileList = ref([]);


        const afterReadAvatar = (file, detail) => {
            let formData = new FormData()
            formData.append('file', file.file)
            console.log(formData)
            uploadImage(formData).then((response) => {
                userInfo.avatar = response;
                updateUser(userInfo).then(() => {
                    showSuccessToast('上传成功');
                    fetchUserInfo();
                    fileList.value = [];
                }).catch((err) => {
                    showFailToast('上传失败');
                    console.log(err)
                })
            }).catch((err) => {
                showFailToast('上传失败');
                console.log(err);
            });
        }


        const seller = reactive({
            avatar: '',
            nickname: '',
            comments: [],
            products: []
        });

        const sellerInfo = reactive({
            productIds: [],
            commentIds: [],
            commentContents: []
        });

        const updateView = () => {
            getUser(window.localStorage.getItem('studentId')).then((response) => { // 获取用户信息
                seller.studentId = response.userId;
                seller.nickname = response.nickName;
                seller.avatar = 'http://franky.pro:7301/api/image?imageURI=' + response.avatar;
                seller.contact = response.contact;
                console.log(seller)

                getUserPublishedProductIds(seller.studentId).then((response) => { // 拿到该用户发布的商品id列表
                    sellerInfo.productIds = response;

                    seller.products = [];
                    sellerInfo.productIds.forEach((productId) => {
                        getDetail(productId).then((response) => { // 获取每个商品的详情
                            seller.products.push({
                                id: productId,
                                name: response.title,
                                description: response.content,
                                image: 'http://franky.pro:7301/api/image?imageURI=' + response.imageURIs[0]
                            })
                        })
                    })

                    getUserCommentIds(seller.studentId).then((response) => {
                        sellerInfo.commentIds = response;

                        seller.comments = [];
                        sellerInfo.commentIds.forEach((commentId) => {
                            getCommentDetail(commentId).then((response) => {
                                sellerInfo.commentContents.push({
                                    id: commentId,
                                    text: response.content
                                })
                                seller.comments.push({
                                    id: commentId,
                                    text: response.content,
                                    stars: response.stars,
                                    commenter: {
                                        id: response.fromId,
                                        nickname: response.nickName,
                                        avatar: 'http://franky.pro:7301/api/image?imageURI=' + response.avatar
                                    }
                                });
                            }).catch((err) => {
                                console.log(err);
                            });
                        })
                    }).catch((err) => {
                        console.log(err);
                    })


                }).catch((err) => {
                    console.log(err);
                });

                getUserCommentIds(seller.studentId).then((response) => {
                    sellerInfo.commentIds = response;
                    console.log(response)
                }).catch((err) => {
                    console.log(err);
                })

            }).catch((err) => {
                console.log(err);
            })
        }
        updateView();

        return {
            tologout,
            userInfo,
            infoEditing,
            updateContact,
            updateNickName,
            saveUserInfo,
            fileList,
            clickAvatar,
            afterReadAvatar,
            seller,
            updateAddress,
            router
        };
    },
};
</script>

<style scoped>
#profile {
    text-align: center;
}

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

#photograph {

    margin-left: -30px;
}
</style>