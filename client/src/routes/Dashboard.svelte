<script lang="ts">
    import { onMount } from "svelte";
    import { api } from "../lib/services/api";
    import { user } from "../lib/stores/user";

    let quizzes: any[] = [];
    let error = "";

    onMount(async () => {
        await fetchQuizzes();
    });

    async function fetchQuizzes() {
        const res = await api("GET", "/api/v1/quizzes", null, $user.token);
        if (res.error) {
            error = res.error;
        } else {
            quizzes = res; // Assuming backend returns array directly or {data: []}?
            // Checked handler: `c.JSON(http.StatusOK, quizzes)` -> returns array directly.
        }
    }

    async function startGame(quizId: string) {
        // 1. Create Game Session
        // We don't need to send body (backend generates PIN)
        // Wait, backend logic: `func (h *Handler) CreateGame` -> generates PIN.
        // AND it needs to associate the game with a quiz?
        // Let's re-read `server/internal/handlers/game/handler.go`
        // It creates a room with `QuizID: ""` initially?
        // Ah, `start_game` WS event sets the quiz ID.
        // So `CreateGame` just reserves a PIN.

        const res = await api("POST", "/api/v1/games", {}, $user.token);

        if (res.error) {
            alert("Failed to start game: " + res.error);
            return;
        }

        const pin = res.game_pin;

        // 2. Redirect to Host Lobby
        // We pass the QuizID via URL hash param so the Lobby knows which quiz to "Start" later
        window.location.hash = `#host/lobby/${pin}?quizId=${quizId}`;
    }
</script>

<div class="dashboard">
    <header>
        <h2>My Quizzes</h2>
        <a href="#create-quiz" class="btn primary">Create New Quiz</a>
    </header>

    {#if error}
        <p class="error">{error}</p>
    {/if}

    <div class="quiz-grid">
        {#each quizzes as quiz}
            <div class="quiz-card">
                <h3>{quiz.title}</h3>
                <p>{quiz.description}</p>
                <div class="actions">
                    <button class="btn play" on:click={() => startGame(quiz.ID)}
                        >Play</button
                    >
                    <!-- Add Edit/Delete later -->
                </div>
            </div>
        {/each}

        {#if quizzes.length === 0}
            <p>No quizzes found. Create one to get started!</p>
        {/if}
    </div>
</div>

<style>
    .dashboard {
        padding: 2rem;
        max-width: 1200px;
        margin: 0 auto;
    }
    header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 2rem;
    }
    .quiz-grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
        gap: 1.5rem;
    }
    .quiz-card {
        background: white;
        padding: 1.5rem;
        border-radius: 8px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }
    .btn {
        padding: 0.5rem 1rem;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        text-decoration: none;
        display: inline-block;
    }
    .btn.primary {
        background: #333;
        color: white;
    }
    .btn.play {
        background: #46178f;
        color: white;
        width: 100%;
        margin-top: 1rem;
    }
</style>
