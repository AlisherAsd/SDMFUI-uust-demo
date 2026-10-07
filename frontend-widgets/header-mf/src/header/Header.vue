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


interface ItemsLink {
    id: number
    title: string
    value: string
}

const props = defineProps<{
    context?: Context
    links?: ItemsLink[]
}>()
</script>

<template>
    <div :class="s.container">
        <div :class="s.logo">
            <a href="/home">
                <h1>SDMFUI-uust (demo)</h1>
            </a>
        </div>
        <div>
            <ul v-if="links?.length">
                <li v-for="l in links" :key="l.id">
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