<template>
  <div id="HeadBack">
      <div id="HeadImage">
          <!-- <img :src="headSrc"> -->
      </div>
      <a href="#">注销</a>
  </div>
</template>

<script>
export default {
    data() {
        return{ 
            headSrc:"",
        }
    }
}
</script>

<style>

#HeadImage {
    border-radius: 100px;
}

#HeadBack {
    height: 70px;
    background-color: rgb(221, 234, 245);
    box-shadow: inset -1 -1 5px black;
}

</style>