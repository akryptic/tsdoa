<script lang="ts">
  import { Check, Plus, X } from "@lucide/svelte";
  import {
    AddStepToTask,
    CreateTask,
    EditStep,
    EditTask,
  } from "@wails/go/services/TaskService";
  import { onMount } from "svelte";
  import { toastManager } from "../lib/stores/toasts.svelte";
  import { withPreventDefault } from "../lib/utils/functions";

  interface Props {
    mode: "add-task" | "edit-task" | "add-step" | "edit-step";
    taskId: string | undefined;
    stepId: string | undefined;
    old: string;
    onCancel: () => void;
  }

  const {
    mode = $bindable(),
    taskId = $bindable(),
    stepId = $bindable(),
    old = $bindable(),
    onCancel,
  }: Props = $props();

  let isOpen = $state(false);
  let inputTask = $state("");
  let inpElement: HTMLInputElement | null = null;

  function getPlaceholder() {
    switch (mode) {
      case "edit-task":
        return "Edit task . . .";
      case "add-step":
        return "Add a new step . . .";
      case "edit-step":
        return "Edit step . . .";
      default:
        return "Add a new task . . .";
    }
  }

  const addTask = async () => {
    if (!isOpen) {
      isOpen = true;
      inpElement?.focus();
      return;
    }
    try {
      switch (mode) {
        case "add-task":
          await CreateTask(inputTask.trim());
          break;
        case "edit-task":
          await EditTask(taskId!, inputTask.trim());
          break;
        case "add-step":
          await AddStepToTask(taskId!, inputTask.trim());
          break;
        case "edit-step":
          await EditStep(taskId!, stepId!, inputTask.trim());
          break;
      }
    } catch (err) {
      toastManager.error(`There was an error: ${err}`);
    }
    closeBar();
  };

  $effect(() => {
    if (mode === "add-step") {
      isOpen = true;
      inputTask = "";
      inpElement?.focus();
    } else if (mode === "edit-task" || mode === "edit-step") {
      isOpen = true;
      inputTask = old;
      inpElement?.focus();
    } else {
      isOpen = false;
    }
  });

  const showCheckIcon = $derived(
    mode === "edit-task" || mode === "add-step" || mode === "edit-step"
  );

  const closeBar = () => {
    inputTask = "";
    isOpen = false;
    inpElement?.blur();
    onCancel();
  };

  const handleShortcuts = (event: KeyboardEvent) => {
    if (event.key === "Escape" && isOpen) {
      closeBar();
    }

    if (event.ctrlKey && event.key === "n") {
      if (!isOpen) {
        isOpen = true;
        inpElement?.focus();
      }
    }
  };

  onMount(() => {
    window.addEventListener("keydown", handleShortcuts);

    return () => {
      window.removeEventListener("keydown", handleShortcuts);
    };
  });
</script>

<add-tasks>
  <form onsubmit={withPreventDefault(addTask)} class:open={isOpen}>
    <button id="close" onclick={closeBar} type="button">
      <X size="20" />
    </button>
    <input
      type="text"
      bind:value={inputTask}
      bind:this={inpElement}
      placeholder={getPlaceholder()}
    />
    <button id="trigger" type="submit">
      {#if showCheckIcon}
        <Check size="20" />
      {:else}
        <Plus size="20" />
      {/if}
    </button>
  </form>
</add-tasks>

<style>
  * {
    display: block;
  }

  add-tasks {
    flex-shrink: 0;
    padding: 0 0 8px 0;
  }

  form {
    display: flex;
    align-items: center;
    justify-content: end;
    width: 36px;
    border-radius: 8px;
    margin: 0 auto;
    overflow: hidden;
    transition:
      width 0.4s cubic-bezier(0.4, 0, 0.2, 1),
      border-radius 0.4s linear;
  }

  form:hover {
    border-radius: 18px;
  }

  form.open {
    width: 640px;
    border-radius: 8px;
    background-color: #ceffee1a;
  }

  input {
    width: 100%;
    color: #ceffee;
    padding: 0 8px;
    font-weight: 400;
    font-size: 14px;
  }

  button {
    padding: 8px;
    flex-shrink: 0;
    cursor: pointer;
  }

  button#close {
    background: rgba(255, 95, 95, 0.25);
    color: #ff5f5f;
  }

  button#trigger {
    background-color: rgba(96, 216, 174, 0.25);
    color: #60d8ae;
  }
</style>
