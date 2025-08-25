<template>
  <div class="max-w-7xl mx-auto p-6">
    <p v-if="loading">Loading. Please wait...</p>
    <p v-else-if="error">{{ error }}</p>
    <div v-else>
      <h2 class="text-3xl font-bold text-gray-800 mb-6">Filters</h2>

      <!--Filters-->
      <div class="bg-gray-50 rounded-lg border border-gray-200 p-6 mb-6">
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-5 gap-4">
          <div class="space-y-2">
            <label for="ticker" class="block text-sm font-semibold text-gray-700">
              Ticker/Company
            </label>
            <input id="ticker" v-model="filters.ticker" type="text" placeholder="Company name/ticker..."
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors" />
          </div>

          <div class="space-y-2">
            <label for="brokerage" class="block text-sm font-semibold text-gray-700">
              Brokerage
            </label>
            <input id="brokerage" v-model="filters.brokerage" type="text" placeholder="Search by brokerage..."
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors" />
          </div>

          <div class="space-y-2">
            <label for="action" class="block text-sm font-semibold text-gray-700">
              Action
            </label>
            <select id="action" v-model="filters.action"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors bg-white">
              <option value="">All actions</option>
              <option v-for="action in UniqueActions" :key="action" :value="action">
                {{ action }}
              </option>
            </select>
          </div>

          <div class="space-y-2">
            <label for="ratingFrom" class="block text-sm font-semibold text-gray-700">
              Rating From
            </label>
            <select id="ratingFrom" v-model="filters.ratingFrom"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors bg-white">
              <option value="">All ratings</option>
              <option v-for="rating in UniqueRatings" :key="rating" :value="rating">
                {{ rating }}
              </option>
            </select>
            <label for="ratingTo" class="block text-sm font-semibold text-gray-700">
              to
            </label>
            <select id="ratingTo" v-model="filters.ratingTo"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors bg-white">
              <option value="">All ratings</option>
              <option v-for="rating in UniqueRatings" :key="rating" :value="rating">
                {{ rating }}
              </option>
            </select>
          </div>

          <div class="space-y-2">
            <label for="targetFrom" class="block text-sm font-semibold text-gray-700">
              Target From $
            </label>
            <input id="targetFrom" v-model.number="filters.targetFrom" type="number" min="0" step="0.01" placeholder="0"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors" />
            <label for="targetTo" class="block text-sm font-semibold text-gray-700">
              to $
            </label>
            <input id="targetTo" v-model.number="filters.targetTo" type="number" min="0" step="0.01"
              placeholder="999999"
              class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-colors" />
          </div>

          <div class="flex items-end">
            <button @click="cleanFilters"
              class="w-full px-4 py-2 bg-gray-600 hover:bg-gray-700 text-white font-medium rounded-md shadow-sm transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2">
              Reset filters
            </button>
          </div>
        </div>
      </div>

      <!--Sumarizer-->
      <div class="mb-4 text-sm text-gray-600 font-medium">
        Showing {{ filteredData.length }} out of {{ stocksData.length }} elements
      </div>

      <!--Table-->
      <div class="bg-white rounded-lg shadow-lg overflow-hidden border border-gray-200">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-800">
              <tr>
                <th class="px-6 py-4 text-left text-xs font-bold text-white uppercase tracking-wider">
                  Ticker
                </th>
                <th class="px-6 py-4 text-left text-xs font-bold text-white uppercase tracking-wider">
                  Company
                </th>
                <th class="px-6 py-4 text-left text-xs font-bold text-white uppercase tracking-wider">
                  Brokerage
                </th>
                <th class="px-6 py-4 text-left text-xs font-bold text-white uppercase tracking-wider">
                  Action
                </th>
                <th class="px-6 py-4 text-left text-xs font-bold text-white uppercase tracking-wider">
                  Rating From
                </th>
                <th class="px-6 py-4 text-left text-xs font-bold text-white uppercase tracking-wider">
                  Rating To
                </th>
                <th class="px-6 py-4 text-left text-xs font-bold text-white uppercase tracking-wider">
                  Target From
                </th>
                <th class="px-6 py-4 text-left text-xs font-bold text-white uppercase tracking-wider">
                  Target To
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="(item, index) in filteredData" :key="index"
                class="hover:bg-gray-50 transition-colors duration-150" :class="{ 'bg-gray-25': index % 2 === 1 }">
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                  {{ item.ticker || 'N/A' }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700">
                  <span
                    class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                    {{ item.company || 'N/A' }}
                  </span>
                </td>
                <td class="px-6 py-4 text-sm text-gray-700 max-w-xs truncate">
                  {{ item.brokerage || 'Sin descripción' }}
                </td>
                <td class="px-6 py-4 text-sm text-gray-700 max-w-xs truncate">
                  {{ item.action || 'Sin descripción' }}
                </td>
                <td class="px-6 py-4 text-sm text-gray-700 max-w-xs truncate">
                  {{ item.rating_from || 'Sin descripción' }}
                </td>
                <td class="px-6 py-4 text-sm text-gray-700 max-w-xs truncate">
                  {{ item.rating_to || 'Sin descripción' }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-bold text-green-600">
                  {{ item.target_from }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-bold text-green-600">
                  {{ item.target_to }}
                </td>
              </tr>
              <tr v-if="filteredData.length === 0">
                <td colspan="12" class="px-6 py-12 text-center text-gray-500 italic">
                  <div class="flex flex-col items-center space-y-2">
                    <span class="text-lg font-medium">:( </span>
                    <span class="text-lg font-medium">No related data </span>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Best stock section -->
      <div class="mt-6 grid grid-cols-1 md:grid-cols-1 gap-4">
        <div class="bg-grey-50 border border-blue-200 rounded-lg p-4">
          <div class="text-xl font-bold text-gray-800">Best Stock</div>
          <div class="max-w-md mx-auto">

            <!--Card-->
            <div
              class="bg-white shadow-lg rounded-lg border border-gray-200 p-5 cursor-pointer transition hover:shadow-xl"
              @click="expanded = !expanded">
              <!--Header-->
              <div class="flex justify-between items-center">
                <div>
                  <h2 class="text-xl font-bold text-gray-800">{{ bestStock.company }}</h2>
                  <p class="text-sm text-gray-500">{{ bestStock.ticker }}</p>
                </div>
                <div class="text-right">
                  <span class="text-2xl font-bold text-green-600">
                    +{{ formatPercentage(stockStats.percent_increase) }}%
                  </span>
                  <p class="text-sm text-gray-500">Target ↑ {{ stockStats.target_increase }}</p>
                </div>
              </div>
              <!--Expandable body-->
              <transition name="fade">
                <div v-if="expanded" class="mt-4 space-y-2 text-gray-700">
                  <p><span class="font-semibold">Action:</span> {{ bestStock.action }}</p>
                  <p><span class="font-semibold">Brokerage:</span> {{ bestStock.brokerage }}</p>
                  <p><span class="font-semibold">Rating:</span> {{ bestStock.rating_from }} → {{ bestStock.rating_to }}
                  </p>
                  <p><span class="font-semibold">Target:</span> {{ bestStock.target_from }} → {{ bestStock.target_to }}
                  </p>
                  <p><span class="font-semibold">Date:</span> {{ formatDate(bestStock.time) }}</p>
                  <p><span class="font-semibold">Criteria:</span> {{ stockStats.criteria }}</p>
                </div>
              </transition>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import axios from "axios";

interface Filters {
  ticker: string
  brokerage: string
  action: string
  ratingTo: string
  ratingFrom: string
  targetFrom: number | null
  targetTo: number | null
}

const stocksData = ref<Record<string, any>[]>([])
const bestStock = ref<Record<string, any>[]>([])
const stockStats = ref<Record<string, any>[]>([])
const expanded = ref<boolean>(false)
const loading = ref<boolean>(true)
const error = ref<string | null>(null);

const filters = ref<Filters>({
  ticker: '',
  brokerage: '',
  action: '',
  ratingTo: '',
  ratingFrom: '',
  targetFrom: null,
  targetTo: null
})

//Get actions
const UniqueActions = computed(() => {
  const actions = stocksData.value
    .map(item => item.action)
    .filter(action => action && typeof action === 'string')
  return [...new Set(actions)].sort()
})

//Get ratings
const UniqueRatings = computed(() => {
  const ratings = stocksData.value
    .map(item => item.rating_from)
    .filter(rating_from => rating_from && typeof rating_from === 'string')
  return [...new Set(ratings)].sort()
})

//Filters
const filteredData = computed(() => {
  return stocksData.value.filter(item => {

    //Ticker & Company
    const sameName = !filters.value.ticker || (
      (item.ticker && item.ticker.toString().toLowerCase().includes(filters.value.ticker.toLowerCase())) ||
      (item.company && item.company.toString().toLowerCase().includes(filters.value.ticker.toLowerCase()))
    );

    const sameBrokerage = !filters.value.brokerage || (
      (item.brokerage && item.brokerage.toString().toLowerCase().includes(filters.value.brokerage.toLowerCase()))
    );

    const sameAction = !filters.value.action || item.action === filters.value.action
    const sameRatingFrom = !filters.value.ratingFrom || item.rating_from === filters.value.ratingFrom
    const sameRatingTo = !filters.value.ratingTo || item.rating_to === filters.value.ratingTo

    const minPrice = filters.value.targetFrom === null ||
      (item.target_from !== undefined && Number(item.target_from.slice(1)) >= filters.value.targetFrom)

    const maxPrice = filters.value.targetTo === null ||
      (item.target_to !== undefined && Number(item.target_to.slice(1)) <= filters.value.targetTo)

    return sameName && sameBrokerage && sameAction && sameRatingFrom && sameRatingTo && minPrice && maxPrice
  })
})

//Methods & Format operations
const cleanFilters = (): void => {
  filters.value = {
    ticker: '',
    brokerage: '',
    action: '',
    ratingFrom: '',
    ratingTo: '',
    targetFrom: null,
    targetTo: null
  }
}

function formatDate(dateStr: string): string {
  return new Date(dateStr).toLocaleString("en-US", {
    dateStyle: "medium",
    timeStyle: "short"
  })
}
function formatPercentage(percentage: number): string {
  return Number(percentage).toFixed(2);
}

onMounted(async () => {
  try {
    const response = await axios.get("http://localhost:8080/api/external");
    stocksData.value = response.data.data.items;
    bestStock.value = response.data.best.best_item;
    stockStats.value = response.data.best;

  } catch (err: unknown) {
    if (err instanceof Error) {
      error.value = "Error loading data: " + err.message;
    } else {
      error.value = "Error loading data: Unknown error"
    }
  } finally {
    loading.value = false;
  }
});
</script>
