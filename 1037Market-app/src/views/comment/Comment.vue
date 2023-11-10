<script setup>
import {ref} from "vue"
import {useRouter, useRoute} from "vue-router";
import {uploadImage} from "@/network/image";
import {postComment} from "@/network/comment";
import {showSuccessToast, showFailToast} from "vant";

const route = useRoute()
const router = useRouter()
const score = ref(0)
const message = ref('')
const images = ref([])

const onSubmit = () => {
    postComment(window.localStorage.getItem('token'),{
        toId: route.params.studentId,
        content: message.value,
        stars: score.value
    }).then(() => {
        showSuccessToast('发布成功')
    }).catch(() => {
        showFailToast('发布失败\n请重新尝试')
    })
}

const readImage = (file) => {
    let formData = new FormData();
    file.status = 'uploading';
    file.message = '上传中';
    formData.append('file', file.file);
    console.log(formData)
    uploadImage(formData).then((response) => {
        images.value.push(response);
        file.status = '';
        file.message = '';
    }).catch((err) => {
        file.status = 'failed';
        file.message = '上传失败';
        images.value.push(err);
        console.log(err);
    })
}

const deleteImage = (index) => {
    images.value.splice(index, 1)
}

</script>

<template>
    <van-nav-bar title="1037集市" fixed placeholder
                 style="--van-nav-bar-background: linear-gradient(rgba(66, 185, 131, 0.9),rgba(66,185,131,0.45));--van-nav-bar-title-text-color: rgba(255,255,255,1);"
                 left-arrow @click-left="router.go(-1)"
    />

    <van-form @submit="onSubmit" style="margin-top: 20%;">

        <van-field required name="rate" label="评分">
            <template #input>
                <van-rate v-model="score"/>
            </template>
        </van-field>

        <van-field name="comment" label="评论" v-model="message"
                   type="textarea" placeholder="留下本次交易的感受吧，你的评论能够帮到其他人~"
                   rows="7" autosize
        />

<!--        <van-field name="images" label="图片">-->
<!--            <template #input>-->
<!--                <van-uploader v-model="images" :after-read="readImage" @delete="deleteImage"-->
<!--                              multiple max-count="4" accept="image/*" deletable/>-->
<!--            </template>-->
<!--        </van-field>-->

        <div style="margin-top: 16px;">
            <van-button round block type="success" color="#42b983" native-type="submit">发布评论</van-button>
        </div>
    </van-form>
</template>

<style scoped>

</style>