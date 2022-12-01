<template>
  <div id="payDiv" ref="PayDiv" @click="closePay()">
    <div id="payPage" v-show="showPay" @click.stop="">
        <div id="ChoosePayDiv" ref="ChoosePayDiv">

            <div style="height: 40px"></div>
            <span style="font-size: 20px; margin-left: 40px;">微信支付</span>
            <div style="height: 40px;"></div>
            <img src="../../public/static/支付码.png" height="180px" width="180px"> <br>
            <div style="height: 30px"></div>
            <button style="width: 50px; height: 50px;padding: 0; background-color: rgba(0, 0, 0, 0); border: 0; cursor: pointer; margin-left: 30px;"  @click="switchToWeichat()"> <img ref="WeichatPay" src="../assets/greenWeichat.png" width="40px" height="40px"></button>
            <button style="width: 50px; height: 50px; padding: 0; background-color: rgba(0, 0, 0, 0); border: 0; cursor: pointer;margin-left: 10px" @click="switchToAliPay()"><img ref="Alipay" width="40px" height="40px" src="../assets/grayAlipay.png"></button>
        </div>
        <div id="payListDiv">
        <div style="height: 50px"></div>
        <span>请在剩余时间内支付订单</span> <br/>
        <div style="height: 5px"></div>
        <span style="color: orange; font-size: 25px; text-align: center; margin-left: 18%;"> {{min}}:{{second}}</span>
        <div style="margin-top: 10px;"></div>
        <span style="font-size: 15px;">订单号：158489489489484894489</span>
    </div>
    </div>
  </div>
</template>

<script>
export default {
    data() {
        return {
            showPay: false,
            chooseWeichat: true,
            chooseAlipay: false,
            min:0,
            second:0,
            action: false,
        }
    },
    methods: {
        closePay() {
            this.showPay = false
            let PayDiv = this.$refs.PayDiv
            let _this = this
            let op = 0.5
            setTimeout( function closing() {
                op -= 0.02
                if (op > 0) {

                    PayDiv.style["background-color"] = `rgba(0,0,0,${op})`
                    setTimeout(closing, 20)
                } else {
                    _this.closeEvet()
                }
            }, 20)
            this.action = false
        },
        showThisPay() {
            let PayDiv = this.$refs.PayDiv
            let op = 0
            setTimeout(() => {
                this.showPay = true
                this.$refs.ChoosePayDiv.style["background-color"] = "rgb(255,255,255)"
            }, 120)
            setTimeout( function showing() {
                op += 0.02
                if (op < 0.5) {
                    PayDiv.style["background-color"] = `rgba(0,0,0,${op})`
                    setTimeout(showing, 20)
                }
            }, 20)
            this.$refs.WeichatPay.src = "./static/greenWeichat.png"
            this.$store.commit("resetPayPrice")
            this.actionTime()
        },
        switchToWeichat() {
            if (this.chooseWeichat == true) {
                return
            }
            this.$refs.WeichatPay.src= "./static/greenWeichat.png"
            this.chooseWeichat = true
            if (this.chooseAlipay == true) {
                this.chooseAlipay = false
                this.$refs.Alipay.src = "./static/grayAlipay.png"
            }
        },
        switchToAliPay() {
            if (this.chooseAlipay == true) {
                return
            }
            this.$refs.Alipay.src = "./static/blueAlipay.png"
            this.chooseAlipay = true
            if (this.chooseWeichat == true) {
                this.chooseWeichat = false
                this.$refs.WeichatPay.src = "./static/grayWeichat.png"
            }
        },
        actionTime() {
            this.action = false
            setTimeout(()=>{
                this.action = true
                this.min = 15
                this.second = 0
                this.changeSecond()
            }, 1200)

        },
        changeSecond() {
            this.second--
            if (this.second == -1) {
                this.second = 59
                this.min--
            }
            if(this.min == -1) {
                this.closePay()
                return
            }
            if (this.action) {
                setTimeout(this.changeSecond, 1000)
            }
        }
    },
    props:{
        closeEvet: Function,
    }
}
</script>

<style>

#payDiv {
    top: 0;
    position: fixed;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    z-index: 90;
}

#payPage {
    position: absolute;
    left: 26%;
    top: 30%;
    height: 400px;
    width: 700px;
    border-radius: 5px;
    box-shadow: 0 0 2px;
    background-color: rgba(255, 255, 255, 1);
    z-index: 99;
}

d{
    background-image: url(../assets/grayWeichat.png);
}

#ChoosePayDiv {
    padding-left: 70px;
    display: inline-block;
    width: 300px;
}   

#payListDiv {
    display: inline-block;
    width: 330px;
    vertical-align: top;
}

</style>