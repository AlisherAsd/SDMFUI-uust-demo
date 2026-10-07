import type { CreateWidget } from "@/admin/AdminPanel.vue";
import type { Widget } from "@/core/useLoadRemotes";
import axios from "axios";
import { toast } from "vue3-toastify";

export interface Page {
  id: number;
  title: string;
  value: string;
}

export type CreatePage = Omit<Page, "id">;
export interface WidgetLayout extends Widget {
  level: number;
}

export interface WidgetPage {
  pageId: number;
  widgetId: number;
  level: number;
}

export const getWidgets = async (): Promise<Widget[]> => {
  return axios
    .get(`http://localhost:8080/widgets`)
    .then((r) => r.data.widgets || [])
    .catch((err) => {
      toast("Ошибка при получении виджетов: " + err.error, {
        type: "error",
      });
    });
};

export const createWidget = async (widget: CreateWidget) => {
  return axios
    .post(`http://localhost:8080/widgets`, widget)
    .then(() => {
      toast("Виджет успешно добавлен", {
        type: "success",
      });
    })
    .catch((err) => {
      toast("Ошибка при построении страницы: " + err.error, {
        type: "error",
      });
    });
};

export const getLayout = async (pageName: string): Promise<WidgetLayout[]> => {
  return axios
    .get(`http://localhost:8080/widgets/${pageName}`)
    .then((r) => r.data.widgets || [])
    .catch((err) => {
      toast("Ошибка при получении виджетов: " + err.error, {
        type: "error",
      });
    });
};

export const getPages = async (): Promise<Page[]> => {
  return axios
    .get(`http://localhost:8080/pages`)
    .then((r) => r.data.pages || [])
    .catch((err) => {
      toast("Ошибка при получении страниц: " + err.error, {
        type: "error",
      });
    });
};

export const createPage = async (page: CreatePage) => {
  return axios
    .post(`http://localhost:8080/pages`, page)
    .then(() => {
      toast("Страница успешно добавлена", {
        type: "success",
      });
    })
    .catch((err) => {
      toast("Ошибка при создании страницы: " + err.error, {
        type: "error",
      });
    });
};

export const addWidgetOnPage = async (widgetPage: WidgetPage) => {
  return axios
    .post(`http://localhost:8080/pages/widget`, widgetPage)
    .then(() => {
      toast("Виджет успешно добавлен на страницу", {
        type: "success",
      });
    })
    .catch((err) => {
      toast("Ошибка при получении страниц: " + err.error, {
        type: "error",
      });
    });
};

export const deleteWidgetFromPage = async (
  pageId: number,
  widgetId: number,
) => {
  return axios
    .delete(`http://localhost:8080/pages/widget/${pageId}/${widgetId}`)
    .then(() => {
      toast("Виджет успешно удален со страницы", {
        type: "success",
      });
    })
    .catch((err) => {
      toast("Ошибка при удалении виджета: " + err.error, {
        type: "error",
      });
    });
};
