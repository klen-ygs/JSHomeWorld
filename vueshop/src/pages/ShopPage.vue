<template>
  <div>
        <HeadShow></HeadShow>
        <div id="box"></div>
        <div id="ShopBack">
            <div id="ShopImagesDiv" ref="ShopImage">
                <button ref="ToLast" id="toLastImageButton" @click="toLastImage()"><img  id="leftForward" src="../assets/forward.png" width="15px"></button>
                <button ref="ToNext" id="toNextImageButton" @click="toNextImage()"><img src="../assets/forward.png" width="15px"></button>
            </div>
            <div id="shopTextDiv">

            </div>
            <button id="addListButton" @click="addToList()">加入购物车</button>
        </div>
  </div>
</template>

<script>
import axios from 'axios'
import HeadShow from '../components/HeadShow.vue'
export default {
  components: { HeadShow },
    mounted() {
        this.$refs.ToLast.disabled = "disabled"
        this.$refs.ToNext.disabled = "disabled"
        this.ShopId = this.$route.query.ShopId
        console.log(this.ShopId)
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
            if (this.imageIdx == this.imageCount) {
                this.$refs.ToLast.disabled = ""
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
            if (this.imageIdx == 1) {
                this.$refs.ToLast.disabled = "disabled"
                this.$refs.ToNext.disabled = ""
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


        }
    },
    data() {
        return{
            imageCount: 0,
            ShopId: 0,
            imageUrls:[],
            imageIdx:1,
        }
    },
}
</script>

<style>

#ShopBack {
    width: 90%;
    margin-left: 5%;
    height: 1000px;
    background-color: rgb(246, 250, 255);
    box-shadow: 0px 0px 2px;
    border-radius: 5px;
}

#addListButton {
    width: 100px;
    height: 30px;
    cursor: pointer;
    background-color: rgb(225, 126, 109);
    border-radius: 15px;
    border: 0px;
    margin-left: 72%;
}

#addListButton:hover {
    width: 100px;
    height: 30px;
    cursor: pointer;
    background-color: rgb(217, 92, 69);
    border-radius: 15px;
    border: 0px;
    margin-left: 72%;
}


#ShopImagesDiv {
    margin-top: 10px;
    margin-left: 13%;
    margin-top: 5%;
    width: 70%;
    height: 500px;
    border-radius: 10px;
    background-color: rgb(224, 237, 236);
    background-repeat: no-repeat;
    background-position-x: center;
    background-position-y: center;
}

#toLastImageButton {
   height: 50px;
   width: 30px;
   margin-top: 22%;
   margin-left: 5px;
}

#toNextImageButton {
   margin-left: 92%;
   background-repeat: no-repeat;
   height: 50px;
   width: 30px;
}

#leftForward {
    transform: rotate(180deg);
}

</style>