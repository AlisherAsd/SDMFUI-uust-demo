<script setup lang="ts">
import { computed } from 'vue';
import ItemList from './items-list/ItemsList.vue';
import AdminPanel from './admin/AdminPanel.vue';

const currentPath = window.location.pathname
const currentPage = currentPath.split('/')[1] || 'home'

const currentPageIsAdmin = computed(() => currentPage === 'admin')
</script>

<template>
  <ItemList v-if="!currentPageIsAdmin" />
  <AdminPanel v-else />
</template>