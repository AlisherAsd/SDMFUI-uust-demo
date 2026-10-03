<script setup lang="ts">
import { computed} from 'vue';
import RenderMachine from './render-machine/RenderMachine.vue';
import AdminPanel from './admin/AdminPanel.vue';

const currentPath = window.location.pathname 
const currentPage = currentPath.split('/')[1] || 'home'  

const currentPageIsAdmin = computed(() => currentPage === 'admin')
</script>

<template>
  <div>
    <RenderMachine v-if="!currentPageIsAdmin" :current-page="currentPage" />
    <AdminPanel v-else />
  </div>
</template>
