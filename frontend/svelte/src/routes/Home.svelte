<script>
    // import { onMount } from "svelte"; // já importado acima
    import { setupOnlineCheck } from "../utils/onlineCheck.js";

    // let isLoggedIn = false; // já declarado acima
    let showOnlinePrompt = false;
    let onlineCheckCleanup;

    function showPrompt() {
        showOnlinePrompt = true;
    }
    function hidePrompt() {
        showOnlinePrompt = false;
    }

    function confirmOnline() {
        hidePrompt();
        // onlineCheck util já limpa o timeout
    }

    onMount(() => {
        isLoggedIn = !!localStorage.getItem("jwt_token");
        onlineCheckCleanup = setupOnlineCheck({
            isLoggedInFn: () => isLoggedIn,
            onLogout: logout,
            showPrompt,
            hidePrompt,
        });
    });

    import { onDestroy } from "svelte";
    onDestroy(() => {
        if (onlineCheckCleanup) onlineCheckCleanup.cleanup();
    });
    import { onMount } from "svelte";

    let isLoggedIn = false;

    // Checa se o usuário está logado ao carregar
    onMount(() => {
        isLoggedIn = !!localStorage.getItem("jwt_token");
    });

    function logout() {
        localStorage.removeItem("jwt_token");
        // Remove cookie do token
        document.cookie =
            "token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
        isLoggedIn = false;
        notify("Logout realizado!", "success");
    }
    import { registerUser, loginUser, setTokenCookie } from "../lib/api.js";
    import { notify } from "../lib/toast.js";

    let reg = { name: "", email: "", password: "" };
    let log = { email: "", password: "" };
    let loadingRegister = false;
    let loadingLogin = false;
    let showRegister = false;
    let showLogin = false;

    async function doRegister() {
        loadingRegister = true;
        try {
            const res = await registerUser(reg);
            notify("Registro concluído!", "success");
            reg = { name: "", email: "", password: "" };
            showRegister = false;
        } catch (e) {
            notify(`Falha no registro: ${e.message}`, "error");
        } finally {
            loadingRegister = false;
        }
    }

    async function doLogin() {
        loadingLogin = true;
        try {
            const res = await loginUser(log);
            if (res.token) {
                setTokenCookie(res.token);
                localStorage.setItem("jwt_token", res.token);
                isLoggedIn = true;
            }
            notify("Login efetuado!", "success");
            showLogin = false;
        } catch (e) {
            notify(`Falha no login: ${e.message}`, "error");
        } finally {
            loadingLogin = false;
        }
    }
</script>

