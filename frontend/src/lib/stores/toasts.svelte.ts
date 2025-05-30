interface Toast {
  id: string;
  message: string;
  type: "error" | "info";
}

function createToastStore() {
  let toastStore = $state<Toast[]>([]);
  let currentTimeout: NodeJS.Timeout | null = null;

  const removeToast = (id: string) => {
    toastStore = toastStore.filter((toast) => toast.id !== id);
  };

  const addToast = (toast: Toast) => {
    if (currentTimeout) {
      clearTimeout(currentTimeout);
    }

    toastStore = [toast];

    currentTimeout = setTimeout(() => removeToast(toast.id), 2000);
  };

  return {
    get toasts() {
      return toastStore;
    },

    error(message: string) {
      addToast({ id: crypto.randomUUID(), message, type: "error" });
    },

    info(message: string) {
      addToast({ id: crypto.randomUUID(), message, type: "info" });
    },
  };
}

export const toastManager = createToastStore();
export type { Toast };
