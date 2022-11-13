<template>
  <div>
      <HeadShow/>
      <div id="box">

      </div>
      <SelectInput :select="select" ></SelectInput>
      <div id="shopDiv">
        <ShopWindown v-for="shop in shops" :key="shop.Id" :Id="shop.Id" 
        :Price="shop.Price" :ShopTitleText="shop.ShopTitleText" :Name="shop.Name"
        :getPhone="getPhone" />
      </div>
      <div id="pageDiv">
        <button class="pageButton" @click="toLast()">上一页</button>
        <button class="pageButton" @click="toNext()">下一页</button>
      </div>
  </div>
</template>

<script>
import HeadShow from '../components/HeadShow.vue'
import SelectInput from '../components/SelectInput.vue'
import ShopWindown from '../components/ShopWindown.vue'
import axios from 'axios'
export default {
    name:"SelectPage",
    methods:{
        select(data) {
            this.shops = data
        },
        toLast() {
            this.Page--
            axios.get("http://127.0.0.1:8000/getShop",
            {
                params :{
                    MaxId: this.MaxId,
                    MinId: this.MinId,
                    Page:this.Page,
                    FindMothod: "Last"
                    
                },
            }
            ).then(res => {
                if (this.noData(res.data)) {
                    return
                }
                this.shops = res.data
                this.MinId = res.data[0].Id
                this.MaxId = res.data[res.data.length - 1].Id
            })

        },
        toNext() {
            this.Page++
            axios.get("http://127.0.0.1:8000/getShop",
            {
                params :{
                    MaxId: this.MaxId,
                    MinId: this.MinId,
                    page:this.Page,
                    FindMothod: "Next",
                },
            }
            ).then(res => {
                if (this.noData(res.data)) {
                    return
                }
                this.shops = res.data
                this.MinId = res.data[0].Id
                this.MaxId = res.data[res.data.length - 1].Id
            })

        },
        getPhone() {
            return this.Phone
        },
        noData(data) {
            if (data.length == 0) {


                return true                
            }
            return false
        }
    },
    mounted() {
            this.Phone = this.$route.params.Phone
            console.log(this.Phone)
            axios.get("http://127.0.0.1:8000/getShop",
                {
                    params: {
                        MaxId: this.MaxId,
                        MinId: this.MinId,
                        Page:this.Page,
                        FindMothod: "Next",
                    }
                }
            ).then(res => {
                console.log(res.data)
                if (this.noData(res.data)) {
                    return
                }
                this.shops = res.data
                this.MinId = res.data[0].Id
                this.MaxId = res.data[res.data.length - 1].Id
            })
    },
    data() {
        return {
            shops: [],
            MinId: 0,
            MaxId: 0,
            findMothod: "",
            Page: 1,
            Phone:""
        }
    },
    components:{
        HeadShow,
        SelectInput,
        ShopWindown,
    }
}
</script>

<style>

#shopDiv {
    margin-top: 20px;
    width: 90%;
    height: 1200px;
    background-color: rgb(248, 248, 248);
    margin-left: 5%;
    box-shadow: 0px 0px 2px  black;
}

#pageDiv {
    margin-top: 30px;
    margin-left: 5%;
    padding-left: 20%;
    width: 90%;
    height: 100px;
    background-color: rgb(255, 255, 255);
}

.pageButton {
    margin-left: 10%;
    margin-right: 10%;
}

#box {
    height: 80px;
}

</style>