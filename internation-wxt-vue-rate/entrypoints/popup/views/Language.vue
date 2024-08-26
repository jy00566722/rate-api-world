<template>
    <div class="lan">
    <Nav darp-e="no"></Nav>
    <h1>{{t('selectLanguage')}}</h1>
    <div style="margin-bottom: 320px;width: 100%;">
        <el-select v-model="locale"  placeholder="Select" size="large"
        style="width: 260px;margin: 0 auto;"
        @change="languageChange()">
            <template #prefix>
                <img :src="'icons/flags/4x3/' + cu + '.svg'" alt="" width="32">
            </template>
            <el-option
              v-for="item in languageList"
              :key="item.code"
              :label="item.translated"
              :value="item.code"
            >
            <img :src="'icons/flags/4x3/' + item.cu + '.svg'" alt="" width="32">  <span>{{item.translated}}</span>
            
            </el-option>
          </el-select>
          <br>
          <div class="back-button-warp">
          <span class="back-button"  @click="$router.push('/')">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><g fill="none"><path d="M24 0v24H0V0zM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035c-.01-.004-.019-.001-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427c-.002-.01-.009-.017-.017-.018m.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093c.012.004.023 0 .029-.008l.004-.014l-.034-.614c-.003-.012-.01-.02-.02-.022m-.715.002a.023.023 0 0 0-.027.006l-.006.014l-.034.614c0 .012.007.02.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01z"/><path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2m0 2a8 8 0 1 0 0 16a8 8 0 0 0 0-16M9.879 8.464L12 10.586l2.121-2.122a1 1 0 1 1 1.415 1.415l-2.122 2.12l2.122 2.122a1 1 0 0 1-1.415 1.415L12 13.414l-2.121 2.122a1 1 0 0 1-1.415-1.415L10.586 12L8.465 9.879a1 1 0 0 1 1.414-1.415"/></g></svg>
          Back</span></div>
          
          <br>
          Powered by <a href="https://currency.oeo.li/" target="_blank" style=" text-decoration: none;"> <span class="lan">https://currency.oeo.li/</span> </a>
          <br><br>
          <Note></Note>
        </div>
    </div>
</template>

