<template>
  <div class="nav">
    <div class="nav_up">

      <div class="image-shadow">
        <!-- <div class="image-blur"></div> -->
        <img src="/icons/logo.svg" width="45" class="head_start logo" @click="$router.push('/')">
      </div>

      <el-tooltip class="box-item" effect="dark" :content="t('facebookTip')" placement="bottom">
        <span class="head_start" 
        style="padding: 0 5px;"
       ><a href="https://twitter.com/Currency_plus" target="_blank">
        <svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" viewBox="0 0 1536 1536"
        style="vertical-align: middle;">
          <path class="svg-icon" d="M1280 482q-56 25-121 34q68-40 93-117q-65 38-134 51q-61-66-153-66q-87 0-148.5 61.5T755 594q0 29 5 48q-129-7-242-65T326 422q-29 50-29 106q0 114 91 175q-47-1-100-26v2q0 75 50 133.5T461 885q-29 8-51 8q-13 0-39-4q21 63 74.5 104t121.5 42q-116 90-261 90q-26 0-50-3q148 94 322 94q112 0 210-35.5t168-95t120.5-137t75-162T1176 618q0-18-1-27q63-45 105-109m256-194v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288" />
        </svg>
       </a></span>
      </el-tooltip>

        <span  @click="fetchData()" v-if="darpE=='yes'" class="head_start head_start_m" 
          >{{t('dragAndSort')}}</span>

      <!-- </el-tooltip> -->
      <el-tooltip v-if="darpE=='no'" class="box-item" effect="dark" :content="t('backHomeTip')" placement="bottom">
        <span class="head_start" @click="$router.push('/')"><el-icon style="vertical-align: middle" :size="iconsSize">
            <!-- <HomeFilled /> -->
            <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 48 48"><path fill="currentColor" d="M21.606 5.855a3.778 3.778 0 0 1 4.788 0l13.175 10.742A3.91 3.91 0 0 1 41 19.628v19.485C41 41.26 39.288 43 37.175 43h-5.35C29.712 43 28 41.26 28 39.113v-9.226c0-.716-.57-1.296-1.275-1.296h-5.45c-.704 0-1.275.58-1.275 1.296v9.226C20 41.26 18.288 43 16.175 43h-5.35C8.713 43 7 41.26 7 39.113V19.628a3.91 3.91 0 0 1 1.431-3.031L21.606 5.855Zm3.192 2.021a1.26 1.26 0 0 0-1.596 0L10.027 18.617a1.304 1.304 0 0 0-.477 1.011v19.485c0 .716.57 1.296 1.275 1.296h5.35c.704 0 1.275-.58 1.275-1.296v-9.226c0-2.147 1.713-3.887 3.825-3.887h5.45c2.112 0 3.825 1.74 3.825 3.887v9.226c0 .716.57 1.296 1.275 1.296h5.35c.704 0 1.275-.58 1.275-1.296V19.628a1.3 1.3 0 0 0-.477-1.01L24.798 7.875Z"/></svg>
          </el-icon></span>
      </el-tooltip>

      <el-tooltip class="box-item" effect="dark" :content="t('languageTip')" placement="bottom">
        <span @click="$router.push('/language')" class="head_start"><el-icon style="vertical-align: middle"
            :size="iconsSize"
           >
            <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="M12 21q-1.858 0-3.5-.71t-2.86-1.93q-1.22-1.217-1.93-2.86Q3 13.858 3 12q0-1.863.71-3.503q.711-1.64 1.93-2.858Q6.857 4.421 8.5 3.711Q10.142 3 12 3q1.863 0 3.503.71q1.64.711 2.858 1.93q1.218 1.218 1.928 2.857Q21 10.137 21 12q0 1.858-.71 3.5t-1.93 2.86q-1.218 1.219-2.857 1.93Q13.863 21 12 21Zm0-.992q.88-1.131 1.452-2.221q.571-1.09.929-2.44H9.619q.396 1.426.948 2.516q.552 1.09 1.433 2.145Zm-1.273-.15q-.7-.825-1.278-2.04q-.578-1.214-.86-2.472H4.753q.86 1.866 2.437 3.06q1.578 1.194 3.536 1.452Zm2.546 0q1.958-.258 3.536-1.452t2.437-3.06h-3.834q-.38 1.277-.957 2.492q-.578 1.214-1.182 2.02Zm-8.927-5.512h4.035q-.114-.615-.16-1.199q-.048-.584-.048-1.147q0-.563.047-1.147q.047-.584.16-1.2H4.347q-.163.52-.255 1.133T4 12q0 .602.091 1.214q.092.613.255 1.132Zm5.035 0h5.238q.114-.615.16-1.18q.048-.564.048-1.166t-.047-1.166q-.047-.565-.16-1.18H9.38q-.113.615-.16 1.18q-.047.564-.047 1.166t.047 1.166q.047.565.16 1.18Zm6.238 0h4.035q.163-.52.255-1.132T20 12q0-.602-.091-1.214q-.092-.613-.255-1.132h-4.035q.114.615.16 1.199q.048.584.048 1.147q0 .563-.047 1.147q-.047.584-.16 1.2Zm-.207-5.692h3.834q-.879-1.904-2.408-3.06q-1.53-1.156-3.565-1.47q.7.92 1.259 2.106q.558 1.185.88 2.424Zm-5.793 0h4.762q-.396-1.408-.977-2.546q-.58-1.139-1.404-2.116q-.823.977-1.404 2.116q-.58 1.138-.977 2.546Zm-4.865 0h3.834q.322-1.239.88-2.424q.559-1.186 1.259-2.107q-2.054.315-3.574 1.48q-1.52 1.166-2.4 3.05Z"/></svg>
          </el-icon></span>
      </el-tooltip>
      <el-tooltip class="box-item" effect="dark" :content="t('settingTip')" placement="bottom">
        <span @click="$router.push('/setting')" class="head_start"><el-icon style="vertical-align: middle"
            :size="iconsSize">
            <!-- <Setting /> -->
            <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="M10.125 22q-.375 0-.65-.25t-.325-.625l-.3-2.325q-.325-.125-.613-.3t-.562-.375l-2.175.9q-.35.15-.7.038t-.55-.438L2.4 15.4q-.2-.325-.125-.7t.375-.6l1.875-1.425Q4.5 12.5 4.5 12.337v-.674q0-.163.025-.338L2.65 9.9q-.3-.225-.375-.6t.125-.7l1.85-3.225q.2-.325.55-.438t.7.038l2.175.9q.275-.2.575-.375t.6-.3l.3-2.325q.05-.375.325-.625t.65-.25h3.75q.375 0 .65.25t.325.625l.3 2.325q.325.125.613.3t.562.375l2.175-.9q.35-.15.7-.038t.55.438L21.6 8.6q.2.325.125.7t-.375.6l-1.875 1.425q.025.175.025.338v.674q0 .163-.05.338l1.875 1.425q.3.225.375.6t-.125.7l-1.85 3.2q-.2.325-.563.45t-.712-.025l-2.125-.9q-.275.2-.575.375t-.6.3l-.3 2.325q-.05.375-.325.625t-.65.25h-3.75ZM11 20h1.975l.35-2.65q.775-.2 1.438-.588t1.212-.937l2.475 1.025l.975-1.7l-2.15-1.625q.125-.35.175-.737T17.5 12q0-.4-.05-.787t-.175-.738l2.15-1.625l-.975-1.7l-2.475 1.05q-.55-.575-1.212-.962t-1.438-.588L13 4h-1.975l-.35 2.65q-.775.2-1.437.588t-1.213.937L5.55 7.15l-.975 1.7l2.15 1.6q-.125.375-.175.75t-.05.8q0 .4.05.775t.175.75l-2.15 1.625l.975 1.7l2.475-1.05q.55.575 1.213.963t1.437.587L11 20Zm1.05-4.5q1.45 0 2.475-1.025T15.55 12q0-1.45-1.025-2.475T12.05 8.5q-1.475 0-2.488 1.025T8.55 12q0 1.45 1.012 2.475T12.05 15.5ZM12 12Z"/></svg>
          </el-icon></span>
      </el-tooltip>
      <el-tooltip class="box-item" effect="dark" content="Like" placement="bottom">
        <span @click="goLike" class="head_start"><el-icon style="vertical-align: middle" :size="iconsSize">

            <svg xmlns="http://www.w3.org/2000/svg" width="128" height="128" viewBox="0 0 24 24">
              <path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M14 10h4.764a2 2 0 0 1 1.789 2.894l-3.5 7A2 2 0 0 1 15.263 21h-4.017c-.163 0-.326-.02-.485-.06L7 20m7-10V5a2 2 0 0 0-2-2h-.095c-.5 0-.905.405-.905.905a3.61 3.61 0 0 1-.608 2.006L7 11v9m7-10h-2M7 20H5a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h2.5" />
            </svg>
          </el-icon></span>
      </el-tooltip>
      <el-tooltip class="box-item" effect="dark" :content="t('popupTip')" placement="bottom">
        <span @click="popWindows" class="head_start"><el-icon style="vertical-align: middle" :size="iconsSize">
            <!-- <TopRight /> -->
            <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 32 32"><path fill="currentColor" d="M28 4H10a2.006 2.006 0 0 0-2 2v14a2.006 2.006 0 0 0 2 2h18a2.006 2.006 0 0 0 2-2V6a2.006 2.006 0 0 0-2-2Zm0 16H10V6h18Z"/><path fill="currentColor" d="M18 26H4V16h2v-2H4a2.006 2.006 0 0 0-2 2v10a2.006 2.006 0 0 0 2 2h14a2.006 2.006 0 0 0 2-2v-2h-2Z"/></svg>
          </el-icon></span>
      </el-tooltip>
    </div>
  </div>
