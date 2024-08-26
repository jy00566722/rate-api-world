export default defineBackground(() => {
    console.log('%c%s', 'color: #0000ff; background: #ffffff', ' backgroud-server worker start...')
    //用红色字黄色底打印当前时间,按中国的时间规格
    console.log('%c%s', 'color: #ff0000; background: #ffff00', (new Date()).toLocaleString('zh-CN', { hour12: false }))

    //获取汇率 url='https://rate-api.oeo.li/v1/exchange'
    const getExchageRate = async ()=>{

        let url='https://rate-api.oeo.li/v1/exchange'
        const a = await fetch(url)
        const b = await a.json()
        // console.log(b)
        if(b.code === 200){
          chrome.storage.local.set({ "exchange_rate": b.data }, function (s) {  
            // console.log("汇率更新成功");
          });
          let exchange_rate_obj = {}
          b.data.forEach((item)=>{
            exchange_rate_obj[item.Code] = item
          })
          chrome.storage.local.set({ "exchange_rate_obj": exchange_rate_obj }, function (s) {  
            // console.log("汇率更新成功");
          });
        }
        return b
    }
        //获取加密货币信息 url='https://rate-api.oeo.li/v1/exchange_crypto'
        const getExchageRate_crypto = async ()=>{
          let url='https://rate-api.oeo.li/v1/exchange_crypto'
          const a = await fetch(url)
          const b = await a.json()
          if(b.code === 200){
            chrome.storage.local.set({ "exchange_rate_crypto": b.data }, function (s) {  
              // console.log("汇率更新成功");
            });
            let exchange_rate_crypto_obj = {}
            b.data.forEach((item)=>{
              exchange_rate_crypto_obj[item.Code] = item
            })
            chrome.storage.local.set({ "exchange_rate_crypto_obj": exchange_rate_crypto_obj }, function (s) {  
              // console.log("汇率更新成功");
            });
          }
          return b
      }
      

      
      //向激活的窗口，发送消息，弹出主界面
      function sendMessageToContentScript(message, callback) {
        chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
          chrome.tabs.sendMessage(tabs[0].id, message, function (response) {
            if (callback) callback(response);
          });
        });
      }
      //生成uuid
      function generateUUID() {
        let d = new Date().getTime();
        let uuid = 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
            let r = (d + Math.random() * 16) % 16 | 0;
            d = Math.floor(d / 16);
            return (c == 'x' ? r : (r & 0x3 | 0x8)).toString(16);
        });
        return uuid
      }
      
      
      chrome.runtime.onInstalled.addListener(async function(e){
      
        if(e.reason === 'install'){
                // console.log('插件安装') 
                //初始化货币列表中的国家
                let countries = ['USD','BTC','ETH','EUR','JPY','GBP','KRW','CAD','CNY']
                chrome.storage.local.set({ "currency_list_now": countries }, function (s) {  
                  // console.log("初始化::货币列表在currency_list_now中...");
                });
                //设置小数点位数
                chrome.storage.local.set({ "decimal": 2 }, function (s) {  
                  // console.log("初始化::小数点位数在decimal_point中...");
                });
      
                start_all_function()
                chrome.storage.local.get(['my_id'],function(s){
                  if(s.my_id){
                    // console.log('install:存在my_id')
                  }else{
                    // console.log('install:不存在my_id')
                    chrome.storage.local.set({"my_id":generateUUID()},function(){console.log('good two')})
                  }
                })
        }
        if(e.reason === 'update'){
                start_all_function()    
                //加入BTC提示用户加入了加密货币
                let {currency_list_now} = await chrome.storage.local.get('currency_list_now')
                if(!currency_list_now.includes('BTC')){
                  currency_list_now.unshift('BTC')
                  await chrome.storage.local.set({currency_list_now})
                }  
                
                if (chrome.runtime.openOptionsPage) {
                  // chrome.runtime.openOptionsPage();
                } else {
                  // window.open(chrome.runtime.getURL('options.html'));
                }
                chrome.storage.local.get(['my_id'],function(s){

                  if(s.my_id){
                    // console.log('update:存在my_id')
                  }else{
                    // console.log('update:不存在my_id')
                    chrome.storage.local.set({"my_id":generateUUID()},function(){console.log('good one')})
                  }
                })
  
        }
          if(e.reason === 'chrome_update'){
                  start_all_function()
          }
      
      })
      
      const start_all_function = ()=>{

        getExchageRate()
        getExchageRate_crypto()

        chrome.storage.local.set({'get_data_time':(new Date()).toString()},function(){
        })
      }
      
      // chrome.alarms.create('start_all_function',{ periodInMinutes: 5 })

      async function checkAlarmState() {     
          console.log('check alarm status...')
          const alarm = await chrome.alarms.get("mihu-alarm");
          if (!alarm) {
            await chrome.alarms.create("mihu-alarm",{ periodInMinutes: 10 });
          }
      }
      
      checkAlarmState();
      
      chrome.alarms.onAlarm.addListener((r) => {
        start_all_function()
      })
      
      
    //设置卸载反馈
    // chrome.runtime.setUninstallURL(
    //     'https://rate-api.oeo.li/v1/uninstall',()=>{
    //         console.log('设置反馈URL成功')
    //     }
    // ) 

    //设置接收消息
    chrome.runtime.onMessage.addListener(
        function(request, sender, sendResponse) {
          if (request.do === "getStart")
            sendResponse({status: "good"});
            start_all_function()
        }
      );

    //chrome语言检测相关
    const testLanguage = async ()=>{
      let acceptLanguages = await chrome.i18n.getAcceptLanguages()
      console.log(acceptLanguages,':acceptLanguages\n')
      let uiLanguage = await chrome.i18n.getUILanguage()
      console.log(uiLanguage,':uiLanguage\n')
      let a = await chrome.i18n.getMessage('actionTitle')
      console.log('a:',a)
      console.log("someKey:",i18n.global.t('someKey'));
    }
    // testLanguage()
  });