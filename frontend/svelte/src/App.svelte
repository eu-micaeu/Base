<script>
  import { onMount } from 'svelte';
  import { getUsers, registerUser, loginUser, baseUrl } from './lib/api.js';
  import Toasts from './lib/Toasts.svelte';
  import { notify } from './lib/toast.js';
  import Router from 'svelte-spa-router';
  import Navbar from './components/Navbar.svelte';
  import Home from './routes/Home.svelte';
  import Sobre from './routes/Sobre.svelte';
  export let apiUrl = baseUrl; // still allow override if parent passes

  // Forms
  let reg = { name: '', email: '', password: '' };
  let log = { email: '', password: '' };
  let loadingRegister = false;
  let loadingLogin = false;
  let showRegister = false;

  async function doRegister() {
    loadingRegister = true;
    try {
      const res = await registerUser(reg);
      notify('Registro concluído!', 'success');
      // opcional: limpar formulário
      reg = { name: '', email: '', password: '' };
      showRegister = false;
    } catch (e) {
      notify(`Falha no registro: ${e.message}`, 'error');
    } finally {
      loadingRegister = false;
    }
  }

  async function doLogin() {
    loadingLogin = true;
    try {
      const res = await loginUser(log);
      if (res.token) {
        localStorage.setItem('jwt_token', res.token);
      }
      notify('Login efetuado!', 'success');
    } catch (e) {
      notify(`Falha no login: ${e.message}`, 'error');
    } finally {
      loadingLogin = false;
    }
  }

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
