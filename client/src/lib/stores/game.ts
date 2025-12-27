import { writable } from 'svelte/store';

export interface Question {
    id: string;
    text: string;
    options: string[];
    time_limit: number;
}

export interface GameState {
    gamePin: string | null;
    isStarted: boolean;
    currentQuestion: Question | null;
    leaderboard: any[] | null;
    players: string[]; // for lobby
    score: number; // for player
}

const initialState: GameState = {
    gamePin: null,
    isStarted: false,
    currentQuestion: null,
    leaderboard: null,
    players: [],
    score: 0
};

export const game = writable<GameState>(initialState);

export const resetGame = () => game.set(initialState);
