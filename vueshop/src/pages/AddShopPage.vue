<template>
  <div>
      <table>
          <tr>
              <td>商品名字</td>
              <td><input type="text" v-model="name"></td>
          </tr>
          <tr>
              <td><label>宣传标语</label></td>
              <td><input type="text" v-model="shopTitleText" /></td>
          </tr>
          <tr>
            <td>
                <label>商品宣传图</label>
                <input type="file" ref="TitleImage" @change="addImage($event)" />
            </td>
          </tr>
          <tr>
            <td>
                <label>店内简介</label>
            </td>
            <td>
                <textarea v-model="text"></textarea>
            </td>
          </tr>
          <tr>
              <td><label>店内宣传图</label></td>
              <td><input type="file" multiple="multiple" @change="addImages($event)"></td>
          </tr>
          <tr>
              <td><label>商品价格</label></td>
              <td><input type="text" v-model.number="price" /></td>
          </tr>
          <tr>
              <td><label>商品数量</label></td>
              <td><input type="text" v-model.number="having" /></td>
          </tr>
          <tr>
              <img ref="Image" />
          </tr>
      </table>
      <button @click="update()">上传</button>
      <button @click="get()">get</button>
      
  </div>
</template>

<script>
import axios from 'axios'

export default {
    data() {
        return {
            name:"",
            shopTitleText:"",
            price:0,
            having:0,
            text:"",
            titleImageData: null,
            imagesData: null,
        }
    },
    methods: {
        get(){
            axios({
                method:"Get",
                url:"http://127.0.0.1:8000/addImageTest",
                responseType: "blob"
            }).then( res=> {
                
                this.$refs.Image.src =URL.createObjectURL(new Blob([res.data],{type:"image/*"}))
            } )

        },
        update() {
            axios({
                method:"Post",
                url:"http://127.0.0.1:8000/addShop",
                data: {
                    name: this.name,
                    shopTitleText: this.shopTitleText,
                    text: this.text,
                    having: this.having,
                    price: this.price
                }
            }).then( res => {
                let shopId = res.data.id
                let data = new FormData()
                console.log(this.titleImageData)
                data.append("ImageData",this.titleImageData)
                data.append("flag","true")
                data.append("shopId",shopId)
                axios.post("http://127.0.0.1:8000/addImage",
                data,
                 )
                 let imageDatas = new FormData()
                 imageDatas.append("flag","false")
                 imageDatas.append("shopId",shopId)
                 for (let a = 0; a < this.imagesData.length; a++) {
                     imageDatas.set("ImageData",this.imagesData[a])
                     axios.post("http://127.0.0.1:8000/addImage",imageDatas)
                 }
            } )

        },
        addImage(event) {
           this.titleImageData = event.target.files[0]
        },
        addImages(event) {
            this.imagesData = event.target.files
        }
    }
}
</script>

<style>

</style>