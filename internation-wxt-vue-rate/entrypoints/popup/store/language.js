import { ref ,watch,computed} from 'vue'
import { defineStore } from 'pinia'

export const useLanguageStore = defineStore('language', ()=>{
    const language = ref('en_US') //语言
    const cu = ref('us') //国家代码,用于获取国旗
    const currencyNames = ref({}) //货币翻译的语言列表
    const goodList = ['bg','cs','da','de','en','es','fr','hi','hu','id','it','ja','ko','nl','no','pl','pt','ru','sv','th','tr','uk','vi','zh']
    const isGoodCurrency = computed(()=>{
      let cu = language.value.slice(0,2)
      return goodList.includes(cu)
    })
    const $reset=()=>{
      language.value = 'en_US'
      cu.value = 'us'
      currencyNames.value = {}
    }
    const setCurrencyNames = async ()=>{
      // console.log('setCurrencyNames:',language.value)
      if(isGoodCurrency.value){
        // console.log('好语言')
        const response = await fetch('/language/locale_'+ language.value.slice(0,2) + '.json')
        const newTranslation = await response.json();
        currencyNames.value = newTranslation;
      }else{
        // console.log('坏语言')
        const response = await fetch('/language/locale_en.json')
        const newTranslation = await response.json();
        currencyNames.value = newTranslation;
      }
    }
    watch(
      language,
        (state) => {
          // setCurrencyNames()
        },
        { deep: true }
      )
    return {
      language,$reset,cu,currencyNames,setCurrencyNames
    }
  })