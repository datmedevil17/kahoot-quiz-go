<script lang="ts">
  import { api } from '../lib/services/api';
  import { user } from '../lib/stores/user';
  
  let email = '';
  let password = '';
  let error = '';

  async function handleLogin() {
    error = '';
    const res = await api('POST', '/auth/login', { email, password });
    
    if (res.error) {
        error = res.error;
    } else {
        // res is the token string directly based on backend
        // Wait, backend returns "c.String(http.StatusOK, token)"
        // so res is likely just the string if parsed as text, but api.ts parses JSON.
        // Let's check api.ts implementation again.
        // api.ts tries JSON.parse, if fails returns { data: text }.
        
        const token = res.data || res; // handle if it came back as object or raw string
        
        // We need to fetch user details or just store token?
        // Requirements say "Save JWT".
        // Let's fetch user details too so we know who they are.
        
        const userRes = await api('GET', '/api/v1/users/me', null, token as string);
        if (userRes.error) {
            error = 'Failed to fetch profile';
        } else {
            user.set({ ...userRes, token });
            window.location.hash = '#dashboard';
        }
    }
  }
</script>

<div class="auth-container">
    <h2>Login</h2>
    {#if error}
        <p class="error">{error}</p>
    {/if}
    <form on:submit|preventDefault={handleLogin}>
        <div class="form-group">
            <label for="email">Email</label>
            <input type="email" id="email" bind:value={email} required />
        </div>
        <div class="form-group">
            <label for="password">Password</label>
            <input type="password" id="password" bind:value={password} required />
        </div>
        <button type="submit">Login</button>
    </form>
    <p>Don't have an account? <a href="#signup">Sign up</a></p>
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
    input {
        width: 100%;
        padding: 0.5rem;
    }
    .error {
        color: red;
    }
</style>
