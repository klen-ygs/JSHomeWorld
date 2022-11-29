<template>
    <div>
        <HeadShow/>
        <div id="box">
        </div>
        <SelectInput ref="selectInput" :select="select" ></SelectInput>
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
      name:"SelectedPage",
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
              this.SearchText = this.$route.query.search
              this.$refs.selectInput.setSearchWord(this.SearchText)
              axios.get("http://127.0.0.1:8000/getShop",
                  {
                      params: {
                          MaxId: this.MaxId,
                          MinId: this.MinId,
                          Page: 1,
                          SearchWord: this.SearchText,
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