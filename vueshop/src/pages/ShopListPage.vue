<template>
  <div>
      <HeadShow/>
      <div style="height:150px"></div>
      <div id="pay">
        <div style=" height: 25px"></div>
        <span style="font-size: 20px;margin-left: 10%;height: 20px;">
          总计：
        </span>
        <span style="font-size: 20px">
          ￥{{$store.state.AllPrice}}
        </span>
        <span style="margin-left: 30px">
          <button id="payButton" ref="payButton" @mouseenter="moveInPayButton()" @mouseleave="leavePayButton()" @click="openPay()">支付</button>
        </span>
      </div>
      <div id="ListDiv">
        <ShopTip v-for="(shop, index) in shopLists" :key="Number(shop.ShopId)" :Phone="Phone" :ShopId="Number(shop.ShopId)" :Index="index" :PayNum="Number(shop.PayNum)" :Choise="shop.Choise"  />
      </div>
      <div v-show="showMessage" ref="listMessage" id="listMessage">
        删除成功
      </div>
      <PayWidget v-show="showPay" ref="Pay" :closeEvet="PayClose"></PayWidget>
  </div>
</template>

<script>
import axios from 'axios'


import ShopTip from '../components/ShopTip'
import HeadShow from '../components/HeadShow.vue'
import PayWidget from '@/components/PayWidget.vue'
export default {
  name:"ShopListPage",
  mounted() {
    this.Phone = this.$route.params.Phone
    axios.get("http://127.0.0.1:8000/addTip",
      {
        params:{
          Phone : this.Phone
        },
        withCredentials:true
      }
      ).then(res => {
            this.shopLists = JSON.parse(res.data.List)
            console.log(this.shopLists)
          })

  },
  data() {
    return {
      shopLists:[],
      Phone:"",
      showMessage: false,
      showPay: false,
    }

  },
  methods: {
    showMessageFunction() {
      this.showMessage = true
      let height = 50
      let count = 0
      let _this = this
      let v = 1
      setTimeout(function show() {
          _this.$refs.listMessage.style["top"] = `${height}px`
          height += v
          v -= 0.02
          count++ 
          if (count < 50) {
            setTimeout(show, 16.7)
          } else {
            setTimeout( ()=> {
              _this.showMessage = false
            },250)
          }
        },16.7)
    },
    openPay() {
      this.showPay = true
      this.$refs.Pay.showThisPay()
    },
    PayClose() {
      this.showPay = false
    },
    moveInPayButton() {
      this.$refs.payButton.style["background-color"] = "rgb(130, 185, 255)"

    },
    leavePayButton() {
      this.$refs.payButton.style["background-color"] = "rgb(109, 165, 244)"
    }
  },
  components: {
    ShopTip,
    HeadShow,
    PayWidget,
  }
}
</script>

<style>

#ListDiv {
  margin-top: 0px;
  min-height: 800px;
  width: 90%;
  margin-left: 5%;
  background-color: rgb(245, 249, 253);
  border-radius: 5px;
  box-shadow: 0px 0px 2px gray;
}

#listMessage {
    position: fixed;
    width: 400px;
    height: 100px;
    left: 35%;
    border-radius: 10px;
    background-color: white;
    display: inline-block;
    box-shadow: 0px 0px 4px grey;
}

#pay {
  width: 500px;
  height: 80px;
  margin-left: 5%;
  box-shadow: 0 0 4px grey;
  margin-bottom: 0px;
}

#payButton {
  cursor: pointer;
  background-color: rgb(109, 165, 244);
  height: 30px;
  width: 70px;
  border-radius: 20px;
  box-shadow: 0 0 2px gray;
  outline: none;
  border: 0 black;
}


</style>