<template>
  <div class="min-h-screen bg-black text-white flex items-center justify-center p-8">
    <div class="container max-w-xl w-full">
      <h1 class="text-4xl font-bold mb-8 tracking-tight">
        Wiki<span class="text-[#00ADD8] italic">GO</span>lfer
      </h1>
      <form @submit.prevent="searchPath" class="flex flex-col gap-2 mb-8">
        <input
          v-model="start"
          type="text"
          placeholder="開始記事タイトルを入力"
          required
          class="px-5 py-3 text-black bg-white border-2 border-white rounded-lg outline-none focus:border-gray-500"
        />
        <input
          v-model="end"
          type="text"
          placeholder="終点記事タイトルを入力"
          required
          class="px-5 py-3 text-black bg-white border-2 border-white rounded-lg outline-none focus:border-gray-500"
        />
        <button
          type="submit"
          class="px-8 py-3 font-semibold bg-white text-black rounded-lg cursor-pointer"
        >
          検索
        </button>
      </form>
      <div
        class="p-6 bg-[#111] border border-[#333] rounded-lg min-h-[60px] text-base leading-relaxed"
      >
        <template v-if="loading">
          検索中...
        </template>
        <template v-else-if="error">
          <span class="text-red-500">{{ error }}</span>
        </template>
        <template v-else-if="path.length">
          <div v-for="(title, index) in path" :key="index" class="flex items-center gap-4 py-2">
            <span class="text-gray-500 font-semibold min-w-[2rem] text-right">{{ index + 1 }}</span>
            <a
              :href="`https://ja.wikipedia.org/wiki/${encodeURIComponent(title)}`"
              target="_blank"
              class="text-white border-b border-gray-500 hover:underline"
            >{{ title }}</a>
          </div>
        </template>
        <template v-else>
          結果がここに表示されます
        </template>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const start = ref('')
const end = ref('')
const path = ref([])
const loading = ref(false)
const error = ref('')

const searchPath = async () => {
  loading.value = true
  error.value = ''
  path.value = []
  try {
    const res = await fetch(
      `http://localhost:8080/api/path?start=${encodeURIComponent(start.value)}&end=${encodeURIComponent(end.value)}`
    )
    const data = await res.json()
    if (!res.ok) {
      error.value = data.error || data.message || 'エラーが発生しました'
    } else {
      path.value = data.path || []
    }
  } catch (e) {
    error.value = '通信エラー'
  } finally {
    loading.value = false
  }
}
</script>