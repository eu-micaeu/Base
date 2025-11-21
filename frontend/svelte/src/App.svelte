<script>
  import { onMount } from 'svelte';
  import { baseUrl } from './lib/api.js';
  import Toasts from './lib/Toasts.svelte';
  import Router from 'svelte-spa-router';
  import Navbar from './components/Navbar.svelte';
  import Footer from './components/Footer.svelte';
  import Home from './routes/Home.svelte';
  import Sobre from './routes/Sobre.svelte';
  export let apiUrl = baseUrl; // still allow override if parent passes

  const routes = {
    '/': Home,
    '/sobre': Sobre
  };

  onMount(() => {
    // Inicialização sem verificação de saúde
  });
</script>

<Toasts position="top-right" />
<Navbar />
<main>
  <Router {routes} />
</main>
<Footer />

<style>
  main { font-family: system-ui, sans-serif; padding: 2rem; }
  .actions { margin-bottom: 1rem; }
  button { margin-right: 0.5rem; }
  ul { margin-top: 1rem; padding-left: 1.25rem; }
  body { margin: 0; }
  .error { color: #b00020; font-size: 0.9rem; }
  button[disabled] { opacity: 0.6; cursor: not-allowed; }
  form { display: grid; gap: 0.5rem; }
</style>
