<script setup lang="ts">
import { onMounted, ref } from 'vue';
import s from './ItemsList.module.css'
import { getItems, type Item } from '@/api/items';

const items = ref<Item[]>([])

onMounted(async () => {
  items.value = await getItems()
})
</script>

<template>
  <div :class="s.container">
    <div :class="s.list">
      <div :class="s.item" v-for="item in items" :key="item.description">
        <img
          src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR-rSLeZKd7EPxtx6lbzxF2iej1uBCCTFRjVmWDxsI0jA&s=10" />
        <b>{{ item.title }}</b>
        <p>{{ item.description }}</p>
        <b>{{ item.price }}</b>
      </div>
    </div>
  </div>
</template>