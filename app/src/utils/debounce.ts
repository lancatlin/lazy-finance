export function debounce<T extends (...args: Parameters<T>) => ReturnType<T>>(
  callback: T,
  wait: number
): (...args: Parameters<T>) => Promise<ReturnType<T>> {
  let timeoutId: ReturnType<typeof setTimeout> | null = null;
  return async (...args: Parameters<T>): Promise<ReturnType<T>> => {
    clearTimeout(timeoutId as ReturnType<typeof setTimeout>);
    return new Promise((resolve) => {
      timeoutId = setTimeout(() => resolve(callback(...args)), wait);
    });
  };
}
