<script lang="ts">
  import { Pencil, Square, SquareCheckBig, Trash2 } from "@lucide/svelte";
  import type { models } from "@wails/go/models";
  import { DeleteStep, ToggleStep } from "@wails/go/services/TaskService";
  import { getUserConfirmation } from "../lib/stores/modal.svelte";
  import { toastManager } from "../lib/stores/toasts.svelte";

  const iconSize = 16;

  interface Props {
    step: models.StepResponse;
    taskId: string;
    onEdit: () => void;
  }

  const { step, onEdit, taskId }: Props = $props();

  const toggleStep = async () => {
    try {
      await ToggleStep(taskId, step.id);
    } catch (err) {
      toastManager.error(`Failed to toggle step: ${err}`);
    }
  };

  const deleteStep = async () => {
    const ok = await getUserConfirmation("This will delete the step.");
    if (!ok) return;
    try {
      await DeleteStep(taskId, step.id);
    } catch (err) {
      toastManager.error(`Failed to delete step: ${err}`);
    }
  };
</script>

<step class:done={step.done}>
  <button onclick={toggleStep}>
    {#if step.done}
      <SquareCheckBig size={iconSize} />
    {:else}
      <Square size={iconSize} />
    {/if}
  </button>
  <p>{step.title}</p>
  {#if !step.done}
    <step-actions>
      <button onclick={onEdit}>
        <Pencil size={iconSize} />
      </button>
      <button id="delete" onclick={deleteStep}>
        <Trash2 size={iconSize} />
      </button>
    </step-actions>
  {/if}
</step>

<style>
  * {
    display: block;
  }

  step {
    font-size: 16px;
    color: #d3d3d3;
  }

  step:hover {
    color: whitesmoke;
  }

  step,
  step-actions {
    display: flex;
    gap: 8px;
    align-items: center;
  }

  step-actions {
    visibility: hidden;
  }

  step:hover step-actions {
    visibility: visible;
  }

  step > * {
    flex-shrink: 0;
  }

  step p {
    flex-shrink: 1;
    width: 100%;
  }

  step.done {
    color: #c2fec2bf;
  }

  step.done p {
    text-decoration: line-through;
  }

  button {
    cursor: pointer;
  }

  button#delete:hover {
    color: lightcoral;
  }
</style>
