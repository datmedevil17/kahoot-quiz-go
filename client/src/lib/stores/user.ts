import { writable } from 'svelte/store';

export const user = writable(JSON.parse(localStorage.getItem('user') || 'null'));

user.subscribe((value) => {
	if (value) {
		localStorage.setItem('user', JSON.stringify(value));
	} else {
		localStorage.removeItem('user');
	}
});
