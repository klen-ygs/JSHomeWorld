<template>
  <div>
        <HeadShow></HeadShow>
        <div id="box"></div>
        <div id="ShopBack">
            <div id="ShopImagesDiv" ref="ShopImage" > 
                <video v-show="showVideo" ref="shopVideo" id="shopVideo" controls="controls" preload=" metadata" width="85%">
                </video>
                <button ref="ToLast" id="toLastImageButton" @click="toLastImage()"><img  id="leftForward" src="../assets/forward.png" width="15px"></button>
                <button ref="ToNext" id="toNextImageButton" @click="toNextImage()"><img src="../assets/forward.png" width="15px"></button>
            </div>
            <div id="shopInfo">
                <div style="margin-top: 10px; margin-left:40px; font-size: 20px; color: black; height: 60px; width: 85%; background-color: rgb(184, 184, 184); padding: 10px 0px 0px 10px ;">
                    {{TitleText}}
                </div>
                <div style="margin-top:15px; margin-left:50px; font-size: 25px; color: brown;">
                    ￥{{Price}}
                </div>
                <div id="moreInfo" ref="moreInfo" @mouseenter="mouseEnterMoreInfo()" @mouseleave="mouseLeaveMoreInfo()">
                    详情
                    <div v-show="moreInfoShow" style="width:500px; height: 200px; position: absolute; background-color: rgb(248, 248, 248); top: 300px; box-shadow: 0 0 4px gray;border-radius: 5px;">
                        {{Text}}
                    </div>
                </div>
                <div style="margin-top: 20px; min-height: 30px;">
                    <span style=" color:gray; margin-left: 30px; font-size: 14px;">发货：</span>
                    {{Addr}}
                </div>
                <div id="chooseList">
                    <span style="margin-left: 30px; color: gray; font-size: 15px;">
                        {{chooseTip}}:
                        <button class="TipButtons" ref="chooseTips" v-for="(name, index) in chooseTipList" :key="name" @click="chooseATip(name, index)">
                            {{name}}
                        </button>
                    </span>
                </div>
                <div style="height:50px; margin-top: 30px; width: 100%; float: left;"> 
                    <span style="margin-left:25px; font-size:15px;color: gray;">数量:</span>
                    <button id="shopSubButton" ref="shopSubButton" @click="subPayNum()">-</button>
                    <input type="text" v-model.number="PayNum" style="width: 60px; height: 26px; margin-left: 5px; margin-right: 5px; outline: none;background-color: rgb(242, 242, 242); border: 1px solid gray; text-align: center; border-radius: 3px; font-size: 15px;">
                    <button id="shopAddButton" ref="shopAddButton" @click="addPayNum()" >+</button>
                </div>
                <div >
                    <span style="color:gray; margin-left: 25px; font-size: 14px">库存：</span>
                    <span style="margin-left: 20px;">{{Having + "件"}}</span>
                </div>
                <div style="width: 80%; height: 80px; margin-top: 20px; position: relative;">
                    <button id="payNowButton" @click="PayNow()">立即购买</button>
                    <button id="addListButton" @click="addToList()">加入购物车</button>
                    <div v-show="addListACShow"  ref="AddListACTip" style="position: absolute; right: 190px; top: -28px;; height: 25px; width: 120px; border-radius: 5px; box-shadow: 0 0 4px gray; background-color: white; overflow: hidden;">
                        <img src="../../public/static/点赞.png" width="20px">
                        <span>
                            收藏成功
                        </span>
                    </div>
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
            document.title = this.TitleText
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
                console.log("666")
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
                    if (this.imageCount > 1) {
                        this.$refs.ToNext.disabled = ""
                    }
                    this.$refs.ToLast.disabled = "disabled"
                    if (this.hasVideo == false) {
                        this.toNextImage()
                    }
                })
                this.$refs.shopSubButton.disabled = "disabled"
                axios.get("http://127.0.0.1:8000/shopChoose",{
                    params: {
                        ShopId: this.ShopId,
                    }
                }).then( res => {
                    this.chooseTip = res.data.ChooseTip
                    if (this.chooseTip == "") {
                        return
                    }
                    this.ChooseTip = res.data.ChooseTip
                    this.chooseTipList = JSON.parse(res.data.ChooseTipList)
                } )
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
                    Id: this.imageIdx ,
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
            if (this.imageIdx == 0) {
                this.$refs.ToLast.disabled = "disabled"
                this.$refs.ShopImage.style["background-image"] = ""
                if (this.hasVideo == true) {
                    this.showVideo = true
                }
                return
            }
            if (this.imageIdx == 1 && this.hasVideo == false) {
                this.$refs.ToLast.disabled = "disabled"
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
                list.push({
                    ShopId: this.ShopId,
                    Choise: this.choise,
                    PayNum: this.PayNum,
                })
                console.log(JSON.stringify(list))
                axios({
                    method:"Post",
                    url: "http://127.0.0.1:8000/addTip",
                    data: {
                        List:  JSON.stringify(list)
                    },
                    withCredentials:true,
                })
            } ).then(()=> {
                this.addListAC()
            })
        },
        addListAC() {
            this.addListACShow = true
            setTimeout(()=> {
                this.addListACShow = false
            }, 800)
        },
        mouseEnterMoreInfo() {
            this.$refs.moreInfo.style["color"] = "rgb(238, 91, 91)"
            this.$refs.moreInfo.style["border"] = "1px solid rgb(238, 91, 91)"
            this.$refs.moreInfo.style["background-color"] = "rgb(245, 245, 245)"
            this.moreInfoShow = true

        },
        mouseLeaveMoreInfo() {
            this.$refs.moreInfo.style["color"] = ""
            this.$refs.moreInfo.style["border"] = ""
            this.$refs.moreInfo.style["background-color"] = ""
            this.moreInfoShow = false

        },
        PayNow() {
            this.showPay = true
            this.$refs.paywidget.showThisPay()
        },
        closePay() {
            this.showPay = false
        },
        subPayNum() {
            this.PayNum--
            this.$refs.shopSubButton.disabled = ""
            if (this.PayNum == 0) {
                this.$refs.shopSubButton.disabled = "disabled"
            }
        },
        addPayNum() {
            this.PayNum++
            this.$refs.shopSubButton.disabled = ""
            if (this.PayNum >= this.Having) {
                this.$refs.shopAddButton.disabled = "disabled"
            }
        },
        chooseATip(name, index) {
            this.choise = name
            console.log(name)
            for (let a = 0; a < this.$refs.chooseTips.length; a++) {
                if (a != index) {
                    this.$refs.chooseTips[a].style["border"] = ""
                    this.$refs.chooseTips[a].style["color"] = ""
                    this.$refs.chooseTips[a].style["background-color"] = ""
                } else {
                    this.$refs.chooseTips[a].style["border"] = "1px solid orange"
                    this.$refs.chooseTips[a].style["color"] = "orange"
                    this.$refs.chooseTips[a].style["background-color"] = "rgb(240, 236, 229)"
                }
            }
        },
    },
    data() {
        return{
            imageCount: 0,
            ShopId: 0,
            imageUrls:[],
            imageIdx:0,
            showVideo: false,
            hasVideo : false,
            Text: "",
            Price: 0,
            Having: 0,
            TitleText: "",
            showPay: false,
            Addr:"衡阳",
            PayNum: 0,
            chooseTip:"无",
            chooseTipList: [],
            choise:"",
            moreInfoShow: false,
            addListACShow: false 
        }
    },
}
</script>

