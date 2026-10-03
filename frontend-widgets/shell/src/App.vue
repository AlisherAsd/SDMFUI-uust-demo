<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import RenderMachine from './render-machine/RenderMachine.vue';
import AdminPanel from './admin/AdminPanel.vue';
import { getContext, type Context } from './api/bff.ts';

const context = ref<Context>({})

const currentPath = window.location.pathname
const currentPage = currentPath.split('/')[1] || 'home'

const currentPageIsAdmin = computed(() => currentPage === 'admin')

onMounted(async () => {
  const res = await getContext()
  context.value = res
})
</script>

<template>
  <div>
    <RenderMachine v-if="!currentPageIsAdmin" :current-page="currentPage" :context="context" />
    <AdminPanel v-else :context="context" />
  </div>
</template>
