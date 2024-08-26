<template>
    <div class="setting">
        <Nav darp-e="no"></Nav>
        <el-icon :size="40" aria-colcount="setting-icon">
            <UserFilled />
        </el-icon>
        <h2 class="setting-title">{{t('setting')}}</h2>
        <el-row :gutter="6">
            <el-col :span="12">
                <el-card class="setting_c1" style="width: 100%;margin-top: 5px;">
                    <template #header>
                        <div class="card-header">
                            <span style="font-size: 14px;">{{t('settingText1')}}</span>
                        </div>
                    </template>
        
                    <span style="font-size: 14px;">{{t('settingText2')}}</span>
                    <el-input-number v-model="decimal" :step="1" step-strictly  :min="2" :max="6"
                        @change="handlChange" />
        
                </el-card>
            </el-col>
            <el-col :span="12">
                <el-card class="setting_c2" style="width: 100%;margin-top: 5px;margin-bottom: 5px;">
                    <template #header>
                        <div class="card-header">
                            <span style="font-size: 14px;">{{t('settingText3')}}</span>
                        </div>
                    </template>
                    <el-radio-group v-model="currencyFormat" @change="handlChange" style="justify-content: start;">
                        <el-radio label="./"  border>1234567.89</el-radio>
                        <el-radio label="./,"  border>1,234,567.89</el-radio>
                        <el-radio label="./ "  border>1 234 567.89</el-radio>
                        <!-- <el-radio label=",/."  border> 1.234.567,89</el-radio>
                        <el-radio label=",/ "  border>1 234 567,89</el-radio> -->
                        <div> <span class="setting-exp">{{currency}}</span></div>
                       
                    </el-radio-group>
        
            </el-card>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="24">
                <el-card class="setting_c3" style="width: 100%;margin-top: 0px;margin-bottom: 10px;">
                    <template #header>
                        <div class="card-header">
                           
                                <span style="font-size: 18px;">Theme Color</span>
                                <svg xmlns="http://www.w3.org/2000/svg" width="64" height="64" viewBox="0 0 512 512">
                                <defs>
                                    <linearGradient id="settinggrad1" x1="0%" y1="0%" x2="100%" y2="0%">
                                        <stop offset="0%" style="stop-color:var(--theme-icon-color1);stop-opacity:1" />
                                        <stop offset="100%" style="stop-color:var(--theme-icon-color2);stop-opacity:1" />
                                        <!-- <stop offset="100%" style="stop-color:var(--theme-icon-color3);stop-opacity:1" /> -->
                                    </linearGradient>
                                  </defs>
                                <path fill="url(#settinggrad1)" d="M416 352c-12.6-.84-21-4-28-12c-14-16-14-36 5.49-52.48l32.82-29.14c50.27-44.41 50.27-117.21 0-161.63C389.26 64.14 339.54 48 287.86 48c-60.34 0-123.39 22-172 65.11c-90.46 80-90.46 210.92 0 290.87c45 39.76 105.63 59.59 165.64 60h1.84c60 0 119.07-19.5 161.2-56.77C464 390 464 385 444.62 355.56C440 348 431 353 416 352M112 208a32 32 0 1 1 32 32a32 32 0 0 1-32-32m40 135a32 32 0 1 1 32-32a32 32 0 0 1-32 32m40-199a32 32 0 1 1 32 32a32 32 0 0 1-32-32m64 271a48 48 0 1 1 48-48a48 48 0 0 1-48 48m72-239a32 32 0 1 1 32-32a32 32 0 0 1-32 32" />
                            </svg>
                        </div>
                    </template>
                    <el-radio-group v-model="theme" @change="themeChange">
                        <el-radio-button label="auto" value="auto" />
                        <el-radio-button label="light" value="light" />
                        <el-radio-button label="dark" value="dark" />
                        <el-radio-button label="navy" value="navy" />
                        <el-radio-button label="green" value="green" />
                      
                      </el-radio-group>
        
            </el-card>
            </el-col>
            <!-- <el-col :span="12">2</el-col> -->
          </el-row>

    <div class="back-button-warp">
    <span class="back-button" style="margin-top: 10px;" @click="$router.push('/')">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><g fill="none"><path d="M24 0v24H0V0zM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035c-.01-.004-.019-.001-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427c-.002-.01-.009-.017-.017-.018m.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093c.012.004.023 0 .029-.008l.004-.014l-.034-.614c-.003-.012-.01-.02-.02-.022m-.715.002a.023.023 0 0 0-.027.006l-.006.014l-.034.614c0 .012.007.02.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01z"/><path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2m0 2a8 8 0 1 0 0 16a8 8 0 0 0 0-16M9.879 8.464L12 10.586l2.121-2.122a1 1 0 1 1 1.415 1.415l-2.122 2.12l2.122 2.122a1 1 0 0 1-1.415 1.415L12 13.414l-2.121 2.122a1 1 0 0 1-1.415-1.415L10.586 12L8.465 9.879a1 1 0 0 1 1.414-1.415"/></g></svg>
        Back</span></div><br>
   
        <span class="setting-title">Powered by</span> <a href="https://currency.oeo.li/" target="_blank" style=" text-decoration: none;"> <span class="setting-title">https://currency.oeo.li/</span> </a>
    </div>
