<template>
  <van-nav-bar title="发布商品"
               fixed
               placeholder

  />
  <product-publish :dialog="dialog" :form="form" :update="false" />
</template>

<script>
import { ref, reactive } from 'vue';
import {Dialog, showFailToast, showSuccessToast, Toast} from "vant";
import {getCategoryData} from "@/network/category";
import {uploadImage} from "@/network/image";
import {updateUser} from "@/network/user";
import {publishProduct} from "@/network/publish";
import PriceDisplay from "@/components/content/goods/PriceDisplay.vue";
import ProductPublish from "@/components/content/goods/ProductPublish.vue";

export default {
  components: {
    ProductPublish,
    'BackgroundSurround': PriceDisplay
  },

  setup() {
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
          form
        }

    },
};
</script>

<style scoped>

</style>