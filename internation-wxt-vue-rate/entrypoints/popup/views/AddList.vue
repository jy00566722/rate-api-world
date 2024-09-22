<template>
  <div class="add_list">
    <Nav darp-e="no"></Nav>
    <el-row justify="space-between" align="middle">
      <el-col :span="18"> <h3 class="add_list_m">{{t('SelectCurrencyToAddToCalculationList')}}</h3></el-col>
      <el-col :span="6"><span 
        :class="selectedClassObject"
        @click="lookSelected"
        >{{selected_text}}</span></el-col>
    </el-row>
    <div class="add_list_search">
      <el-input size="default" v-model="searchKey" :placeholder="t('SearchForCurrencyCountryCode')" >
          <template #append><span>
            <el-icon><Search /></el-icon>
          
              </span>
            </template>
        </el-input>
      </div>

    <RecycleScroller
      class="add_list_all"
      :items="filteredCountryList"
      :item-size="58"
      key-field="tcur"
    >
      <template v-slot="{ item }">
        <div class="add_list_country_item" @click="changStatus(item)">
          <span><img :src="item.flagURL" alt="" width="40" height="40"></span>
          <span class="add_list_all_s">
            {{item.tcur}}{{" - "}}{{item.currencyName}} <span v-if="item.symbol" class="add_list_symbol">{{" "+item.symbol}}</span>
          </span>
          <span>
            <el-icon v-if="item.selected" :size="20" class="add_list_selected"><Select /></el-icon>
          </span>
        </div>
      </template>
    </RecycleScroller>

    <div class="add_country_button">
      <el-button type="danger"  circle @click="$router.push('/')">Back</el-button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeMount, reactive, computed } from 'vue'
import Nav from './Nav.vue'
import { storeToRefs } from 'pinia'
import { useLanguageStore } from '../store/language.js'
import { useI18n } from 'vue-i18n'
import { RecycleScroller } from 'vue-virtual-scroller'
import 'vue-virtual-scroller/dist/vue-virtual-scroller.css'

const languageStore = useLanguageStore()
const { currencyNames } = storeToRefs(languageStore)
const { t, locale } = useI18n()
const searchKey = ref('')
const country_list = ref([])
const country_selected = ref([])
const country_list_status = ref([])
const selected_status = ref(false)
const selected_text = ref(t('ViewSelected'))
const selectedClassObject = computed(() => ({
  'selected_country': selected_status.value,
  'selected_country_not': !selected_status.value,
}))

const filteredCountryList = computed(() => {
  const baseList = selected_status.value 
    ? country_list_status.value.filter(item => item.selected)
    : country_list_status.value

  if (!searchKey.value) {
    return baseList
  }
  
  const searchRegex = new RegExp(escapeRegExp(searchKey.value), 'i')
  return baseList.filter(item => searchRegex.test(item.searchText))
})

onBeforeMount(() => {
  chrome.storage.local.get(["currency_list_now","exchange_rate_obj","exchange_rate_crypto_obj"], (result) => {
    let { currency_list_now, exchange_rate_obj, exchange_rate_crypto_obj } = result
    let c = []
    let exchange_rate_obj_keys = Object.keys(exchange_rate_obj)
    exchange_rate_obj_keys.forEach((el) => {
      if (exchange_rate_obj[el].IsVirtual) {
        return
      }
      let rate_country = {}
      let currencyNamesSearchText = ''
      if (currencyNames.value[exchange_rate_obj[el].Code]) {
        rate_country.currencyName = currencyNames.value[exchange_rate_obj[el].Code]
        currencyNamesSearchText = currencyNames.value[exchange_rate_obj[el].Code]
      } else {
        rate_country.currencyName = exchange_rate_obj[el].Currency
      }
      rate_country.flagURL = 'icons/flags/4x3/' + exchange_rate_obj[el].Country_code + '.svg'
      rate_country.name = exchange_rate_obj[el].Currency
      rate_country.country = exchange_rate_obj[el].Country
      rate_country.searchText = exchange_rate_obj[el].Currency + " " + exchange_rate_obj[el].Country_code + " " + exchange_rate_obj[el].Country + " " + exchange_rate_obj[el].Code + " " + currencyNamesSearchText
      rate_country.scur = ""
      rate_country.status = "ALREADY"
      rate_country.tcur = exchange_rate_obj[el].Code
      rate_country.symbol = exchange_rate_obj[el].Symbol
      rate_country.selected = currency_list_now.includes(el)
      c.push(rate_country)
    })
    let exchange_rate_crypto_obj_keys = Object.keys(exchange_rate_crypto_obj)
    exchange_rate_crypto_obj_keys.forEach((el) => {
      if (!exchange_rate_crypto_obj[el].IsVirtual) {
        return
      }
      let rate_country = {}
      let currencyNamesSearchText = ''
      if (currencyNames.value[exchange_rate_crypto_obj[el].Code]) {
        rate_country.currencyName = currencyNames.value[exchange_rate_crypto_obj[el].Code]
        currencyNamesSearchText = currencyNames.value[exchange_rate_crypto_obj[el].Code]
      } else {
        rate_country.currencyName = exchange_rate_crypto_obj[el].Currency
      }
      rate_country.flagURL = exchange_rate_crypto_obj[el].FlagURL
      rate_country.name = exchange_rate_crypto_obj[el].Currency
      rate_country.country = exchange_rate_crypto_obj[el].Country
      rate_country.searchText = exchange_rate_crypto_obj[el].Currency + " " + exchange_rate_crypto_obj[el].Country_code + " " + exchange_rate_crypto_obj[el].Country + " " + exchange_rate_crypto_obj[el].Code + " " + currencyNamesSearchText
      rate_country.scur = ""
      rate_country.status = "ALREADY"
      rate_country.tcur = exchange_rate_crypto_obj[el].Code
      rate_country.symbol = exchange_rate_crypto_obj[el].Symbol
      rate_country.selected = currency_list_now.includes(el)
      c.push(rate_country)
    })
    country_list_status.value = c
    country_selected.value = currency_list_now
  })
})

