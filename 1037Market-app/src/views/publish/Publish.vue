<template>
    <van-form @submit="onSubmit">
        <van-field
            v-model="form.name"
            label="商品名称"
            name="name"
            placeholder="请输入商品名称"
            required
        />

        <van-uploader
            v-model="form.images"
            :after-read="afterRead"
            multiple
            :max-count="5"
            accept="image/*"
            :deletable="true"
            @delete="onDelete"
        />

        <van-field
            v-model="form.description"
            label="商品描述"
            name="description"
            type="textarea"
            placeholder="请输入商品描述"
            required
        />

        <van-field
            v-model="form.price"
            label="商品价格"
            name="price"
            type="number"
            placeholder="请输入商品价格"
            required
        />

        <van-radio-group v-model="form.class1">
            <van-radio name="二手书">二手书</van-radio>
            <van-radio name="闲置物品">闲置物品</van-radio>
        </van-radio-group>

        <van-checkbox-group v-model="form.class2">
            <van-checkbox name="数学">数学</van-checkbox>
            <van-checkbox name="英语">英语</van-checkbox>
            <van-checkbox name="计算机">计算机</van-checkbox>
        </van-checkbox-group>

        <van-button round block type="info" native-type="submit">提交</van-button>
    </van-form>
</template>

<script>
import { ref, reactive } from 'vue';

export default {
    setup() {
        const form = reactive({
            name: '',
            images: [],
            description: '',
            price: null,
            class1: 'books',
            class2: []
        });

        const afterRead = (file) => {
            // 此处应实现上传逻辑，并将返回的图片URL添加到form.images中
            const reader = new FileReader();
            reader.readAsDataURL(file.file);
            reader.onload = () => {
                // 假设上传成功，将图片URL添加到form.images
                form.images.push(reader.result);
            };
        };

        const onDelete = (index) => {
            // 删除指定索引的图片
            form.images.splice(index, 1);
        };

        const onSubmit = () => {
            console.log('提交的表单数据', form);
            // 在这里处理表单的提交，比如将数据发送到服务器
        };

        return {
            form,
            afterRead,
            onDelete,
            onSubmit,
        };
    },
};
</script>
