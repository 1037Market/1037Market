import { createRouter, createWebHistory } from 'vue-router';
import store from '../store/index'
const Home = () => import('@/views/home/Home.vue');
const Category = () => import('@/views/category/Category.vue');
const Detail = () => import('@/views/detail/Detail.vue');
const ShopCart = () => import('@/views/shopcart/ShopCart.vue');
const Profile = () => import('@/views/profile/Profile.vue');
const Register = () => import('@/views/profile/Register.vue');
const Login = () => import('@/views/profile/Login.vue')
const Publish = () => import('@/views/publish/Publish.vue')
const Seller = () => import('@/views/seller/Seller.vue');
import { showNotify } from "vant";


const routes = [
  {
    path: '',
    name: 'defaultHome',
    component: Home,
    //添加浏览器导航栏标题
    meta: {
      title: '1037集市'
    }
  },
  {
    path: '/category',
    name: 'Category',
    component: Category,
    meta: {
      title: '1037集市--分类'
    }
  },
  {
    path: '/publish',
    name: 'Publish',
    component: Publish,
    meta: {
      title: '发布新商品'
    }
  },
  {
    path: '/detail/:id',
    name: 'Detail',
    component: Detail,
    meta: {
      title: '商品详情'
    }
  },
  {
    path: '/seller/:studentId',
    name: 'Seller',
    component: Seller,
    meta: {
      title: '卖家详情'
    }
  },
  {
    path: '/shopcart',
    name: 'ShopCart',
    component: ShopCart,
    meta: {
      title: '购物车',
      isAuthRequired: true
    }
  },
  {
    path: '/profile',
    name: 'profile',
    component: Profile,
    meta: {
      title: '个人中心',
      isAuthRequired: true
    }
  },
  {
    path: '/register',
    name: 'register',
    component: Register,
    meta: {
      title: '用户注册'
    }
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
    meta: {
      title: '用户登录'
    }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})


//导航守卫
router.beforeEach((to, from, next) => {
  // 如果没有登录， 在这里到login
  if (to.meta.isAuthRequired && store.state.user.isLogin == false) {
    showNotify({message:'请先登录'})
    return next('/login')

  } else {
    next();

  }
  //浏览器导航栏标题变化
  document.title = to.meta.title
})

export default router
