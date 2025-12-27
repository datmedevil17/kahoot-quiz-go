import { writable, get } from 'svelte/store';
import { game } from '../stores/game';

export const ws = writable<WebSocket | null>(null);

export function connect(token: string) {
    if (get(ws)) return; // Already connected

    const socket = new WebSocket(`ws://localhost:8080/api/v1/ws?token=${token}`);

    socket.onopen = () => {
        console.log('Connected to WS');
    };

    socket.onmessage = (msg) => {
        try {
            const { event, data } = JSON.parse(msg.data);
            console.log('WS Event:', event, data);

            game.update(state => {
                switch(event) {
                    case 'player_joined':
                        return { ...state, players: [...state.players, data.username] };
                    case 'NEXT_QUESTION':
                        return { ...state, isStarted: true, currentQuestion: data, leaderboard: null };
                    case 'leaderboard':
                        return { ...state, leaderboard: data };
                    case 'GAME_OVER':
                         // potentially handle game over state
                        return { ...state, currentQuestion: null, leaderboard: data }; // data often leaderboard in game over
                    default:
                        return state;
                }
            });

        } catch (e) {
            console.error('Failed to parse WS message', e);
        }
    };

    socket.onclose = () => {
        console.log('WS Disconnected');
        ws.set(null);
    };

    ws.set(socket);
}

export function sendMessage(event: string, data: any) {
    const socket = get(ws);
    if (socket && socket.readyState === WebSocket.OPEN) {
        socket.send(JSON.stringify({ event, data }));
    } else {
        console.error('WebSocket not connected');
    }
}

export function disconnect() {
    const socket = get(ws);
    if (socket) {
        socket.close();
        ws.set(null);
    }
}
