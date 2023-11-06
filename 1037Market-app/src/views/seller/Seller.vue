<template>
    <van-nav-bar title="卖家信息" fixed
                 left-arrow @click-left="router.go(-1)"
    />

    <div class="detail-container">
        <!-- Seller Details -->
        <van-barrage v-model="sellerInfo.commentContents">
            <div class="display seller-details">
                <van-image :src="seller.avatar" alt="Seller's Avatar" class="seller-avatar" radius="45px"/>
                <van-field left-icon="contact" v-model="seller.nickname" label="卖家昵称" size="large" readonly/>
                <van-field left-icon="phone" v-model="seller.contact" label="卖家电话" size="large" readonly/>
                <van-field left-icon="map-marked" v-model="seller.address" label="卖家地址" size="large" readonly/>
                <!-- Rating could be included here if needed -->
            </div>
        </van-barrage>

        <profile-comments :seller="seller" title="TA收到的评论"/>

        <profile-products :seller="seller" title="TA出售的商品"/>

    </div>
</template>

<script setup>

import {nextTick, onMounted, reactive, ref} from 'vue';
import {getUser, getUserPublishedProductIds} from "@/network/user";
import {useRoute, useRouter} from "vue-router/dist/vue-router";
import {useStore} from "vuex";
import {getCommentDetail, getUserCommentIds} from "@/network/comment";
import {getDetail} from "@/network/detail";
import ProfileComments from "@/components/content/profile/Comments.vue";
import ProfileProducts from "@/components/content/profile/Products.vue";

const route = useRoute();
const router = useRouter();
const store = useStore();

const seller = reactive({
    avatar: '', // 卖家头像路径
    nickname: '', // 卖家昵称
    comments: [
        // {
        //   id: 1,
        //   text: 'Great service, very satisfied with the purchase!',
        //   commenter: {
        //     nickname: 'Alice',
        //     avatar: 'https://i.pravatar.cc/150?img=3'
        //   }
        // },
    ],
    products: [
        // {
        //   id: 1,
        //   name: 'Vintage Leather Bag',
        //   description: 'High-quality leather bag with unique design.',
        //   image: 'https://i.pravatar.cc/300?u=bag',
        // },
    ]
});

const sellerInfo = reactive({
    productIds: [],
    commentIds: [],
    commentContents: []
});

const updateView = () => {
    getUser(route.params.studentId).then((response) => { // 获取用户信息
        seller.studentId = response.userId;
        seller.nickname = response.nickName;
        seller.avatar = 'http://franky.pro:7301/api/image?imageURI=' + response.avatar;
        seller.contact = response.contact;
        seller.address = response.address;
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
                            text: response.content,
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

onMounted(() => {
    updateView();
})

</script>


<style scoped>
.detail-container {
    margin-top: 1px;
}

.seller-details {
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
}

.seller-avatar {
    width: 90px;
    height: 90px;
    border-radius: 50%;
    margin-bottom: 10px;
}

.seller-info h2 {
    margin: 5px 0;
}

.seller-rating .rating-icon {
    color: #ffcc00;
}

</style>

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

.van-field {
    text-align: center;
}
</style>