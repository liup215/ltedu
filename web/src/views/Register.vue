<template>
  <div class="flex items-center justify-center min-h-screen bg-gray-100">
    <div class="w-full max-w-md p-8 space-y-6 bg-white rounded-lg shadow-md">
      <h2 class="text-2xl font-bold text-center text-gray-700">{{ $t('register.title') }}</h2>
      <form @submit.prevent="handleRegister" class="space-y-6">
        <div>
          <label for="username" class="block text-sm font-medium text-gray-700">{{ $t('register.username') }}</label>
          <input id="username" v-model="username" type="text" required
            class="w-full px-3 py-2 mt-1 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            :placeholder="$t('register.usernamePlaceholder')" />
        </div>
        <div>
          <label for="email" class="block text-sm font-medium text-gray-700">{{ $t('register.email') }}</label>
          <input id="email" v-model="email" type="email" required
            class="w-full px-3 py-2 mt-1 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            :placeholder="$t('register.emailPlaceholder')" />
        </div>
        <div class="flex items-center space-x-2">
          <label for="verificationCode" class="block text-sm font-medium text-gray-700">{{ $t('register.verificationCode') }}</label>
          <input id="verificationCode" v-model="verificationCode" type="text" required
            class="w-full px-3 py-2 mt-1 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            :placeholder="$t('register.verificationCodePlaceholder')" />
<button
  type="button"
  @click="handleSendCode"
  :disabled="isSendingCode || countdown > 0 || !email || !captchaValue"
  class="px-3 py-2 bg-indigo-500 text-white rounded hover:bg-indigo-600 disabled:opacity-50"
>
  <span v-if="isSendingCode">{{ $t('register.sending') }}</span>
  <span v-else-if="countdown > 0">{{ countdown }}{{ $t('register.secondsLater') }}</span>
  <span v-else>{{ $t('register.sendCode') }}</span>
</button>
        </div>
        <div class="flex items-center space-x-2 mt-2">
          <img :src="captchaImg" :alt="$t('register.captcha')" class="h-10 w-32 border rounded" @click="refreshCaptcha" style="cursor:pointer;" />
          <input id="captchaValue" v-model="captchaValue" type="text"
          :disabled="verificationCode !== ''"
            class="w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            :placeholder="$t('register.captchaPlaceholder')" />
        </div>
        <div>
          <label for="password" class="block text-sm font-medium text-gray-700">{{ $t('register.password') }}</label>
          <input id="password" v-model="password" type="password" required
            class="w-full px-3 py-2 mt-1 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            :placeholder="$t('register.passwordPlaceholder')" />
        </div>
        <div>
          <label for="confirmPassword" class="block text-sm font-medium text-gray-700">{{ $t('register.confirmPassword') }}</label>
          <input id="confirmPassword" v-model="confirmPassword" type="password" required
            class="w-full px-3 py-2 mt-1 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
            :placeholder="$t('register.confirmPasswordPlaceholder')" />
        </div>
        <div v-if="error" class="text-sm text-red-600">
          {{ error }}
        </div>
        <div v-if="successMessage" class="text-sm text-green-600">
          {{ successMessage }}
        </div>
        <div>
          <button type="submit" :disabled="isLoading"
            class="w-full px-4 py-2 font-medium text-white bg-indigo-600 rounded-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50">
            <span v-if="isLoading">{{ $t('register.registering') }}</span>
            <span v-else>{{ $t('register.register') }}</span>
          </button>
        </div>
      </form>
      <p class="text-sm text-center text-gray-600">
        {{ $t('register.alreadyHaveAccount') }}
        <router-link to="/login" class="font-medium text-indigo-600 hover:text-indigo-500">
          {{ $t('register.loginHere') }}
        </router-link>
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { authService } from '../services/authService'; // Import authService
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
// import { useUserStore } from '../stores/userStore'; // May not be needed directly for registration

const username = ref('');
const email = ref('');
const verificationCode = ref('');
const password = ref('');
const confirmPassword = ref('');
const captchaValue = ref('');
const captchaId = ref('');
const captchaImg = ref('');
const error = ref<string | null>(null);
const successMessage = ref<string | null>(null);
const isLoading = ref(false);
const isSendingCode = ref(false);
const countdown = ref(0);
let timer: number | undefined;
const router = useRouter();

const handleRegister = async () => {
  error.value = null;
  successMessage.value = null;

  if (password.value !== confirmPassword.value) {
    error.value = t('register.passwordsDoNotMatch');
    return;
  }

  if (password.value.length < 6) {
    error.value = t('register.passwordTooShort');
    return;
  }

  if (!verificationCode.value) {
    error.value = t('register.enterVerificationCode');
    return;
  }

  isLoading.value = true;
  try {
    const registrationData = {
      username: username.value,
      password: password.value,
      passwordConfirm: confirmPassword.value,
      email: email.value,
      verificationCode: verificationCode.value,
    };
    
    await authService.register(registrationData);
    successMessage.value = t('register.success');
    setTimeout(() => {
      router.push('/login');
    }, 2000);
  } catch (err: any) {
    refreshCaptcha()
    error.value = err.message || t('register.registrationFailed');
  } finally {
    isLoading.value = false;
  }
};

const handleSendCode = async () => {
  error.value = null;
  successMessage.value = null;
  isSendingCode.value = true;
  try {
    await authService.sendVerificationCode({
      email: email.value,
      captchaId: captchaId.value,
      captchaValue: captchaValue.value,
    });
    successMessage.value = t('register.verificationCodeSent');
    refreshCaptcha();
    captchaValue.value = '';
    // Start countdown
    countdown.value = 60;
    timer = window.setInterval(() => {
      if (countdown.value > 0) {
        countdown.value--;
      } else {
        clearInterval(timer);
        timer = undefined;
      }
    }, 1000);
  } catch (err: any) {
    error.value = err.message || t('register.failedToSendCode');
    refreshCaptcha();
    captchaValue.value = '';
  } finally {
    isSendingCode.value = false;
  }
};

import captchaService from '../services/captchaService';

const fetchCaptcha = async () => {
  try {
    const response = await captchaService.getCaptchaImage();
    if (response.data) {
      captchaId.value = response.data.key;
      captchaImg.value = response.data.img;
    }
  } catch {
    captchaImg.value = '';
    captchaId.value = '';
  }
};

const refreshCaptcha = () => {
  fetchCaptcha();
};

onMounted(() => {
  fetchCaptcha();
});

import { onUnmounted } from 'vue';
onUnmounted(() => {
  if (timer) {
    clearInterval(timer);
    timer = undefined;
  }
});
</script>

<style scoped>
/* Add any specific styles for the registration page here */
</style>
