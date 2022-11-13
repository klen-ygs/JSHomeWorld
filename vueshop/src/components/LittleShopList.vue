<template>
  <div id="littleList" @mouseenter="mouseEnter()" @mouseleave="mouseLeave()" >
        <h3 id="addListTitle">收藏</h3>
        <div id="listsDiv" onmousewheel="wheel()">
            <LittleShop v-for="shop in lists" :key="shop" :ShopId="Number(shop)"/>
        </div>
  </div>
</template>

<script>
import LittleShop from './LittleShop.vue'
import axios from 'axios'
export default {
    data() {
        return {
            mouseLeaveObservers:[],
            mouseEnterObservers:[],
            lists: [],
        }
    },
    methods: {
        show() {
            axios.get("http://127.0.0.1:8000/addTip",
      {
        params:{
          Phone : this.Phone
        },
        withCredentials:true
      }
      ).then(res => {
            this.lists = JSON.parse(res.data.List)
          })

        },
        mouseLeave() {
            for (let a = 0; a < this.mouseLeaveObservers.length; a++) {
                this.mouseLeaveObservers[a]()
            }
        },
        mouseEnter() {
            for (let a = 0; a < this.mouseEnterObservers.length; a++) {
                    this.mouseEnterObservers[a]()
            }
        },
        addMouseEnter(func) {
            this.mouseEnterObservers.push(func)
        },
        addMouseLeave(func) {
            this.mouseLeaveObservers.push(func)
        },
        wheel() {
            return false
        }
    },
    props: {
        Phone: String,
    },
    components: {
        LittleShop,
    }
}
</script>

<style>

#littleList {
    height: 450px;
    width: 300px;
    border-radius: 10px;
    box-shadow: 0 0 4px gray;
    position: absolute;
    top: 60px;
    right: 270px;
    background-color: rgb(232, 235, 238);
}

#listsDiv {
    overflow: visible;
    overflow-x: hidden;
    margin-top: 20px;
    height: 85%;
}

#listsDiv::-webkit-scrollbar {
    width: 4px;
}

#listsDiv::-webkit-scrollbar-thumb {
    border-radius: 10px;
    background-color: gray;
}

#addListTitle {
    text-align: center;

}
</style>