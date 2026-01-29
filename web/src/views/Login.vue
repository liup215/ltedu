<template>
  <div class="flex flex-col items-center justify-center flex-1 bg-white pt-16">
    <div class="bg-white rounded-xl shadow-lg p-8 max-w-md w-full">
      <h2 class="text-2xl font-bold text-indigo-700 mb-6 text-center">{{ $t('login.title') }}</h2>
      <form @submit.prevent="handleLogin" class="flex flex-col gap-4">
        <input v-model="credentials.username" type="text" :placeholder="$t('login.username')"
          class="border rounded px-4 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-400" required />
        <input v-model="credentials.password" type="password" :placeholder="$t('login.password')"
          class="border rounded px-4 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-400" required />
        <div class="flex items-center gap-4">
          <input v-model="credentials.captchaValue" type="text" :placeholder="$t('login.captcha')"
            class="border rounded px-4 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-400 w-full" required />
          <img :src="captchaImage" @click="fetchCaptcha" class="cursor-pointer h-10 rounded" :alt="$t('login.captcha')" />
        </div>
        <button type="submit"
          class="bg-indigo-600 text-white py-2 px-4 rounded hover:bg-indigo-700 transition disabled:opacity-50"
          :disabled="appStore.isLoading">
          {{ appStore.isLoading ? $t('login.login') + '...' : $t('login.login') }}
        </button>
      </form>
      <div v-if="errorMessage" class="text-red-500 text-sm mt-4 text-center">{{ errorMessage }}</div>
      <p class="text-sm text-center text-gray-600 mt-6">
        {{ $t('login.register') }}?
        <router-link to="/register" class="font-medium text-indigo-600 hover:text-indigo-500">
          {{ $t('login.register') }}
        </router-link>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { authService } from '../services/authService';
import captchaService from '../services/captchaService';
import { useUserStore } from '../stores/userStore';
import { useAppStore } from '../stores/appStore';
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const router = useRouter();
const userStore = useUserStore();
const appStore = useAppStore();

const credentials = reactive({
  username: '',
  password: '',
  captchaId: '',
  captchaValue: '',
});
const captchaImage = ref('');
const errorMessage = ref('');

async function fetchCaptcha() {
  try {
    const response = await captchaService.getCaptchaImage();
    if (response.data) {
      credentials.captchaId = response.data.key;
      captchaImage.value = response.data.img;
    }
  } catch (error) {
    errorMessage.value = t('login.loginFailed');
  }
}

onMounted(() => {
  fetchCaptcha();
});

async function handleLogin() {
  errorMessage.value = '';
  if (!credentials.username || !credentials.password || !credentials.captchaValue) {
    errorMessage.value = t('login.loginFailed');
    return;
  }

  try {
    const loginData = await authService.login(credentials);
    await userStore.login(loginData.data.token);
    router.push('/');
  } catch (error: any) {
    errorMessage.value = error.message || t('login.loginFailed');
    fetchCaptcha();
  }
}
</script>

<style scoped>
/* Tailwind handles all styling */
</style>
