<template>
    <div>
        <span class="has-text-light">Date update time:{{TimeLastUpdate}},Base currency:{{BaseCurrency}}</span>
    </div>
   <div>
    <!-- <div v-for="item in currencyList" :key="item.Code">
        <span>{{item.Currency}}--{{item.Code}}--{{item.ExchangeRate}}</span>
        <img :src="'https://file.oeoli.org/flags/'+item.Country_code+'.svg'" alt="" width="40" height="30">
        </div> -->
        <table class="table">
            <thead>
              <tr>
                <th><abbr title="Position">No.</abbr></th>
                <th>Country</th>
                <th>Currency</th>
                <th>Country_code</th>
                <th>Symbol</th>
                <th>ExchangeRate</th>
                <th>Flag</th>
              </tr>
            </thead>

            <tbody>
                <tr v-for="(item ,index) in currencyList" :key="item.Code">
                    <th>{{index+1}}</th>
                    <td>{{item.Country}}</td>
                    <td>{{item.Currency}}</td>
                    <td>{{item.Code}}</td>
                    <td>{{item.Symbol}}</td>
                    <td>{{item.ExchangeRate}}</td>
                    <td><img :src="'https://file.oeoli.org/flags/'+item.Country_code+'.svg'" alt="" width="40" height="30"></td>
              
              </tr>
            </tbody>
          </table>
   </div>
</template>

<script setup>
    import { ref,onMounted,onBeforeMount } from 'vue'

    const currencyList = ref([])
    const TimeLastUpdate = ref("")
    const BaseCurrency = ref("")
    onMounted(async()=>{
        let a = await fetch('https://rate-api.oeo.li/v1/exchange?from=web')
        let b = await a.json()
        currencyList.value = b.data
        TimeLastUpdate.value = b.TimeLastUpdate
        BaseCurrency.value = b.base
        console.log(TimeLastUpdate.value,currencyList.value)
    })


</script>