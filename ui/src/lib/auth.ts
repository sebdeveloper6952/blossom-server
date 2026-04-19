import { writable } from 'svelte/store';
import type { Signer } from './nostr';

export type SessionState = {
	pubkey: string;
	isAdmin: boolean;
	signer: Signer;
} | null;

export const session = writable<SessionState>(null);

export function logout() {
	session.set(null);
}
