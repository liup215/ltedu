<template>
  <div class="p-6">
    <header class="mb-6">
      <h1 class="text-3xl font-bold text-gray-900">{{ $t('systemSettings.title') }}</h1>
      <p class="mt-1 text-sm text-gray-500">{{ $t('systemSettings.subtitle') }}</p>
    </header>

    <div class="bg-white shadow sm:rounded-lg p-6">
      <h2 class="text-xl font-semibold text-gray-900 mb-4">Image Storage Settings</h2>
      
      <div v-if="!loading">
        <form @submit.prevent="saveConfig">
          <!-- Disk Type Selection -->
          <div class="mb-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">Storage Provider</label>
          <div class="flex space-x-4">
            <label class="inline-flex items-center">
              <input type="radio" v-model="config.disk" :value="DISK_PUBLIC" class="form-radio text-indigo-600">
              <span class="ml-2">Local Storage</span>
            </label>
            <label class="inline-flex items-center">
              <input type="radio" v-model="config.disk" :value="DISK_OSS" class="form-radio text-indigo-600">
              <span class="ml-2">Aliyun OSS</span>
            </label>
            <label class="inline-flex items-center">
              <input type="radio" v-model="config.disk" :value="DISK_COS" class="form-radio text-indigo-600">
              <span class="ml-2">Tencent COS</span>
            </label>
            <label class="inline-flex items-center">
              <input type="radio" v-model="config.disk" :value="DISK_QINIU" class="form-radio text-indigo-600">
              <span class="ml-2">Qiniu Kodo</span>
            </label>
          </div>
        </div>

        <!-- Aliyun OSS Settings -->
        <div v-if="config.disk === DISK_OSS" class="space-y-4 mb-6 border-t pt-4">
          <h3 class="text-lg font-medium text-gray-900">Aliyun OSS Config</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700">Access Key ID</label>
              <input type="text" v-model="config.ossAccessKeyId" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">Access Key Secret</label>
              <input type="password" v-model="config.ossAccessKeySecret" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">Bucket Name</label>
              <input type="text" v-model="config.ossBucket" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">Endpoint</label>
              <input type="text" v-model="config.ossEndpoint" placeholder="oss-cn-hangzhou.aliyuncs.com" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
            </div>
            <div class="md:col-span-2">
              <label class="block text-sm font-medium text-gray-700">CDN URL (Optional)</label>
              <input type="text" v-model="config.ossCDNUrl" placeholder="https://cdn.example.com" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
            </div>
          </div>
        </div>

        <!-- Tencent COS Settings -->
        <div v-if="config.disk === DISK_COS" class="space-y-4 mb-6 border-t pt-4">
          <h3 class="text-lg font-medium text-gray-900">Tencent COS Config</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700">Secret ID</label>
              <input type="text" v-model="config.cosSecretId" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">Secret Key</label>
              <input type="password" v-model="config.cosSecretKey" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">Bucket Name</label>
              <input type="text" v-model="config.cosBucket" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">Region</label>
              <input type="text" v-model="config.cosRegion" placeholder="ap-guangzhou" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">App ID</label>
              <input type="text" v-model="config.cosAppId" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
            </div>
            <div class="md:col-span-2">
              <label class="block text-sm font-medium text-gray-700">CDN URL (Optional)</label>
              <input type="text" v-model="config.cosCDNUrl" placeholder="https://cdn.example.com" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
            </div>
          </div>
        </div>

        <!-- Qiniu Kodo Settings -->
        <div v-if="config.disk === DISK_QINIU" class="space-y-4 mb-6 border-t pt-4">
          <h3 class="text-lg font-medium text-gray-900">Qiniu Kodo Config</h3>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700">Access Key</label>
              <input type="text" v-model="config.qiniuAccessKey" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">Secret Key</label>
              <input type="password" v-model="config.qiniuSecretKey" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700">Bucket Name</label>
              <input type="text" v-model="config.qiniuBucket" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
            </div>
            <div class="md:col-span-2">
              <label class="block text-sm font-medium text-gray-700">CDN URL</label>
              <input type="text" v-model="config.qiniuCDNUrl" placeholder="https://cdn.example.com" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm">
            </div>
          </div>
        </div>

        <div class="flex justify-end pt-4">
          <button 
            type="submit" 
            class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            :disabled="saving"
          >
            {{ saving ? 'Saving...' : 'Save Configuration' }}
          </button>
        </div>
        </form>

        <!-- Migration Tool -->
        <div class="mt-8 border-t pt-6">
          <h3 class="text-lg font-medium text-gray-900 mb-4">Maintenance Tools</h3>
          <div class="bg-gray-50 p-4 rounded-md">
            <h4 class="text-md font-medium text-gray-900">Migrate Question Images</h4>
            <p class="text-sm text-gray-500 mt-1 mb-4">
              Scan all questions for Base64 encoded images and migrate them to the currently configured storage provider. 
              This process may take a while depending on the number of questions.
            </p>
            <button 
              type="button"
              @click="handleMigrate"
              class="inline-flex justify-center py-2 px-4 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              :disabled="migrating"
            >
              {{ migrating ? 'Migrating...' : 'Start Migration' }}
            </button>
          </div>
        </div>
      </div>

      <div v-else class="text-center py-4">
        <p class="text-gray-500">Loading configuration...</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import configService from '../../services/configService';
import { 
  type ImageUploadConfig, 
  DISK_PUBLIC, 
  DISK_OSS, 
  DISK_COS, 
  DISK_QINIU 
} from '../../models/config.model';
import { showSuccess } from '../../utils/notification';

const loading = ref(true);
const saving = ref(false);
const migrating = ref(false);
const config = ref<ImageUploadConfig>({
  disk: DISK_PUBLIC,
});

const loadConfig = async () => {
  try {
    const response = await configService.getImageUploadConfig();
    if (response.data) {
      config.value = response.data;
    }
  } catch (error) {
    console.error('Failed to load config:', error);
  } finally {
    loading.value = false;
  }
};

const saveConfig = async () => {
  saving.value = true;
  try {
    await configService.saveImageUploadConfig(config.value);
    showSuccess('Settings saved successfully');
  } catch (error) {
    console.error('Failed to save config:', error);
  } finally {
    saving.value = false;
  }
};

const handleMigrate = async () => {
  if (!confirm('Are you sure you want to start the migration? This action will modify question content fields.')) {
    return;
  }
  
  migrating.value = true;
  try {
    await configService.migrateImages();
    showSuccess('Image migration completed successfully');
  } catch (error) {
    console.error('Failed to migrate images:', error);
    // Error is handled by global error handler
  } finally {
    migrating.value = false;
  }
};

onMounted(() => {
  loadConfig();
});
</script>