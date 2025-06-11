<script lang="ts">
  import { Settings as SettingsIcon } from "@lucide/svelte";
  import {
    Quit,
    WindowMinimise,
    WindowToggleMaximise,
  } from "@wails/runtime/runtime";
  import { showCustomModal } from "../lib/stores/modal.svelte";
  import Settings from "./Settings.svelte";
</script>

<titlebar style="--wails-draggable:drag">
  <span>TSdoA</span>
  <button id="settings" onclick={() => showCustomModal("Settings", Settings)}>
    <SettingsIcon size="14" />
  </button>

  <wrap>
    <button id="minimize" onclick={WindowMinimise} aria-label="Minimize"
    ></button>
    <button id="maximize" onclick={WindowToggleMaximise} aria-label="Maximize"
    ></button>
    <button id="quit" onclick={Quit} aria-label="Close"></button>
  </wrap>
</titlebar>

<style>
  * {
    display: block;
  }

  titlebar {
    background-color: rgba(0, 0, 0, 0.25);
    display: flex;
    justify-content: space-between;
    padding: 4px 8px;
    flex-shrink: 0;
    font-weight: 600;
    align-items: center;
    font-size: 14px;
    line-height: 16px;
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  }

  wrap {
    display: flex;
    gap: 8px;
    justify-content: center;
    align-items: center;
  }

  button {
    width: 10px;
    height: 10px;
    border-radius: 50%;
  }

  #settings {
    cursor: pointer;
    width: fit-content;
    height: fit-content;
    color: rgb(158, 221, 242);
    transition: transform 0.2s ease-in-out;
  }

  #settings:hover {
    transform: scale(0.9) rotate(45deg);
  }

  button:hover {
    filter: brightness(0.75);
  }

  button#minimize {
    background-color: #51da48;
  }

  button#maximize {
    background-color: #dccc71;
  }

  button#quit {
    background-color: #e66c6c;
  }
</style>
