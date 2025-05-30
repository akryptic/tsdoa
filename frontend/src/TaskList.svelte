<script lang="ts">
  import { models } from "@wails/go/models";
  import { GetStepsByTaskId, GetTasks } from "@wails/go/services/TaskService";
  import { EventsOff, EventsOn } from "@wails/runtime/runtime";
  import { onMount } from "svelte";
  import InputBar from "./components/InputBar.svelte";
  import LoadMore from "./components/LoadMore.svelte";
  import Modal from "./components/modals/Modal.svelte";
  import Task from "./components/Task.svelte";
  import Titlebar from "./components/Titlebar.svelte";
  import Toasts from "./components/toasts/Toasts.svelte";
  import { Events } from "./lib/utils/const";

  let tasks = $state<models.TaskResponse[]>([]);
  let mode = $state<"add-task" | "edit-task" | "add-step" | "edit-step">(
    "add-task"
  );
  let taskId = $state<string | undefined>(undefined);
  let stepId = $state<string | undefined>(undefined);
  let old = $state<string>("");
  let hasMore = $state<boolean>(true);

  const handleTaskEdit = (id: string, oldTitle: string) => {
    mode = "edit-task";
    old = oldTitle;
    taskId = id;
  };

  const handleAddStep = (taskID: string) => {
    mode = "add-step";
    taskId = taskID;
  };

  const handleEditStep = (taskID: string, id: string, oldTitle: string) => {
    mode = "edit-step";
    old = oldTitle;
    taskId = taskID;
    stepId = id;
  };

  const handleCancel = () => {
    mode = "add-task";
    taskId = undefined;
    stepId = undefined;
    old = "";
  };

  const handleMoreStepLoad = async (taskId: string, offset: number) => {
    try {
      const res = await GetStepsByTaskId(taskId, offset);
      const taskIndex = tasks.findIndex((task) => task.id === res.task_id);
      if (taskIndex !== -1) {
        tasks[taskIndex].steps.push(...(res.steps || []));
        tasks[taskIndex].has_more_steps = res.more_available;
      }
    } catch (err) {
      console.error("Error loading more steps:", err);
    }
  };

  const handleMoreTaskLoad = async () => {
    try {
      const res = await GetTasks(tasks.length);
      tasks.push(...(res.tasks || []));
      hasMore = res.has_more_tasks;
    } catch (err) {
      console.error("Error loading more tasks:", err);
    }
  };

  onMount(() => {
    GetTasks(0)
      .then((res) => {
        tasks = res.tasks || [];
        hasMore = res.has_more_tasks;
      })
      .catch((err) => console.error(err)); // Initial fetch of tasks

    // Backend events to handle task updates
    EventsOn(Events.TaskCreated, (data) => {
      tasks.push(data);
    });

    EventsOn(Events.TaskEdited, (data) => {
      const taskIndex = tasks.findIndex((task) => task.id === data.taskId);

      if (taskIndex !== -1) {
        tasks[taskIndex].title = data.title;
      }
    });

    EventsOn(Events.TaskDeleted, (data) => {
      const taskId = data.taskId;
      tasks = tasks.filter((task) => task.id !== taskId);
    });

    EventsOn(Events.StepAdded, (data) => {
      const taskIndex = tasks.findIndex((task) => task.id === data.taskId);

      if (taskIndex !== -1) {
        tasks[taskIndex].steps.push(data.step);
      }
    });

    EventsOn(Events.StepEdited, (data) => {
      const taskIndex = tasks.findIndex((task) => task.id === data.taskId);

      if (taskIndex !== -1) {
        const stepIndex = tasks[taskIndex].steps.findIndex(
          (step) => step.id === data.step.id
        );

        if (stepIndex !== -1) {
          tasks[taskIndex].steps[stepIndex].title = data.step.title;
          tasks[taskIndex].steps[stepIndex].done = data.step.done;
        }
      }
    });

    EventsOn(Events.StepDeleted, (data) => {
      const taskIndex = tasks.findIndex((task) => task.id === data.taskId);

      if (taskIndex !== -1) {
        tasks[taskIndex].steps = tasks[taskIndex].steps.filter(
          (step) => step.id !== data.stepId
        );
      }
    });

    EventsOn(Events.DBWiped, () => {
      tasks = [];
      hasMore = false;
    });
    return () => {
      // Cleanup event listeners
      EventsOff(Events.TaskCreated);
      EventsOff(Events.TaskEdited);
      EventsOff(Events.TaskDeleted);
      EventsOff(Events.StepAdded);
      EventsOff(Events.StepEdited);
      EventsOff(Events.DBWiped);
    };
  });
</script>

<Titlebar />
<main>
  {#each tasks as task (task.id)}
    <Task
      onEdit={handleTaskEdit}
      onAddStep={handleAddStep}
      onEditStep={handleEditStep}
      onMoreStepLoad={handleMoreStepLoad}
      {task}
    />
  {/each}

  {#if hasMore}
    <LoadMore type="task" showBars onLoad={handleMoreTaskLoad} />
  {/if}

  {#if tasks.length === 0}
    <p>No tasks available. Add a new task to get started!</p>
  {/if}
</main>
<InputBar bind:mode bind:taskId bind:stepId bind:old onCancel={handleCancel} />
<Toasts />
<Modal />

<style>
  main {
    height: 100%;
    overflow-x: hidden;
    overflow-y: scroll;
    padding: 0 32px;
    margin: 8px 0;
    display: flex;
    flex-direction: column;
  }

  main::-webkit-scrollbar {
    width: 6px;
  }

  main::-webkit-scrollbar-track {
    background-color: rgba(255, 255, 255, 0.1);
    border-radius: 3px;
  }

  main::-webkit-scrollbar-thumb {
    background-color: rgba(255, 255, 255, 0.15);
    border-radius: 3px;
  }

  p {
    text-align: center;
    color: rgba(255, 255, 255, 0.5);
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }
</style>
