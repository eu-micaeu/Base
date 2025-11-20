const baseUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080';

async function request(path, options = {}) {
    const url = `${baseUrl}${path}`;
    try {
        const res = await fetch(url, options);
        if (!res.ok) {
            const text = await res.text();
            throw new Error(`HTTP ${res.status} - ${text}`);
        }
        const contentType = res.headers.get('content-type') || '';
        if (contentType.includes('application/json')) {
            return await res.json();
        }
        return await res.text();
    } catch (e) {
        throw e;
    }
}

export async function getHealth() {
    return request('/health');
}

export async function getUsers() {
    return request('/users');
}

export async function registerUser(payload) {
    return request('/auth/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload)
    });
}

export async function loginUser(payload) {
    return request('/auth/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload)
    });
}

export { baseUrl };
