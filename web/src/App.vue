<template>
  <Notivue v-slot="item">
    <Notification :item="item" />
  </Notivue>
  <DonationPopup ref="donationPopupRef" />
  <router-view />
  <FeedbackWidget v-if="userStore.user" />
</template>

<script setup lang="ts">
import { Notivue, Notification } from 'notivue'
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import DonationPopup from './components/DonationPopup.vue'
import FeedbackWidget from './components/FeedbackWidget.vue'
import { useUserStore } from './stores/userStore'

const donationPopupRef = ref()
const router = useRouter()
const userStore = useUserStore()

onMounted(() => {
  router.afterEach(() => {
    const user = userStore?.user
    const now = Date.now()
    const isVip = user?.vipExpireAt && new Date(user.vipExpireAt).getTime() > now
    const isAdmin = user?.isAdmin === true
    if (Math.random() < 0.05 && donationPopupRef.value && !isVip && !isAdmin && user) {
      donationPopupRef.value.show()
    }
  })
})
</script>

<style>
/* No additional styles needed; Tailwind handles all styling */
</style>
