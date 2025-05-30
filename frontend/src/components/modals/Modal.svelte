<script>
  import { X } from "@lucide/svelte";
  import { fade } from "svelte/transition";
  import { closeModal, modal } from "../../lib/stores/modal.svelte";
</script>

{#if $modal}
  <overlay transition:fade={{ duration: 100 }}>
    <box>
      <h3>{$modal.title}</h3>
      {#if $modal.type === "confirm"}
        <p>{$modal.message}</p>
        <actions>
          <button
            id="no"
            onclick={() => {
              $modal.resolve(false);
              closeModal();
            }}>No</button
          >
          <button
            id="yes"
            onclick={() => {
              $modal.resolve(true);
              closeModal();
            }}>Yes</button
          >
        </actions>
      {:else if $modal.type === "custom"}
        <button onclick={closeModal} id="close"><X size="16" /></button>
        <svelte:component this={$modal.component} />
      {/if}
    </box>
  </overlay>
{/if}

<style>
  * {
    display: block;
  }

  overlay {
    position: fixed;
    top: 24px;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.65);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
  }

  box {
    position: relative;
    background-color: #2d2f30;
    border-radius: 8px;
    padding: 12px;
    width: 400px;
    max-width: 90%;
    box-shadow: 0 4px 8px 1px rgba(0, 0, 0, 0.25);
    z-index: 1001;
  }

  h3 {
    color: #dadada;
    border-bottom: 1px solid darkgray;
    padding-bottom: 8px;
  }

  p {
    color: #d3d3d3;
    margin: 4px 8px 0 8px;
    font-size: 14px;
    line-height: 1.5;
  }

  button#close {
    position: absolute;
    top: 8px;
    right: 8px;
    background: none;
    border: none;
    color: #d3d3d3;
    cursor: pointer;
  }

  button#yes,
  button#no {
    padding: 4px 12px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 14px;
    font-weight: 500;
    letter-spacing: 0.1ch;
    transition: background-color 0.2s ease-in-out;
  }

  button#no {
    background-color: #994646;
    color: #ffc1c1;
  }

  button#yes {
    background-color: #3b836b;
    color: #a7f7dc;
  }

  actions {
    display: flex;
    justify-content: flex-end;
    gap: 16px;
    margin-top: 16px;
  }
</style>
