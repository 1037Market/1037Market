<template>
  <div class="display seller-comments">
    <h3>{{ title }}</h3>
    <div class="comments-container">
      <div
          v-for="(comment, index) in displayedComments"
          :key="comment.id"
          class="comment"
          @click="clickComment(comment.commenter.id)"
      >
        <div class="commenter-details">
          <van-image
              :src="comment.commenter.avatar"
              alt="Commenter's Avatar"
              class="commenter-avatar"
              radius="15px"
          />
          <span class="commenter-nickname">{{ comment.commenter.nickname }}</span>
        </div>
        <p class="comment-text">{{ comment.text }}</p>
      </div>
      <!-- "View More" Button -->
      <button
          v-if="seller.comments.length > 2 && !showAllComments"
          class="view-more-btn"
          @click="toggleComments"
      >
        查看更多
      </button>
      <!-- "Collapse" Button -->
      <button
          v-if="showAllComments"
          class="collapse-btn"
          @click="toggleComments"
      >
        收起
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  title: String,
  seller: {
    type: Object,
    default: () => ({
      avatar: '',
      nickname: '',
      comments: [],
      products: [],
    }),
  },
})

const showAllComments = ref(false)

function toggleComments() {
  showAllComments.value = !showAllComments.value
}

function clickComment(commenterId) {
  // Handle click on comment
}

const displayedComments = computed(() => {
  return showAllComments.value
      ? props.seller.comments
      : props.seller.comments.slice(0, 2)
})
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

.view-more-btn,  .collapse-btn {
  background-color: transparent;
  border: none;
  color: #1f8efa;
  padding: 10px;
  font-size: 14px;
  cursor: pointer;
  display: block;
  margin: 10px auto; /* Center the button */
}

.view-more-btn:hover, .collapse-btn:hover {
  text-decoration: underline;
}
</style>
