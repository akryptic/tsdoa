<script lang="ts">
  import { ChevronDown, Pencil, Plus, Trash2 } from "@lucide/svelte";
  import { models } from "@wails/go/models";
  import { DeleteTask } from "@wails/go/services/TaskService";
  import { getUserConfirmation } from "../lib/stores/modal.svelte";
  import { toastManager } from "../lib/stores/toasts.svelte";
  import LoadMore from "./LoadMore.svelte";
  import Step from "./Step.svelte";

  const iconSize = 16;

  interface Props {
    task: models.TaskResponse;
    onEdit: (id: string, oldTitle: string) => void;
    onAddStep: (taskID: string) => void;
    onEditStep: (taskID: string, id: string, oldTitle: string) => void;
    onMoreStepLoad: (taskId: string, offset: number) => void;
  }

  let { task, onEdit, onAddStep, onEditStep, onMoreStepLoad }: Props = $props();

  const deleteTask = async () => {
    const ok = await getUserConfirmation(
      "Do you really want to delete this task?"
    );
    if (!ok) return;
    try {
      await DeleteTask(task.id);
    } catch (err) {
      toastManager.error(`Failed to delete task: ${err}`);
    }
  };
</script>

<task>
  <task-title>
    <ChevronDown size="24" />
    <p>{task.title}</p>
    <task-actions>
      <button onclick={() => onEdit(task.id, task.title)}
        ><Pencil size={iconSize} /></button
      >
      <button id="delete" onclick={deleteTask}
        ><Trash2 size={iconSize} /></button
      >
    </task-actions>
  </task-title>

  <steps-container>
    {#each task.steps as step (step.id)}
      <Step
        {step}
        taskId={task.id}
        onEdit={() => onEditStep(task.id, step.id, step.title)}
      />
    {/each}
    {#if task.has_more_steps}
      <wrap>
        <LoadMore
          type="step"
          onLoad={() => onMoreStepLoad(task.id, task.steps.length)}
        />
      </wrap>
    {:else}
      <button onclick={() => onAddStep(task.id)}>
        <Plus size={iconSize} />
        Add Step
      </button>
    {/if}
  </steps-container>
</task>

<style>
  * {
    display: block;
  }

  task-title,
  task-actions {
    display: flex;
    align-items: center;
    gap: 8px;
    color: #ffffffbf;
  }

  task:hover task-title,
  task:hover task-actions {
    color: whitesmoke;
  }

  task-title > * {
    flex-shrink: 0;
  }

  task-title p {
    flex-shrink: 1;
    text-transform: capitalize;
    width: 100%;
    font-size: 20px;
    font-weight: 500;
  }

  task-actions {
    visibility: hidden;
  }

  task:hover task-actions,
  task:hover steps-container button {
    visibility: visible;
  }

  button {
    cursor: pointer;
  }

  button#delete:hover {
    color: lightcoral;
  }

  steps-container {
    display: flex;
    flex-direction: column;
    gap: 4px;
    padding: 0 28px;
  }

  steps-container button {
    display: flex;
    width: fit-content;
    font-size: 12px;
    color: lightblue;
    align-items: center;
    gap: 4px;
    margin-bottom: 4px;
    visibility: hidden;
  }

  wrap {
    width: fit-content;
    margin-left: 24px;
    margin-bottom: 12px;
  }
</style>
