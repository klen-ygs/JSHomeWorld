<template>
  <div id="SelectInputDiv">
      <div id="SelectInputBack">
          <input type="text" id="SelectInput">
          <button id="SelectButton">搜索</button>
      </div>
  </div>
</template>

<script>
export default {

}
</script>

<style>

#SelectInputDiv {
    padding-top: 10px;
    padding-left: 450px;
}

#SelectInput {
    width: 400px;
    height: 30px;
    border: 0px groove black;
}



#SelectInputBack {
    width: 500px;
    height: 30px;
    border: 5px solid  rgb(170, 220, 238);
}

#SelectButton {
    background-color: rgb(170, 220, 238);
    height: 30px;
    width: 96px;
    border: 2px solid rgb(170, 220, 238);
}

#SelectButton:hover {
    background-color: rgb(120, 199, 228);
    height: 30px;
    width: 96px;
    border-left:5px solid rgb(170, 220, 238) ;
}

</style>