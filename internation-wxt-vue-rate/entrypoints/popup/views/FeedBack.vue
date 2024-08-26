<template>
    <div class="feedback">
    <Nav darp-e="no"></Nav>
    <h1>{{t('QuestionsAndSuggestions')}}
      <el-icon color="#9911ff"><MagicStick /></el-icon>
    </h1>
    <h2>{{t('YourFeedbackIsVeryImportant')}}</h2>
    <div class="feedback-form">
    <el-form
    label-position="top"
    label-width="100px"
    :model="form"
    style="max-width: 460px"
  >
    <el-form-item :label="t('YourNameandContactInformation')">
      <el-input v-model="form.name" />
    </el-form-item>
    <el-form-item :label="t('YourQuestion')">
      <el-input v-model="form.url" 
      :rows="2"
      maxlength="300"
      show-word-limit
      type="textarea"/>
    </el-form-item>
    <el-form-item :label="t('OrYourSuggestion')">
      <el-input v-model="form.context" 
      :rows="3"
      maxlength="500"
      show-word-limit
      type="textarea"
      />
    </el-form-item>
    <el-row>
        <el-col :span="12"><el-button type="primary" text plain @click="postFeedBack">

            <svg xmlns="http://www.w3.org/2000/svg"  width="28" height="28" viewBox="0 0 19 19"><path fill="currentColor" d="M19 5.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0m-4.146-2.353l-.003-.003A.498.498 0 0 0 14.503 3h-.006a.499.499 0 0 0-.35.147l-2 2a.5.5 0 0 0 .707.707L14 4.707V7.5a.5.5 0 1 0 1 0V4.707l1.146 1.147a.5.5 0 1 0 .707-.708zM17 14.5v-4.1a5.507 5.507 0 0 0 1-.657V14.5a2.5 2.5 0 0 1-2.5 2.5h-11A2.5 2.5 0 0 1 2 14.5v-8A2.5 2.5 0 0 1 4.5 4h4.707a5.48 5.48 0 0 0-.185 1H4.5A1.5 1.5 0 0 0 3 6.5v.302l7 4.118l1.441-.848a5.49 5.49 0 0 0 1.043.547l-2.23 1.312a.5.5 0 0 1-.426.038l-.082-.038L3 7.963V14.5A1.5 1.5 0 0 0 4.5 16h11a1.5 1.5 0 0 0 1.5-1.5"/></svg>
       
            Submit
        </el-button></el-col>
        <el-col :span="12"><el-button type="primary"  text plain @click="$router.push('/')">
          <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24"><g fill="none"><path d="M24 0v24H0V0zM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035c-.01-.004-.019-.001-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427c-.002-.01-.009-.017-.017-.018m.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093c.012.004.023 0 .029-.008l.004-.014l-.034-.614c-.003-.012-.01-.02-.02-.022m-.715.002a.023.023 0 0 0-.027.006l-.006.014l-.034.614c0 .012.007.02.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01z"/><path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2m0 2a8 8 0 1 0 0 16a8 8 0 0 0 0-16M9.879 8.464L12 10.586l2.121-2.122a1 1 0 1 1 1.415 1.415l-2.122 2.12l2.122 2.122a1 1 0 0 1-1.415 1.415L12 13.414l-2.121 2.122a1 1 0 0 1-1.415-1.415L10.586 12L8.465 9.879a1 1 0 0 1 1.414-1.415"/></g></svg>
          Back</el-button></el-col>
      </el-row>


  </el-form>
</div>
</div>
</template>

<script setup>
    import { ref,onMounted,onBeforeMount ,reactive} from 'vue'
    import Nav from './Nav.vue'
    import { ElMessage } from 'element-plus'
    import {useI18n} from 'vue-i18n';
    const { t,locale } = useI18n();

    const form = reactive({
        name:"",//用户名称
        url:"", //问题URL
        context:"",
        time:""
    })
    const iconsSize = ref(30) 
    const postFeedBack=async ()=>{
      if(form.name==""){
        ElMessage.error(t('provideCompleteInformation'))
        return
      }
      form.time = new Date().toLocaleString()
      let url = 'https://rate-api.oeo.li/v2/feedback'
      try{
        const a = await fetch(url,{
                method:'POST',
                body:JSON.stringify(form),
                headers: {
                  'content-type': 'application/json'
                }
          })
      const b = await a.json()
      if(b.code===20000&b.msg==='success'){
          form.name = ''
          form.url = ''
          form.context= ''
          form.time = ''
          ElMessage({message:t('FeedbackSuccessful'),type:'success'})
        }else{
          form.name = ''
          form.url = ''
          form.context= ''
          form.time = ''
          ElMessage({message:t('FeedbackFailedPleaseTryAgain'),type:'warning'})
        }
      }catch(e){
          ElMessage.error(t('PleaseCheckYourNetwork'))
        return
      }

    }

</script>

<style>
    .feedback{
        background-color: rgb(47, 64, 93);
        color: rgb(194, 194, 194);
    }
    .feedback h2{
        color: rgb(75, 171, 171);
    }
    .feedback-form {
        padding: 2px 20px 30px 20px;
    }
    .feedback-form .el-form-item__label{
        color: rgb(224, 135, 18);
    }
</style>