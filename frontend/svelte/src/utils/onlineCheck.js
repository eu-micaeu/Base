// utils/onlineCheck.js
import { tick } from "svelte";

/**
 * Periodically checks if user is online and triggers a callback if not confirmed.
 * @param {function} isLoggedInFn - Function that returns if user is logged in
 * @param {function} onLogout - Function to call for logout
 * @param {function} showPrompt - Function to show prompt (should set a reactive var)
 * @param {function} hidePrompt - Function to hide prompt (should clear prompt var)
 * @param {number} intervalMs - Interval in ms (default: 5min)
 * @param {number} timeoutMs - Timeout for prompt in ms (default: 20s)
 * @returns {function} cleanup - Call to clear intervals/timeouts
 */
export function setupOnlineCheck({
    isLoggedInFn,
    onLogout,
    showPrompt,
    hidePrompt,
    intervalMs = 5 * 60 * 1000,
    timeoutMs = 20000
}) {
    let intervalId, onlineTimeout;

    async function checkOnline() {
        if (!isLoggedInFn()) return;
        showPrompt();
        await tick();
        onlineTimeout = setTimeout(() => {
            hidePrompt();
            onLogout();
        }, timeoutMs);
    }

    intervalId = setInterval(checkOnline, intervalMs);

    function confirmOnline() {
        hidePrompt();
        if (onlineTimeout) clearTimeout(onlineTimeout);
    }

    function cleanup() {
        if (intervalId) clearInterval(intervalId);
        if (onlineTimeout) clearTimeout(onlineTimeout);
    }

    return { confirmOnline, cleanup };
}
