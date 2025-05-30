import { writable } from "svelte/store";

type Modal =
  | {
      type: "confirm";
      title: string;
      message: string;
      resolve: (value: boolean) => void;
    }
  | {
      type: "custom";
      title: string;
      component: any;
      props?: Record<string, any>;
    };

export const modal = writable<Modal | null>(null);

export function getUserConfirmation(
  message: string,
  title: string = "Are you sure?",
): Promise<boolean> {
  return new Promise((resolve) => {
    modal.set({
      type: "confirm",
      title,
      message,
      resolve,
    });
  });
}

export function showCustomModal(
  title: string,
  component: any,
  props: Record<string, any> = {}
): void {
  modal.set({
    type: "custom",
    title,
    component,
    props,
  });
}

export function closeModal(): void {
  modal.set(null);
}
