<script>
    // Footer animado e moderno
    const year = new Date().getFullYear();
</script>

<footer>
    <div class="footer-content">
        <span>© {year} Base App</span>
        <span class="made"
            >Feito com <span class="heart">❤</span> por micael</span
        >
    </div>
</footer>

<style>
    footer {
        width: 100%;
        padding: 1.2rem 0;
        background: linear-gradient(90deg, #2575fc 0%, #6a11cb 100%);
        color: #fff;
        text-align: center;
        position: fixed;
        left: 0;
        bottom: 0;
        z-index: 100;
        box-shadow: 0 -2px 16px rgba(80, 80, 160, 0.08);
    }
    .footer-content {
        display: flex;
        justify-content: center;
        align-items: center;
        gap: 2rem;
        font-size: 1rem;
    }
    .made {
        font-size: 0.95rem;
        opacity: 0.85;
        display: flex;
        align-items: center;
        gap: 0.3rem;
    }
    .heart {
        color: #ff4b4b;
        animation: pulse 1.2s infinite;
    }
    @keyframes pulse {
        0%,
        100% {
            transform: scale(1);
        }
        50% {
            transform: scale(1.2);
        }
    }
</style>
