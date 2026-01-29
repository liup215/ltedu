<template>
  <div class="chapter-tree">
    <div v-if="!chapters || chapters.length === 0" class="text-gray-400 text-sm px-3 py-2">
      {{ $t('chapterTree.noChapters') }}
    </div>
    <div v-for="chapter in chapters" :key="chapter.id" class="chapter-item">
      <div 
        :class="[
          'chapter-row flex items-center py-2 px-3 hover:bg-gray-50 cursor-pointer rounded',
          selectedId === chapter.id ? 'bg-indigo-50 text-indigo-700' : 'text-gray-700'
        ]"
        @click.stop="onSelect(chapter.id)"
      >
        <button 
          v-if="chapter.children?.length"
          class="mr-2 w-4 h-4 flex items-center justify-center text-gray-400 hover:text-gray-600"
          @click.stop="onToggleExpand(chapter.id)"
        >
          <svg 
            :class="[
              'w-4 h-4 transform transition-transform',
              expandedChapters.includes(chapter.id) ? 'rotate-90' : ''
            ]"
            xmlns="http://www.w3.org/2000/svg" 
            viewBox="0 0 20 20" 
            fill="currentColor"
          >
            <path 
              fill-rule="evenodd" 
              d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" 
              clip-rule="evenodd" 
            />
          </svg>
        </button>
        <span v-else class="w-6"></span>
        <span class="truncate">{{ chapter.name }}</span>
      </div>
      <div 
        v-if="chapter.children?.length && expandedChapters.includes(chapter.id)"
        class="ml-6 mt-1"
      >
        <ChapterTree
          :chapters="chapter.children"
          :selected-id="selectedId"
          :expanded-chapters="expandedChapters"
          @select="id => $emit('select', id)"
          @toggle-expand="id => $emit('toggle-expand', id)"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Chapter } from '../../models/chapter.model';

defineProps<{
  chapters: Chapter[];
  selectedId: number | null;
  expandedChapters: number[];
}>();

const emit = defineEmits<{
  (event: 'select', id: number): void;
  (event: 'toggle-expand', id: number): void;
}>();

const onSelect = (id: number) => {
  emit('select', id);
};

const onToggleExpand = (id: number) => {
  emit('toggle-expand', id);
};
</script>

<style scoped>
.chapter-tree {
  user-select: none;
}
</style>
