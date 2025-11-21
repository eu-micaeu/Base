<script>
    // Navegação por hash links
    // Logo simples
    const logo = "⚡";
</script>

<nav>
    <div class="logo">{logo}</div>
    <a href="#/">Home</a>
    <a href="#/sobre">Sobre</a>
</nav>

<style>
    nav {
        display: flex;
        align-items: center;
        gap: 1.5rem;
        margin-bottom: 2rem;
        padding: 0.75rem 2rem;
        background: linear-gradient(270deg, #6a11cb, #2575fc, #6a11cb);
        background-size: 400% 400%;
        box-shadow: 0 2px 16px rgba(80, 80, 160, 0.08);
    }

    .logo {
        font-size: 2rem;
        font-weight: bold;
        margin-right: 1.5rem;
        color: #fff;
        text-shadow: 0 2px 8px #0002;
        letter-spacing: 2px;
        user-select: none;
        transition: transform 0.3s;
    }
    .logo:hover {
        transform: scale(1.1) rotate(-8deg);
    }
    a {
        position: relative;
        background: transparent;
        border: none;
        padding: 0.5rem 1.2rem;
        border-radius: 8px;
        cursor: pointer;
        text-decoration: none;
        color: #fff;
        font-weight: 500;
        font-size: 1.1rem;
        transition:
            background 0.3s,
            color 0.3s,
            box-shadow 0.3s;
        overflow: hidden;
    }
    a:hover {
        background: rgba(255, 255, 255, 0.12);
        color: #ffe;
        box-shadow: 0 2px 8px #0002;
    }
    a:hover::after {
        transform: scaleX(1);
    }
</style>
