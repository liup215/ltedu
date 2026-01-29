<template>
  <div v-if="visible" class="fixed inset-0 z-[999] flex items-center justify-center bg-black bg-opacity-30">
    <div class="bg-white rounded-lg shadow-lg p-6 max-w-sm w-full relative">
      <button @click="close" class="absolute top-2 right-2 text-gray-500 hover:text-gray-700 text-xl font-bold">&times;</button>
      <h2 class="text-xl font-bold text-center mb-2">{{ $t('donation.title') }}</h2>
      <div class="bg-yellow-100 border-l-4 border-yellow-500 text-yellow-800 font-bold px-4 py-2 mb-3 text-center rounded">
        <span>{{ $t('donation.importantNotice') }}</span>
      </div>
      <p class="text-gray-700 text-center mb-4 text-sm">
        {{ $t('donation.popupDescription1') }}<br>
        {{ $t('donation.popupDescription2') }}
      </p>
      <div class="flex flex-row justify-center items-start gap-6 mb-2">
        <div class="flex flex-col items-center">
          <span class="text-sm font-semibold mb-1">{{ $t('donation.wechatPay') }}</span>
          <img src="/src/assets/img/weixin_pay_qrcode.png" alt="WeChat Pay QR Code" class="w-24 h-auto rounded border" />
        </div>
        <div class="flex flex-col items-center">
          <span class="text-sm font-semibold mb-1">{{ $t('donation.qqGroup') }}</span>
          <img src="/src/assets/img/qq_group_qrcode.jpg" alt="QQ Group QR Code" class="w-24 h-auto rounded border" />
        </div>
      </div>
      <div class="flex justify-center">
        <a href="/donate" class="text-blue-600 underline text-sm">{{ $t('donation.moreOptions') }}</a>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, defineExpose } from 'vue'
const visible = ref(false)
function show() { visible.value = true }
function close() { visible.value = false }
defineExpose({ show, close })
</script>