</template>

<script setup>
    import { ref, onMounted, onBeforeMount,computed } from 'vue'
    import Nav from './Nav.vue'
    import {useI18n} from 'vue-i18n';
    const { t } = useI18n();
    const theme = ref('auto')
    const decimal = ref(2) //POP页面价格转换时的小数点位数
    const currencyFormat = ref("./,") //POP页面价格转换时货币格式
    const temp = ref(1234567.8912324567)

    //用计算属性实现货币格式转换，传入temp的值，用customFormatCurrency处理后返回
    const currency = computed(() => {
        return customFormatCurrency(temp.value, currencyFormat.value, decimal.value)
    })

    function customFormatCurrency(value, formatString, decimalPlaces) {
    if (isNaN(value)) {
        return 'Invalid number';
    }

    // 提取小数分隔符和千分位分隔符
    const [decimalSeparator, thousandsSeparator] = formatString.split('/');

    // 四舍五入到指定的小数位数
    const factor = Math.pow(10, decimalPlaces);
    value = Math.round(value * factor) / factor;

    // 将数字转换为字符串，并用小数点分割整数和小数部分
    let [integerPart, decimalPart] = value.toFixed(decimalPlaces).split('.');

    // 添加千分位分隔符（如果定义了千分位分隔符）
    if (thousandsSeparator && thousandsSeparator !== ' ') {
        const regex = new RegExp(`\\B(?=(\\d{3})+(?!\\d))`, 'g');
        integerPart = integerPart.replace(regex, thousandsSeparator);
    } else if (thousandsSeparator === ' ') {
        // 如果千分位分隔符是空格，我们需要稍微调整正则表达式
        integerPart = integerPart.replace(/\B(?=(\d{3})+(?!\d))/g, '\u00A0');
    }

    // 将整数部分和小数部分重新组合，并替换小数点符号
    return `${integerPart}${decimalSeparator}${decimalPart}`;
}

    const saveDecimal = () => {
        chrome.storage.local.set({ decimal: decimal.value }, function () {
            console.log('decimal : ' + decimal.value);
        });
    }
    const handlChange = () => {

        chrome.storage.local.set({ decimal: decimal.value ,currencyFormat:currencyFormat.value}, function () {
            // console.log('decimal位数是 : ' + decimal.value);
        });
    }
    const fetchData = () => {
        //给后台发消息
        chrome.runtime.sendMessage({ do: "getStart" }, function (response) {
            // console.log(response.farewell);
        });
    }
    const themeChange=()=>{
            chrome.storage.local.set({theme: theme.value}, function() {
                console.log('theme is ' + theme.value);
            });
            //根据theme.value设置主题
            if(theme.value==='auto'){
               
                let theme_ = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
                document.documentElement.setAttribute('data-theme',theme_);
            }else{
                document.documentElement.setAttribute('data-theme',theme.value)
            }
    }
    onMounted(async () => {
        // selectedColor.value = document.querySelector('#selectedColor')
        const response = await chrome.storage.local.get(['decimal','currencyFormat','theme']);

        if (typeof response.decimal == 'number' && response.decimal >= 2 && response.decimal <= 6) {
            decimal.value = response.decimal
        } else {
            decimal.value = 2
            //删除原来的decimal
            await chrome.storage.local.remove('decimal')
        }
        if (response.currencyFormat) {
            currencyFormat.value = response.currencyFormat
        } else {
            currencyFormat.value = "./,"
            //删除原来的decimal
            // await chrome.storage.local.remove('currencyFormat')
        }
        if (response.theme) {
            theme.value = response.theme
        } else {
            theme.value = 'auto'
        }

    })
</script>

<style>
    .setting {
        margin-bottom: 10px;
        color: var(--setting-icon-color);
    }
    .setting-exp{
        font-size: 18px;
        font-weight: bold;
        color: var(--setting-exp-text-color);
    }
    .setting-title{
        color:var(--setting-title-text-color);
        margin:0px 1px;
    }
    .setting .el-card{

        background: var(--main-backgroud-color);
        border-radius: 10px;
        border:var(--setting-card-border-color) 1px solid;
        color: var(--setting-card-border-color);
        --el-card-border-color: var(--setting-card-border-color);
        --el-card-border-radius: 10px;
        /* --el-card-bg-color: #000000; */
        --el-card-padding: 20px;
    }
    .setting_c3 .el-card__header{
        padding: 0;
    }
    .setting_c3 .card-header{
        font-weight: 900;
        display: flex;
        align-items: center;
        justify-content: center; /* 如果需要水平居中 */
        /* height: 200px; 设置容器高度 */
    }
</style>