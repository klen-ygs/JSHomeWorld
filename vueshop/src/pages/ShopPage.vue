<template>
  <div>
        <HeadShow></HeadShow>
        <div id="box"></div>
        <div id="ShopBack">
            <div style="height: 40px"></div>
            <span style="font-size: 35px;margin-left: 100px;color: rgb(225, 96, 73);">{{TitleText}}</span>
            <div id="ShopImagesDiv" ref="ShopImage" > 
                <video v-show="showVideo" ref="shopVideo" id="shopVideo" controls="controls" preload=" auto" width="870px">
                </video>
                <button ref="ToLast" id="toLastImageButton" @click="toLastImage()"><img  id="leftForward" src="../assets/forward.png" width="15px"></button>
                <button ref="ToNext" id="toNextImageButton" @click="toNextImage()"><img src="../assets/forward.png" width="15px"></button>
            </div>
            <div id="infoDiv">
                <div id="shopTextDiv">
                {{Text}}
            </div>
            <div id="ShopPriceDiv">
                <label>价格:</label>
                <label style="font-size: 20px;color: rgb(233, 80, 25); margin-left: 50px;">￥{{Price}}</label>
            </div>
            <div id="HavingSpan">
                <label>库存:</label>
                <label>{{Having}}</label>
            </div>
            <div style="margin-top: 30px" >
                <button id="payNowButton" @click="PayNow()">立即购买</button>
                <button id="addListButton" @click="addToList()">加入购物车</button>
            </div>
            </div>
        </div>
        <pay-widget ref="paywidget" v-show="showPay" :closeEvet="closePay"></pay-widget>
  </div>
</template>

<script>
import axios from 'axios'
import HeadShow from '../components/HeadShow.vue'
import PayWidget from '../components/PayWidget.vue'

export default {
  components: { HeadShow, PayWidget },
    mounted() {
        this.$refs.ToLast.disabled = "disabled"
        this.$refs.ToNext.disabled = "disabled"
        this.ShopId = this.$route.query.ShopId
        axios.get("http://127.0.0.1:8000/getShop",
        {
            params: {
                ShopId: this.ShopId,
            }
        }
        ).then( res => {
            this.Text = res.data.Text
            this.Price = res.data.Price
            this.Having = res.data.Having
            this.TitleText = res.data.ShopTitleText
        })

        axios.get("http://127.0.0.1:8000/video/findVideo",
        {
            params: {
                ShopId: this.ShopId,
            }
        }
        ).then( res=> {
            if (res.data.Request == false) {
                this.showVideo = false
                this.hasVideo = false
                return
            }
            this.showVideo = true
            this.imageIdx = 0
            this.hasVideo = true
            this.$refs.shopVideo.src = "http://127.0.0.1:8000/video/getVideo?path=" + res.data.Path
        } )
        axios.get("http://127.0.0.1:8000/getImages",
                    {
                        params: {
                            ShopId: this.ShopId,
                        }
                    }
                ).then(res => {
                    this.imageCount = res.data.count
                    if (this.imageCount > 0) {
                        axios({
                            url:"http://127.0.0.1:8000/getImages",
                            method:"Get",
                            params: {
                                ShopId: this.ShopId,
                                Id: 1,
                            },
                            responseType:"blob"
                        }).then( res=> {
                            this.imageUrls.push(URL.createObjectURL(new Blob([res.data],{type:"image/*"})))
                            this.$refs.ShopImage.style["background-image"] = `url(${this.imageUrls[0]})`
                            if (this.imageCount > 1) {
                                this.$refs.ToNext.disabled = ""
                            }
                            this.$refs.ToLast.disabled = "disabled"
                        } )
                    }
                })
    },
    methods:{
        toNextImage() {
            this.imageIdx++
            if (this.imageIdx == 1) {
                this.showVideo = false
            }
            this.$refs.ToLast.disabled = ""
            if (this.imageIdx == this.imageCount) {
                this.$refs.ToNext.disabled = "disabled"
            }
            if (this.imageUrls.length >= this.imageIdx) {
                this.$refs.ShopImage.style["background-image"] = `url(${this.imageUrls[this.imageIdx - 1]})`
                return
            }
            axios({
                url:"http://127.0.0.1:8000/getImages",
                method:"Get",
                params: {
                    ShopId : this.ShopId,
                    Id: this.imageIdx,
                },
                responseType:"blob",
            }).then( res=> {
                this.imageUrls.push(URL.createObjectURL(new Blob([res.data],{type:"image/*"})))
                this.$refs.ShopImage.style["background-image"] = `url(${this.imageUrls[this.imageIdx - 1]})`
            })
        },
        toLastImage(){
            this.imageIdx--
            if (this.imageCount > this.imageIdx) {
                this.$refs.ToNext.disabled = ""
            }
            if ((this.imageIdx == 1 && this.hasVideo == false) || (this.imageIdx == 0)) {
                this.$refs.ToLast.disabled = "disabled"
                if (this.hasVideo == true) {
                    this.showVideo = true
                }
                return
            }
            this.$refs.ShopImage.style["background-image"] = `url(${this.imageUrls[this.imageIdx - 1]})`
        },
        addToList() {
            axios.get("http://127.0.0.1:8000/addTip",
            {
                params :{
                },
                withCredentials:true
            }
            ).then( res => {
                console.log(res.data)
                let list = JSON.parse(res.data.List)
                list.push(this.ShopId)
                console.log(JSON.stringify(list))
                axios({
                    method:"Post",
                    url: "http://127.0.0.1:8000/addTip",
                    data: {
                        List:  JSON.stringify(list)
                    },
                    withCredentials:true,
                })
            } )


        },
        PayNow() {
            this.showPay = true
            this.$refs.paywidget.showThisPay()
        },
        closePay() {
            this.showPay = false
        }
    },
    data() {
        return{
            imageCount: 0,
            ShopId: 0,
            imageUrls:[],
            imageIdx:1,
            showVideo: false,
            hasVideo : false,
            Text: "",
            Price: 0,
            Having: 0,
            TitleText: "",
            showPay: false
        }
    },
}
</script>

