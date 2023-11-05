<template>
  <!-- Customer Reviews -->
  <div class="display seller-comments">
    <h3>{{title}}</h3>
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
</template>

<script>
import {reactive} from "vue";

const clickComment = (commenterId) => {

}


export default {
  name: "ProfileComments",
  props: {
    title: '',
    seller: {
        avatar: '',
        nickname: '',
        comments: [
        ],
        products: [
        ]
    }
  },
  methods: {
    clickComment
  }
}
</script>

<style scoped>

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
</style>