<template>
  <div>
      <HeadShow/>
      <SelectInput></SelectInput>
      <div id="shopDiv" v-for="shop in shops" :key="shop.id">
        <ShopWindown :imageId="shop.ImageId" :price="shop.Price" :text="shop.ShopTitleText"></ShopWindown>
      </div>
  </div>
</template>

<script>
import HeadShow from '../components/HeadShow.vue'
import SelectInput from '../components/SelectInput.vue'
import ShopWindown from '../components/ShopWindown.vue'
import axios from 'axios'
export default {
    mounted() {
        for (let a = 1; a <= 20; a++ ) {
            axios.get(`http://127.0.0.1:8000/getShop`,
                {
                    params: {
                        id: a,
                    }
                }
            ).then(res => {
                this.shops.push(res.data)
            })
        }
    },
    data() {
        return {
            shops: [],
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
    background-color: rgb(243, 235, 235);
    margin-left: 5%;
    box-shadow: 0px 0px 2px  black;
}

</style>