<style>

#ShopBack {
    width: 90%;
    margin-left: 5%;
    height: 1000px;
    background-color: rgb(252, 252, 252);
    box-shadow: 0px 0px 2px;
    border-radius: 5px;
}

#addListButton {
    width: 100px;
    height: 30px;
    cursor: pointer;
    background-color: rgb(54, 144, 209);
    border-radius: 15px;
    border: 0px;
    margin-left: 30%;
}

#addListButton:hover {
    width: 100px;
    height: 30px;
    cursor: pointer;
    background-color: rgb(69, 163, 217);
    border-radius: 15px;
    border: 0px;
    margin-left: 30%;
}


#ShopImagesDiv {
    position: relative;
    top: 10px;
    margin-left: 13%;
    margin-top: 2%;
    width: 70%;
    height: 500px;
    border-radius: 10px;
    background-color: rgb(224, 237, 236);
    background-repeat: no-repeat;
    background-position-x: center;
    background-position-y: center;
}

#toLastImageButton {
    position: absolute;
   height: 50px;
   width: 30px;
   margin-top: 22%;
   margin-left: 5px;
}

#toNextImageButton {
    position: absolute;
   right: 5px;
   margin-top: 22%;
   background-repeat: no-repeat;
   height: 50px;
   width: 30px;
}

#leftForward {
    transform: rotate(180deg);
}
 
#HavingSpan {
    height: 30px;
    margin-top: 0px;
    margin-left: 10%;
    margin-bottom: 20px;
    background-color: rgb(169, 172, 174);
    width: 100px;
}

#shopVideo {
    position: absolute;
    margin-left: 4%;
    margin-top: 5px;
    border: 0 0 0 black;
}

#shopTextDiv {
    margin-top: 20px;
    margin-left: 10%;
    width: 800px;
    height: 100px;
    background-color: rgb(234, 240, 245);
    border-radius: 10px;
    box-shadow: 0 0 2px gray;
}


#ShopPriceDiv {
    margin-top: 5px;
    border-radius: 5px;
    margin-left: 10%;
    height: 30px;
    width: 400px;
    background-color: rgb(196, 192, 192);
    color: rgb(106, 108, 110);
    font-size: 12px;
    margin-bottom: 5px;
}

#infoDiv {
    margin-top: 20px;
    padding-top: 10px;
    background-color: rgb(225, 233, 240);
    box-shadow: 0 0 2px gray;
    width: 950px;
    margin-left: 13%;
    height: 300px;
}

#payNowButton {
    width: 100px;
    height: 30px;
    cursor: pointer;
    background-color: rgb(225, 126, 109);
    border-radius: 15px;
    border: 0px;
    margin-left: 20%;
}

#payNowButton:hover {
    width: 100px;
    height: 30px;
    cursor: pointer;
    background-color: rgb(217, 92, 69);
    border-radius: 15px;
    border: 0px;
    margin-left: 20%;
}

</style>