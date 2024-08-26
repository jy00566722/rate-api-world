<template>
    <div>
      <UForm :state="form" @submit="onSubmit">
        <div class="flex items-center space-x-4">
          <USelectMenu
            v-model="form.fromCurrency"
            :options="currencyOptions"
            value-attribute="Code"
            class="w-1/3"
            @change="onCurrencyChange('from')"
          >
            <template #label>
              <div v-if="form.fromCurrency" class="flex items-center">
                <img :src="`/assets/4x3/${findCurrency(form.fromCurrency).Country_code.toLowerCase()}.svg`" :alt="findCurrency(form.fromCurrency).Country" class="w-6 h-4 mr-2" />
                {{ findCurrency(form.fromCurrency).Currency }} / {{ form.fromCurrency }}
              </div>
            </template>
            <template #option="{ option }">
              <div class="flex items-center">
                <img :src="`/assets/4x3/${option.Country_code.toLowerCase()}.svg`" :alt="option.Country" class="w-6 h-4 mr-2" />
                {{ option.Currency }} / {{ option.Code }}
              </div>
            </template>
          </USelectMenu>
  
          <UInput
            v-model="form.fromAmount"
            type="text"
            inputmode="decimal"
            placeholder="Enter amount"
            class="w-1/4"
            @input="onAmountChange('from')"
          />
  
          <UButton @click="swapCurrencies" class="p-2">
            <UIcon name="i-heroicons-arrow-right-left" class="w-5 h-5" />
          </UButton>
  
          <UInput
            v-model="form.toAmount"
            type="text"
            inputmode="decimal"
            placeholder="Converted amount"
            class="w-1/4"
            @input="onAmountChange('to')"
          />
  
          <USelectMenu
            v-model="form.toCurrency"
            :options="currencyOptions"
            value-attribute="Code"
            class="w-1/3"
            @change="onCurrencyChange('to')"
          >
            <template #label>
              <div v-if="form.toCurrency" class="flex items-center">
                <img :src="`/4x3/${findCurrency(form.toCurrency).Country_code.toLowerCase()}.svg`" :alt="findCurrency(form.toCurrency).Country" class="w-6 h-4 mr-2" />
                {{ findCurrency(form.toCurrency).Currency }} / {{ form.toCurrency }}
              </div>
            </template>
            <template #option="{ option }">
              <div class="flex items-center">
                <img :src="`/4x3/${option.Country_code.toLowerCase()}.svg`" :alt="option.Country" class="w-6 h-4 mr-2" />
                {{ option.Currency }} / {{ option.Code }}
              </div>
            </template>
          </USelectMenu>
        </div>
      </UForm>
    </div>
  </template>
  
  <script setup>
  import { ref, computed, onMounted } from 'vue'
  import { useAsyncData } from '#app'
  
  const form = ref({
    fromCurrency: 'USD',
    toCurrency: 'CNY',
    fromAmount: '1',
    toAmount: ''
  })
  
  const { data: currencyData } = await useAsyncData('currencies', () => 
    $fetch('https://rate-api.oeo.li/v1/exchange')
  )
  
  const currencyOptions = computed(() => currencyData.value?.data || [])
  
  const findCurrency = (code) => currencyOptions.value.find(c => c.Code === code) || {}
  
  const convertCurrency = (direction) => {
    const fromCurrency = findCurrency(form.value.fromCurrency)
    const toCurrency = findCurrency(form.value.toCurrency)
  
    if (!fromCurrency.ExchangeRate || !toCurrency.ExchangeRate) return
  
    const fromAmount = parseFloat(form.value.fromAmount)
    const toAmount = parseFloat(form.value.toAmount)
  
    if (direction === 'from' && !isNaN(fromAmount)) {
      form.value.toAmount = ((fromAmount * toCurrency.ExchangeRate) / fromCurrency.ExchangeRate).toFixed(3)
    } else if (direction === 'to' && !isNaN(toAmount)) {
      form.value.fromAmount = ((toAmount * fromCurrency.ExchangeRate) / toCurrency.ExchangeRate).toFixed(3)
    }
  }
  
  const swapCurrencies = () => {
    [form.value.fromCurrency, form.value.toCurrency] = [form.value.toCurrency, form.value.fromCurrency]
    [form.value.fromAmount, form.value.toAmount] = [form.value.toAmount, form.value.fromAmount]
  }
  
  const onSubmit = () => {
    // 可以在这里添加表单提交逻辑
  }
  
  const onCurrencyChange = () => {
    convertCurrency('from')
  }
  
  const onAmountChange = (direction) => {
    convertCurrency(direction)
  }
  
  onMounted(() => {
    convertCurrency('from')
  })
  </script>