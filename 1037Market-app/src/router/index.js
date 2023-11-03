import { createRouter, createWebHistory } from 'vue-router';
import store from '../store/index'
const Home = () => import('views/home/Home');
const Category = () => import('views/category/Category');
const Detail = () => import('views/detail/Detail');
const ShopCart = () => import('views/shopcart/ShopCart');
const Profile = () => import('views/profile/Profile');
const Register = () => import('views/profile/Register');
const Login = () => import('views/profile/Login')
const Publish = () => import('views/publish/Publish')
import { Notify } from "vant";


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
  history: createWebHistory(process.env.BASE_URL),
  routes
})


//导航守卫
router.beforeEach((to, from, next) => {
  // 如果没有登录， 在这里到login
  if (to.meta.isAuthRequired && store.state.user.isLogin == false) {
    Notify('请先登录')
    return next('/login')

  } else {
    next();

  }
  //浏览器导航栏标题变化
  document.title = to.meta.title
})

export default router
