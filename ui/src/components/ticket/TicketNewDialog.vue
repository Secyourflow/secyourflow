<script setup lang="ts">
import JSONSchemaFormFields from '@/components/form/JSONSchemaFormFields.vue'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Calendar } from '@/components/ui/v-calendar'
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger
} from '@/components/ui/dialog'
import { FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue
} from '@/components/ui/select'

import { CalendarIcon } from 'lucide-vue-next'

import { useMutation, useQueryClient } from '@tanstack/vue-query'
import { format } from 'date-fns'
import { defineRule, useForm } from 'vee-validate'
import { computed, ref, watch } from 'vue'
import { useRouter } from 'vue-router'

import { pb } from '@/lib/pocketbase'
import type { Ticket, Type } from '@/lib/types'
import { cn, handleError } from '@/lib/utils'

const queryClient = useQueryClient()
const router = useRouter()

const props = defineProps<{
  selectedType: Type
}>()

const isOpen = ref(false)

const addTicketMutation = useMutation({
  mutationFn: (): Promise<Ticket> => {
    if (!pb.authStore.model) return Promise.reject('Not logged in')

    return pb.collection('tickets').create({
      name: name.value,
      description: description.value,
      open: true,
      type: props.selectedType.id,
      schema: props.selectedType.schema,
      state: state.value,
      owner: pb.authStore.model.id,
      severity: severity.value || '',
      deadline: deadline.value ? deadline.value.toISOString() : ''
    })
  },
  onSuccess: (data: Ticket) => {
    router.push(`/tickets/${props.selectedType.id}/${data.id}`)
    queryClient.invalidateQueries({ queryKey: ['tickets'] })
    isOpen.value = false
  },
  onError: handleError
})

defineRule('required', (value: string) => {
  if (!value || !value.length) {
    return 'This field is required'
  }

  return true
})

const validationSchema = computed(() => {
  const fields: Record<string, any> = {
    name: 'required'
  }

  Object.keys(props.selectedType.schema.properties).forEach((key) => {
    const property = props.selectedType.schema.properties[key]
    if (property.type === 'string') {
      fields[key] =
        props.selectedType.schema.required && props.selectedType.schema.required.includes(key)
          ? 'required'
          : ''
    } else if (property.type === 'boolean') {
      fields[key] = ''
    }
  })

  return fields
})

const { handleSubmit, validate } = useForm({
  validationSchema: validationSchema.value
})

const onSubmit = handleSubmit((values: any) => {
  validate().then((res) => {
    if (res.valid) {
      addTicketMutation.mutate(values)
    }
  })
})

const state = ref({})
const name = ref('')
const description = ref('')
const severity = ref<'low' | 'medium' | 'high' | 'critical' | ''>('')
const deadline = ref<Date | undefined>(undefined)

const getSeverityColor = (sev: string) => {
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

watch(
  () => isOpen.value,
  () => {
    if (isOpen.value) {
      name.value = ''
      description.value = ''
      state.value = {}
      severity.value = ''
      deadline.value = undefined
    }
  }
)
</script>

<template>
  <Dialog v-model:open="isOpen">
    <DialogTrigger as-child>
      <Button variant="ghost"> New Ticket</Button>
    </DialogTrigger>
    <DialogContent>
      <DialogHeader>
        <DialogTitle>New Ticket</DialogTitle>
        <DialogDescription>
          Create a new {{ props.selectedType.singular }} ticket.
        </DialogDescription>
      </DialogHeader>

      <form @submit="onSubmit">
        <FormField name="name" v-slot="{ componentField }" v-model="name">
          <FormItem>
            <FormLabel for="name" class="text-right">Name</FormLabel>
            <Input id="name" class="col-span-3" v-bind="componentField" />
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="description" v-slot="{ componentField }" v-model="description">
          <FormItem>
            <FormLabel for="description" class="text-right">Description</FormLabel>
            <Input id="description" class="col-span-3" v-bind="componentField" />
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="severity" v-model="severity">
          <FormItem>
            <FormLabel for="severity" class="text-right">Severity</FormLabel>
            <Select v-model="severity">
              <SelectTrigger>
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
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="deadline" v-model="deadline">
          <FormItem>
            <FormLabel for="deadline" class="text-right">Deadline</FormLabel>
            <Popover>
              <PopoverTrigger as-child>
                <Button
                  variant="outline"
                  :class="
                    cn(
                      'w-full justify-start text-left font-normal',
                      !deadline && 'text-muted-foreground'
                    )
                  "
                >
                  <CalendarIcon class="mr-2 h-4 w-4" />
                  {{ deadline ? format(deadline, 'PPP') : 'Pick a date' }}
                </Button>
              </PopoverTrigger>
              <PopoverContent class="w-auto p-0">
                <Calendar v-model="deadline" />
              </PopoverContent>
            </Popover>
            <FormMessage />
          </FormItem>
        </FormField>

        <JSONSchemaFormFields v-model="state" :schema="selectedType.schema" />

        <DialogFooter class="mt-4 sm:justify-start">
          <Button type="submit"> Save </Button>
          <DialogClose as-child>
            <Button type="button" variant="secondary">Cancel</Button>
          </DialogClose>
        </DialogFooter>
      </form>
    </DialogContent>
  </Dialog>
</template>
