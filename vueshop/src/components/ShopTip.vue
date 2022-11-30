<template>
  <div id="tip" ref="tip" @mouseenter="MouseIn()" @mouseleave="MouseOut()" @click="JumptoShop()">
      <div id="tipImageDiv">
        <img ref="tipImage" id="tipImage" height="95%" >
      </div>
        <div id="ShopTitleTextDiv">
            {{ ShopTitleText }}
        </div>
        <div style="display: inline-block; vertical-align: top; margin-top: 40px; width: 50px; font-size: 20px; color: orange;">
            {{Choise}}  
        </div>
        <div style="display: inline-block; vertical-align: top; margin-top: 40px;">
            <button id="shopSubButton" ref="shopSubButton" @click.stop="subPayNum()">-</button>
            <span><input type="text" @click.stop="0" v-model.number="InputPayNum" style="width: 60px; height: 26px; margin-left: 5px; margin-right: 5px; outline: none;background-color: rgb(242, 242, 242); border: 1px solid gray; text-align: center; border-radius: 3px; font-size: 15px;"></span>
            <button id="shopAddButton" ref="shopAddButton" @click.stop="addPayNum()">+</button>
        </div>
        <div id="PriceDiv">
            ￥{{ totlePrice }}
        </div>
        <button id="deleteButton" @click.stop="deleteTip()">删除</button>
        <button v-show="changePayNum" id="submitChange" @click.stop="saveChange()">保存修改</button>
  </div>
</template>

<script>
import axios from 'axios'
export default {
    data() {
        return{
            ShopTitleText:"",
            Price:"",
            InputPayNum: 0,
            Having: 0,
            changePayNum: false,
            totlePrice: 0,
        }
    },
    watch: {
        InputPayNum(newval, oldval) {
            if (newval < 0) {
                this.InputPayNum = 0
            }
            this.$store.commit("subPrice", oldval * this.Price)
            this.$store.commit("addPrice", newval * this.Price)
            this.totlePrice = this.InputPayNum * this.Price

        }
    },
    mounted() {
        axios.get("http://127.0.0.1:8000/getImage",
        {
            params:{
                ShopId: this.ShopId,
                Flag: false,
            },
            responseType:"blob"
        }).then( res => {
            this.$refs.tipImage.src = URL.createObjectURL(new Blob([res.data],{type:"image/*"}))
        })
        axios.get("http://127.0.0.1:8000/getShop",
            {
                params: {
                    ShopId: this.ShopId
                }
            }
        ).then(res => {
            this.ShopTitleText = res.data.ShopTitleText
            this.Price = res.data.Price
            this.Having = res.data.Having
            if (this.Having < this.PayNum) {
                this.$refs.shopAddButton.disabled = "disabled"
            }
            this.totlePrice = this.InputPayNum * this.Price
            this.$store.commit("addPrice", this.totlePrice)

        })
        this.$refs.shopSubButton.disabled = "disabled"
        if (!isNaN(this.PayNum)) {
            this.InputPayNum = this.PayNum
        }
        if (this.InputPayNum > 0) {
            this.$refs.shopSubButton.disabled = ""
        }
    },
    methods: {
        deleteTip() {
            axios.get("http://127.0.0.1:8000/addTip",
            {
                data: {
                    Phone :this.Phone,
                },
                withCredentials: true
            }
            ).then( res => {
                let list = JSON.parse(res.data.List)
                list.splice(this.Index, 1)
                //this.$parent.showMessageFunction() //废弃的提示方法
                axios({
                    withCredentials:true,
                    method:"Post",
                    url:"http://127.0.0.1:8000/addTip",
                    data: {
                        Phone: this.Phone,
                        List: JSON.stringify(list),
                    }
                }).then(()=> {
                    this.DeleteListener()
                })
                
            } )
            this.$store.commit("subPrice", this.InputPayNum * this.Price)

        },
        MouseIn() {
            let r = 253
            let g = 248
            let b = 248
            for (let a = 0; a < 10; a++) {
                setTimeout(()=> {
                    this.$refs.tip.style["background-color"] = `rgb(${r--}, ${g--}, ${b--})`
                }, 30)
            }
        },
        MouseOut() {
            let r = 243
            let g = 238
            let b = 238
            for (let a = 0; a < 10; a++) {
                setTimeout(()=> {
                    this.$refs.tip.style["background-color"] = `rgb(${r++}, ${g++}, ${b++})`
                }, 50)
            }
        },
        JumptoShop() {
            let url = this.$router.resolve({path:`/Shop`,query:{ShopId: String(this.ShopId)}})
            window.open(url.href,"_blank")
        },
        addPayNum() {
            this.changePayNum = true
            this.$refs.shopSubButton.disabled = ""
            this.InputPayNum++
            this.totlePrice += this.Price
            this.$store.commit("addPrice", this.Price)
        },
        subPayNum() {
            this.changePayNum = true
            this.InputPayNum--
            if (this.InputPayNum == 0) {
                this.$refs.shopSubButton.disabled = "disabled"
            }
            this.totlePrice -= this.Price
            this.$store.commit("subPrice", this.Price)
        },
        saveChange() {
            axios.get("http://127.0.0.1:8000/addTip",
            {
                data: {
                    Phone :this.Phone,
                },
                withCredentials: true
            }
            ).then( res => {
                let list = JSON.parse(res.data.List)
                list[this.Index].PayNum = this.InputPayNum
                axios({
                    withCredentials:true,
                    method:"Post",
                    url:"http://127.0.0.1:8000/addTip",
                    data: {
                        Phone: this.Phone,
                        List: JSON.stringify(list),
                    }
                }).then(() => {
                    this.changePayNum = false
                })
                
            } )
        }
    },
    props: {
        ShopId:Number,
        Phone: String,
        PayNum: Number,
        Choise: String,
        Index: Number,
        DeleteListener: Function,
    }
}
</script>

