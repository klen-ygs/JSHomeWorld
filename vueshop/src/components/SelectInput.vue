<template>
  <div id="SelectInputDiv">
      <div id="SelectInputBack">
          <input type="text" id="SelectInput" v-model="SearchWord">
      </div>
      <button id="SelectButton" @click="search()"><strong>搜索</strong></button>
  </div>
</template>

<script>
import axios from 'axios'
export default {
    data() {
        return {
            SearchWord:"",
            page: 1
        }
    },
    methods:{
        search() {
            this.page = 1
            axios.get("http://127.0.0.1:8000/getShop",
            {
                params: {
                    Page: 1,
                    SearchWord: this.SearchWord,
                }
            }).then(res => {
                this.select(res.data)
            })
        }
    },
    props:{
        select:Function,
    }
}
</script>

<style>

#SelectInputDiv {
    padding-top: 10px;
    padding-left: 400px;
}

#SelectInput {
    width: 500px;
    height: 30px;
    border: 0px groove black;
    outline: none;
}

#SelectInput:focus {
    width: 500px;
    height: 30px;
    border: 0px groove black;
    outline: none
}

#SelectInputBack {
    width: 500px;
    height: 30px;
    border: 5px solid  rgb(170, 220, 238);
    display: inline-block;
}

#SelectButton {
    background-color: rgb(170, 220, 238);
    height: 40px;
    width: 96px;
    border: 2px solid rgb(170, 220, 238);
    font-size: 15px;
    font-family:Arial, Helvetica, sans-serif;
    color: rgb(255, 255, 255);
}

#SelectButton:hover {
    cursor: pointer;
    background-color: rgb(170, 220, 238);
    height: 40px;
    width: 96px;
    border-left:2px solid rgb(170, 220, 238) ;
    font-size: 15px;
    font-family:Arial, Helvetica, sans-serif;
}

</style>