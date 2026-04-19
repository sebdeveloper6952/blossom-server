import { finalizeEvent, getPublicKey, nip19, type EventTemplate, type Event } from 'nostr-tools';

export type Verb = 'list' | 'delete' | 'upload';

export type Signer = (template: EventTemplate) => Promise<Event>;

export type Session = {
	pubkey: string;
	signer: Signer;
};

export async function loginWithNip07(): Promise<Session> {
	if (typeof window === 'undefined' || !window.nostr) {
		throw new Error('No NIP-07 extension detected');
	}
	const nostr = window.nostr;
	const pubkey = await nostr.getPublicKey();
	const signer: Signer = async (template) => nostr.signEvent(template) as Promise<Event>;
	return { pubkey, signer };
}

export function loginWithNsec(nsec: string): Session {
	const trimmed = nsec.trim();
	const decoded = nip19.decode(trimmed);
	if (decoded.type !== 'nsec') throw new Error('Not an nsec');
	const sk = decoded.data as Uint8Array;
	const pubkey = getPublicKey(sk);
	const signer: Signer = async (template) => finalizeEvent(template, sk);
	return { pubkey, signer };
}

export function buildAuthTemplate(verb: Verb, hash?: string): EventTemplate {
	const tags: string[][] = [
		['t', verb],
		['expiration', String(Math.floor(Date.now() / 1000) + 60)]
	];
	if (hash) tags.push(['x', hash]);
	return {
		kind: 24242,
		created_at: Math.floor(Date.now() / 1000),
		tags,
		content: `admin ui ${verb}`
	};
}

export async function authHeader(signer: Signer, verb: Verb, hash?: string): Promise<string> {
	const signed = await signer(buildAuthTemplate(verb, hash));
	const b64 = btoa(JSON.stringify(signed));
	return `Nostr ${b64}`;
}
