<template>
  <nav-bar>
    <template v-slot:center>卖家详情</template>
  </nav-bar>
  <div class="detail-container">
    <!-- Seller Details -->
    <div class="display seller-details">
      <van-image :src="seller.avatar" alt="Seller's Avatar" class="seller-avatar" radius="45px" />
      <h2 class="seller-nickname">{{ seller.nickname }}</h2>
      <h4>{{seller.contact}}</h4>
      <!-- Rating could be included here if needed -->
    </div>

    <!-- Customer Reviews -->
    <div class="display seller-comments">
      <h3>TA收到的评论</h3>
      <div class="comments-container">
        <div v-for="comment in seller.comments" :key="comment.id" class="comment" @click="clickComment(comment.commenter.id)">
          <div class="commenter-details">
            <van-image :src="comment.commenter.avatar" alt="Commenter's Avatar" class="commenter-avatar" radius="15px" />
            <span class="commenter-nickname">{{ comment.commenter.nickname }}</span>
          </div>
          <p class="comment-text">{{ comment.text }}</p>
        </div>
      </div>
    </div>

    <!-- Published Products -->
    <div class="display seller-products">
      <h3>TA的商品</h3>
      <div class="products-grid">
        <div class="product-card" v-for="product in seller.products" :key="product.id" @click="clickProduct(product.id)">
          <van-image :src="product.image" alt="Product Image" class="product-image" />
          <div class="product-info">
            <h4 class="product-name">{{ product.name }}</h4>
            <p class="product-description">{{ product.description }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import NavBar from "@/components/common/navbar/NavBar.vue";

import {nextTick, onMounted, reactive, ref} from 'vue';
import {getUser, getUserPublishedProductIds} from "@/network/user";
import {useRoute, useRouter} from "vue-router/dist/vue-router";
import {useStore} from "vuex";
import {getCommentDetail, getUserCommentIds} from "@/network/comment";
import {getDetail} from "@/network/detail";

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
  commentIds: []
});

const updateView = () => {
  getUser(route.params.studentId).then((response) => { // 获取用户信息
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
            seller.comments.push({
              id: commentId,
              text: response.content,
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

const clickProduct = (productId) => {
  router.push({path: `/detail/${productId}`})
}

const clickComment = (commenterId) => {

}

</script>


<style scoped>
.detail-container {
  margin-top: 50px;
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

.seller-comments {
  margin-top: 20px; /* Additional spacing if needed */
}

.comments-container {
  border-radius: 8px;
  overflow: hidden;
}

.comment {
  padding: 10px 0;
  border-top: 1px solid #eee;
}

.comment:first-child {
  border-top: none; /* Remove border for the first item */
}

.commenter-details {
  display: flex;
  align-items: center;
  margin-bottom: 5px;
}

.commenter-avatar {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  margin-right: 10px;
}

.commenter-nickname {
  font-weight: bold;
  font-size: 14px;
}

.comment-text {
  font-size: 13px;
  line-break: anywhere; /* Ensure long words do not break the layout */
}
.seller-products h3 {
  margin-top: 20px;
  margin-bottom: 20px;
  text-align: center;
}

.products-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
  gap: 10px;
  margin-bottom: 45px;
}

.product-card {
  border: 1px solid #e1e1e1;
  border-radius: 6px;
  overflow: hidden;
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
  height: 200px;
}

.product-image {
  width: 100%;
  height: 70%;
  overflow: hidden;
  object-fit: contain;
}

.product-info {
  padding: 8px;
}

.product-name {
  font-size: 16px;
  font-weight: bold;
  color: #333;
  margin: 0 0 5px 0;
}

.product-description {
  font-size: 12px;
  color: #666;
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
</style>