<script setup lang="ts">
/**
 * RatingBadge — displays a movie rating as a colored badge.
 * - Green for 7+, Yellow for 5-7, Red below 5.
 */
const props = defineProps<{
  rating: number;
}>();

const formattedRating = computed(() => props.rating.toFixed(1));

const badgeColor = computed(() => {
  if (props.rating >= 7) return "bg-green-500";
  if (props.rating >= 5) return "bg-yellow-500";
  return "bg-red-500";
});
</script>

<template>
  <span
    :class="badgeColor"
    class="inline-flex items-center gap-1 rounded-md px-2 py-0.5 text-xs font-bold text-white shadow-sm"
  >
    <svg
      class="h-3 w-3"
      fill="currentColor"
      viewBox="0 0 20 20"
      aria-hidden="true"
    >
      <path
        d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"
      />
    </svg>
    {{ formattedRating }}
  </span>
</template>
