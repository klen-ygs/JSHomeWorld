<template>
  <div id="littleTip" @click="JumpShop()">
    <span id="littleImage">
        <img ref="tipImage" height="85%">
    </span>
    <span id="LittleText">
        {{ ShopTitleText }}
    </span>
    <span id="LittlePrice">
        ￥{{ Price }}
    </span>
  </div>
</template>

<script>
import axios from 'axios'
export default {
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
    data() {
        return {
            ShopTitleText: "",
            Price: 0,
        }
    },
    props: {
        ShopId: Number

    },
    methods: {
        JumpShop() {
            let url = this.$router.resolve({path:`/Shop`,query:{ShopId: String(this.ShopId)}})
            window.open(url.href,"_blank")
        },

    },
}
</script>

<style>
#littleTip {
    height: 80px;
    width: 100%;
}

#littleTip:hover {
    height: 80px;
    width: 100%;
    background-color: rgb(186, 186, 186);
}

#LittleText {
    margin-left: 10px;
    display: inline-block;
    font-size: 10px;
    text-align: center;
    width: 30%;
    vertical-align: middle;
    margin-right: 15px;
}

#LittlePrice {
    display: inline-block;
    font-size: 15px;
    color: rgb(191, 89, 89);
}

#littleImage {
    display: inline-block;
    height: 100%;
    width: 100px;
    padding-top: 5px;
}

</style>