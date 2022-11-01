import LoginPage from '../pages/LoginPage'
import SelectPage from '../pages/SelectPage'
import vueRouter from 'vue-router'


export default new vueRouter({
    routes: [
        {
            path:"/",
            component:LoginPage,
        },
        {
            path:"/select",
            component:SelectPage,
        },
    ],
})