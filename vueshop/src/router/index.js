import LoginPage from '../pages/LoginPage'
import SelectPage from '../pages/SelectPage'
import vueRouter from 'vue-router'
import AddShopPage from '../pages/AddShopPage'
import ShopListPage from '../pages/ShopListPage'
import ShopPage from '../pages/ShopPage'

export default new vueRouter({
    mode: "history",
    routes: [
        {
            path:"/",
            component:LoginPage,
        },
        {
            path:"/select",
            name: "SelectPage",
            component:SelectPage,
        },
        {
            path:"/addShop",
            component:AddShopPage,
        },
        {
            name:"ShopListPage",
            path:"/ShopList",
            component:ShopListPage,
        },
        {
            path:"/Shop",
            component: ShopPage,
        }
    ],
})