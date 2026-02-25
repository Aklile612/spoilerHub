<script setup lang="ts">
/**
 * SpoilerToggle — expandable spoiler section with warning + reveal button.
 * Content is hidden by default and smoothly expands.
 */
defineProps<{
  spoiler: string;
}>();

const isRevealed = ref(false);

function toggle() {
  isRevealed.value = !isRevealed.value;
}

/**
 * Very lightweight markdown → HTML renderer.
 * Handles headings, bold, italic, lists, and paragraphs.
 */
function renderMarkdown(text: string): string {
  return text
    // Headings
    .replace(/^### (.+)$/gm, '<h3>$1</h3>')
    .replace(/^## (.+)$/gm, '<h2>$1</h2>')
    .replace(/^# (.+)$/gm, '<h2>$1</h2>')
    // Bold
    .replace(/\*\*(.+?)\*\*/g, '<strong>$1</strong>')
    // Italic
    .replace(/\*(.+?)\*/g, '<em>$1</em>')
    // Unordered list items
    .replace(/^- (.+)$/gm, '<li>$1</li>')
    // Wrap consecutive <li> in <ul>
    .replace(/((?:<li>.*<\/li>\n?)+)/g, '<ul>$1</ul>')
    // Paragraphs (double newline)
    .replace(/\n\n/g, '</p><p>')
    // Single newlines → <br> inside paragraphs
    .replace(/\n/g, '<br>')
    // Wrap in paragraph
    .replace(/^(.+)/, '<p>$1</p>');
}
</script>

<template>
  <div class="mt-8">
    <!-- Warning box -->
    <div
      class="rounded-lg border border-red-200 bg-red-50 px-5 py-4"
    >
      <div class="flex items-start gap-3">
        <span class="text-2xl" aria-hidden="true">⚠️</span>
        <div>
          <p class="font-semibold text-red-800">Spoiler Warning</p>
          <p class="mt-1 text-sm text-red-600">
            This section contains major plot spoilers, twists, and ending details.
          </p>
        </div>
      </div>

      <!-- Reveal / Hide button -->
      <button
        @click="toggle"
        :aria-expanded="isRevealed"
        aria-controls="spoiler-content"
        class="mt-4 inline-flex items-center gap-2 rounded-lg px-5 py-2.5 text-sm font-semibold transition-colors"
        :class="
          isRevealed
            ? 'bg-gray-200 text-gray-700 hover:bg-gray-300'
            : 'bg-red-600 text-white hover:bg-red-700'
        "
      >
        <svg
          class="h-4 w-4 transition-transform duration-300"
          :class="{ 'rotate-180': isRevealed }"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M19 9l-7 7-7-7"
          />
        </svg>
        {{ isRevealed ? "Hide Spoiler" : "Reveal Spoiler" }}
      </button>
    </div>

    <!-- Expandable spoiler content -->
    <Transition name="spoiler">
      <div
        v-if="isRevealed"
        id="spoiler-content"
        class="mt-4 rounded-lg border border-gray-200 bg-white px-6 py-5"
      >
        <div
          class="spoiler-content prose prose-sm max-w-none"
          v-html="renderMarkdown(spoiler)"
        />
      </div>
    </Transition>
  </div>
</template>
