<script lang="ts">
    import { api } from "../lib/services/api";

    let name = "";
    let email = "";
    let password = "";
    let role = "player";
    let error = "";
    let message = "";

    async function handleSignup() {
        error = "";
        message = "";
        const res = await api("POST", "/auth/signup", {
            name,
            email,
            password,
            role,
        });

        if (res.error) {
            error = res.error;
        } else {
            message = "Signup successful! Please login.";
            setTimeout(() => {
                window.location.hash = "#login";
            }, 1500);
        }
    }
</script>

<div class="auth-container">
    <h2>Sign Up</h2>
    {#if error}
        <p class="error">{error}</p>
    {/if}
    {#if message}
        <p class="success">{message}</p>
    {/if}
    <form on:submit|preventDefault={handleSignup}>
        <div class="form-group">
            <label for="name">Name</label>
            <input type="text" id="name" bind:value={name} required />
        </div>
        <div class="form-group">
            <label for="email">Email</label>
            <input type="email" id="email" bind:value={email} required />
        </div>
        <div class="form-group">
            <label for="password">Password</label>
            <input
                type="password"
                id="password"
                bind:value={password}
                required
            />
        </div>
        <div class="form-group">
            <label for="role">Role</label>
            <select id="role" bind:value={role}>
                <option value="player">Player</option>
                <option value="host">Host</option>
            </select>
        </div>
        <button type="submit">Sign Up</button>
    </form>
    <p>Already have an account? <a href="#login">Login</a></p>
</div>

<style>
    .auth-container {
        max-width: 400px;
        margin: 2rem auto;
        padding: 2rem;
        border: 1px solid #ccc;
        border-radius: 8px;
    }
    .form-group {
        margin-bottom: 1rem;
    }
    label {
        display: block;
        margin-bottom: 0.5rem;
    }
    input,
    select {
        width: 100%;
        padding: 0.5rem;
    }
    .error {
        color: red;
    }
    .success {
        color: green;
    }
</style>
