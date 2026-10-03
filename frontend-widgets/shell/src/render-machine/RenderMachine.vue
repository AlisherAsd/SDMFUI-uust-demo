<script setup lang="ts">
import type { Context } from '@/api/bff';
import { getLayout } from '@/api/remotes';
import { useLoadWidgets } from '@/core/useLoadRemotes';
import { onBeforeMount, shallowRef, type Component } from 'vue';

const props = defineProps<{
  currentPage: string
  context: Context
}>()

const components = shallowRef<Component[]>([])

onBeforeMount(async () => {
  const widgets = await getLayout(props.currentPage)
  const result = await useLoadWidgets(widgets)
  components.value = result
})
</script>

<template>
  <div v-for="c in components">
    <component :is="c" :key="c.name" :context="props.context" />
  </div>
</template>
