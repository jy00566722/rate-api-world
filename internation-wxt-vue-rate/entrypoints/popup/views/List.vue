<template>
  <div class="list">
    <div ref="el">
      <TransitionGroup name="rateList">

        <div class="rate_item" v-for="item in country_list_rate" :key="item.tcur">
          <img :src="item.country_img" alt="" width="50" style="border-radius: 5px" class="handle1 flage_img" />
          <div class="rate_item_in">
            <div>
            <span>
                  {{ item.currency_symbol }}
                </span>
            <el-input size="default" v-model="item.num" placeholder="0" clearable :formatter="formatter"
              :parser="parser" @input="numChange(item)" @clear="clear()">
             
            </el-input>
          </div>
            <!-- <span class="currency_text">{{item.tcur}} {{" - "}} {{item.currency_name}}</span> -->
            <span class="currency_text">{{item.tcur}} {{" - "}} {{item.currencyName}}</span>
          </div>
          <span class="del_span" @click="delItem(item)">
            <el-icon :size="18">
              <CloseBold />
            </el-icon>
            <!-- <svg viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" data-v-ea893728=""><path fill="currentColor" d="M764.288 214.592 512 466.88 259.712 214.592a31.936 31.936 0 0 0-45.12 45.12L466.752 512 214.528 764.224a31.936 31.936 0 1 0 45.12 45.184L512 557.184l252.288 252.288a31.936 31.936 0 0 0 45.12-45.12L557.12 512.064l252.288-252.352a31.936 31.936 0 1 0-45.12-45.184z"></path></svg> -->
          </span>
        </div>

      </TransitionGroup>
    </div>
  </div>
</template>

