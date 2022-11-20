import Vue from "vue";
import vuex from 'vuex'


Vue.use(vuex)

export default new vuex.Store({
    state() {
        return {
            AllPrice: 0,
            ChoosePay: 0,
        }
    },
    mutations: {
        addPrice(status ,price) {
            status.AllPrice += price
        },
        resetPayPrice(status) {
            status.ChoosePay = status.AllPrice
        },
        subPrice(status, price) {
            status.AllPrice -= price
        },
        subPayPrice(status, price) {
            status.ChoosePay -= price
        }
    }
    
})