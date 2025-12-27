<script lang="ts">
  import { onMount } from "svelte";
  import { user } from "./lib/stores/user";

  import Login from "./routes/Login.svelte";
  import Signup from "./routes/Signup.svelte";
  import Dashboard from "./routes/Dashboard.svelte";
  import CreateQuiz from "./routes/CreateQuiz.svelte";
  import GameLobby from "./routes/GameLobby.svelte";
  import Game from "./routes/Game.svelte";

  let currentRoute = "#";
  let pinInput = "";

  function hashchange() {
    currentRoute = window.location.hash;
  }

  function enterGame() {
    if (pinInput) {
      window.location.hash = `#play/lobby/${pinInput}`;
    }
  }

  onMount(() => {
    hashchange();
    window.addEventListener("hashchange", hashchange);
    return () => {
      window.removeEventListener("hashchange", hashchange);
    };
  });
</script>

<main>
  <nav>
    <a href="/">Home</a>
    {#if $user}
      <span>Welcome, {$user.name || $user.email}</span>
      <a href="#dashboard">Dashboard</a>
      <button
        on:click={() => {
          user.set(null);
          window.location.hash = "#";
        }}>Logout</button
      >
    {:else}
      <a href="#login">Login</a>
      <a href="#signup">Sign Up</a>
    {/if}
  </nav>

  {#if currentRoute === "#login"}
    <Login />
  {:else if currentRoute === "#signup"}
    <Signup />
  {:else if currentRoute === "#dashboard"}
    <Dashboard />
  {:else if currentRoute === "#create-quiz"}
    <CreateQuiz />
  {:else if currentRoute.startsWith("#host/lobby/")}
    <GameLobby pin={currentRoute.split("/")[2].split("?")[0]} isHost={true} />
  {:else if currentRoute.startsWith("#play/lobby/")}
    <!-- Matches #play/lobby/12345 -->
    <GameLobby
      pin={currentRoute.split("/")[3] || currentRoute.split("/")[2]}
      isHost={false}
    />
  {:else if currentRoute.startsWith("#game/")}
    <Game isHost={$user?.role === "host" || $user?.role === "Host"} />
  {:else}
    <!-- Home -->
    <h1 class="title">Welcome to Kahoot Clone</h1>
    <div class="hero fa-center">
      <input
        type="text"
        placeholder="Game PIN"
        class="pin-input"
        bind:value={pinInput}
      />
      <button class="enter-btn" on:click={enterGame}>Enter</button>
    </div>
  {/if}
</main>

<style>
  nav {
    display: flex;
    gap: 1rem;
    padding: 1rem;
    background: #f4f4f4;
    align-items: center;
  }
  .title {
    text-align: center;
    margin-top: 3rem;
  }
  .hero {
    display: flex;
    justify-content: center;
    gap: 1rem;
    margin-top: 2rem;
  }
  .pin-input {
    padding: 1rem;
    font-size: 1.2rem;
    border: 2px solid #ccc;
    border-radius: 4px;
  }
  .enter-btn {
    padding: 1rem 2rem;
    background: #333;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 1.2rem;
  }
  /* Simple global reset for nice box model */
  :global(body) {
    margin: 0;
    font-family: sans-serif;
  }
  :global(*) {
    box-sizing: border-box;
  }
</style>