<script setup lang="ts">
  import { ref, onMounted, onBeforeMount,onBeforeUnmount } from 'vue'
  import { storeToRefs } from 'pinia'
  import { useDraggable, type UseDraggableReturn } from 'vue-draggable-plus'
  import {useLanguageStore} from '../store/language.js'
  import {useI18n} from 'vue-i18n';

  const languageStore = useLanguageStore()
  const {currencyNames,language} = storeToRefs(languageStore)


  //i18n
  const { t,locale,fallbackLocale,mergeLocaleMessage } = useI18n();
  const country_list = ref([])
  const country_list_rate = ref([]) 
  const el = ref() //拖动列表的ref

  const decimal = ref(2)    //定义小数位置
  const currencyFormat = ref("./,") //POP页面价格转换时货币格式

  //定义记住的当前货币入数字
  const currentCurrency = ref({
    currency:'',
    num:0
  })
  onBeforeMount(() => {
    fetchData()
  })
  onMounted(() => {
    initData()
    chrome.storage.onChanged.addListener(storageListenerFunction)
  })
  onBeforeUnmount(()=>{
    chrome.storage.onChanged.removeListener(storageListenerFunction);
  })
  
  const initData = ()=>{
    chrome.storage.local.get(
      ["currency_list_now", "exchange_rate_obj", "decimal", "currencyFormat","exchange_rate_crypto_obj","current_currency","language"],
      async function (result) {
        let { currency_list_now,  exchange_rate_obj,exchange_rate_crypto_obj } = result;
        if (result.decimal) { //小数点位数
          if (typeof result.decimal == 'number' && result.decimal >= 2 && result.decimal <= 6) {
            decimal.value = result.decimal
          } else {
            //删除原来的decimal
            await chrome.storage.local.remove('decimal')
          }
        }
        if (result.currencyFormat) { //货币格式
          currencyFormat.value = result.currencyFormat
        }

        if (!currency_list_now) {
          currency_list_now = ['USD', 'BTC','ETH','EUR', 'JPY', 'GBP', 'KRW', 'CAD', 'CNY'];
          chrome.storage.local.set(
            { currency_list_now },
            function (s) {}
          );
        }
        // console.log("obj:",exchange_rate_obj)
        country_list.value = currency_list_now;
        //合并两个对象来获取当前要转换的货币
        let all_exchange = {...exchange_rate_obj,...exchange_rate_crypto_obj}
        //初始化货币名称对照表
        if(result.language){
            language.value = result.language
            await languageStore.setCurrencyNames()
          }else{
            language.value = 'en_US'
            await languageStore.setCurrencyNames()
          }
        let c = [];
        country_list.value.forEach((el) => {
          let rate_country = {};
          if(all_exchange[el] && all_exchange[el].IsVirtual){ //如果是加密货币
            rate_country.country_img = all_exchange[el].FlagURL;
            rate_country.IsVirtual = true
            rate_country.rate = 1/parseFloat(all_exchange[el].ExchangeRate);
 
          }else{  //如果是传统货币
            rate_country.country_img = 'icons/flags/4x3/' +all_exchange[el].Code.slice(0, 2).toLowerCase() + '.svg'
            rate_country.IsVirtual = false
            rate_country.rate = all_exchange[el].ExchangeRate;
            rate_country.currency_symbol = all_exchange[el].Symbol;
          }
          

          rate_country.num = "";
          rate_country.currency_name = all_exchange[el].Currency;
          rate_country.currency_status = "";
          rate_country.tcur = all_exchange[el].Code;

          //加入货币对应语言的翻译locale
          if(currencyNames.value[all_exchange[el].Code]){
            // console.log('111')
            rate_country.currencyName = currencyNames.value[all_exchange[el].Code]
          }else{
            // console.log('222')
            // console.log(currencyNames,all_exchange[el].Code)
            rate_country.currencyName = all_exchange[el].Currency;
          }
          c.push(rate_country);
        });
        country_list_rate.value = c;
      
    
        //从之前保存的当前货币中恢复数字
      if(result.current_currency&&result.current_currency.currency&&result.current_currency.num){
        if(country_list.value.includes(result.current_currency.currency)){
          let num = parseFloat(result.current_currency.num)
          if(num>0){
            // console.log('num',num)
            country_list_rate.value.forEach((el)=>{
              if(el.tcur===result.current_currency.currency){
                el.num = result.current_currency.num
                numChange(el)
                currentCurrency.value = {
                  currency:result.current_currency.currency,
                  num:result.current_currency.num
                }
              }
            })
            
          }else{
            // console.log('NO,3')
          }
        }else{
          // console.log('NO,2')
        }
      }else{
        // console.log('NO,1')
      }
      }
    );
  
  }
  const numChange = (val) => {
    country_list_rate.value.forEach((el) => {
      if (val.tcur === el.tcur) {
      } else {
        el.num = ((Number(el.rate) / Number(val.rate)) * Number(val.num)).toFixed(decimal.value);
      }
    });
    //记录输入的货币与数字
      currentCurrency.value.currency = val.tcur,
      currentCurrency.value.num = val.num
      chrome.storage.local.set({current_currency:currentCurrency.value},function(){})
  }
  //删除列表中的某一项
  const delItem = (val) => {
    country_list_rate.value.forEach((el, index) => {
      if (val.tcur === el.tcur) {
        country_list_rate.value.splice(index, 1)
      }
    });
    //把当前的列表存入storage
    let c = country_list_rate.value.map((el) => {
      return el.tcur
    })
    chrome.storage.local.set({ currency_list_now: c }, function () {})
  }
  //拖动列表
  const draggable = useDraggable < UseDraggableReturn > (el, country_list_rate, {
    animation: 150,
    handle: '.handle1',
    onStart() {

    },
    onUpdate() {
      //把当前拖动后的列表保存到storage
      let c = country_list_rate.value.map((el) => {
        return el.tcur
      })
      chrome.storage.local.set({ currency_list_now: c }, function () {
        // console.log('修改成功')
      })
    }
  })
  //格式化input输入框的值
  const formatter = (value) => {
    if (!value) return "";
    let [decimalSeparator, thousandsSeparator] = currencyFormat.value.split("/");
    if (!thousandsSeparator) {
      thousandsSeparator = "";
    }
    if (thousandsSeparator === " ") {
      thousandsSeparator = "\u00A0";
    }
    const parts = value.toString().split(".");
    parts[0] = parts[0]
      .replaceAll(",", "")
      .replaceAll(/(\d)(?=(?:\d{3})+$)/g, "$1"+thousandsSeparator);
    return parts.join(".");
  }
  const parser = (value, options = {}) => {
    if (!value)
      return options.defaultReturn === undefined ? "" : options.defaultReturn;
    // 处理以.开头的情况
    value = value.startsWith(".") ? `0${value}` : value;

    // 只允许数字和小数点，且只允许一个小数点
    const result = value.replaceAll(/[^\d.]/g, "").replaceAll(/(\..*)\./g, "$1");
    if (result === "")
      return options.defaultReturn === undefined ? "" : options.defaultReturn;

    const parts = result.split(".");
    if (
      parts[1] &&
      options.decimalPlaces &&
      parts[1].length > options.decimalPlaces
    ) {
      parts[1] = parts[1].slice(0, options.decimalPlaces);
    }

    // 处理数字以0开头的情况，但不包括0后面跟着小数点的情况
    if (
      parts[0].startsWith("0") &&
      !parts[0].startsWith("0.") &&
      parts[0].length > 1
    ) {
      parts[0] = parts[0].slice(1);
    }

    return parts.join(".");
  }
  const fetchData = () => {
        //给后台发消息
        chrome.runtime.sendMessage({ do: "getStart" }, function (response) {
            // console.log('fetch data ok.');
        });
    }
  const storageListenerFunction = (changes, namespace)=>{
    // console.log(changes,namespace)
    let eventName = Object.keys(changes)
    if(eventName.includes('exchange_rate_crypto_obj') || eventName.includes('exchange_rate_obj')){
      // console.log('update exchange')
      updateExchange()
    }

  }
  const updateExchange = ()=>{
    chrome.storage.local.get(
      [ "exchange_rate_obj", "exchange_rate_crypto_obj"],
      async function (result) {
        let {  exchange_rate_obj,exchange_rate_crypto_obj } = result;

        //合并两个对象来获取当前要转换的货币
        let all_exchange = {...exchange_rate_obj,...exchange_rate_crypto_obj}
        // console.log(all_exchange)
        country_list_rate.value.map((el)=>{
          if(el.IsVirtual){//加密货币
            if(el.rate != 1/parseFloat(all_exchange[el.tcur].ExchangeRate)){
              // console.log('need update:',el.tcur,el.rate,1/parseFloat(all_exchange[el.tcur].ExchangeRate))
              el.rate = 1/parseFloat(all_exchange[el.tcur].ExchangeRate)
            }
          }else{ //传统考货币
            if(el.rate != all_exchange[el.tcur].ExchangeRate){
                // console.log('need update:',el.tcur,el.rate,all_exchange[el.tcur].ExchangeRate)
                el.rate = all_exchange[el.tcur].ExchangeRate
              }
          }


        })
      }
    );
  
  }
  //清空输入框时,去掉保存的当前货币与数字
  const clear = ()=>{
      chrome.storage.local.set({current_currency:{
        currency:'',
        num:0
      }},function(){
        // console.log('清空current_currency成功')
      })
}
 

