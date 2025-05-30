<script lang="ts">
  import { blur } from "svelte/transition";
  import { toastManager } from "../../lib/stores/toasts.svelte";
</script>

<toast-stack>
  {#each toastManager.toasts as toast (toast.id)}
    <toast in:blur data-type={toast.type}>
      {toast.message}
    </toast>
  {/each}
</toast-stack>

<style>
  * {
    display: block;
  }

  toast-stack {
    position: fixed;
    top: 40px;
    left: 50%;
    transform: translateX(-50%);
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 8px;
    height: 0px;
    z-index: 1000; /* Ensure it appears above other content */
  }

  toast {
    height: fit-content;
    width: fit-content;
    flex-shrink: 0;
    padding: 6px 12px;
    font-size: 18px;
    font-weight: 700;
    border-radius: 4px;
  }

  toast[data-type="info"] {
    background-color: #334452; /* Green for success */
    color: #99b8d2;
  }

  toast[data-type="error"] {
    background-color: #6c4646; /* Red for error */
    color: #db7171;
  }
</style>
