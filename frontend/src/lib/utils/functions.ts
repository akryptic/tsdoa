export function withPreventDefault<T extends Event>(handler: (e: T) => void) {
  return (e: T) => {
    e.preventDefault();
    handler(e);
  };
}