
<script setup lang="ts">
import { onMounted, ref } from 'vue';
import type { Widget } from '@/core/useLoadRemotes';
import { addWidgetOnPage, createPage, createWidget, getPages, getWidgets, type CreatePage, type Page, type WidgetPage } from '@/api/remotes';

export type CreateWidget = Omit<Widget, 'id'>

const widgets = ref<Widget[]>([])
const pages = ref<Page[]>([])

const getWidget = (): CreateWidget => ({
  mfName: '',
  componentName: '',
  entryUrl : '',
})

type PartialWidgetPage = Partial<WidgetPage>
type PartialPage = Partial<Page>

const widget = ref<CreateWidget>(getWidget())
const widgetPage = ref<PartialWidgetPage>({})
const page = ref<PartialPage>({})

const cleanForm = () => {
  widget.value = getWidget()
}

const handleCreateWidget = async () => {
  if (!widget.value.componentName || !widget.value.entryUrl || !widget.value.mfName) return
  await createWidget(widget.value)
}

const handleCreatePage = async () => {
    if (!page.value.name) return
    await createPage(page.value as Page)
}


const handleAddWidgetOnPage = async() => {
    if (!widgetPage.value.level || !widgetPage.value.pageId || !widgetPage.value.widgetId) return
    await addWidgetOnPage(widgetPage.value as WidgetPage)
}

onMounted(async () => {
    widgets.value = await getWidgets()
    pages.value = await getPages()
} )
</script>

<template>
  <h1>LAYOUT ADMIN</h1>
  <div class="container">
    <div>
      <h3>Доступные виджеты</h3>
      <div>
        <div v-if="widgets.length">
          <div v-for="r in widgets" :key="r.entryUrl">
          {{ JSON.stringify(r, null, 2) }}
          </div>
        </div>
        <h3 v-else>На данной странице не найдено виджетов</h3>
      </div>
    </div>
    <div>
      <h3>Создать виджет</h3>
      <div class="form">
        <b>mfName</b>
        <input placeholder="Введите название микросервиса виджета (mf-xxx-xxx)" v-model="widget.mfName" />
        <b>componentName</b>
        <input placeholder="Введите название компонента с большой буквы" v-model="widget.componentName" />
        <b>entryUrl</b>
        <input placeholder="Введите адрес remoteEntry файла (xxxxx/remoteEntry.js)" v-model="widget.entryUrl" />
      </div>
      <div class="formBtn">
        <button class="clear" @click="cleanForm">Очистить</button>
        <button class="save" @click="handleCreateWidget">Создать</button>
      </div>      
    </div>
    <div>
      <h3>Создать страницу</h3>
      <div class="form">
        <b>Название</b>
        <input placeholder="Введите название страницы" v-model="page.name" />       
      </div>
      <div class="formBtn">
        <button class="save" @click="handleCreatePage">Создать</button>
      </div>      
    </div>
    <div>
      <h3>Добавить виджет на страницу виджет</h3>
      <div class="form">
        <b>Страница на которую надо добавить виджет</b>
        <select v-model="widgetPage.pageId">
            <option v-for="page in pages" :value="page.id">{{ page.name }}</option>
        </select>
        <b>Виджет который надо добавить</b>
        <select v-model="widgetPage.widgetId">
            <option v-for="widget in widgets" :value="widget.id">{{ `${widget.mfName}/${widget.componentName}` }}</option>
        </select>
        <b>Место виджета</b>
        <input type="number" placeholder="Введите порядковый номер" v-model="widgetPage.level" />
      </div>
      <div class="formBtn">
        <button class="save" @click="handleAddWidgetOnPage">Добавить</button>
      </div>      
    </div>
  </div>
</template>