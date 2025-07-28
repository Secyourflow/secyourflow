<script setup lang="ts">
import Secyourflow from '@/components/common/SecYourFlowLogo.vue'
import IncidentNav from '@/components/sidebar/IncidentNav.vue'
import NavList from '@/components/sidebar/NavList.vue'
import UserDropDown from '@/components/sidebar/UserDropDown.vue'
import { Button } from '@/components/ui/button'
import { Separator } from '@/components/ui/separator'

import { Menu } from 'lucide-vue-next'

import { cn } from '@/lib/utils'
import { useSecyourflowStore } from '@/store/secyourflow'

const secyourflowStore = useSecyourflowStore()
</script>

<template>
  <div
    :class="
      cn(
        'flex min-w-48 shrink-0 flex-col border-r bg-popover',
        secyourflowStore.sidebarCollapsed && 'min-w-[50px]'
      )
    "
  >
    <div class="flex h-[57px] items-center border-b bg-background">
      <Secyourflow :class="cn('size-12', secyourflowStore.sidebarCollapsed ? 'flex-1' : '')" />

      <h1 class="text-xl font-bold" v-if="!secyourflowStore.sidebarCollapsed">SecYourFlow</h1>
    </div>
    <NavList
      :is-collapsed="secyourflowStore.sidebarCollapsed"
      :links="[
        {
          title: 'Dashboard',
          icon: 'PanelsTopLeft',
          variant: 'ghost',
          to: '/dashboard'
        }
      ]"
    />
    <Separator />
    <IncidentNav :is-collapsed="secyourflowStore.sidebarCollapsed" />

    <div class="flex-1" />

    <Separator />
    <NavList
      :is-collapsed="secyourflowStore.sidebarCollapsed"
      :links="[
        {
          title: 'Reactions',
          icon: 'Zap',
          variant: 'ghost',
          to: '/reactions'
        }
      ]"
    />
    <Separator />
    <UserDropDown :is-collapsed="secyourflowStore.sidebarCollapsed" />
    <Separator />
    <div :class="cn('flex h-14 items-center px-3', !secyourflowStore.sidebarCollapsed && 'px-2')">
      <Button
        variant="ghost"
        @click="secyourflowStore.toggleSidebar()"
        size="default"
        :class="
          cn(
            'p-0',
            secyourflowStore.sidebarCollapsed && 'w-9',
            !secyourflowStore.sidebarCollapsed && 'w-full justify-start px-3'
          )
        "
      >
        <Menu class="size-4" />
        <span v-if="!secyourflowStore.sidebarCollapsed" class="ml-2">Toggle Sidebar</span>
      </Button>
    </div>
  </div>
</template>