</template>
<script setup>
  import { ref, onMounted, onUnmounted, watch, computed } from 'vue'
  import { ElMessage } from 'element-plus'
  import {useI18n} from 'vue-i18n';


  const props = defineProps({
    darpE: {
      type: String,
      default: 'yes'
      // required: true props.darpE!='no'
    }
  }) //接受是否显示拖动提示的flag
  const iconsSize = ref(30) 
  //i18n
  const { t,locale,fallbackLocale,mergeLocaleMessage,setLocaleMessage } = useI18n();

  //打开option页面
  const showOptionPage = () => {
    if (chrome.runtime.openOptionsPage) {
      chrome.runtime.openOptionsPage();
    } else {
      window.open(chrome.runtime.getURL('options.html'));
    }
  }

  const fetchData = () => {
        //给后台发消息
        chrome.runtime.sendMessage({ do: "getStart" }, function (response) {
            console.log('fetch data ok.');
        });
    }
  const goLike = ()=>{
    let url ="https://chromewebstore.google.com/detail/loeilaonggnalkaiiaepbegccilkmjjp";
        if (navigator.userAgent.includes("Edg")) {
          url ="https://microsoftedge.microsoft.com/addons/detail/bahgpcdmknmcfgebpbebpmkmkhidiikh";
        }
        window.open(url);
  }
  //弹出桌面
  const popWindows = () => {
    chrome.windows.create({
      url: chrome.runtime.getURL("popup.html"),
      type: "popup",
      width: 390,
      height: 620,
    });
  }
