<template>

  <router-view/>
    
</template>
<script setup>
import { ref,onMounted,onBeforeMount } from 'vue'
import {useLanguageStore} from './store/language.js'
import {useI18n} from 'vue-i18n';


const { t,locale,fallbackLocale} = useI18n();
// const languageStore = useLanguageStore()
// console.log('App中的数据:',languageStore.language,languageStore.setCurrencyNames)


//定义默认语言 //正式发布时改为en_US
const defaultLanguage = 'en_US'
const defaultCu = 'us'
const theme = ref('')
onBeforeMount(()=>{
//从storage中获取主题,theme.theme可能的值为light,dark,auto,如果为auto,则根据系统主题来设置,如果为light或者dark,则直接设置.如果没有获取到，则默认为auto.为auto时用window.matchMedia来获取系统主题。
//按获取的主题设置主题,设置方法为获取html上的data-theme属于，然后设置其值为light或dark
chrome.storage.local.get(["theme","layout"],(result)=>{
  if(result.theme){
    theme.value = result.theme
  }else{
    theme.value = 'auto'
  }
  // console.log(theme.value)
  if(theme.value==='auto'){
    // console.log('auto')
    let theme = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
    // console.log(theme)
    document.documentElement.setAttribute('data-theme',theme);
  }else{
    // console.log('light or dark')
    document.documentElement.setAttribute('data-theme',theme.value)
  }
  if(result.layout){
    document.documentElement.setAttribute('data-layout',result.layout);
    messageStore.layout = result.layout
  }else{
    document.documentElement.setAttribute('data-layout','default');

  }
  })
})

onMounted(async() => {
      //从stroage中获取language设置 //这里是总初始化i18n的地方
      chrome.storage.local.get(["language","cu"], function (result) {
          if (result.language&&result.cu) {
            locale.value = result.language;
            fallbackLocale.value = result.language;
            // languageStore.$patch({
            //   language:result.language,
            //   cu :result.cu
            // })
            // languageStore.setCurrencyNames()
          }else{
            locale.value = defaultLanguage
            fallbackLocale.value = defaultLanguage
            
            // languageStore.$patch({
            //   language:defaultLanguage,
            //   cu:defaultCu
            // })
            // setCurrencyNames(locale.value)
            // languageStore.setCurrencyNames()
            chrome.storage.local.set({ language: defaultLanguage ,cu:defaultCu}, function () {});
          }
        });
    })

//按language请求货币的翻译，合并到locale
// const setCurrencyNames = async (language)=>{
//   try{
//     const response = await fetch('/language/locale_'+ language.slice(0,2) + '.json')
//     const newTranslation = await response.json();
//     console.log(newTranslation)
//     languageStore.setCurrencyNames
//   }catch(e){
//     console.log('no currency name:',language)
//     // console.log(e)
//   }
// }
</script>
<style>
body{
  margin: 0;
  padding: 0;
  /* display: flex; */
  align-items:center;
  justify-content:center;
  background:var( --main-backgroud-color);
  width:385px;
}
#app {
  width: 380px;
  background: var( --main-backgroud-color);
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  /* color: #2c3e50; */
  /* color: var(--main-text-color); */
}

body::-webkit-scrollbar {
    width: 5px;
    height: 10px;
  }

  body::-webkit-scrollbar-thumb {
    border-radius: 10px;
    -webkit-box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.2);
    box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.2);
    background: var(--main-scrollbar-thumb-color);
  }

  body::-webkit-scrollbar-track {
    -webkit-box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.2);
    box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.2);
    border-radius: 2px;
    background: var(--main-scrollbar-track-color);
  }


</style>
