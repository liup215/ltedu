<template>
  <div class="flex items-center text-sm text-gray-900">
    <span class="font-mono whitespace-pre text-gray-500">{{ indent }}{{ prefix }}</span>
    <label class="flex items-center gap-2 hover:bg-gray-50 px-2 py-1 rounded cursor-pointer group">
      <input 
        type="checkbox" 
        :checked="isSelected"
        :indeterminate="isIndeterminate"
        @change="toggleChapter"
        class="rounded text-indigo-600 focus:ring-indigo-500 group-hover:border-indigo-500 transition-colors"
      >
      <span :class="[
        'truncate',
        isIndeterminate ? 'text-gray-600' : '',
        'group-hover:text-gray-900 transition-colors'
      ]">
        {{ chapter.name }}
        <span v-if="isIndeterminate" class="text-xs text-gray-500">
          ({{ $t('chapterOption.selectedCount', { count: getSelectedCount() }) }})
        </span>
      </span>
    </label>
  </div>
  <!-- Recursive rendering for child chapters -->
  <template v-if="chapter.children && chapter.children.length > 0">
    <div 
      v-for="(child, index) in chapter.children" 
      :key="child.id" 
      class="ml-6"
    >
      <ChapterOption 
        :chapter="child"
        :level="props.level + 1"
        :is-last="index === chapter.children.length - 1"
        :parent-tree="getChildParentTree(index, chapter.children.length)"
        :selected-chapters="selectedChapters"
        @update:selected="updateSelected"
      />
    </div>
  </template>
</template>

<script setup lang="ts">
/**
 * ChapterOption Component
 * 
 * A recursive component that renders a chapter tree structure with checkboxes
 * for multiple selection. Supports unlimited nesting levels with proper tree
 * visualization using ASCII characters.
 * 
 * Features:
 * - Tree visualization with proper branch lines and indentation
 * - Multiple chapter selection with checkboxes
 * - Two-way binding for selected chapters
 * - Recursive rendering of nested chapter structures
 * - Proper handling of last items in each branch
 * 
 * Usage:
 * <ChapterOption
 *   :chapter="chapterData"
 *   :level="0"
 *   :is-last="false"
 *   :selected-chapters="selectedIds"
 *   @update:selected="updateSelection"
 * />
 * 
 * The component emits 'update:selected' events when chapters are selected/deselected,
 * providing an array of selected chapter IDs.
 * 
 * Tree visualization example:
 * Chapter 1
 * ├── Section 1.1
 * │   ├── Topic 1.1.1
 * │   └── Topic 1.1.2
 * └── Section 1.2
 *     └── Topic 1.2.1
 */

import { computed } from 'vue';
import type { Chapter } from '../../models/chapter.model';

/**
 * Component props
 * @prop chapter - Chapter data object containing name, id, and optional children
 * @prop level - Nesting level, 0 for root nodes (optional, default: 0)
 * @prop isLast - Whether this is the last item in its sibling group (optional, default: false)
 * @prop parentTree - Array tracking which parent levels need vertical lines (optional, default: [])
 * @prop selectedChapters - Array of currently selected chapter IDs (optional, default: [])
 */
const props = withDefaults(defineProps<{
  chapter: Chapter;
  level?: number;
  isLast?: boolean;
  parentTree?: boolean[];
  selectedChapters?: number[];
}>(), {
  level: 0,
  isLast: false,
  parentTree: () => [],
  selectedChapters: () => []
});

/**
 * Emit chapter selection changes to parent
 */
const emit = defineEmits<{
  (e: 'update:selected', chapters: number[]): void
}>();

/**
 * Computed value to check if current chapter is selected
 * A chapter is considered selected if:
 * - It is directly selected OR
 * - All its children are selected (if it has children)
 */
const isSelected = computed(() => {
  if (!props.chapter.children?.length) {
    return props.selectedChapters?.includes(props.chapter.id) ?? false;
  }

  const descendantIds = getAllDescendantIds(props.chapter);
  return descendantIds.every(id => props.selectedChapters?.includes(id));
});

/**
 * Computed value to track indeterminate state
 * True when some but not all descendants are selected
 */
const isIndeterminate = computed(() => {
  if (!props.chapter.children?.length) return false;

  const descendantIds = getAllDescendantIds(props.chapter);
  const selectedCount = descendantIds.filter(id => props.selectedChapters?.includes(id)).length;
  
  return selectedCount > 0 && selectedCount < descendantIds.length;
});

/**
 * Get all descendant chapter IDs including the current chapter
 */
const getAllDescendantIds = (chapter: Chapter): number[] => {
  const ids = [chapter.id];
  if (chapter.children) {
    chapter.children.forEach(child => {
      ids.push(...getAllDescendantIds(child));
    });
  }
  return ids;
};

/**
 * Toggle chapter selection and emit changes
 * When a chapter is selected/deselected, all its descendants are also selected/deselected
 */
const toggleChapter = () => {
  const newSelection = [...(props.selectedChapters || [])];
  const currentSelected = isSelected.value;
  const allIds = getAllDescendantIds(props.chapter);
  
  if (!currentSelected) {
    // Select this chapter and all its descendants
    allIds.forEach(id => {
      if (!newSelection.includes(id)) {
        newSelection.push(id);
      }
    });
  } else {
    // Deselect this chapter and all its descendants
    allIds.forEach(id => {
      const index = newSelection.indexOf(id);
      if (index !== -1) {
        newSelection.splice(index, 1);
      }
    });
  }
  
  emit('update:selected', newSelection);
};

/**
 * Handle selection updates from child components
 */
const updateSelected = (chapters: number[]) => {
  emit('update:selected', chapters);
};

/**
 * Gets the count of selected descendants for the current chapter
 * Used to display selection count in indeterminate state
 */
const getSelectedCount = () => {
  const descendantIds = getAllDescendantIds(props.chapter);
  return descendantIds.filter(id => props.selectedChapters?.includes(id)).length;
};

/**
 * Computed tree context for the current node
 * Manages the state of vertical lines for the current level and its children
 */
const treeContext = computed(() => {
  // Use inherited parent tree or create empty one
  const currentTree = [...props.parentTree];
  
  // Add current level's line state based on isLast
  if (props.level > 0) {
    currentTree.push(!props.isLast);
  }
  
  return {
    isLast: props.isLast,
    level: props.level,
    parentTree: currentTree
  };
});

// Compute the indentation based on parent context
const indent = computed(() => {
  if (props.level <= 0) return '';

  // Create indentation based on parent levels only
  const parentLevels = props.parentTree;
  const indentParts = parentLevels.map(hasLine => hasLine ? '│' : '');
  return indentParts.join('');
});

// Compute the prefix for current node
const prefix = computed(() => {
  if (props.level <= 0) return '';
  return props.isLast ? '└──' : '├──';
});

// Method to update parent tree for child nodes
const getChildParentTree = (index: number, total: number) => {
  // Get the parent tree state up to the current level
  const childTree = [...treeContext.value.parentTree];
  
  // Add information about siblings for this new level
  // When this child isn't the last one, add a vertical line for its children
  const hasNextSibling = index < total - 1;
  childTree.push(hasNextSibling);
  
  return childTree;
};
</script>
