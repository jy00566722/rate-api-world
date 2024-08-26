<template>
    <div class="note" v-if="noteInfo.main">
          <el-alert
          :title="noteInfo.context"
          type="success"
          center
          >
        </el-alert>
        </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch, computed ,onBeforeMount} from 'vue'
import { useMessageStore } from '../store/message'
import {useLanguageStore} from '../store/language.js'
import { ElMessage } from 'element-plus'


const noteInfo=ref({
    "context": "Thank you for using the extension!",
    "index": 1,
    "main": false,
    "type": "info"
})
const messageStore = useMessageStore()
const languageStore = useLanguageStore()

onBeforeMount(async()=>{

        let url = 'https://rate-api.oeo.li/v1/message?language='+languageStore.language
        let a = await fetch(url)
        let b = await a.json()
        if(b.code===20000){
          b.data.forEach(item => {
            if(item.main){
              noteInfo.value.main = true
              noteInfo.value.context = item.context
              // ElMessage({
              //   message: '这是一个提示哈哈哈=======',
              //   type: 'success',
              // })
            }
          });
        //保存进入pinia
        messageStore.message = b.data
        }else{
          console.log('some thing is not Good',b)
        }
  
})

//返回浏览器Chrome版本是否大于等于102
const getChromeVersion=()=>{
let  ua = navigator.userAgent;
let chrome = /Chrome\/([\d]+)./i.exec(ua);
let version = parseInt(chrome[1])
// console.log(version)
return version < 102
}
</script>

<style>
    .note {
      margin: 5px 10px 5px 10px;
      /* background-color: rgb(39, 46, 56); */
      
    }
    .note .el-alert{
      padding: 4px 12px 4px 12px;
    }
</style>