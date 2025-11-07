<script setup lang="ts">
import TanView from '@/components/TanView.vue'
import { DonutChart } from '@/components/ui/chart-donut'

import { useQuery } from '@tanstack/vue-query'
import { computed } from 'vue'

import { pb } from '@/lib/pocketbase'
import type { Ticket } from '@/lib/types'

const {
  isPending,
  isError,
  data: tickets,
  error
} = useQuery({
  queryKey: ['tickets', 'severity'],
  queryFn: (): Promise<Array<Ticket>> =>
    pb.collection('tickets').getFullList({
      filter: 'open = true && severity != ""'
    })
})

const severityCounts = computed(() => {
  if (!tickets.value) return []

  const counts: Record<string, number> = {
    low: 0,
    medium: 0,
    high: 0,
    critical: 0
  }

  tickets.value.forEach((ticket) => {
    if (ticket.severity) {
      counts[ticket.severity]++
    }
  })

  return Object.entries(counts)
    .filter(([_, count]) => count > 0)
    .map(([severity, count]) => ({
      name: severity.charAt(0).toUpperCase() + severity.slice(1),
      count
    }))
})
</script>

<template>
  <TanView :isError="isError" :isPending="isPending" :error="error">
    <div v-if="severityCounts.length > 0" class="flex flex-1 items-center">
      <DonutChart index="name" type="donut" category="count" :data="severityCounts" />
    </div>
    <div v-else class="flex flex-1 items-center justify-center text-sm text-muted-foreground">
      No tickets with severity
    </div>
  </TanView>
</template>
