<template>
  <div>
      <HeadShow/>
      <div id="box">

      </div>
      <SelectInput ref="selectInput" :select="select" ></SelectInput>
      <div v-show="!onSearch" id="recommend">
        <div id="sortList">
            <div ref="recommedList" class="recommedList" v-for="(name, index) in sortList" :key="index" @mouseenter="enterRecommed(index, name)">
                {{name}}
            </div>
        </div>
        <div id="recommendImg" ref="recommendImg">

        </div>
      </div>
      <div v-show="!onSearch" id="hotPoint">
        <div style=" width=100%; height: 40px; text-align: center; font-size: 20px; color:rgb(229, 91, 88); border-bottom: 2px groove black; padding-top: 10px;">热搜榜</div>
        <div v-for="(select, index) in hotSelect" :key="select" class="hotSelect"  @click="searchFromHotPoint(select.SearchText)">
            <span style="color: orange">{{index + 1}}</span> {{"  " +select.SearchText}}
        </div>
    </div>
      <div id="shopDiv">
        <ShopWindown v-for="shop in shops" :key="shop.Id" :Id="shop.Id" 
        :Price="shop.Price" :ShopTitleText="shop.ShopTitleText" :Name="shop.Name"
        :getPhone="getPhone" />
      </div>
      <div id="pageDiv">
        <button ref="toLastPage" class="pageButton" @click="toLast()">上一页</button>
        <button ref="toNextPage" class="pageButton" @click="toNext()">下一页</button>
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
            if (this.Page == 1) {
                this.$refs.toLastPage.disabled = "disabled"
            }
            this.$refs.toNextPage.disabled = ""
            axios.get("http://127.0.0.1:8000/getShop",
            {
                params :{
                    MaxId: this.MaxId,
                    MinId: this.MinId,
                    Page:this.Page,
                    FindMethod: "Last"
                    
                },
            }
            ).then(res => {
                if (this.noData(res.data)) {
                    return
                }
                console.log(res.data)
                this.shops = res.data
                this.MinId = res.data[0].Id
                this.MaxId = res.data[res.data.length - 1].Id
            })

        },
        toNext() {
            this.Page++
            this.$refs.toLastPage.disabled = ""
            axios.get("http://127.0.0.1:8000/getShop",
            {
                params :{
                    MaxId: this.MaxId,
                    MinId: this.MinId,
                    page:this.Page,
                    FindMethod: "Next",
                },
            }
            ).then(res => {
                if (this.noData(res.data)) {
                    this.shops = null
                    this.$refs.toNextPage.disabled = "disabled"
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
        },
        enterRecommed(index, name) {
            this.$refs.recommedList[this.chooseRecommed].style["background-color"] = ""
            this.chooseRecommed = index
            this.$refs.recommedList[index].style["background-color"] = "rgb(238, 80, 48)"
            this.$refs.recommendImg.src = name + ".png"
        },
        getHotSelect() {
            axios({
                method:"Get",
                url:"http://127.0.0.1:8000/hotSearch",
            }).then( res => {
                this.hotSelect = res.data
            } )

        },
       searchFromHotPoint(name) {
            this.$refs.selectInput.setSearchWord(name)
            this.$refs.selectInput.search()
        }
    },
    mounted() {
            this.Phone = this.$route.params.Phone
            axios.get("http://127.0.0.1:8000/getShop",
                {
                    params: {
                        MaxId: this.MaxId,
                        MinId: this.MinId,
                        Page:this.Page,
                        FindMethod: "Next",
                    }
                }
            ).then(res => {
                if (this.noData(res.data)) {
                    return
                }
                this.shops = res.data
                this.MinId = res.data[0].Id
                this.MaxId = res.data[res.data.length - 1].Id
            })
            this.$refs.toLastPage.disabled= "disabled"
            this.getHotSelect()
            let _this = this
            setTimeout(function getHot() {
                _this.getHotSelect()
                setTimeout(getHot, 10000)
            }, 10000)
            this.$refs.selectInput.addSearchListener(()=> {
                this.onSearch = true
            })
    },
    data() {
        return {
            shops: [],
            MinId: 0,
            MaxId: 0,
            findMothod: "",
            Page: 1,
            Phone:"",
            sortList: ["衣服","电器","玩具","日用品","婴儿","零食","厨具"],
            hotSelect: [ {SearchText:"当前无热搜"},{SearchText:"显卡"}],
            chooseRecommed: 0,
            onSearch: false,
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

d{
    color: rgb(229, 91, 88);
}

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

#recommend {
    display: inline-block;
    margin-top: 30px;
    margin-left: 10%;
    width: 60%;
    height: 500px;
    background-color: rgb(158, 158, 158);
    border-radius: 2px;
}

#sortList {
    float: left;
    width: 15%;
    height: 100%;
    background-color: rgb(55, 55, 55);
}

#recommendImg {
    display: inline-block;
    width: 85%;
    height: 100%;
    background-color: rgb(217, 217, 217);
}

#hotPoint {
    float: right;
    margin-right: 10%;
    margin-top: 30px;
    width: 15%;
    height: 500px;
    background-color: rgb(255, 255, 255);
    box-shadow: 0 1px 2px gray;
    border-radius: 4px;
}

.recommedList {
    margin: 0 0 0 0;
    display: inline-block;
    width: 100%;
    height: 8%;
    font-size: 16px;
    text-align: center;
    padding-top: 20px;
    color: white;
}

.recommedList:hover {
    margin: 0 0 0 0;
    display: inline-block;
    width: 100%;
    height: 8%;
    font-size: 16px;
    text-align: center;
    color: white;
    padding-top: 20px;
    background-color: rgb(238, 80, 48);
}

.hotSelect {
    cursor: pointer;
    max-lines: 1;
    word-break:keep-all;
	white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
    padding-left: 20px;
    text-align: left;
    height: 25px;
    padding-top: 10px;
    font-size: 15px;
    color: blue;
}

.hotSelect:hover {
    padding-left: 20px;
    text-align: left;
    height: 25px;
    max-lines: 1;
    text-decoration: underline;
    padding-top: 10px;
    font-size: 15px;
    color: rgb(83, 83, 249);
}

</style>