</script>

<style>
  .ghost {
    opacity: 0.5;
    background: #c8ebfb;
  }

  .list {
    margin: 5px 0 20px 0;
    background: var(--main-backgroud-color);
    height: 120px;
  }

  .rate_item {
    margin: 10px auto;
    width: 360px;
    height: 85px;
    background: var(--list-item-bg-color);
    /* color: var(--list-text-color); */
    border-radius: 14px;
    display: grid;
    grid-template-columns: 80px auto 35px;
    grid-template-rows: 1fr;
    place-items: center center;
  }

  .rate_item:hover {
    background: var(--list-item-hover-bg-color);
  }

  .rate_item_in {
    display: grid;
    grid-template-rows: 2fr 1fr;
    grid-template-columns: 1fr;
    place-items: center start;
  }

  .flage_img{
    border-radius: 5px;
    transition: transform 0.5s ease;
    /* width: 50px; */
    /* height: 50px; */
    overflow: hidden;
    box-shadow:  -3px 3px 5px 0px rgba(0, 0, 0, 0.4);
  }

  .flage_img:hover {
    /* transform: scale(1.1); */
  }
  .rate_item:hover .flage_img {
    transform: scale(1.1);
  }

  .rate_item_in>div{
    display: grid;
    grid-template-columns: 1fr 4fr;
    align-items: center;
  }
  .rate_item_in>div>span{
    color: var(--list-item-symbol-text-color);
    font-size: 24px;
    padding-left: 0px;
    font-weight: bold;
    text-shadow: 1px 1px 0px var(--list-item-symbol-text-shadow-color);
    font-family:'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  }
  .rate_item_in .el-input-group__prepend {
    background: rgba(10, 31, 61, 0);
    box-shadow: none;
  }

  .rate_item_in .el-input__wrapper {
    background: rgba(53, 75, 105, 0);
    box-shadow: none;
  }

  .rate_item_in .el-input-group__prepend+.el-input__wrapper {
    border-radius: 5px;
  }

  .rate_item_in .el-input__inner {
    color: var(--list-item-input-text-color);
    font-family: Verdana, Geneva, Tahoma, sans-serif;
    font-size: 20px;
  }

  .rate_item_in>span {
    color: var(--list-item-span-text-color);
    font-size: 14px;
    /* margin-left: 30px; */

    font-family: Verdana, Geneva, Tahoma, sans-serif;
    white-space: nowrap;  /* 防止内容换行 */
    overflow: hidden;     /* 隐藏溢出的内容 */
    display: inline-block; /* 或者使用 block，取决于布局需求 */
    width: 100%;           /* 指定一个宽度，根据实际需要调整 */
  }

  .del_span {
    width: 18px;
    height: 18px;
    /* font-size: 16px; */
    color: var(--list-item-del-span-text-color);
  }

  .del_span:hover {
    width: 20px;
    height: 20px;
    color: var(--list-item-del-span-text-hover-color);
    cursor: pointer;
  }

  .rateList-enter-active,
  .rateList-leave-active {
    transition: all 0.5s ease;
  }

  .rateList-enter-from,
  .rateList-leave-to {
    opacity: 0;
    transform: translateX(30px);
  }
</style>