<template>
  <div id="ShopWindowsBack" @click="JumpToShop()">
      <div id="imageDiv"> 
          <img ref="ShopImage" id="ShopImg" width="90%" />
      </div>
      <div id="info">
        <p id="titleText">{{ ShopTitleText }}</p>
        <h2 id="price">
            ￥{{Price}}
        </h2>
        <h4 id="shopName">
            {{ Name }}
        </h4>
      </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
    methods:{
        JumpToShop() {
            let url = this.$router.resolve({path:`/Shop?`,query:{ShopId: this.Id}})
            window.open(url.href,"_blank")
        }
    },
    data() {
        return{
           
        }
    },
    mounted() {
        axios({
            method:"Get",
            url:"http://127.0.0.1:8000/getImage",
            params:{
                ShopId: this.Id
            },
            responseType:"blob"
        }).then(res => {
            this.$refs.ShopImage.src = URL.createObjectURL(new Blob([res.data],{type:"image/*"}))
        })
    },
    props:{
        Price:Number,
        Id: Number,
        ShopTitleText: String,
        Name: String,
        getPhone: Function,
    }
}
</script>

<style>

#ShopWindowsBack {
    margin-left: 0.5%;
    width: 19%;
    display: inline-block;
    height: 30%;
    background-color: rgb(244, 250, 253);
    box-shadow: 0px 0px 2px;
    border-radius: 2px;
    border: 1px solid rgb(239, 232, 232);
}

#info {
    margin-left: 6%;
}

#titleText {
    color: rgb(146, 136, 136);
    font-size: 10;
    overflow: hidden;
    max-height: 20px;
}

#ShopWindowsBack:hover {
    margin-left: 0.5%;
    width: 19%;
    display: inline-block;
    height: 30%;
    background-color: rgb(244, 250, 253);
    box-shadow: 0px 0px 2px;
    border-radius: 2px;
    border: 1px solid rgb(204, 81, 81);
}

#shopImg {
    padding-left: 20px;
    width: 100px;
    margin-left: 5%;
    height: 50%;
    margin-bottom: 10%;
}

#imageDiv {
    margin-top: 5%;
    margin-left: 6%;
    height: 60%;
    margin-bottom: 2%;
}

#price {
    color: rgb(195, 58, 58);

}

#shopName {
    color: rgb(131, 123, 123);
    font-family: 'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
}

</style>