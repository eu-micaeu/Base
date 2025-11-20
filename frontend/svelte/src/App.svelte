<script>
  import { onMount } from 'svelte';
  import { getHealth, getUsers, registerUser, loginUser, baseUrl } from './lib/api.js';
  import Toasts from './lib/Toasts.svelte';
  import { notify } from './lib/toast.js';
  export let apiUrl = baseUrl; // still allow override if parent passes

  let message = 'Pronto para consultar API';
  let users = [];
  let loadingHealth = false;
  let loadingUsers = false;
  let errorHealth = '';
  let errorUsers = '';

  // Forms
  let reg = { name: '', email: '', password: '' };
  let log = { email: '', password: '' };
  let loadingRegister = false;
  let loadingLogin = false;

  async function loadUsers() {
    loadingUsers = true;
    errorUsers = '';
    try {
      users = await getUsers();
    } catch (e) {
      errorUsers = e.message;
    } finally {
      loadingUsers = false;
    }
  }

  async function health() {
    loadingHealth = true;
    errorHealth = '';
    try {
      message = await getHealth();
    } catch (e) {
      errorHealth = e.message;
    } finally {
      loadingHealth = false;
    }
  }

  async function doRegister() {
    loadingRegister = true;
    try {
      const res = await registerUser(reg);
      notify('Registro concluído!', 'success');
      // opcional: limpar formulário
      reg = { name: '', email: '', password: '' };
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
      notify('Login efetuado!', 'success');
    } catch (e) {
      notify(`Falha no login: ${e.message}`, 'error');
    } finally {
      loadingLogin = false;
    }
  }

  onMount(() => {
    // Opcional: verificar saúde automaticamente
    health();
  });
</script>

<Toasts position="top-right" />
<main>
  <h1>{message}</h1>
  <p><small>Backend: {apiUrl}</small></p>
  <div class="actions">
    <button on:click={health} disabled={loadingHealth}>{loadingHealth ? 'Verificando...' : 'Ping API'}</button>
    <button on:click={loadUsers} disabled={loadingUsers}>{loadingUsers ? 'Carregando...' : 'Carregar Usuários'}</button>
  </div>
  <section class="auth">
    <div class="card">
      <h2>Registrar</h2>
      <form on:submit|preventDefault={doRegister}>
        <label>Nome <input bind:value={reg.name} placeholder="Seu nome" required /></label>
        <label>Email <input type="email" bind:value={reg.email} placeholder="voce@exemplo.com" required /></label>
        <label>Senha <input type="password" bind:value={reg.password} minlength="6" required /></label>
        <button type="submit" disabled={loadingRegister}>{loadingRegister ? 'Registrando...' : 'Registrar'}</button>
      </form>
    </div>
    <div class="card">
      <h2>Login</h2>
      <form on:submit|preventDefault={doLogin}>
        <label>Email <input type="email" bind:value={log.email} placeholder="voce@exemplo.com" required /></label>
        <label>Senha <input type="password" bind:value={log.password} required /></label>
        <button type="submit" disabled={loadingLogin}>{loadingLogin ? 'Entrando...' : 'Entrar'}</button>
      </form>
    </div>
  </section>
  {#if errorHealth}
    <p class="error">Erro saúde: {errorHealth}</p>
  {/if}
  {#if errorUsers}
    <p class="error">Erro usuários: {errorUsers}</p>
  {/if}
  {#if users.length}
    <ul>
      {#each users as u}
        <li>{u.name} - {u.email}</li>
      {/each}
    </ul>
  {:else if !loadingUsers && !errorUsers}
    <p>Nenhum usuário carregado ainda.</p>
  {/if}
</main>

<style>
  main { font-family: system-ui, sans-serif; padding: 2rem; }
  .actions { margin-bottom: 1rem; }
  button { margin-right: 0.5rem; }
  ul { margin-top: 1rem; padding-left: 1.25rem; }
  body { margin: 0; }
  .error { color: #b00020; font-size: 0.9rem; }
  button[disabled] { opacity: 0.6; cursor: not-allowed; }
  .auth { display: grid; grid-template-columns: repeat(auto-fit, minmax(240px, 1fr)); gap: 1rem; margin-top: 1rem; }
  .card { border: 1px solid #e0e0e0; border-radius: 8px; padding: 1rem; }
  form { display: grid; gap: 0.5rem; }
  label { display: grid; gap: 0.25rem; font-size: 0.9rem; }
  input { padding: 0.5rem; border: 1px solid #ccc; border-radius: 4px; }
 </style>
