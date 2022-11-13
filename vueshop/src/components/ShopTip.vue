<template>
  <div id="tip" ref="tip" @mouseenter="MouseIn()" @mouseleave="MouseOut()" @click="JumptoShop()">
      <div id="tipImageDiv">
        <img ref="tipImage" id="tipImage" height="95%" >
      </div>
        <div id="ShopTitleTextDiv">
            {{ ShopTitleText }}
        </div>
        <div id="PriceDiv">
            ￥{{ Price }}
        </div>
        <button id="deleteButton" @click.stop="deleteTip()">删除</button>
  </div>
</template>

<script>
import axios from 'axios'
export default {
    data() {
        return{
            ShopTitleText:"",
            Price:""
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
        })

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
                for (let a = 0; a < list.length; a++) {
                    if (list[a] == this.ShopId) {
                        list.splice(a, 1)
                        this.$parent.showMessageFunction()
                        break
                    }
                }
                axios({
                    withCredentials:true,
                    method:"Post",
                    url:"http://127.0.0.1:8000/addTip",
                    data: {
                        Phone: this.Phone,
                        List: JSON.stringify(list),
                    }
                }).then(() => {
                    let list = this.$parent.shopLists
                    for (let a = 0; a < list.length; a++) {
                        if (list[a] == this.ShopId) {
                            list.splice(a, 1)

                            break
                        }
                    }

                })
                
            } )

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
        }
    },
    props: {
        ShopId:Number,
        Phone: String,
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
    margin-bottom: 60px;
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
    margin-bottom: 60px;
    background-color: rgb(196, 227, 246);
    box-shadow: 0px 0px 1px;
    border: 0px ;
}

#deleteButton:hover {
    cursor: pointer;
    margin-left: 50px;
    width: 100px;
    height: 30px;
    width: 100px;
    height: 30px;
    vertical-align: middle;
    border-radius: 20px;
    margin-bottom: 60px;
    background-color: rgb(126, 183, 218);
    box-shadow: 0px 0px 1px;
}

</style>