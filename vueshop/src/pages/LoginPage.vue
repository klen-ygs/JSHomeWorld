<template>
  <div class="loginBack">
    <div id="inputPage">
        <div class="inputs">
            <form>
                <table id="registerPage">
                    <tr class="loginTr">
                        <td class="loginTd"><label>手机号</label></td>
                        <td><input type="text" v-model="phone" class="loginInput"> </td>
                    </tr>
                    <tr class="loginTr">
                        <td class="loginTd"><label>昵称</label></td>
                        <td><input type="text" v-model="name" class="loginInput"></td>
                    </tr>
                    <tr class="loginTr">
                        <td class="loginTd"><label>密码</label></td>
                        <input type="password" v-model="password" class="loginInput">
                    </tr>
                    <tr class="loginTr">
                        <td class="loginTd"><label>确认密码</label></td>
                        <td><input type="password" v-model="password1" class="loginInput"></td>
                    </tr>
                </table>
            </form>
            <div id="registerButtonDiv">
                <input type="radio"><label>同意<a href="#">《用户使用》</a>协议</label><br/>
                <br/>
                <button class="firstPageButton" @click="register()">注册</button>
            </div>
        </div>
        <div class="inputs">

        </div>
        <!-- <div id="cover">

        </div> -->
    </div>

  </div>
</template>

<script>
import axios from 'axios'
export default {
    data(){
        return{
            phone:"",
            name:"",
            password:"",
            password1:"",
        }
    },
    methods: {
        register() {
            axios.post("127.0.0.1:8000/register",
            {
                phone:this.phone,
                name:this.name,
                password: this.password,
            }
            ).then(res=>{   

            })

        }
    }
}
</script>

<style>
.loginBack {
    width: 100%;
    height: 800px;
    /*background-image: url(../assets/leaf.jpeg);*/
}

#inputPage {
    position: relative;
    width: 750px;
    height: 450px;
    top: 150px;
    left: 350px;
    background-color: rgb(245, 239, 239);
    border-radius: 10px;
}

.inputs {
    position: absolute;
    display: inline-block;
    width: 375px;
    height: 450px;
}

#cover {
    width: 375px;
    height: 450px;
    position: absolute;
    background-color: rgb(236, 223, 223);
}

#registerPage {
    margin-left: 80px;
}
.loginTd {
    text-align: right;
}

.loginInput {
    margin-top: 10px;
    margin-bottom: 10px;
    margin-left: 10px;
    height: 30px;
    width: 200px;
    border-radius: 12px;
    border: 2px groove gray;
}

#registerButtonDiv {
    margin-top: 30px;
    margin-left: 100px;
}

.firstPageButton {
    margin-left: 80px;
    width: 120px;
    height: 45px;
    background-color: rgb(141, 181, 200);
    border-radius: 15px;
    border: 1px groove rgb(191, 220, 226);
}


</style>