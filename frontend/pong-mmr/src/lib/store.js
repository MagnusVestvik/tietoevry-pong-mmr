import { writable } from 'svelte/store';

export const opponent = writable(null);

export const submitedMatch = writable(false);