</script>
<style>
  .nav {
    /* position:fixed; */
    top: 0px;
    /* z-index: 1900; */
    display: block;
    /* grid-template-rows: 40% 60%; */
    width: 100%;
    margin-right: 8px;
  }

  .nav_up {
    padding: 5px 0 0 0;
    /* height: 25px; */
    display: flex;
    align-items: center;
    justify-content: space-around;

    background: var(--main-backgroud-color);
  }
  .nav .nav_up .logo{
    /* box-shadow: 3px 3px 10px rgba(107, 74, 237, 0.4); */
  }

  .head_start {
    font-size: 14px;
    color: var(--nav-head-start-text-color);
  }

  .head_start:hover {
    font-weight: 900;
    color: var(--nav-head-start-text-hover-color);
    cursor: pointer;
  }
  .svg-icon {
    fill: var( --nav-head-start-text-color);
  }
  .head_start.head_start_m{
    color: var(--nav-head-start-m-text-color);
    font-family:Verdana, Geneva, Tahoma, sans-serif;
    font-style: oblique;
  }

  .image-shadow {
    position: relative;
    display: inline-block;
  }

  .image-shadow .image-blur {
    position: absolute;
    top: -3px;
    /* 根据需要调整阴影偏移量 */
    left: -3px;
    /* 根据需要调整阴影偏移量 */
    right: -3px;
    /* 根据需要调整阴影偏移量 */
    bottom: -3px;
    /* 根据需要调整阴影偏移量 */
    border-radius: 50%;
    z-index: 5;
    /* background-image: linear-gradient(100deg, #fbfbfb, #F15A24); */
    filter: blur(30px);
    /* 模糊滤镜的程度 */
    transform: scale(0.2);
    /* 根据需要调整缩放比例 */
  }

  /* //F15A24 logo的颜色 */
  .image-shadow img {
    display: block;
    /* width: 50px; */
    /* height: auto; */
    position: relative;
    z-index: 10;
  }
</style>