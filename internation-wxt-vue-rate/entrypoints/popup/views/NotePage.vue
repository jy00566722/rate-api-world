<template >
    <div class="note-page">
    <Nav darp-e="no"></Nav>
    <el-icon :size="40" color="#cc9911"><Message /></el-icon>
    <h2 style="color: rgb(157, 157, 157);margin:0px 1px">{{t('messageList')}}</h2>
    <div class="note-list">
        <div class="note-item" v-for="item in messageStore.message" :key="item.index">
            <el-alert
            :title="item.context"
            :type="item.type"
            show-icon
            >
            </el-alert>
        </div>
    </div>
    <el-button type="primary" text plain @click="$router.push('/')">Back</el-button>

    </div>
</template>

<script setup>
import { ref,onMounted,onBeforeMount } from 'vue'
import Nav from './Nav.vue'
import { useMessageStore } from '../store/message'
import {useI18n} from 'vue-i18n';

const { t} = useI18n();
// locale.value = 'zh_CN'


 
</script>

<style>
    .note-page{
        /* background-color: aquamarine; */
        margin-bottom: 30px;
    }
    .note-list{
        padding: 10px 10px;
    }
    .note-item{
        margin: 10px 5px 10px 5px;
    }
    
</style>
