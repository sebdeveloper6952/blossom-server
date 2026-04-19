declare global {
	interface Window {
		nostr?: {
			getPublicKey(): Promise<string>;
			signEvent(event: {
				kind: number;
				created_at: number;
				tags: string[][];
				content: string;
			}): Promise<{
				id: string;
				pubkey: string;
				kind: number;
				created_at: number;
				tags: string[][];
				content: string;
				sig: string;
			}>;
		};
	}
	namespace App {}
}

export {};
