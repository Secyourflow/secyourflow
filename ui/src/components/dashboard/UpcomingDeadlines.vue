<script setup lang="ts">
import PanelListElement from '@/components/layout/PanelListElement.vue'
import { buttonVariants } from '@/components/ui/button'
import { Card } from '@/components/ui/card'
import { Separator } from '@/components/ui/separator'
import { Badge } from '@/components/ui/badge'

import { ChevronRight } from 'lucide-vue-next'

import { useQuery } from '@tanstack/vue-query'
import { formatDistanceToNow, isPast, parseISO } from 'date-fns'

import { pb } from '@/lib/pocketbase'
import type { Ticket } from '@/lib/types'
import { cn } from '@/lib/utils'

const {
  isPending,
  isError,
  data: tickets,
  error
} = useQuery({
  queryKey: ['tickets', 'deadlines'],
  queryFn: (): Promise<Array<Ticket>> =>
    pb.collection('tickets').getFullList({
      sort: 'deadline',
      filter: 'open = true && deadline != ""',
      expand: 'owner,type'
    })
})

const formatDeadline = (deadline: string) => {
  const date = parseISO(deadline)
  const distance = formatDistanceToNow(date, { addSuffix: true })
  return distance
}

const isOverdue = (deadline: string) => {
  return isPast(parseISO(deadline))
}

const getSeverityColor = (severity?: string) => {
  switch (severity) {
    case 'critical':
      return 'destructive'
    case 'high':
      return 'default'
    case 'medium':
      return 'secondary'
    case 'low':
      return 'outline'
    default:
      return 'outline'
  }
}
</script>

<template>
  <div class="flex flex-col gap-2">
    <Card>
      <div v-if="tickets && tickets.length === 0" class="p-2 text-center text-sm text-gray-500">
        No tickets with deadlines
      </div>
      <PanelListElement v-else v-for="ticket in tickets" :key="ticket.id" class="gap-2 pr-1">
        <span>{{ ticket.name }}</span>
        <Separator orientation="vertical" class="hidden h-4 sm:block" />
        <span class="text-sm text-muted-foreground">{{ ticket.expand.type.singular }}</span>
        <Separator orientation="vertical" class="hidden h-4 sm:block" />
        <Badge v-if="ticket.severity" :variant="getSeverityColor(ticket.severity)" class="text-xs">
          {{ ticket.severity }}
        </Badge>
        <Separator orientation="vertical" class="hidden h-4 sm:block" />
        <span
          :class="
            cn('text-sm', isOverdue(ticket.deadline!) ? 'font-semibold text-red-600' : 'text-muted-foreground')
          "
        >
          {{ isOverdue(ticket.deadline!) ? 'Overdue' : 'Due' }}
          {{ formatDeadline(ticket.deadline!) }}
        </span>
        <RouterLink
          :to="{
            name: 'tickets',
            params: { type: ticket.type, id: ticket.id }
          }"
          :class="
            cn(
              buttonVariants({ variant: 'outline', size: 'sm' }),
              'h-8 w-full sm:ml-auto sm:w-auto'
            )
          "
        >
          <span class="flex flex-row items-center text-sm text-gray-500">
            View
            <ChevronRight class="ml-2 h-4 w-4" />
          </span>
        </RouterLink>
      </PanelListElement>
    </Card>
  </div>
</template>