<style>

d{
    background-color: rgb(248, 248, 248)
}

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
    height: 45px;
    cursor: pointer;
    background-color: rgb(54, 144, 209);
    background-image: linear-gradient(to right, rgb(37, 106, 244), rgb(67, 170, 243));
    box-shadow: 2px 0 4px rgb(37, 106, 244);
    border-radius: 25px;
    border-top-left-radius: 0px;
    border-bottom-left-radius: 0px;
    border: 0px;
}

#addListButton:hover {
    width: 100px;
    height: 45px;
    cursor: pointer;
    background-color: rgb(69, 163, 217);
    background-image: linear-gradient(to right, rgb(15, 91, 242), rgb(25, 153, 244));
    box-shadow: 2px 0 4px rgb(37, 106, 244);
    border-radius: 25px;
    border-top-left-radius: 0px;
    border-bottom-left-radius: 0px;
    border: 0px;
}


#ShopImagesDiv {
    display: inline-block;
    position: relative;
    top: 10px;
    margin-left: 2%;
    margin-top: 2%;
    width: 45%;
    height: 500px;
    border-radius: 10px;
    background-color: rgb(238, 238, 238);
    background-repeat: no-repeat;
    background-position-x: center;
    background-position-y: center;
}

#toLastImageButton {
    position: absolute;
   height: 50px;
   width: 30px;
   margin-top: 35%;
   margin-left: 5px;
}

