<script setup lang="ts">
import { loadRemote, registerRemotes } from '@module-federation/enhanced/runtime';
import { computed, onBeforeMount, ref, shallowRef, watch, type Component } from 'vue';
import { useLoadRemotes } from './core/useLoadRemotes';
import { getRemotes } from './api/remotes';
import { useRoute, useRouter } from 'vue-router';

const components = shallowRef<Component[]>([])

const currentPath = window.location.pathname 
const currentPage = currentPath.split('/')[1] || 'home'  

onBeforeMount(async () => {
  const remotes = await getRemotes(currentPage)
  const res = await useLoadRemotes(remotes)
  components.value = res
})
</script>

<template>
  <div v-for="c in components">
    <component :is="c" />
  </div>
</template>
