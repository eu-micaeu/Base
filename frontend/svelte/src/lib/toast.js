import { writable } from 'svelte/store';

export const toasts = writable([]);

let id = 0;
export function notify(message, type = 'info', timeout = 3000) {
    const t = { id: ++id, message, type };
    toasts.update((all) => [t, ...all]);
    if (timeout > 0) {
        setTimeout(() => dismiss(t.id), timeout);
    }
}

export function dismiss(id) {
    toasts.update((all) => all.filter((t) => t.id !== id));
}
