<script lang="ts">
    import { api } from "../lib/services/api";
    import { user } from "../lib/stores/user";

    let title = "";
    let description = "";

    // Questions State
    let questions: any[] = [
        { text: "", options: ["", "", "", ""], answer: 0, time_limit: 30 },
    ];

    let error = "";
    let success = "";

    function addQuestion() {
        questions = [
            ...questions,
            { text: "", options: ["", "", "", ""], answer: 0, time_limit: 30 },
        ];
    }

    async function handleCreate() {
        error = "";
        success = "";

        if (!title) {
            error = "Title is required";
            return;
        }

        // 1. Create Quiz
        const quizRes = await api(
            "POST",
            "/api/v1/quizzes",
            { title, description },
            $user.token,
        );

        if (quizRes.error) {
            error = quizRes.error;
            return;
        }

        const quizId = quizRes.data?.ID || quizRes.ID; // Handle different potential response shapes

        // 2. Add Questions
        for (const [index, q] of questions.entries()) {
            if (!q.text) continue; // skip empty

            const qPayload = {
                text: q.text,
                options: q.options, // array of strings
                answer: parseInt(q.answer),
                time_limit: parseInt(q.time_limit),
            };

            const qRes = await api(
                "POST",
                `/api/v1/quizzes/${quizId}/questions`,
                qPayload,
                $user.token,
            );
            if (qRes.error) {
                console.error(
                    `Failed to save question ${index + 1}`,
                    qRes.error,
                );
            }
        }

        success = "Quiz created successfully!";
        setTimeout(() => {
            window.location.hash = "#dashboard";
        }, 1500);
    }
</script>

<div class="create-container">
    <h2>Create New Quiz</h2>

    {#if error}
        <p class="error">{error}</p>
    {/if}
    {#if success}
        <p class="success">{success}</p>
    {/if}

    <div class="section">
        <label>
            Quiz Title
            <input
                type="text"
                bind:value={title}
                placeholder="e.g. Science Trivia"
            />
        </label>

        <label>
            Description
            <input
                type="text"
                bind:value={description}
                placeholder="Short description..."
            />
        </label>
    </div>

    <hr />

    <h3>Questions</h3>
    {#each questions as q, i}
        <div class="question-card">
            <h4>Question {i + 1}</h4>
            <input
                type="text"
                bind:value={q.text}
                placeholder="Question Text"
                class="q-input"
            />

            <div class="options-grid">
                {#each q.options as opt, optIndex}
                    <input
                        type="text"
                        bind:value={q.options[optIndex]}
                        placeholder="Option {optIndex + 1}"
                    />
                {/each}
            </div>

            <div class="meta-row">
                <label
                    >Correct Answer (0-3): <input
                        type="number"
                        min="0"
                        max="3"
                        bind:value={q.answer}
                    /></label
                >
                <label
                    >Time (sec): <input
                        type="number"
                        bind:value={q.time_limit}
                    /></label
                >
            </div>
        </div>
    {/each}

    <button type="button" on:click={addQuestion} class="btn secondary"
        >+ Add Question</button
    >

    <div class="actions">
        <button on:click={handleCreate} class="btn primary"
            >Save & Finish</button
        >
        <a href="#dashboard" class="btn cancel">Cancel</a>
    </div>
</div>

<style>
    .create-container {
        max-width: 800px;
        margin: 2rem auto;
        padding: 2rem;
    }
    .section,
    .question-card {
        background: white;
        padding: 1.5rem;
        margin-bottom: 1.5rem;
        border: 1px solid #ddd;
        border-radius: 8px;
    }
    input {
        display: block;
        width: 100%;
        padding: 0.5rem;
        margin-bottom: 1rem;
        box-sizing: border-box;
    }
    .options-grid {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 1rem;
        margin-bottom: 1rem;
    }
    .meta-row {
        display: flex;
        gap: 2rem;
    }
    .meta-row label {
        display: flex;
        align-items: center;
        gap: 0.5rem;
    }
    .meta-row input {
        width: 60px;
        margin-bottom: 0;
    }
    .btn {
        padding: 0.75rem 1.5rem;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        font-size: 1rem;
    }
    .primary {
        background: #333;
        color: white;
    }
    .secondary {
        background: #eee;
        color: #333;
    }
    .cancel {
        text-decoration: none;
        color: red;
        margin-left: 1rem;
    }
    .error {
        color: red;
    }
    .success {
        color: green;
    }
</style>
