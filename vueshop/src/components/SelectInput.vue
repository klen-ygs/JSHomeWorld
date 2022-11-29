<template>
  <div id="SelectInputDiv">
      <div id="SelectInputBack">
          <input type="text" id="SelectInput" v-model="SearchWord">
          <button id="SelectButton" @click="search()"><strong>搜索</strong></button>
      </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
    data() {
        return {
            SearchWord:"",
            page: 1,
            searchListeners: [],
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
            for (let a = 0; a < this.searchListeners.length; a++) {
                this.searchListeners[a]()
            }
        },
        addSearchListener(listener) {
            this.searchListeners.push(listener)
        },
        setSearchWord(name) {
            this.SearchWord = name
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
    width: 600px;
    height: 30px;
    border: 6px solid  rgb(170, 220, 238);
    display: inline-block;
    padding: 0 0 0 0;
}

#SelectButton {
    float: right;
    background-color: rgb(170, 220, 238);
    height: 32px;
    width: 96px;
    border: 2px solid rgb(170, 220, 238);
    padding-bottom: 3px;
    font-size: 15px;
    font-family:Arial, Helvetica, sans-serif;
    color: rgb(255, 255, 255);
}

#SelectButton:hover {
    padding-bottom: 3px;
    cursor: pointer;
    background-color: rgb(170, 220, 238);
    height: 32px;
    width: 96px;
    border-left:2px solid rgb(170, 220, 238) ;
    font-size: 15px;
    font-family:Arial, Helvetica, sans-serif;
}

</style>