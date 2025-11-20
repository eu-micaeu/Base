<script>
  import { toasts, dismiss } from './toast.js';
  export let position = 'top-right';
</script>

<div class="toasts {position}">
  {#each $toasts as t (t.id)}
    <div class="toast {t.type}">
      <span>{t.message}</span>
      <button on:click={() => dismiss(t.id)} aria-label="dismiss">×</button>
    </div>
  {/each}
  <slot />
  <slot name="extra" />
</div>

<style>
  .toasts { position: fixed; z-index: 9999; pointer-events: none; }
  .top-right { top: 1rem; right: 1rem; }
  .top-left { top: 1rem; left: 1rem; }
  .bottom-right { bottom: 1rem; right: 1rem; }
  .bottom-left { bottom: 1rem; left: 1rem; }
  .toast { pointer-events: auto; background: #333; color: #fff; padding: 0.5rem 0.75rem; border-radius: 6px; margin-bottom: 0.5rem; display: flex; align-items: center; gap: 0.5rem; box-shadow: 0 2px 8px rgba(0,0,0,.2); }
  .toast.success { background: #2e7d32; }
  .toast.error { background: #c62828; }
  .toast.info { background: #1565c0; }
  .toast button { background: transparent; border: none; color: inherit; font-size: 1rem; cursor: pointer; }
</style>