<main>
    {#if isLoggedIn}
        <h1>Bem-vindo de volta!</h1>
        <p>Você está logado.</p>
        <div class="actions">
            <button on:click={logout}>Logout</button>
        </div>
        {#if showOnlinePrompt}
            <div class="online-prompt">
                <span>Confirme que você está online!</span>
                <button on:click={confirmOnline}>Estou online</button>
            </div>
        {/if}
    {:else}
        <h1>Bem-vindo à Home!</h1>
        <p>Esta é a página inicial do app.</p>
        <div class="actions">
            <button on:click={() => (showRegister = true)}>Registrar</button>
            <button on:click={() => (showLogin = true)}>Login</button>
        </div>
    {/if}
    {#if showRegister}
        <div
            class="modal-overlay"
            role="button"
            tabindex="0"
            aria-label="Fechar modal de registro"
            on:click={() => !loadingRegister && (showRegister = false)}
            on:keydown={(e) => {
                if ((e.key === "Enter" || e.key === " ") && !loadingRegister) {
                    showRegister = false;
                }
            }}
        >
            <div
                class="modal"
                role="dialog"
                aria-modal="true"
                on:click|stopPropagation
            >
                <button
                    class="close"
                    on:click={() => (showRegister = false)}
                    aria-label="Fechar">×</button
                >
                <h2>Registrar</h2>
                <form on:submit|preventDefault={doRegister}>
                    <label
                        >Nome <input
                            bind:value={reg.name}
                            placeholder="Seu nome"
                            required
                        /></label
                    >
                    <label
                        >Email <input
                            type="email"
                            bind:value={reg.email}
                            placeholder="voce@exemplo.com"
                            required
                        /></label
                    >
                    <label
                        >Senha <input
                            type="password"
                            bind:value={reg.password}
                            minlength="6"
                            required
                        /></label
                    >
                    <div class="form-actions">
                        <button type="submit" disabled={loadingRegister}
                            >{loadingRegister
                                ? "Registrando..."
                                : "Registrar"}</button
                        >
                        <button
                            type="button"
                            on:click={() => (showRegister = false)}
                            disabled={loadingRegister}>Cancelar</button
                        >
                    </div>
                </form>
            </div>
        </div>
    {/if}
    {#if showLogin}
        <div
            class="modal-overlay"
            role="button"
            tabindex="0"
            aria-label="Fechar modal de login"
            on:click={() => !loadingLogin && (showLogin = false)}
            on:keydown={(e) => {
                if ((e.key === "Enter" || e.key === " ") && !loadingLogin) {
                    showLogin = false;
                }
            }}
        >
            <div
                class="modal"
                role="dialog"
                aria-modal="true"
                on:click|stopPropagation
            >
                <button
                    class="close"
                    on:click={() => (showLogin = false)}
                    aria-label="Fechar">×</button
                >
                <h2>Login</h2>
                <form on:submit|preventDefault={doLogin}>
                    <label
                        >Email <input
                            type="email"
                            bind:value={log.email}
                            placeholder="voce@exemplo.com"
                            required
                        /></label
                    >
                    <label
                        >Senha <input
                            type="password"
                            bind:value={log.password}
                            required
                        /></label
                    >
                    <div class="form-actions">
                        <button type="submit" disabled={loadingLogin}
                            >{loadingLogin ? "Entrando..." : "Entrar"}</button
                        >
                        <button
                            type="button"
                            on:click={() => (showLogin = false)}
                            disabled={loadingLogin}>Cancelar</button
                        >
                    </div>
                </form>
            </div>
        </div>
    {/if}
</main>

<style>
    .online-prompt {
        position: fixed;
        left: 50%;
        top: 10%;
        transform: translateX(-50%);
        background: #2575fc;
        color: #fff;
        padding: 1.2rem 2rem;
        border-radius: 12px;
        box-shadow: 0 2px 16px rgba(80, 80, 160, 0.18);
        z-index: 2000;
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 1rem;
        animation: fadeInPrompt 0.5s;
    }
    @keyframes fadeInPrompt {
        from {
            opacity: 0;
            transform: translateX(-50%) scale(0.8);
        }
        to {
            opacity: 1;
            transform: translateX(-50%) scale(1);
        }
    }
    .online-prompt button {
        background: #fff;
        color: #2575fc;
        border: none;
        border-radius: 8px;
        padding: 0.5rem 1.2rem;
        font-weight: bold;
        cursor: pointer;
        box-shadow: 0 2px 8px #0001;
        transition:
            background 0.2s,
            color 0.2s;
    }
    .online-prompt button:hover {
        background: #2575fc;
        color: #fff;
    }
    .actions {
        margin-bottom: 1rem;
    }
    button {
        margin-right: 0.5rem;
    }
    button[disabled] {
        opacity: 0.6;
        cursor: not-allowed;
    }
    form {
        display: grid;
        gap: 0.5rem;
    }
    label {
        display: grid;
        gap: 0.25rem;
        font-size: 0.9rem;
    }
    input {
        padding: 0.5rem;
        border: 1px solid #ccc;
        border-radius: 4px;
    }
    .modal-overlay {
        position: fixed;
        inset: 0;
        background: rgba(0, 0, 0, 0.45);
        display: flex;
        align-items: center;
        justify-content: center;
        z-index: 1000;
    }
    .modal {
        background: #fff;
        width: min(400px, 90%);
        border-radius: 12px;
        padding: 1.5rem 1.25rem 1.25rem;
        box-shadow: 0 6px 24px rgba(0, 0, 0, 0.25);
        position: relative;
    }
    .modal h2 {
        margin-top: 0;
    }
    .close {
        position: absolute;
        top: 0.5rem;
        right: 0.75rem;
        background: transparent;
        border: none;
        font-size: 1.25rem;
        cursor: pointer;
    }
    .form-actions {
        display: flex;
        gap: 0.5rem;
        margin-top: 0.25rem;
    }
    .form-actions button:last-child {
        background: #666;
        color: #fff;
    }
</style>
