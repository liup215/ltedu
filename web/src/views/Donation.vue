<template>
  <div class="max-w-xl mx-auto p-8 bg-white rounded-lg shadow mt-12">
    <h1 class="text-2xl font-bold text-center mb-4">{{ $t('donation.title', { appTitle }) }}</h1>
    <div class="bg-yellow-100 border-l-4 border-yellow-500 text-yellow-800 font-bold px-4 py-2 mb-3 text-center rounded">
      <span>{{ $t('donation.importantNotice') }}</span>
    </div>
    <p class="text-gray-700 text-center mb-6">
      {{ $t('donation.description1', { appTitle }) }}<br>
      {{ $t('donation.description2', { appTitle }) }}<br>
      {{ $t('donation.description3', { appTitle }) }}<br>
      {{ $t('donation.description4', { appTitle }) }}
    </p>
    <div class="flex flex-row justify-center items-start gap-8">
      <div class="flex flex-col items-center">
        <span class="text-lg font-semibold mb-2">{{ $t('donation.wechatPay') }}</span>
        <img src="/src/assets/img/weixin_pay_qrcode.png" alt="WeChat Pay QR Code" class="w-40 max-w-full h-auto rounded border" />
      </div>
      <div class="flex flex-col items-center">
        <span class="text-lg font-semibold mb-2">{{ $t('donation.joinQQGroup') }}</span>
        <img src="/src/assets/img/qq_group_qrcode.jpg" alt="QQ Group QR Code" class="w-40 max-w-full h-auto rounded border" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { APP_TITLE } from '../const/config'

const appTitle = ref('Nerdlet')
onMounted(async () => {
  appTitle.value = await APP_TITLE()
})
</script>