<script setup>
    import { ref, onMounted, onBeforeMount,watch } from 'vue'
    import Nav from './Nav.vue'
    import Note from './Note.vue'
    import {useLanguageStore} from '../store/language.js'
    import {useI18n} from 'vue-i18n';

    const languageStore = useLanguageStore()
    const cu = ref('us')
    //i18n
    const { t,locale } = useI18n();
    watch(locale, (val) => {
        chrome.storage.local.set({ language: val }, function () { });
    })
    onBeforeMount(async ()=>{
        const result = await chrome.storage.local.get('cu')
        cu.value = result.cu
    })
    const languageChange=()=>{

        //找出cu的值
        let cu1 = languageList.value.find((el)=>el.code == locale.value).cu
        cu.value = cu1
        chrome.storage.local.set({ language: locale.value,cu:cu1 }, function () {  });
        languageStore.cu = cu.value
        languageStore.language = locale.value
    }
    const languageList = ref( 
    [
        // {"code": "ar", "language": "Arabic", "translated": "العربية", "chineseNawme": "阿拉伯语"},
        // {"code": "am", "language": "Amharic", "translated": "አማርኛ", "chineseNawme": "阿姆哈拉语"},
        {"code": "bg", "language": "Bulgarian", "translated": "български", "chineseNawme": "保加利亚语","cu_code":"BGN","cu":"bg"},
        {"code": "bn", "language": "Bengali", "translated": "বাংলা", "chineseNawme": "孟加拉语","cu_code":"BDT","cu":"bd"},
        // {"code": "ca", "language": "Catalan", "translated": "Català", "chineseNawme": "加泰罗尼亚语"},
        {"code": "cs", "language": "Czech", "translated": "Čeština", "chineseNawme": "捷克语","cu_code":"CZK","cu":"cz"},
        {"code": "da", "language": "Danish", "translated": "Dansk", "chineseNawme": "丹麦语","cu_code:":"DKK","cu":"dk"},
        {"code": "de", "language": "German", "translated": "Deutsch", "chineseNawme": "德语","cu_code":"EUR","cu":"de"},
        {"code": "el", "language": "Greek", "translated": "Ελληνικά", "chineseNawme": "希腊语","cu_code":"EUR","cu":"gr"},
        {"code": "en", "language": "English", "translated": "English", "chineseNawme": "英语","cu_code":"USD","cu":"us"},
        {"code": "en_AU", "language": "English (Australia)", "translated": "English (Australia)", "chineseNawme": "英语（澳大利亚）","cu_code":"AUD","cu":"au"},
        {"code": "en_GB", "language": "English (Great Britain)", "translated": "English (Great Britain)", "chineseNawme": "英语（英国）","cu_code":"GBP","cu":"gb"},
        {"code": "en_US", "language": "English (USA)", "translated": "English (USA)", "chineseNawme": "英语（美国）","cu_code":"USD","cu":"us"},
        {"code": "es", "language": "Spanish", "translated": "Español", "chineseNawme": "西班牙语","cu_code":"EUR","cu":"es"},
        // {"code": "es_419", "language": "Spanish (Latin America and Caribbean)", "translated": "Español (América Latina y el Caribe)", "chineseNawme": "西班牙语（拉丁美洲和加勒比海）","cu_code":"EUR"},
        {"code": "et", "language": "Estonian", "translated": "Eesti", "chineseNawme": "爱沙尼亚语","cu_code":"EUR","cu":"ee"},
        {"code": "fa", "language": "Persian", "translated": "فارسی", "chineseNawme": "波斯语","cu_code":"IRR","cu":"ir"},
        {"code": "fi", "language": "Finnish", "translated": "Suomi", "chineseNawme": "芬兰语","cu_code":"EUR","cu":"fi"},
        {"code": "fil", "language": "Filipino", "translated": "Filipino", "chineseNawme": "菲律宾语","cu_code":"PHP","cu":"ph"},
        {"code": "fr", "language": "French", "translated": "Français", "chineseNawme": "法语","cu_code":"EUR","cu":"fr"},
        {"code": "gu", "language": "Gujarati", "translated": "ગુજરાતી", "chineseNawme": "古吉拉特语","cu_code":"INR","cu":"in"},
        {"code": "he", "language": "Hebrew", "translated": "עברית", "chineseNawme": "希伯来语","cu_code":"","cu":"il"},
        {"code": "hi", "language": "Hindi", "translated": "हिन्दी", "chineseNawme": "印地语","cu_code":"INR","cu":"in"},
        {"code": "hr", "language": "Croatian", "translated": "Hrvatski", "chineseNawme": "克罗地亚语","cu_code":"HRK","cu":"hr"},
        {"code": "hu", "language": "Hungarian", "translated": "Magyar", "chineseNawme": "匈牙利语","cu_code":"HUF","cu":"hu"},
        {"code": "id", "language": "Indonesian", "translated": "Bahasa Indonesia", "chineseNawme": "印度尼西亚语","cu_code":"IDR","cu":"id"},
        {"code": "it", "language": "Italian", "translated": "Italiano", "chineseNawme": "意大利语","cu_code":"EUR","cu":"it"},
        {"code": "ja", "language": "Japanese", "translated": "日本語", "chineseNawme": "日语","cu_code":"JPY","cu":"jp"},
        {"code": "kn", "language": "Kannada", "translated": "ಕನ್ನಡ", "chineseNawme": "卡纳达语","cu_code":"INR","cu":"in"},
        {"code": "ko", "language": "Korean", "translated": "한국어", "chineseNawme": "韩语","cu_code":"KRW","cu":"kr"},
        {"code": "lt", "language": "Lithuanian", "translated": "Lietuvių", "chineseNawme": "立陶宛语","cu_code":"EUR","cu":"lt"},
        {"code": "lv", "language": "Latvian", "translated": "Latviešu", "chineseNawme": "拉脱维亚语","cu_code":"LVL","cu":"lv"},
        // {"code": "ml", "language": "Malayalam", "translated": "മലയാളം", "chineseNawme": "马拉雅拉姆语"},
        // {"code": "mr", "language": "Marathi", "translated": "मराठी", "chineseNawme": "马拉地语"},
        {"code": "ms", "language": "Malay", "translated": "Bahasa Melayu", "chineseNawme": "马来语","cu_code":"MYR","cu":"my"},
        {"code": "nl", "language": "Dutch", "translated": "Nederlands", "chineseNawme": "荷兰语","cu_code":"ANG","cu":"an"},
        {"code": "no", "language": "Norwegian", "translated": "Norsk", "chineseNawme": "挪威语","cu_code":"NOK","cu":"no"},
        {"code": "pl", "language": "Polish", "translated": "Polski", "chineseNawme": "波兰语","cu_code":"PLN","cu":"pl"},
        // {"code": "pt_BR", "language": "Portuguese (Brazil)", "translated": "Português (Brasil)", "chineseNawme": "葡萄牙语（巴西）"},
        {"code": "pt_PT", "language": "Portuguese (Portugal)", "translated": "Português (Portugal)", "chineseNawme": "葡萄牙语","cu_code":"EUR","cu":"pt"},
        {"code": "ro", "language": "Romanian", "translated": "Română", "chineseNawme": "罗马尼亚语","cu_code":"RON","cu":"ro"},
        {"code": "ru", "language": "Russian", "translated": "Русский", "chineseNawme": "俄语","cu_code":"RUB","cu":"ru"},
        {"code": "sk", "language": "Slovak", "translated": "Slovenčina", "chineseNawme": "斯洛伐克语","cu_code":"EUR","cu":"sk"},
        {"code": "sl", "language": "Slovenian", "translated": "Slovenščina", "chineseNawme": "斯洛文尼亚语","cu_code":"EUR","cu":"si"},
        {"code": "sr", "language": "Serbian", "translated": "Српски", "chineseNawme": "塞尔维亚语","cu_code":"RSD","cu":"rs"},
        {"code": "sv", "language": "Swedish", "translated": "Svenska", "chineseNawme": "瑞典语","cu_code":"SEK","cu":"se"},
        {"code": "sw", "language": "Swahili", "translated": "Kiswahili", "chineseNawme": "斯瓦希里语","cu_code":"TZS","cu":"tz"},
        {"code": "ta", "language": "Tamil", "translated": "தமிழ்", "chineseNawme": "泰米尔语","cu_code":"INR","cu":"lk"},
        // {"code": "te", "language": "Telugu", "translated": "తెలుగు", "chineseNawme": "泰卢固语","cu_code":"INR","cu":"in"},
        {"code": "th", "language": "Thai", "translated": "ไทย", "chineseNawme": "泰语","cu_code":"THB","cu":"th"},
        {"code": "tr", "language": "Turkish", "translated": "Türkçe", "chineseNawme": "土耳其语","cu_code":"TRY","cu":"tr"},
        {"code": "uk", "language": "Ukrainian", "translated": "Українська", "chineseNawme": "乌克兰语","cu_code":"UAH","cu":"ua"},
        {"code": "vi", "language": "Vietnamese", "translated": "Tiếng Việt", "chineseNawme": "越南语","cu_code":"VND","cu":"vn"},

        {"code": "zh_CN", "language": "Chinese (China)", "translated": "中文(简体)", "chineseNawme": "中文（简体）","cu_code":"CNY","cu":"cn"},
        {"code": "zh_TW", "language": "Chinese (Taiwan)", "translated": "中文(繁體)", "chineseNawme": "中文（繁体）","cu_code":"TWD","cu":"tw"}
    ]
    )
    
    //c = ["ph","us","gb","es","it","in1","bd","br","ru","jp","de","bg","ad","hr","cz","dk","nl","ee","ph","fi","fr","gr","in2","hu","il","ge","tj","np","am","uz","id","lv","lt","ba","my","in3","no","cn","pl","ro","rs","lk","sk","si","in4","se","tr","ws","vn","ua"]
    //只在我的中=["ar", "am", "ca", "cs", "da", "el", "en", "en_AU", "en_GB", "en_US", "es_419", "et", "fa", "fil", "gu", "he", hi", "hu", "ja", "kn", "ko", "ml", "mr", "ms", "pt_BR", "pt_PT", "sk", "sl", "sr", "sv", "sw", "ta", "te", "th","uk", "vi", "zh_CN", "zh_TW"]
    //只在别人中 = ["ph", "us", "in1", "bd", "ad", "cz", "dk", "ee", "gr", "in2", "il", "ge", "tj", "np", "uz", "ba", "my", "in3", "rs", "lk", "si", "in4", "ws"]
    const show_string = ()=>{
        let a = []
        languageList.value.forEach((el)=>{
           let b ={}
              b.code = el.code
              b.language = el.language
              a.push(b)
        })
        console.log("language:",a)

        chrome.storage.local.get(["exchange_rate"], function (r) {
           let exchange_rate = r.exchange_rate
           let c = []
                exchange_rate.forEach((el)=>{
                    if(el.IsVirtual){
                        return
                    }
                let b ={}
                b.Code = el.Code
                b.Currency = el.Currency
                b.Symbol = el.Symbol
                c.push(b)
              })
            console.log("currency:",c)
        });
    }
    //onMounted(show_string)
    onMounted(()=>{
        // show_string()
    })
    

</script>

<style>
    .lan{
        color: var(--lan-text-color);
    }
    .back-button-warp{
        display: flex;
        width: 100%;
        align-items: center;
        justify-content: center;
    }
    .back-button{
        display: flex;
        background:var(--main-backgroud-color);
        align-items: center;
        justify-content: center;
        padding: 5px;
        font-size: medium;
        margin-top: 60px;
    }
    .back-button:hover{
        background: hsl(from var(--main-backgroud-color) h s calc(l * 1.2));
        border-radius: 5px;
        box-shadow:  0px 1px 2px rgba(0, 0, 0, 0.5);
    }
    .back-button:active {
        background: hsl(from var(--main-backgroud-color) h s calc(l * 0.8));
}
</style>