<template>
    <van-nav-bar  title="发布商品" fixed placeholder 
                 left-arrow @click-left="router.go(-1)"
    />
    <product-publish :dialog="dialog" :form="form" :update="false"/>
</template>

<script>
import {ref, reactive} from 'vue';
import {Dialog, showFailToast, showSuccessToast, Toast} from "vant";
import {getCategoryData} from "@/network/category";
import {uploadImage} from "@/network/image";
import {updateUser} from "@/network/user";
import {publishProduct} from "@/network/publish";
import PriceDisplay from "@/components/content/goods/PriceDisplay.vue";
import {useRouter} from "vue-router"
import ProductPublish from "@/components/content/goods/ProductPublish.vue";

export default {
    components: {
        ProductPublish,
        'BackgroundSurround': PriceDisplay
    },

    setup() {
        const router = useRouter()
        const form = reactive({
            name: '',
            images: [],
            imageURIs: [],
            description: '',
            price: null,
            categories: []
        });

        const dialog = reactive({
            show: false,
            value: '添加分类',
            options: []
        });

        getCategoryData().then((categories) => {
            dialog.options = [];
            console.log(categories)
            categories.forEach((element) => {
                dialog.options.push({
                    text: element,
                    value: element
                })
            });
        }).catch((err) => {
            console.log(err);
        });

        return {
            dialog,
            form,
            router
        }

    },
};
</script>

<style scoped>

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

div van-form {
    width: 100%;
}
</style>