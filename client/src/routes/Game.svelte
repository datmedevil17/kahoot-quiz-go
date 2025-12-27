<script lang="ts">
    import { onMount } from "svelte";
    import { game } from "../lib/stores/game";
    import { user } from "../lib/stores/user";
    import { sendMessage } from "../lib/services/ws";

    export let isHost = false;

    // Colors for options
    const colors = ["#e21b3c", "#1368ce", "#d89e00", "#26890c"]; // Red, Blue, Yellow, Green
    const shapes = ["▲", "◆", "●", "■"];

    function sendAnswer(index: number) {
        if (!$game.currentQuestion) return;
        sendMessage("submit_answer", {
            question_id: $game.currentQuestion.id,
            option: index,
        });
    }
</script>

<div class="game-container">
    {#if $game.leaderboard}
        <!-- LEADERBOARD VIEW -->
        <div class="leaderboard">
            <h1>Scoreboard</h1>
            {#each $game.leaderboard.slice(0, 5) as entry, i}
                <div class="entry">
                    <span class="rank">{i + 1}</span>
                    <span class="name">{entry.username}</span>
                    <span class="score">{entry.score}</span>
                </div>
            {/each}

            <p>Next question coming up...</p>
        </div>
    {:else if $game.currentQuestion}
        <!-- QUESTION VIEW -->
        {#if isHost}
            <!-- HOST VIEW (Show Question) -->
            <div class="host-view">
                <div
                    class="timer-bar"
                    style="animation-duration: {$game.currentQuestion
                        .time_limit}s"
                ></div>
                <h2 class="question-text">{$game.currentQuestion.text}</h2>

                <div class="options-grid">
                    {#each $game.currentQuestion.options as opt, i}
                        <div
                            class="option-card"
                            style="background-color: {colors[i]}"
                        >
                            <span class="shape">{shapes[i]}</span>
                            <span class="text">{opt}</span>
                        </div>
                    {/each}
                </div>
            </div>
        {:else}
            <!-- PLAYER VIEW (Show Buttons) -->
            <div class="player-view">
                <div class="buttons-grid">
                    {#each $game.currentQuestion.options as _, i}
                        <button
                            class="answer-btn"
                            style="background-color: {colors[i]}"
                            on:click={() => sendAnswer(i)}
                        >
                            <span class="shape">{shapes[i]}</span>
                        </button>
                    {/each}
                </div>
            </div>
        {/if}
    {:else}
        <!-- LOADING / GAME OVER -->
        <div class="loading">
            <h2>Waiting...</h2>
        </div>
    {/if}
</div>

<style>
    .game-container {
        min-height: 100vh;
        background: #f2f2f2;
        padding: 1rem;
    }

    /* Leaderboard */
    .leaderboard {
        text-align: center;
        max-width: 600px;
        margin: 2rem auto;
    }
    .entry {
        display: flex;
        justify-content: space-between;
        background: white;
        padding: 1rem;
        margin-bottom: 0.5rem;
        border-radius: 4px;
        font-size: 1.2rem;
    }

    /* Host View */
    .host-view {
        display: flex;
        flex-direction: column;
        height: 90vh;
    }
    .question-text {
        text-align: center;
        font-size: 3rem;
        margin: 2rem 0;
        flex-grow: 1;
    }
    .options-grid {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 1rem;
        height: 40vh;
    }
    .option-card {
        display: flex;
        align-items: center;
        padding: 2rem;
        color: white;
        font-size: 1.5rem;
        border-radius: 4px;
        font-weight: bold;
    }
    .option-card .shape {
        margin-right: 1rem;
    }

    /* Player View */
    .player-view {
        height: 90vh;
        display: flex;
        align-items: center;
    }
    .buttons-grid {
        display: grid;
        grid-template-columns: 1fr 1fr;
        grid-template-rows: 1fr 1fr;
        gap: 1rem;
        width: 100%;
        height: 100%;
    }
    .answer-btn {
        border: none;
        border-radius: 8px;
        font-size: 4rem;
        color: white;
        cursor: pointer;
    }
</style>
