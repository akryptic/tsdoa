<script lang="ts">
  import { RefreshCcw } from "@lucide/svelte";

  const {
    type,
    onLoad,
    showBars,
  }: { type: "task" | "step"; onLoad: () => void; showBars?: boolean } =
    $props();
</script>

<flag>
  {#if showBars}
    <bar></bar>
  {/if}
  <span>There are more {type}s available.</span>
  <button onclick={onLoad}><RefreshCcw size="10" />Load</button>
  {#if showBars}
    <bar></bar>
  {/if}
</flag>

<style>
  * {
    display: block;
  }

  flag,
  button {
    display: flex;
    align-items: center;
    justify-content: center;
  }

  flag {
    width: 100%;
    font-size: 12px;
    gap: 8px;
  }

  flag > * {
    flex-shrink: 0;
  }

  button {
    cursor: pointer;
    gap: 4px;
    height: 12px;
  }

  button:hover {
    color: rgb(140, 199, 140);
  }

  bar {
    display: block;
    height: 1px;
    width: 100%;
    flex-shrink: 1;
    background-color: darkgray;
  }
</style>
