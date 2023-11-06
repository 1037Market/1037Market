<template>
  <nav-bar>
    <template v-slot:center>更新商品</template>
  </nav-bar>
  <div style="margin-top: 50px" />
  <product-publish :dialog="dialog" :form="form" :update="true" />
</template>

<script>
import {ref, reactive, onMounted} from 'vue';
import {getCategoryData} from "@/network/category";
import PriceDisplay from "@/components/content/goods/PriceDisplay.vue";
import ProductPublish from "@/components/content/goods/ProductPublish.vue";
import {getDetail} from "@/network/detail";
import NavBar from "@/components/common/navbar/NavBar.vue";
import {useRoute} from "vue-router/dist/vue-router";

export default {
  components: {
    ProductPublish,
    'BackgroundSurround': PriceDisplay,
    NavBar
  },

  setup() {

    const route = useRoute();

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

    onMounted(() => {
      let id = ref(route.params.id);
      getDetail(id.value).then((res) => {
        form.name = res.title;
        form.imageURIs = res.imageURIs;
        form.description = res.content;
        form.price = res.price;
        form.categories = res.categories;

      }).catch((err) => {
        console.log(err);
      });
      getCategoryData().then((categories) => {
        dialog.options = [];
        categories.forEach((element) => {
          dialog.options.push({
            text: element,
            value: element
          })
        });
      }).catch((err) => {
        console.log(err);
      });

    })



    return {
      dialog,
      form
    }

  },
};
</script>

<style scoped>

</style>