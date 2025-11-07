<script setup lang="ts">
import DynamicInput from '@/components/input/DynamicInput.vue'
import { Badge } from '@/components/ui/badge'
import { Separator } from '@/components/ui/separator'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue
} from '@/components/ui/select'
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover'
import { Button } from '@/components/ui/button'
import { Calendar } from '@/components/ui/v-calendar'

import { CalendarIcon } from 'lucide-vue-next'

import { useMutation, useQueryClient } from '@tanstack/vue-query'
import format from 'date-fns/format'
import { ref } from 'vue'

import { pb } from '@/lib/pocketbase'
import type { Ticket } from '@/lib/types'
import { handleError, cn } from '@/lib/utils'

const queryClient = useQueryClient()

const props = defineProps<{
  ticket: Ticket
}>()

const name = ref(props.ticket.name)
const severity = ref(props.ticket.severity)
const deadline = ref(props.ticket.deadline ? new Date(props.ticket.deadline) : undefined)

const editNameMutation = useMutation({
  mutationFn: () =>
    pb.collection('tickets').update(props.ticket.id, {
      name: name.value
    }),
  onSuccess: () => queryClient.invalidateQueries({ queryKey: ['tickets', props.ticket.id] }),
  onError: handleError
})

const editSeverityMutation = useMutation({
  mutationFn: (newSeverity: string) =>
    pb.collection('tickets').update(props.ticket.id, {
      severity: newSeverity
    }),
  onSuccess: () => queryClient.invalidateQueries({ queryKey: ['tickets', props.ticket.id] }),
  onError: handleError
})

const editDeadlineMutation = useMutation({
  mutationFn: (newDeadline: Date | undefined) =>
    pb.collection('tickets').update(props.ticket.id, {
      deadline: newDeadline ? newDeadline.toISOString() : ''
    }),
  onSuccess: () => queryClient.invalidateQueries({ queryKey: ['tickets', props.ticket.id] }),
  onError: handleError
})

const updateName = (value: string) => {
  name.value = value
  editNameMutation.mutate()
}

const updateSeverity = (value: string) => {
  severity.value = value as 'low' | 'medium' | 'high' | 'critical'
  editSeverityMutation.mutate(value)
}

const updateDeadline = (value: Date | undefined) => {
  deadline.value = value
  editDeadlineMutation.mutate(value)
}

const getSeverityColor = (sev?: string) => {
  switch (sev) {
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
  <span class="text-4xl font-bold">
    <DynamicInput :modelValue="ticket.name" @update:modelValue="updateName" class="-mx-1" />
  </span>

  <div class="flex flex-col items-stretch gap-2 text-xs text-muted-foreground md:flex-row md:items-center">
    <div class="flex flex-col gap-1 md:flex-row md:items-center md:gap-2">
      <div>
        Created:
        {{ format(new Date(ticket.created), 'PPpp') }}
      </div>
      <Separator orientation="vertical" class="hidden h-4 md:block" />
      <div>
        Updated:
        {{ format(new Date(ticket.updated), 'PPpp') }}
      </div>
    </div>
    
    <Separator orientation="vertical" class="hidden h-4 md:block" />
    
    <div class="flex flex-wrap items-center gap-2">
      <div class="flex items-center gap-2">
        <span class="text-xs font-medium">Severity:</span>
        <Select :modelValue="severity" @update:modelValue="updateSeverity">
          <SelectTrigger class="h-7 w-32">
            <SelectValue placeholder="Select severity">
              <Badge v-if="severity" :variant="getSeverityColor(severity)" class="text-xs">
                {{ severity }}
              </Badge>
              <span v-else class="text-xs text-muted-foreground">None</span>
            </SelectValue>
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="low">
              <Badge variant="outline" class="text-xs">Low</Badge>
            </SelectItem>
            <SelectItem value="medium">
              <Badge variant="secondary" class="text-xs">Medium</Badge>
            </SelectItem>
            <SelectItem value="high">
              <Badge variant="default" class="text-xs">High</Badge>
            </SelectItem>
            <SelectItem value="critical">
              <Badge variant="destructive" class="text-xs">Critical</Badge>
            </SelectItem>
          </SelectContent>
        </Select>
      </div>

      <Separator orientation="vertical" class="hidden h-4 md:block" />

      <div class="flex items-center gap-2">
        <span class="text-xs font-medium">Deadline:</span>
        <Popover>
          <PopoverTrigger as-child>
            <Button
              variant="outline"
              :class="
                cn(
                  'h-7 w-48 justify-start text-left font-normal',
                  !deadline && 'text-muted-foreground'
                )
              "
            >
              <CalendarIcon class="mr-2 h-4 w-4" />
              {{ deadline ? format(deadline, 'PPP') : 'Pick a date' }}
            </Button>
          </PopoverTrigger>
          <PopoverContent class="w-auto p-0">
            <Calendar :modelValue="deadline" @update:modelValue="updateDeadline" />
          </PopoverContent>
        </Popover>
      </div>
    </div>
  </div>
</template>
