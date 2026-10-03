<script setup lang="ts">
import { getLayout, getWidgets } from '@/api/remotes';
import { useLoadWidgets, type LoadWidget } from '@/core/useLoadRemotes';
import { loadRemote, registerRemotes } from '@module-federation/enhanced/runtime';
import { computed, onBeforeMount, ref, shallowRef, watch, type Component } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const props = defineProps<{
    currentPage: string
}>()

const components = shallowRef<Component[]>([])

onBeforeMount(async () => {
  const widgets = await getLayout(props.currentPage)
  const result = await useLoadWidgets(widgets)
  components.value = result.sort((a, b) => a.level - b.level).map(lw => lw.component)
})
</script>

<template>
  <div v-for="c in components">
    <component :is="c" />
  </div>
</template>
