<script setup lang="ts">
import { addItem, type ItemCreate } from '@/api/items';
import { ref } from 'vue';

const getItem = () => ({
    title: '',
    description: '',
    price: 0,
})

const item = ref<ItemCreate>(getItem())

const handleClearForm = () => item.value = getItem()

const handleAddItem = async () => {
    if (!item.value.description || !item.value.price || !item.value.title) return
    await addItem(item.value)
}
</script>

<template>
    <div>
        <h1>ITEM LIST LAYOUT</h1>
        <b>Введите название товара</b>
        <input placeholder="Введите текст" v-model="item.title" />
        <b>Введите описание товара</b>
        <input placeholder="Введите текст" v-model="item.description" />
        <b>Введите название цену</b>
        <input type="number" placeholder="Введите текст" v-model="item.price" />
        <button @click="handleClearForm">Очистить</button>
        <button @click="handleAddItem">Сохранить</button>
    </div>
</template>