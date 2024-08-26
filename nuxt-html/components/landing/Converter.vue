<template>
    <div class="mt-0 mb-16">
        <h1 class="text-2xl font-bold pb-8 lg:text-3xl"> Currency Converter</h1>
        <div class="flex">
            <USelectMenu
            searchable
            v-model="form"
            :options="D2"
            value-attribute="Code"
            class="w-1/3"
            size="xl"
            @change="onCurrencyChange()"
          >
            <template #label>
              <div v-if="form.Code" class="flex items-center">
                <img :src="`/4x3/${form.Country_code.toLowerCase()}.svg`" :alt="form.Country" class="w-8 h-6 mr-2" />
                {{ form.Currency }} / {{ form.Code }}
              </div>
            </template>
            <template #option="{ option }">
              <div class="flex items-center">
                <img :src="`/4x3/${option.Country_code.toLowerCase()}.svg`" :alt="option.Country" class="w-8 h-6 mr-2" />
                {{ option.Currency }} / {{ option.Code }}
              </div>
            </template>
          </USelectMenu>
  
          <UInput
            v-model="form.num"
            type="text"
            inputmode="decimal"
            placeholder="Enter amount"
            class="w-1/4"
            size="xl"
            @input="onAmountChange('from')"
          />
  
          <UButton @click="swapCurrencies" class="p-2"  size="xl">
            <UIcon name="i-heroicons-arrow-right-left" class="w-5 h-5" />
          </UButton>
  
          <UInput
            v-model="form.num"
            type="text"
            inputmode="decimal"
            placeholder="Converted amount"
            class="w-1/4"
            size="xl"
            @input="onAmountChange('to')"
          />
  
          <USelectMenu
            v-model="to"
            :options="D2"
            value-attribute="Code"
            class="w-1/3"
            size="xl"
            @change="onCurrencyChange()"
          >
            <template #label>
              <div v-if="to.Code" class="flex items-center">
                <img :src="`/4x3/${to.Country_code.toLowerCase()}.svg`" :alt="to.Country" class="w-8 h-6 mr-2" />
                {{ to.Currency }} / {{ to.Code }}
              </div>
            </template>
            <template #option="{ option }">
              <div class="flex items-center">
                <img :src="`/4x3/${option.Country_code.toLowerCase()}.svg`" :alt="option.Country" class="w-8 h-6 mr-2" />
                {{ option.Currency }} / {{ option.Code }}
              </div>
            </template>
          </USelectMenu>
             
        </div>
    </div>
    
</template>

<script setup lang="ts">
//https://rate-api.oeo.li/v1/exchange
//https://rate-api.oeo.li/v1/exchange_crypto
console.log('开始了')
const form = ref({
      "Country": "VIRGIN ISLANDS (U.S.)",
      "Currency": "US Dollar",
      "Code": "USD",
      "Country_code": "us",
      "Symbol": "$",
      "ExchangeRate": 1,
      "IsVirtual": false,
      "FlagURL": "",
      "UpdateTime": "e2024-08-01 08:00:01",
      "num": ""
  })
const to = ref(    {
      "Country": "CHINA",
      "Currency": "Yuan Renminbi",
      "Code": "CNY",
      "Country_code": "cn",
      "Symbol": "¥",
      "ExchangeRate": 7.2106,
      "IsVirtual": false,
      "FlagURL": "",
      "UpdateTime": "g2024-08-01 09:40:20",
      "num": ""
    },)

const D1 = useState("rate_list_raw")
await callOnce(async () => {
  D1.value = await $fetch('https://rate-api.oeo.li/v1/exchange')
})
console.log(D1)
// const res = await useFetch('https://rate-api.oeo.li/v1/exchange')
const D2: any[] = []
D1.data.forEach((el: { IsVirtual: any; num: string; }) => {
    if(!el.IsVirtual){
        el.num = ''
        D2.push(el)
    }
});

const onAmountChange = (p0: string) => {
    
}
const onCurrencyChange = () => {
    
}
const swapCurrencies = () => {

}
// console.log("D2:",D2.value)
if(import.meta.client){
    // console.log("hello-client:",useState("rate_list").value)
    console.log("D2:",D2)
}

</script>