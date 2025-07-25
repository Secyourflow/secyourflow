<script setup lang="ts">
import { cn } from '@/lib/utils'

const props = defineProps<{
  class?: string
}>()
</script>

<template>
  <img
    src="@/assets/svf.png"
    alt="SecYourFlow Logo"
    :class="cn('dark:hidden', props.class)"
  />
  <img
    src="@/assets/svf.png"
    alt="SecYourFlow Logo"
    :class="cn('hidden dark:flex', props.class)"
  />
</template>
