<template>
  <div id="HeadBack">
      <div ref="" class="headTip" @click="toFirstPage()">
          首页
      </div>
      <div id="ShopCar" ref="ShopCar" @click="toList()" @mouseenter="enterLists()" @mouseleave="leaveLists()">
      </div>
      <div id="UserDiv">
        欢迎您 <p id="tag" >{{ Name }}</p> {{"  "}} 手机号： <p id="tag">{{ Phone }}</p>
      </div>
      <button @click="logout()" id="logoutButton" >注销</button>
      <LittleShopList v-show="ShowList" ref="littleList"/>
  </div>
</template>

<script>
import axios from 'axios'
import LittleShopList from './LittleShopList.vue'
export default {
    mounted() {
        axios.get( "http://127.0.0.1:8000/getUserInfo",
            {
                withCredentials: true
            }
        ).then(res => {
            this.Name = res.data.Name
            this.Phone = res.data.Phone
        })
        this.$refs.littleList.addMouseEnter(()=> {
            this.enterList = true
        })
        this.$refs.littleList.addMouseLeave( ()=> {
            this.leaveLists()
        } )
    },
    data() {
        return{ 
            headSrc:"",
            Name:"",
            Phone:"",
            ShowList: false,
            enterList: false,
            bigChaning: false,
            smallChaning: false,
        }
    },
    methods: {
        toList() {
           let url = this.$router.resolve({name:"ShopListPage"})
            window.open(url.href, "_blank")

        },
        logout() {
            axios({
                url:"http://127.0.0.1:8000/login",
                method:"Delete",
                data:{
                    Phone: this.Phone
                },
                withCredentials:true,
            }).then( () => {
                this.$router.replace({
                name:"loginPage",
            })
            })
        },
        enterLists() {
            this.enterList = true
            let _this = this
            setTimeout(() => {
                console.log(this.bigChaning)
                if (this.enterList == true && this.bigChaning == false) {
                    this.bigChaning = true
                    this.ShowList = true
                    this.$refs.littleList.show()
                    let size = 30
                    let margin_right = 20
                    let ShopCar = this.$refs.ShopCar
                    this.smallChaning = false
                    setTimeout( function shopCarBig() {
                        size++
                        margin_right--
                        if (size <= 40 && _this.enterList == true) {
                            ShopCar.style["width"] = String(size) + "px"
                            ShopCar.style["height"] = String(size) + "px"
                            ShopCar.style["margin-right"] = String(margin_right) + "px"
                            setTimeout(shopCarBig, 5)
                        }
                    }, 5)
                }
            },400)

        },
        leaveLists() {
            this.enterList = false
            setTimeout( ()=> {
                if (this.enterList == false && this.smallChaning == false) {
                    this.smallChaning = true
                    this.ShowList = false
                    let _this = this
                    let size = 40
                    let margin_right = 10
                    let ShopCar = this.$refs.ShopCar
                    this.bigChaning = false
                    setTimeout( function shopCarSamll() {
                        size--
                        margin_right++
                        if (size >= 30 && _this.enterList == false) {
                            ShopCar.style["width"] = String(size) + "px"
                            ShopCar.style["height"] = String(size) + "px"
                            ShopCar.style["margin-right"] = String(margin_right) + "px"
                            setTimeout(shopCarSamll, 5)
                        }
                    }, 5)
                }
            },500 )
        },
        toFirstPage() {
            this.$router.push({
                path:"/select"
            })
        }
    },
    components: {
        LittleShopList,
    }
}
</script>

<style>

#HeadImage {
    border-radius: 100px;
}

#HeadBack {
    box-shadow: 0px 0px 3px gray;
    top: 0;
    margin-top: 0;
    position: fixed;
    width: 100%;
    height: 70px;
    background-color: rgb(221, 234, 245);
    z-index: 99;
}

#tag {
    display: inline-block;
    color: rgb(186, 70, 46);
    font-family:'Segoe UI', Tahoma, Geneva, Verdana, sans-serif
}

#UserDiv {
    display: inline-block;
    vertical-align: middle;
    margin-top:15px;
    margin-left: 20px;
}

#ShopCar {
    margin-left: 64%;
    margin-right: 20px;
    cursor: pointer;
    margin-top: 20px;
    vertical-align: middle;
    width: 30px;
    height: 30px;
    float: left;
    background-image: url(../assets/shopcar.png);
    background-size: 100% 100%;
    background-repeat: no-repeat;
}

#logoutButton {
    margin-left: 20px;
    cursor: pointer;
    vertical-align: middle;
    margin-top: 20px;
    width: 55px;
    height: 25px;
    border-radius: 15px;
    background-color: rgb(240, 107, 107);
    border: 0px;
    box-shadow: 0 0 2px rgb(239, 56, 56);
}

#logoutButton:hover {
    margin-left: 20px;
    vertical-align: middle;
    margin-top: 20px;
    width: 55px;
    height: 25px;
    border-radius: 15px;
    background-color: rgb(240, 107, 107);
    border: 0px;
    box-shadow: 0 0 2px rgb(239, 56, 56);
}

.headTip {
    height: 76%;
    width: 80px;
    padding-top: 17px;
    text-align: center;
    float:left;
    font-size: 20px; 
    margin-left: 20px; 
    cursor: pointer;
}

.headTip:hover {
    float:left;
    width: 80px;
    height: 76%;
    padding-top: 17px;
    text-align: center;
    font-size: 20px; 
    margin-left: 20px; 
    background-color: aliceblue;
    background-image: linear-gradient(rgb(136, 198, 243),rgb(165, 216, 233));
}
</style>