<template>
  <div class="home-page">
    <Nav></Nav>
    <!-- <Note></Note> -->
    <List></List>
 
    <div class="add_country_button">
      <el-button type="warning"  circle @click="$router.push('/add_list')">
        <el-icon><Plus /></el-icon>
      </el-button>
    </div>
  </div>
</template>

  <script setup>
    import { ref,onMounted,onBeforeMount } from 'vue'
    import Nav from './Nav.vue'
    import Note from './Note.vue'
    import List from './List.vue'

    onBeforeMount(() => {

    })

  </script>

<style>
  .home-page{
    margin-bottom: 20px;
  }
  .add_country_button{
    position:fixed;
    bottom: 10px;
    right: 10px;
    z-index:99;
  }
  </style>

