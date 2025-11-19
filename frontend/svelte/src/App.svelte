<script>
  export let apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080';
  let message = 'Hello Svelte!';
  let users = [];

  async function loadUsers() {
    try {
      const res = await fetch(`${apiUrl}/users`);
      users = await res.json();
    } catch (e) {
      console.error(e);
    }
  }

  async function health() {
    try {
      const res = await fetch(`${apiUrl}/health`);
      message = await res.text();
    } catch (e) {
      console.error(e);
    }
  }
</script>

<main>
  <h1>{message}</h1>
  <div>
    <button on:click={health}>Ping API</button>
    <button on:click={loadUsers}>Load Users</button>
  </div>
  {#if users.length}
    <ul>
      {#each users as u}
        <li>{u.name} - {u.email}</li>
      {/each}
    </ul>
  {/if}
</main>

<style>
  main { font-family: system-ui, sans-serif; padding: 2rem; }
  button { margin-right: 0.5rem; }
  ul { margin-top: 1rem; }
  body { margin: 0; }
 </style>
