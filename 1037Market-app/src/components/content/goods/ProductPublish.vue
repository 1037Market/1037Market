<template>
  <div class="display">
    <van-form @submit="onSubmit">
      <van-field style="margin-bottom: 10px"
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
          placeholder="负数价格表示你想收该商品"
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
        >
          {{ category }}
        </van-tag>
      </van-cell>
    </van-form>
  </div>

  <div style="margin: 20px">
    <van-button round block type="success" color="#42b983" @click="onSubmit">{{ props.update? "更新":"发布" }}</van-button>
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
import {nextTick, onMounted, reactive, watchEffect} from "vue";
import {showFailToast, showSuccessToast} from "vant";
import {uploadImage} from "../../../network/image";
import {publishProduct, updateProduct} from "../../../network/publish";
import {useRouter} from "vue-router";
import {useRoute} from "vue-router/dist/vue-router";

export default {
  name: "ProductPublish",
  props: {
    form: {
      name: '',
      images: [],
      imageURIs: [],
      description: '',
      price: null,
      categories: []
    },
    dialog: {
      show: false,
      value: '添加分类',
      options: []
    },
    update: false
  },
  setup(props) {
    const form = reactive({
      ...props.form
    });
    const router = useRouter();

    const dialog = reactive({
      ...props.dialog
    });

    watchEffect(() => {
      Object.assign(form, props.form);
      Object.assign(dialog, props.dialog);
      // TODO: 已经存在的图片咋办？
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
      if (form.name === '' || form.imageURIs.length === 0 || form.categories.length === 0 || typeof(form.price) !== 'number' ||
          form.description === '') {
        showFailToast('请填写全部字段')
        return;
      }
      if(props.update) { // update existing product
        updateProduct({
          title: form.name,
          content: form.description,
          categories: form.categories,
          imageURIs: form.imageURIs,
          price: parseFloat(form.price)
        }).then((response) => {
          showSuccessToast("更新成功")
          router.go(-1)
        }).catch((err) => {
          showFailToast('更新失败')
        })
      } else { // publish new product
        publishProduct({
          title: form.name,
          content: form.description,
          categories: form.categories,
          imageURIs: form.imageURIs,
          price: parseFloat(form.price)
        }).then((response) => {
          showSuccessToast("发布成功")
          router.go(-1)
        }).catch((err) => {
          console.log(err)
          showFailToast('发布失败')
        })
      }

    };

    const deleteCategory = (idx) => {
      form.categories.splice(idx, 1);
    }

    const addCategory = () => {
      console.log(dialog.options)
      dialog.show = true;
    }

    const dialogConfirm = () => {
      if (!form.categories.includes(dialog.value)) {
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
      props
    };

  }
}
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
</style>