#toNextImageButton {
    position: absolute;
   right: 5px;
   margin-top: 35%;
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
    margin-left: 8%;
    margin-top: 15%;
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
    height: 45px;
    cursor: pointer;
    background-image: linear-gradient(to right ,rgb(243, 128, 75), rgb(240, 44, 44));
    box-shadow: 0 0 8px rgb(239, 121, 89);
    border-radius: 25px;
    border-top-right-radius: 0px;
    border-bottom-right-radius: 0px;
    border: 0px;
    margin-left: 20%;
}

#payNowButton:hover {
    width: 100px;
    height: 45px;
    cursor: pointer;
    background-image: linear-gradient(to right ,rgb(245, 108, 44), red);
    border-radius: 25px;
    border-top-right-radius: 0px;
    border-bottom-right-radius: 0px;
    border: 0px;
    margin-left: 20%;
}

#shopInfo {
    float: right;
    margin-right: 5%;
    margin-left: 3%;
    margin-top: 3%;
    width: 45%;
    height: 500px;
    background-color: rgb(251, 251, 251);
}

#moreInfo {
    cursor: pointer;
    margin-left: 8%;
    margin-top: 20px;
    width: 80px;
    height: 30px;
    border-radius: 25px;
    border: 1px solid gray;
    text-align: center;
    padding-top: 5px;
}

#chooseList {
    width: 100%;
    height: 150px;
    padding-top: 10px;
}

#shopSubButton {
    cursor: pointer;
    margin-left: 20px;
    width:30px; height:30px; 
    border-top-left-radius: 15px; 
    border-bottom-left-radius: 15px;
    background-color: rgb(242, 242, 242);
    border: 1px solid gray;
}

#shopSubButton:hover {
    margin-left: 20px;
    width:30px; height:30px; 
    border-top-left-radius: 15px; 
    border-bottom-left-radius: 15px;
    background-color: rgb(232, 232, 232);
    border: 1px solid gray;
}

#shopAddButton {
    cursor: pointer;
    width:30px; height:30px; 
    border-top-right-radius: 15px; 
    border-bottom-right-radius: 15px;
    background-color: rgb(242, 242, 242);
    border: 1px solid gray;
}

#shopAddButton:hover {
    cursor: pointer;
    width:30px; height:30px; 
    border-top-right-radius: 15px; 
    border-bottom-right-radius: 15px;
    background-color: rgb(232, 232, 232);
    border: 1px solid gray;
}

.TipButtons {
    cursor: pointer;
    height: 40px;
    border-radius: 10px;
    min-width: 40px;
    background-color: rgb(236, 236, 236);
    border: 1px solid rgb(236, 236, 236);
    margin-left: 8px;
    margin-right: 8px;
}

.TipButtons:hover {
    cursor: pointer;
    height: 40px;
    border-radius: 10px;
    min-width: 40px;
    background-color: rgb(221, 218, 218);
    border: 1px solid rgb(221, 218, 218);
    margin-left: 8px;
    margin-right: 8px;
}

</style>