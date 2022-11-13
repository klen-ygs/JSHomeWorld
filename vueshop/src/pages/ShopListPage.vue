<template>
  <div>
      <HeadShow/>
      <div id="box"></div>
      <div id="ListDiv">
        <ShopTip v-for="tip in shopLists" :key="tip" :Phone="Phone" :ShopId="Number(tip)"/>
      </div>
      <div v-show="showMessage" ref="listMessage" id="listMessage">
        删除成功
      </div>
  </div>
</template>

<script>
import axios from 'axios'


import ShopTip from '../components/ShopTip'
import HeadShow from '../components/HeadShow.vue'
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
          })

  },
  data() {
    return {
      shopLists:[],
      Phone:"",
      showMessage: false
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
    }
    
  },
  components: {
    ShopTip,
    HeadShow

  }
}
</script>

<style>

#ListDiv {
  margin-top: 5px;
  height: 1000px;
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

</style>