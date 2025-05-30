import { mount } from 'svelte'
import './app.css'
import TaskList from './TaskList.svelte'

const app = mount(TaskList, {
  target: document.getElementById('app')!,
})

export default app
