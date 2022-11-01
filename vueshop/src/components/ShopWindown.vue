<template>
  <div id="ShopBack">
      <a :href="shopUrl">
          <img :src="shopImg" id="ShopImg" />
      </a>
      <h2 id="price"></h2>
      <a :href="shopUrl">
          {{shopText}}
      </a>
  </div>
</template>

<script>
export default {
    data() {
        return{
            shopImg:"#",
            shopUrl:"#",
            shopText:"",
        }
    }
}
</script>

<style>

#ShopBack {
    width: 25%;
    display: inline-block;
    height: 30%;
    box-shadow: 0px 0px 1px;
}

#shopImg {
    width: 90%;
    margin-left: 5%;
    height: 50%;
    margin-bottom: 10%;
}

</style>