function escapeRegExp(string) {
  return string.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')
}

function matchIgnoreCase(str, a) {
  const escapedA = escapeRegExp(a)
  const pattern = new RegExp(escapedA, 'i')
  return pattern.test(str)
}

const changStatus = (item) => {
  item.selected = !item.selected
  if (item.selected) {
    country_selected.value.push(item.tcur)
  } else {
    country_selected.value.splice(country_selected.value.indexOf(item.tcur), 1)
  }

  let c = country_selected.value.map((el) => {
    return el
  })
  chrome.storage.local.set({ currency_list_now: c }, function () {})
}

const lookSelected = () => {
  searchKey.value = ''
  if (selected_status.value) {
    country_list_status.value.forEach((el) => {
      el.status = 'ALREADY'
    })
    selected_text.value = t('ViewSelected')
    selected_status.value = false
  } else {
    country_list_status.value.forEach((el) => {
      if (el.selected) {
        el.status = 'ALREADY'
      } else {
        el.status = 'noShow'
      }
      selected_text.value = t('ViewAll')
      selected_status.value = true
    })
  }
}
</script>

<style>
.add_list {
  margin-bottom: 30px;
  padding: 0 8px;
}
.add_list .el-input-group__append {
  background: rgba(10, 31, 61, 0);
  box-shadow: none;
}
.add_list .add_list_m {
  color: var(--add-list-title-text-color);
}
.add_list .el-input__wrapper {
  background: rgba(53, 75, 105, 0);
  box-shadow: none;
  border: 1px solid #adadad;
  border-radius: 5px;
}
.add_list .el-input__inner {
  color: var(--add-list-search-text-color);
  font-family: Verdana, Geneva, Tahoma, sans-serif;
  font-size: 16px;
}
.add_list_search span {
  height: 20px;
  width: 20px;
  font-size: 16px;
  color: var(--add-list-search-icon-color);
}
.add_list_all {
  margin-top: 10px;
  height: 400px;
}
.add_list_all_s {
  color: var(--add-list-item-text-color);
  font-size: 14px;
  justify-self: start;
  padding-left: 20px;
}
.add_list_symbol {
  color: var(--add-list-item-symbol-color);
  font-size: 20px;
  font-weight: bold;
}
.add_list_selected {
  color: var(--add-list-item-selected-d-color);
}
.add_list_country_item {
  background: var(--add-list-item-bg-color);
  width: 100%;
  height: 58px;
  display: grid;
  grid-template-columns: 65px 200px auto;
  place-items: center center;
  border-bottom: 1px solid var(--add-list-item-border-color);
}
.add_list_country_item:hover {
  background: var(--add-list-item-hover-bg-color);
}
.selected_country_not {
  color: var(--add-list-item-selected-d-color);
  font-size: 14px;
  cursor: pointer;
  font-style: oblique;
  font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
}
.selected_country {
  color: var(--add-list-item-selected-not-color);
  font-size: 14px;
  cursor: pointer;
  font-style: initial;
  font-family: 'Times New Roman', Times, serif;
}
</style>