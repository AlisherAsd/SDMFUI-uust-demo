import { loadRemote, registerRemotes } from "@module-federation/enhanced/runtime"
import { ref, type Component } from "vue"

export interface Remotes {
  mfName: string
  componentName: string
  entryUrl :string
}

export const useLoadRemotes = async (remotes: Remotes[]): Promise<Component[]> => {
  const components: Component[] = []
  const formatRemotes = remotes.map(r => ({ name: r.mfName, entry: r.entryUrl, type: 'module' }))
  registerRemotes(formatRemotes)

  for (const r of remotes) {
    const module = await loadRemote(`${r.mfName}/${r.componentName}`) as { default: Component }
    components.push(module.default)
  }

  return components
}