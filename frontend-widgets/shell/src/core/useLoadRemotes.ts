import type { WidgetLayout } from "@/api/remotes"
import { loadRemote, registerRemotes } from "@module-federation/enhanced/runtime"
import { ref, type Component } from "vue"
import { toast } from "vue3-toastify"

export interface Widget {
  id: number
  mfName: string
  componentName: string
  entryUrl :string
}

export interface LoadWidget { 
  component: Component,
  level: number
}

export const useLoadWidgets = async (widgets: WidgetLayout[]): Promise<LoadWidget[]> => {
  if (!widgets.length) {
    window.location.pathname = '/home'
    return []
  }
  const components: LoadWidget[] = []
  const formatRemotes = widgets.map(r => ({ name: r.mfName, entry: r.entryUrl, type: 'module' }))
  registerRemotes(formatRemotes)

  for (const w of widgets) {
    try {
      const module = await loadRemote(`${w.mfName}/${w.componentName}`) as { default: Component }
      if (module.default) {
          components.push({ component: module.default, level: w.level})
      }
    } catch (err) {
       toast("Ошибка при построении страницы, не получилось загрузить виджет", {
            type: "error"
        });
    }
  }

  return components
}