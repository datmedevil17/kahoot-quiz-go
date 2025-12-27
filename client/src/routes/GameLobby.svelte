<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import { ws, connect, disconnect, sendMessage } from "../lib/services/ws";
    import { game } from "../lib/stores/game";
    import { user } from "../lib/stores/user";

    export let pin: string;
    export let isHost = false;

    // Retrieve quizId from URL params if Host
    let quizId = "";

    onMount(() => {
        // Parse QuizID from Hash if host
        if (isHost) {
            const params = new URLSearchParams(
                window.location.hash.split("?")[1],
            );
            quizId = params.get("quizId") || "";
        }

        if ($user && $user.token) {
            connect($user.token);

            // Join Game automatically if not Host?
            // Host creates game via HTTP, but doesn't need to "join" via WS message?
            // Actually, Host MUST join via WS to receive events too?
            // Wait, typical Kahoot: Host screen is a client too.
            // My backend instructions said: "Host: start_game, Players: submit_answer".
            // Host sends `start_game`.

            // Does Host need to send `join_game`?
            // Backend `handleJoinGame` checks `game.GetRoom(pin)`.
            // If Host connects, they are just a connection.
            // If Host wants to be "in" the room to receive broadcasts, they implicitly are?
            // `HandleWS` adds client to Hub.
            // But `room.Clients` only gets populated via `join_game`.
            // So YES, HOST MUST JOIN via `join_game` if we want Host to receive broadcasts.

            // Wait, see `create_game`. It returns PIN.
            // It does NOT auto-join the creator to the WS room.
            // So yes, Host must join.

            // Join Game Message
            setTimeout(() => {
                if ($ws && $ws.readyState === WebSocket.OPEN) {
                    sendMessage("join_game", { game_pin: pin });
                } else {
                    // Retry or wait for open
                    const interval = setInterval(() => {
                        if ($ws && $ws.readyState === WebSocket.OPEN) {
                            sendMessage("join_game", { game_pin: pin });
                            clearInterval(interval);
                        }
                    }, 500);
                }
            }, 500);
        }
    });

    onDestroy(() => {
        // Don't disconnect here, we persist connection to Game
    });

    function startGame() {
        if (!isHost) return;
        sendMessage("start_game", { quiz_id: quizId });
        // Redirect to Game View logic will be handled by store update listening to 'NEXT_QUESTION'
        // But we can also set local state?
        // Better to rely on store `isStarted`
    }

    $: if ($game.isStarted) {
        // Create hash like #game/12345
        window.location.hash = `#game/${pin}`;
    }
</script>

<div class="lobby">
    <h1>Game PIN: <span class="pin">{pin}</span></h1>

    <div class="players-list">
        <h2>Players Joined ({$game.players.length})</h2>
        <div class="grid">
            {#each $game.players as p}
                <div class="player-tag">{p}</div>
            {/each}
        </div>
    </div>

    {#if isHost}
        <div class="actions">
            <button
                class="btn start"
                on:click={startGame}
                disabled={$game.players.length === 0}
            >
                Start Game
            </button>
            <p>Waiting for players...</p>
        </div>
    {:else}
        <div class="waiting-msg">
            <h3>You're in!</h3>
            <p>See your nickname on the screen?</p>
            <div class="loader">Waiting for host to start...</div>
        </div>
    {/if}
</div>

<style>
    .lobby {
        text-align: center;
        padding: 2rem;
        background: #333;
        color: white;
        min-height: 100vh;
    }
    .pin {
        font-size: 4rem;
        font-weight: bold;
        background: white;
        color: #333;
        padding: 0.5rem 2rem;
        border-radius: 8px;
    }
    .grid {
        display: flex;
        flex-wrap: wrap;
        gap: 1rem;
        justify-content: center;
        margin-top: 2rem;
    }
    .player-tag {
        background: rgba(255, 255, 255, 0.2);
        padding: 0.5rem 1rem;
        border-radius: 4px;
        font-size: 1.2rem;
    }
    .actions {
        margin-top: 3rem;
    }
    .btn.start {
        background: #46178f;
        color: white;
        border: none;
        padding: 1rem 3rem;
        font-size: 1.5rem;
        border-radius: 4px;
        cursor: pointer;
    }
    .btn.start:disabled {
        background: #666;
        cursor: not-allowed;
    }
</style>
