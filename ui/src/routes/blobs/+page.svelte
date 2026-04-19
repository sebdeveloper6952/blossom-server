<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { session } from '$lib/auth';
	import { listBlobs, listAllBlobs, deleteBlob, type BlobDescriptor } from '$lib/api';
	import { formatBytes, truncate, timeAgo, isImage } from '$lib/format';

	let blobs = $state<BlobDescriptor[]>([]);
	let loading = $state(true);
	let err = $state<string | null>(null);
	let query = $state('');
	let mimeFilter = $state('');
	let pendingDelete = $state<string | null>(null);

	let sess = $state<typeof $session>(null);
	session.subscribe((s) => (sess = s));

	const filtered = $derived.by(() => {
		const q = query.trim().toLowerCase();
		return blobs.filter((b) => {
			if (mimeFilter && !b.type.startsWith(mimeFilter)) return false;
			if (q) {
				const hay = `${b.sha256} ${b.pubkey ?? ''}`.toLowerCase();
				if (!hay.includes(q)) return false;
			}
			return true;
		});
	});

	const mimeOptions = $derived(Array.from(new Set(blobs.map((b) => b.type.split('/')[0]))).sort());

	onMount(async () => {
		if (!sess) {
			goto(`${base}/`);
			return;
		}
		try {
			blobs = sess.isAdmin
				? await listAllBlobs(sess.signer)
				: await listBlobs(sess.pubkey);
		} catch (e: any) {
			err = e.message;
		} finally {
			loading = false;
		}
	});

	async function confirmDelete(sha: string) {
		if (!sess) return;
		pendingDelete = null;
		try {
			await deleteBlob(sess.signer, sha);
			blobs = blobs.filter((b) => b.sha256 !== sha);
		} catch (e: any) {
			err = e.message;
		}
	}

	async function copy(text: string) {
		try {
			await navigator.clipboard.writeText(text);
		} catch {}
	}
</script>

<div class="flex flex-wrap items-end justify-between gap-4 mb-4">
	<div>
		<h1 class="text-xl font-semibold">
			{sess?.isAdmin ? 'All blobs' : 'Your blobs'}
		</h1>
		<p class="text-sm text-neutral-400">
			{filtered.length} of {blobs.length}
			{blobs.length === 1 ? 'blob' : 'blobs'}
		</p>
	</div>
	<div class="flex gap-2">
		<input
			bind:value={query}
			placeholder="Search hash or pubkey…"
			class="rounded border border-neutral-700 bg-neutral-900 px-3 py-1.5 text-sm font-mono w-64 focus:outline-none focus:border-indigo-400"
		/>
		<select
			bind:value={mimeFilter}
			class="rounded border border-neutral-700 bg-neutral-900 px-2 py-1.5 text-sm focus:outline-none focus:border-indigo-400"
		>
			<option value="">All types</option>
			{#each mimeOptions as m}
				<option value={m}>{m}</option>
			{/each}
		</select>
	</div>
</div>

{#if err}
	<div class="mb-4 rounded border border-red-500/40 bg-red-500/10 px-3 py-2 text-sm text-red-300">
		{err}
	</div>
{/if}

{#if loading}
	<div class="py-16 text-center text-neutral-500">Loading…</div>
{:else if filtered.length === 0}
	<div class="py-16 text-center text-neutral-500">No blobs.</div>
{:else}
	<div class="overflow-x-auto rounded border border-neutral-800">
		<table class="w-full text-sm">
			<thead class="bg-neutral-900 text-xs uppercase text-neutral-400">
				<tr>
					<th class="px-3 py-2 text-left">Preview</th>
					<th class="px-3 py-2 text-left">Hash</th>
					<th class="px-3 py-2 text-left">Size</th>
					<th class="px-3 py-2 text-left">Type</th>
					{#if sess?.isAdmin}
						<th class="px-3 py-2 text-left">Uploader</th>
					{/if}
					<th class="px-3 py-2 text-left">Uploaded</th>
					<th class="px-3 py-2 text-right">Actions</th>
				</tr>
			</thead>
			<tbody class="divide-y divide-neutral-800">
				{#each filtered as b (b.sha256)}
					<tr class="hover:bg-neutral-900/50">
						<td class="px-3 py-2">
							{#if isImage(b.type)}
								<a href={b.url} target="_blank" rel="noopener">
									<img src={b.url} alt="" class="h-10 w-10 rounded object-cover" loading="lazy" />
								</a>
							{:else}
								<div
									class="h-10 w-10 rounded bg-neutral-800 flex items-center justify-center text-xs text-neutral-500"
								>
									{b.type.split('/')[1]?.slice(0, 3) ?? '?'}
								</div>
							{/if}
						</td>
						<td class="px-3 py-2">
							<button
								onclick={() => copy(b.sha256)}
								title="Copy hash"
								class="font-mono text-neutral-300 hover:text-indigo-400"
							>
								{truncate(b.sha256)}
							</button>
						</td>
						<td class="px-3 py-2 text-neutral-300">{formatBytes(b.size)}</td>
						<td class="px-3 py-2 text-neutral-400">{b.type}</td>
						{#if sess?.isAdmin}
							<td class="px-3 py-2">
								<button
									onclick={() => b.pubkey && copy(b.pubkey)}
									title="Copy pubkey"
									class="font-mono text-neutral-300 hover:text-indigo-400"
								>
									{b.pubkey ? truncate(b.pubkey) : '—'}
								</button>
							</td>
						{/if}
						<td class="px-3 py-2 text-neutral-400">{timeAgo(b.uploaded)}</td>
						<td class="px-3 py-2 text-right">
							<div class="inline-flex gap-1">
								<a
									href={b.url}
									target="_blank"
									rel="noopener"
									class="rounded border border-neutral-700 px-2 py-1 text-xs hover:bg-neutral-800"
								>
									Open
								</a>
								<button
									onclick={() => (pendingDelete = b.sha256)}
									class="rounded border border-red-500/40 text-red-300 px-2 py-1 text-xs hover:bg-red-500/10"
								>
									Delete
								</button>
							</div>
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
{/if}

{#if pendingDelete}
	<div
		class="fixed inset-0 bg-black/60 flex items-center justify-center p-4"
		role="dialog"
		aria-modal="true"
	>
		<div class="w-full max-w-md rounded border border-neutral-700 bg-neutral-900 p-5">
			<h2 class="text-lg font-semibold">Delete blob?</h2>
			<p class="mt-2 text-sm text-neutral-400">
				This signs a kind-24242 delete event and removes the blob from the server.
			</p>
			<p class="mt-2 font-mono text-xs text-neutral-500 break-all">{pendingDelete}</p>
			<div class="mt-5 flex justify-end gap-2">
				<button
					onclick={() => (pendingDelete = null)}
					class="rounded border border-neutral-700 px-3 py-1.5 text-sm hover:bg-neutral-800"
				>
					Cancel
				</button>
				<button
					onclick={() => confirmDelete(pendingDelete!)}
					class="rounded bg-red-500 hover:bg-red-400 px-3 py-1.5 text-sm font-medium"
				>
					Delete
				</button>
			</div>
		</div>
	</div>
{/if}
