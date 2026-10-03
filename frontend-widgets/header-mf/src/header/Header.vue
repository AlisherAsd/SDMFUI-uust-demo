<script setup lang="ts">
import { ref } from 'vue';
import s from './Header.module.css'
import User from '@primeicons/vue/user';

interface Context {
    user?: {
        username?: string;
        isAuthorization?: boolean;
    };
}


const props = defineProps<{
    context?: Context
}>()

interface ItemsLink {
    title: string
    value: string
}

const pages = ref<ItemsLink[]>([
    {
        title: 'Главная',
        value: 'home'
    },
    {
        title: 'Пользователи',
        value: '/users'
    },
    {
        title: 'О нас',
        value: '/about'
    },
])
</script>

<template>
    <div :class="s.container">
        <div :class="s.logo">
            <a href="/home">
                <h1>SDMFUI-uust (demo)</h1>
            </a>
        </div>
        <div>
            <ul>
                <li v-for="l in pages">
                    <a :href="l.value">{{ l.title }}</a>
                </li>
                <li v-if="context?.user?.isAuthorization" :class="s.profile">
                    <a href="/profile">
                        <User size="25" />
                    </a>
                </li>
            </ul>
        </div>
    </div>
</template>