<template>
    <section class="hero is-small ">
        <div class="hero-head">
          <div class="container">
            <nav class="navbar" role="navigation" aria-label="main navigation">
    
              <div id="navbarBasicExample" class="navbar-menu">
                <div class="navbar-start">
                  <a class="navbar-item is-active">
                    Home
                  </a>
                  <a class="navbar-item" 
                  target="_blank"
                  href="https://chromewebstore.google.com/detail/%E8%B4%A7%E5%B8%81%E8%BD%AC%E6%8D%A2plus/loeilaonggnalkaiiaepbegccilkmjjp">
                    ChromeWebStore
                  </a>
                  <a class="navbar-item" 
                  target="_blank"
                  href="https://microsoftedge.microsoft.com/addons/detail/%E8%B4%A7%E5%B8%81%E8%BD%AC%E6%8D%A2plus/bahgpcdmknmcfgebpbebpmkmkhidiikh">
                    EdgeWebStore
                  </a>
                </div>
    
                <div class="navbar-end">
                  <div class="navbar-item">
                    <a href="#">
                      <i class="fab fa-facebook"></i>
                    </a>
                    <a href="#">
                      <i class="fab fa-twitter"></i>
                    </a>
                    <a href="#">
                      <i class="fas fa-rss-square"></i>
                    </a>
                  </div>
                </div>
              </div>
            </nav>
          </div>
        </div>
        <div class="hero-body">
          <div class="container has-text-centered">
            <img src="https://file.oeoli.org/256.png" width="100" >
            <h2 class="subtitle">
                Currency Converter Plus is an easy-to-use tool for promptly checking current exchange rates, and free.
            </h2>
            <h2 class="subtitle">
                Make currency conversion a breeze— and it's all for free! Our brand-new Chrome extension is a comprehensive currency conversion tool, designed specifically for users who frequently deal with various currencies. Whether you're traveling around the world, engaged in international trade, or simply interested in different currencies, our currency converter is your ideal companion. This extension utilizes the day's exchange rates to convert between 155 currencies and offers options in 47 languages, enabling you to effortlessly perform currency conversions wherever you are.
            </h2>
          </div>
        </div>
      </section>
      <section class="blog-posts">
        <div class="container">
          <div class="columns">
            <div class="column is-10 is-offset-1">
              <div class="columns featured-post is-multiline">
                <div class="column is-12 post">
                    <hr>
                    <Currenry></Currenry>

                </div>
              </div>
              <hr>
              <div class="columns is-multiline">

              </div>
            </div>
          </div>
        </div>
      </section>
      <Foot></Foot>
</template>

<script setup>
    import { ref,onMounted,onBeforeMount } from 'vue'
    import Currenry from './Currenry.vue'
    import Foot from './Foot.vue'





</script>