<style>
#tip {
    cursor: pointer;
    padding-top: 20px;
    padding-bottom: 20px;
    padding-left: 20px;
    width: 99%;
    height: 100px;
    background-color: rgb(253, 248, 248);
    border: 1px solid gray;
    border-radius: 5px;
}


#ShopTitleTextDiv {
    margin-left: 10px;
    display: inline-block;
    width: 500px;
    height: 20px;
    margin-bottom: 60px;
    vertical-align: middle;
    font-size: 20px;
}

#PriceDiv {
    display: inline-block;
    width: 6%;
    color: rgb(206, 63, 63);
    margin-left: 20px;
    font-size: 30px;
    vertical-align: middle;
    margin-bottom: 70px;
}

#tipImageDiv {
    display: inline-block;
    width: 200px;
    height: 100%;
}

#deleteButton {
    margin-left: 50px;
    width: 100px;
    height: 30px;
    vertical-align: middle;
    border-radius: 20px;
    margin-bottom: 70px;
    background-color: rgb(196, 227, 246);
    box-shadow: 0px 0px 2px rgb(196, 227, 246);
    border: 0px ;
}

#deleteButton:hover {
    cursor: pointer;
    margin-left: 50px;
    width: 100px;
    height: 30px;
    vertical-align: middle;
    border-radius: 20px;
    margin-bottom: 70px;
    background-color: rgb(126, 183, 218);
    box-shadow: 0px 0px 4px rgb(126, 183, 218);
}

#submitChange {
    cursor: pointer;
    margin-left: 30px;
    width: 100px;
    height: 30px;
    vertical-align: middle;
    border-radius: 20px;
    margin-bottom: 70px;
    background-color: rgb(104, 244, 125);
    box-shadow: 0px 0px 4px rgb(104, 244, 125);
    border: 0px;
}

#submitChange:hover {
    margin-left: 30px;
    width: 100px;
    height: 30px;
    vertical-align: middle;
    border-radius: 20px;
    margin-bottom: 70px;
    background-color: rgb(64, 241, 90);
    background-color: rgb(64, 241, 90);
    border: 0px;
}

</style>