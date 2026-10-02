
<script setup lang="ts">
import { ref } from 'vue';
import { createRemotes, getRemotes, type Remote } from './api/api';
import './app.css'

const remotes = ref<Remote[]>([])
const termPage = ref('')
const createRemoteResult = ref<{ok: boolean, mess: string}>({ok: false, mess: ''})

type createRemote = Remote & {page: string}

const getRemote = (): createRemote => ({
  page: '',
  mfName: '',
  componentName: '',
  entryUrl : '',
})

const remote = ref<createRemote>(getRemote())

const cleanForm = () => {
  remote.value = getRemote()
  createRemoteResult.value = {ok: true, mess: ''}
}

const handleGetRemote = async () => {
  if (!termPage.value) return
  remotes.value = await getRemotes(termPage.value)
}

const handleCreateRemote = async () => {
  if (!remote.value.componentName || !remote.value.entryUrl || !remote.value.mfName || !remote.value.page) return
  const res = await createRemotes(remote.value)
  createRemoteResult.value = res
}
</script>

<template>
  <h1>LAYOUT ADMIN</h1>
  <div class="container">
    <div>
      <h3>Посмотреть состав страницы</h3>
      <div>
        <input type="text" v-model="termPage">
        <button @click="handleGetRemote">Посмотреть страницу</button>
      </div>
      <div>
        <div v-if="remotes.length" class="remoteList">
          <div v-for="r in remotes" :key="r.entryUrl">
          {{ JSON.stringify(r, null, 2) }}
          </div>
        </div>
        <div v-else><h3>На данной странице не найдено виджетов</h3></div>
      </div>
    </div>
    <div>
      <h3>Добавить виджет</h3>
      <div class="form">
        <b>Страница</b>
        <input placeholder="Введите страницу на которой должен отображаться виджет" v-model="remote.page" />
        <b>mfName</b>
        <input placeholder="Введите название микросервиса виджета (mf-xxx-xxx)" v-model="remote.mfName" />
        <b>componentName</b>
        <input placeholder="Введите название компонента с большой буквы" v-model="remote.componentName" />
        <b>entryUrl</b>
        <input placeholder="Введите адрес remoteEntry файла (xxxxx/remoteEntry.js)" v-model="remote.entryUrl" />
      </div>
      <div class="formBtn">
        <button class="clear" @click="cleanForm">Очистить</button>
        <button class="save" @click="handleCreateRemote">Сохранить</button>
      </div>
      <div v-if="createRemoteResult.mess">
        <b :class="createRemoteResult.ok ? 'successMessage' : 'errorMessage'">{{ createRemoteResult.mess }}</b>
      </div>
    </div>
  </div>
</template>