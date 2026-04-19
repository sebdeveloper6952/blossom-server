import { authHeader, type Signer } from './nostr';

export type BlobDescriptor = {
	url: string;
	sha256: string;
	size: number;
	type: string;
	uploaded: number;
	pubkey?: string;
};

export type Me = {
	pubkey: string;
	is_admin: boolean;
};

async function jsonOrThrow<T>(res: Response): Promise<T> {
	if (!res.ok) {
		const text = await res.text().catch(() => '');
		throw new Error(text || `HTTP ${res.status}`);
	}
	return res.json() as Promise<T>;
}

export async function getMe(signer: Signer): Promise<Me> {
	const res = await fetch('/api/me', {
		headers: { Authorization: await authHeader(signer, 'list') }
	});
	return jsonOrThrow<Me>(res);
}

export async function listBlobs(
	pubkey: string,
	opts: { since?: number; until?: number } = {}
): Promise<BlobDescriptor[]> {
	const params = new URLSearchParams();
	if (opts.since) params.set('since', String(opts.since));
	if (opts.until) params.set('until', String(opts.until));
	const qs = params.toString();
	const res = await fetch(`/list/${pubkey}${qs ? `?${qs}` : ''}`);
	return jsonOrThrow<BlobDescriptor[]>(res);
}

export async function listAllBlobs(
	signer: Signer,
	opts: { since?: number; until?: number } = {}
): Promise<BlobDescriptor[]> {
	const params = new URLSearchParams();
	if (opts.since) params.set('since', String(opts.since));
	if (opts.until) params.set('until', String(opts.until));
	const qs = params.toString();
	const res = await fetch(`/api/admin/blobs${qs ? `?${qs}` : ''}`, {
		headers: { Authorization: await authHeader(signer, 'list') }
	});
	return jsonOrThrow<BlobDescriptor[]>(res);
}

export async function deleteBlob(signer: Signer, sha256: string): Promise<void> {
	const res = await fetch(`/${sha256}`, {
		method: 'DELETE',
		headers: { Authorization: await authHeader(signer, 'delete', sha256) }
	});
	if (!res.ok) {
		const text = await res.text().catch(() => '');
		throw new Error(text || `HTTP ${res.status}`);
	}
}
