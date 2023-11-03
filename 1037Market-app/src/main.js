import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import 'vant/lib/index.css';
// import 'src/assets/css/iconfont.css';

import { DropdownMenu, DropdownItem, Dialog, NavBar, Uploader, Toast, Cell, CellGroup, PullRefresh, List, Popup, Grid, GridItem, AddressEdit,
    AddressList, Icon,
    SubmitBar, CheckboxGroup, SwipeCell, Stepper, Checkbox, Field, Form, Tag, Button, Image as VanImage, Card, Tab,
    Tabs, Swipe, SwipeItem, Lazyload, Badge, Sidebar, SidebarItem, Collapse, CollapseItem, RadioGroup, Radio }
    from 'vant';




createApp(App)
  .use(Swipe).use(SwipeItem).use(Lazyload, {
    loading: require('./assets/images/default.jpg')
  }).use(Badge).use(Sidebar).use(SidebarItem).use(Collapse).use(CollapseItem)
  .use(Tab).use(Tabs).use(Card).use(VanImage).use(Tag).use(Button).use(Form).use(Field)
  .use(Checkbox).use(Stepper).use(SwipeCell).use(CheckboxGroup).use(SubmitBar).use(Icon).use(AddressList).use(AddressEdit)
  .use(GridItem).use(Grid).use(Popup).use(List).use(PullRefresh).use(store).use(router).use(Cell)
  .use(CellGroup).use(Toast).use(Uploader).use(NavBar).use(Dialog).use(DropdownMenu).use(DropdownItem).mount('#app')
