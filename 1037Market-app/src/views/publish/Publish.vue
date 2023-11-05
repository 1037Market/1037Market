<template>
  <van-nav-bar title="发布商品"
               fixed
               placeholder

  />
    <div class="display">
        <van-form @submit="onSubmit" style="width: 100%;">
          <van-field
              v-model="form.name"
              label="商品名称"
              name="name"
              placeholder="请输入商品名称"
              required
              style="margin-bottom: 10px"
          />

          <van-uploader style="margin-left: 130px;"
              v-model="form.images"
              :after-read="afterRead"
              multiple
              :max-count="4"
              accept="image/*"
              :deletable="true"
              @delete="onDelete"
          />

          <van-field
              v-model="form.price"
              label="商品价格"
              name="price"
              type="number"
              placeholder="请输入商品价格"
              required
          />

          <van-field
              v-model="form.description"
              label="商品描述"
              name="description"
              type="textarea"
              placeholder="请输入商品描述"
              required
              style="max-height: 200px; min-height: 100px; overflow-y: auto"
              autosize
          />

          <van-cell title="添加分类" required style="text-align: left;">
            <van-tag type="primary" size="medium" style="margin-right: 5px" @click="addCategory">添加</van-tag>
            <van-tag type="primary"
                     size="medium"
                     v-for="(category, idx) in form.categories"
                     closeable
                     @close="deleteCategory(idx)"
                     style="margin-right: 5px"
            >
              {{ category }}
            </van-tag>
          </van-cell>
        </van-form>
      </div>

      <div style="margin: 20px">
        <van-button round block type="success" color="#42b983" @click="onSubmit">提交</van-button>
      </div>



    <van-dialog v-model:show="dialog.show" title="添加类别" @confirm="dialogConfirm">
      <div style="margin-bottom: 300px">
        <van-dropdown-menu>
          <van-dropdown-item v-model="dialog.value" :options="dialog.options" style="top: 100px"/>
        </van-dropdown-menu>
      </div>

    </van-dialog>
</template>

<script>
import { ref, reactive } from 'vue';
import {Dialog, showFailToast, showSuccessToast, Toast} from "vant";
import {getCategoryData} from "@/network/category";
import {uploadImage} from "@/network/image";
import {updateUser} from "@/network/user";
import {publishProduct} from "@/network/publish";
import PriceDisplay from "@/components/content/goods/PriceDisplay.vue";

export default {
  components: {
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

        const afterRead = (file) => {
          let formData = new FormData();
          file.status = 'uploading';
          file.message = '上传中';
          formData.append('file', file.file);
          console.log(formData)
          uploadImage(formData).then((response) => {
            form.imageURIs.push(response);
            file.status = '';
            file.message = '';
          }).catch((err) => {
            file.status = 'failed';
            file.message = '上传失败';
            form.imageURIs.push(err);
            console.log(err);
          })
        };

        const onDelete = (index) => {
            // 删除指定索引的图片
            form.imageURIs.splice(index, 1);
        };

        const onSubmit = () => {
            if(form.name === '' || form.imageURIs === [] || form.categories === [] || form.price === [] ||
                form.description === '') {
              showFailToast('请填写全部字段')
              return;
            }
            publishProduct({
              title: form.name,
              content: form.description,
              categories: form.categories,
              imageURIs: form.imageURIs,
              price: parseFloat(form.price)
            }).then((response) => {
              showSuccessToast("发布成功")
            }).catch((err) => {
              showFailToast('发布失败')
            })
        };

        const deleteCategory = (idx) => {
            form.categories.splice(idx, 1);
        }

        const addCategory = () => {
            console.log(dialog.options)
            dialog.show = true;
        }

        const dialogConfirm = () => {
            if(!form.categories.includes(dialog.value)) {
              form.categories.push(dialog.value);
            }
        }

        return {
            form,
            afterRead,
            onDelete,
            onSubmit,
            deleteCategory,
            addCategory,
            dialog,
            dialogConfirm,
        };
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
div van-form{
  width: 100%;
}
</style>