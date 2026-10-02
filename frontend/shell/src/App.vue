<script setup lang="ts">
import { loadRemote, registerRemotes } from '@module-federation/enhanced/runtime';
import { onMounted, ref, shallowRef, type Component } from 'vue';
import { useLoadRemotes, type Remotes } from './core/useLoadRemotes';

const testUrls = ref<Remotes[]>([
  {
    mfName: 'mf-header',
    componentName: 'Header',
    entryUrl: 'http://localhost:5001/remoteEntry.js',
  },
   {
    mfName: 'mf-footer',
    componentName: 'Footer',
    entryUrl: 'http://localhost:5002/remoteEntry.js',
  },
     {
    mfName: 'mf-footer',
    componentName: 'Footer',
    entryUrl: 'http://localhost:5002/remoteEntry.js',
  },
])

const components = shallowRef<Component[]>([])

onMounted(async () => {
  const res = await useLoadRemotes(testUrls.value)
  components.value = res
})
</script>

<template>
  <div v-for="c in components">
    <component :is="c" />
  </div>
</template>
