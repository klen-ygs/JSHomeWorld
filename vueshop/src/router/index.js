import LoginPage from '../pages/LoginPage'
import SelectPage from '../pages/SelectPage'
import vueRouter from 'vue-router'
import AddShopPage from '../pages/AddShopPage'
import ShopListPage from '../pages/ShopListPage'
import ShopPage from '../pages/ShopPage'
import SelectedPage from '../pages/SelectedPage'

export default new vueRouter({
    mode: "history",
    routes: [
        {
            path:"/",
            name:"loginPage",
            component:LoginPage,
        },
        {
            path:"/select",
            name: "SelectPage",
            component:SelectPage,
        },
        {
            path:"/AddShop",
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
        },
        {
            path:"/selected",
            component: SelectedPage,
        }
    ],
})