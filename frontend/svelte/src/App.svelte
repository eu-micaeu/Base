<script>
  import { onMount } from 'svelte';
  import { getUsers, registerUser, loginUser, baseUrl } from './lib/api.js';
  import Toasts from './lib/Toasts.svelte';
  import { notify } from './lib/toast.js';
  export let apiUrl = baseUrl; // still allow override if parent passes

  let users = [];
  let loadingUsers = false;
  let errorUsers = '';

  // Forms
  let reg = { name: '', email: '', password: '' };
  let log = { email: '', password: '' };
  let loadingRegister = false;
  let loadingLogin = false;
  let showRegister = false;
  let showLogin = false;

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

  // Endpoint /health removido

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
      showLogin = false;
    } catch (e) {
      notify(`Falha no login: ${e.message}`, 'error');
    } finally {
      loadingLogin = false;
    }
  }

  onMount(() => {
    // Inicialização sem verificação de saúde
  });
</script>

<Toasts position="top-right" />
<main>
  <div class="actions">
    <button on:click={() => showRegister = true}>Registrar</button>
    <button on:click={() => showLogin = true}>Login</button>
  </div>
  {#if showRegister}
    <div class="modal-overlay" on:click={() => !loadingRegister && (showRegister = false)}>
      <div class="modal" on:click|stopPropagation>
        <button class="close" on:click={() => showRegister = false} aria-label="Fechar">×</button>
        <h2>Registrar</h2>
        <form on:submit|preventDefault={doRegister}>
          <label>Nome <input bind:value={reg.name} placeholder="Seu nome" required /></label>
          <label>Email <input type="email" bind:value={reg.email} placeholder="voce@exemplo.com" required /></label>
          <label>Senha <input type="password" bind:value={reg.password} minlength="6" required /></label>
          <div class="form-actions">
            <button type="submit" disabled={loadingRegister}>{loadingRegister ? 'Registrando...' : 'Registrar'}</button>
            <button type="button" on:click={() => showRegister = false} disabled={loadingRegister}>Cancelar</button>
          </div>
        </form>
      </div>
    </div>
  {/if}
  {#if showLogin}
    <div class="modal-overlay" on:click={() => !loadingLogin && (showLogin = false)}>
      <div class="modal" on:click|stopPropagation>
        <button class="close" on:click={() => showLogin = false} aria-label="Fechar">×</button>
        <h2>Login</h2>
        <form on:submit|preventDefault={doLogin}>
          <label>Email <input type="email" bind:value={log.email} placeholder="voce@exemplo.com" required /></label>
          <label>Senha <input type="password" bind:value={log.password} required /></label>
          <div class="form-actions">
            <button type="submit" disabled={loadingLogin}>{loadingLogin ? 'Entrando...' : 'Entrar'}</button>
            <button type="button" on:click={() => showLogin = false} disabled={loadingLogin}>Cancelar</button>
          </div>
        </form>
      </div>
    </div>
  {/if}
  {#if errorUsers}
    <p class="error">Erro usuários: {errorUsers}</p>
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
  form { display: grid; gap: 0.5rem; }
  label { display: grid; gap: 0.25rem; font-size: 0.9rem; }
  input { padding: 0.5rem; border: 1px solid #ccc; border-radius: 4px; }
  .modal-overlay { position: fixed; inset: 0; background: rgba(0,0,0,0.45); display: flex; align-items: center; justify-content: center; z-index: 1000; }
  .modal { background: #fff; width: min(400px, 90%); border-radius: 12px; padding: 1.5rem 1.25rem 1.25rem; box-shadow: 0 6px 24px rgba(0,0,0,0.25); position: relative; }
  .modal h2 { margin-top: 0; }
  .close { position: absolute; top: 0.5rem; right: 0.75rem; background: transparent; border: none; font-size: 1.25rem; cursor: pointer; }
  .form-actions { display: flex; gap: 0.5rem; margin-top: 0.25rem; }
  .form-actions button:last-child { background: #666; color: #fff; }
 </style>
