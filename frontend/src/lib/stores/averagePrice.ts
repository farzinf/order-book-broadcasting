import { writable } from 'svelte/store';

export const averagePrice = writable<number | null>(null);
