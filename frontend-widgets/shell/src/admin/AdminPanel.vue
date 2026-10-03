<script setup lang="ts">
import { onMounted, ref, watch } from 'vue';
import type { Widget } from '@/core/useLoadRemotes';
import s from './AdminPanel.module.css'
import { addWidgetOnPage, createPage, createWidget, deleteWidgetFromPage, getLayout, getPages, getWidgets, type CreatePage, type Page, type WidgetPage } from '@/api/remotes';
import type { Context } from '@/api/bff';


defineProps<{ context: Context }>()

export type CreateWidget = Omit<Widget, 'id'>

const widgets = ref<Widget[]>([])
const pageWidgets = ref<Widget[]>([])
const pages = ref<Page[]>([])

const getWidget = (): CreateWidget => ({
  mfName: '',
  componentName: '',
  entryUrl: '',
})

type PartialWidgetPage = Partial<WidgetPage>
type PartialPage = Partial<Page>

const widget = ref<CreateWidget>(getWidget())
const widgetPage = ref<PartialWidgetPage>({})
const page = ref<PartialPage>({})

const selectedPageName = ref('home')

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


const handleAddWidgetOnPage = async () => {
  if (!widgetPage.value.level || !widgetPage.value.pageId || !widgetPage.value.widgetId) return
  await addWidgetOnPage(widgetPage.value as WidgetPage)
}

const handleDeleteWidgetFromPage = async (pageId: number, widgetId?: number) => {
  if (!pageId || !widgetId) return
  const ready = confirm('Вы точно хотите удалить виджет?')
  if (!ready) return
  await deleteWidgetFromPage(pageId, widgetId)
}

onMounted(async () => {
  widgets.value = await getWidgets()
  pages.value = await getPages()
})

watch(selectedPageName, async (val: string) => {
  if (val) {
    pageWidgets.value = await getLayout(val)
  }
}, {
  immediate: true
})
</script>

<template>
  <h1 :class="s.logo">LAYOUT ADMIN</h1>
  <div :class="s.container">
    <div>
      <div>
        <h3>Доступные виджеты</h3>
        <div>
          <div v-if="widgets.length">
            <div v-for="w in widgets" :key="w.entryUrl">
              {{ JSON.stringify(w, null, 2) }}
            </div>
          </div>
          <h3 v-else>На данной странице не найдено виджетов</h3>
        </div>
      </div>
      <div>
        <h3>Посмотреть виджеты на странице</h3>
        <b>Выберите страницу</b>
        <select v-model="selectedPageName">
          <option v-for="p in pages" :value="p.name" :key="p.id">{{ p.name }}</option>
        </select>
        <div v-if="pageWidgets.length">
          <div v-for="w in pageWidgets" :key="w.entryUrl">
            {{ JSON.stringify(w, null, 2) }}
            <button :class="s.clear"
              @click="() => handleDeleteWidgetFromPage(w.id, pages.find(p => p.name === selectedPageName)?.id)">
              Удалить</button>
          </div>
        </div>
      </div>
    </div>
    <div>
      <div>
        <h3>Создать виджет</h3>
        <div :class="s.form">
          <b>mfName</b>
          <input placeholder="Введите название микросервиса виджета (mf-xxx-xxx)" v-model="widget.mfName" />
          <b>componentName</b>
          <input placeholder="Введите название компонента с большой буквы" v-model="widget.componentName" />
          <b>entryUrl</b>
          <input placeholder="Введите адрес remoteEntry файла (xxxxx/remoteEntry.js)" v-model="widget.entryUrl" />
        </div>
        <div>
          <button :class="s.save" @click="handleCreateWidget">Создать</button>
          <button :class="s.clear" @click="cleanForm">Очистить</button>
        </div>
      </div>
      <div>
        <h3>Создать страницу</h3>
        <div :class="s.form">
          <b>Название</b>
          <input placeholder="Введите название страницы" v-model="page.name" />
        </div>
        <button :class="s.save" @click="handleCreatePage">Создать</button>
      </div>
      <div>
        <h3>Добавить виджет на страницу</h3>
        <div :class="s.form">
          <b>Страница на которую надо добавить виджет</b>
          <select v-model="widgetPage.pageId">
            <option v-for="page in pages" :value="page.id">{{ page.name }}</option>
          </select>
          <b>Виджет который надо добавить</b>
          <select v-model="widgetPage.widgetId">
            <option v-for="widget in widgets" :value="widget.id">{{ `${widget.mfName}/${widget.componentName}` }}
            </option>
          </select>
          <b>Место виджета</b>
          <input type="number" placeholder="Введите порядковый номер" v-model="widgetPage.level" />
        </div>
        <button :class="s.save" @click="handleAddWidgetOnPage">Добавить</button>
      </div>
    </div>
  </div>
</template>