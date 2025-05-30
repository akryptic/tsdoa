<script>
  import { Plus } from "@lucide/svelte";
  import { DeleteAllTasks } from "@wails/go/services/TaskService";
  import { BrowserOpenURL } from "@wails/runtime/runtime";
  import { getUserConfirmation } from "../lib/stores/modal.svelte";
  import { toastManager } from "../lib/stores/toasts.svelte";

  const handleClearDB = async () => {
    const ok = await getUserConfirmation(
      "This will clear the entire database. Are you sure?"
    );
    if (!ok) return;

    try {
      await DeleteAllTasks();
      toastManager.info("Database cleared successfully!");
    } catch (err) {
      toastManager.error(`Failed to clear database: ${err}`);
    }
  };
</script>

<container>
  <keybind-list>
    <h5>Keybinds</h5>
    <keybind-item>
      <key-combo>
        <key>Ctrl</key>
        <Plus size="18" strokeWidth="3" />
        <key>N</key>
      </key-combo>
      <key-action>New Task</key-action>
    </keybind-item>

    <keybind-item>
      <key-combo>
        <key>Escape</key>
      </key-combo>
      <key-action>Close Input Bar</key-action>
    </keybind-item>
  </keybind-list>

  <db-clearance>
    <h5>Miscellaneous</h5>
    <p>Clear Database <button onclick={handleClearDB}>CLEAR</button></p>
  </db-clearance>
</container>

<a-bar>
  Created with 🧠 by <button
    onclick={() => BrowserOpenURL("https://github.com/akryptic")}
  >
    Akryptic
  </button>
</a-bar>

<style>
  * {
    display: block;
  }

  container {
    display: flex;
    flex-direction: column;
    gap: 16px;
    padding: 0 12px;
  }

  container > * {
    border-bottom: 1px solid gray;
  }

  container > *:last-child {
    border-bottom: none;
  }

  keybind-list {
    display: flex;
    width: 100%;
    flex-direction: column;
    gap: 1rem;
    padding: 12px 0;
  }

  keybind-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  key-combo {
    display: flex;
    gap: 8px;
    flex-wrap: wrap;
    align-items: center;
  }

  key {
    background: rgba(255, 255, 255, 0.15);
    border: 1px solid rgba(255, 255, 255, 0.5);
    border-bottom-width: 3px;
    border-radius: 6px;
    padding: 2px 8px;
    font-family: monospace;
    font-size: 14px;
  }

  key-action {
    font-size: 16px;
    font-weight: 500;
    color: #cccccc;
  }

  db-clearance {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  db-clearance p {
    color: #d3d3d3;
    font-size: 18px;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  db-clearance button {
    background-color: #914444;
    padding: 2px 8px;
    color: #e4b5b5;
    border-radius: 4px;
    font-size: 16px;
    cursor: pointer;
  }

  h5 {
    font-size: 16px;
    font-weight: 500;
    color: darkgray;
  }

  a-bar {
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
    color: whitesmoke;
    padding: 8px;
    text-align: center;
    font-size: 18px;
    font-weight: 500;
    justify-content: center;
    gap: 1ch;
    display: flex;
  }

  a-bar button {
    color: #51da48;
    cursor: pointer;
    text-decoration: underline;
    font-weight: 800;
  }
</style>
