<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { onMount } from 'svelte';
	import { loginWithNip07, loginWithNsec } from '$lib/nostr';
	import { getMe } from '$lib/api';
	import { session } from '$lib/auth';

	let nsec = $state('');
	let reveal = $state(false);
	let err = $state<string | null>(null);
	let busy = $state(false);
	let hasExtension = $state(false);

	onMount(() => {
		hasExtension = typeof window !== 'undefined' && !!window.nostr;
		const unsub = session.subscribe((s) => {
			if (s) goto(`${base}/blobs`);
		});
		return unsub;
	});

	async function finish(s: { pubkey: string; signer: any }) {
		try {
			const me = await getMe(s.signer);
			session.set({ pubkey: me.pubkey, isAdmin: me.is_admin, signer: s.signer });
			goto(`${base}/blobs`);
		} catch (e: any) {
			err = `Auth succeeded but /me failed: ${e.message}`;
		}
	}

	async function signInExt() {
		err = null;
		busy = true;
		try {
			await finish(await loginWithNip07());
		} catch (e: any) {
			err = e.message;
		} finally {
			busy = false;
		}
	}

	async function signInNsec(ev: SubmitEvent) {
		ev.preventDefault();
		err = null;
		busy = true;
		try {
			await finish(loginWithNsec(nsec));
			nsec = '';
		} catch (e: any) {
			err = e.message;
		} finally {
			busy = false;
		}
	}
</script>

<div class="mx-auto max-w-md pt-16">
	<h1 class="text-2xl font-semibold mb-1">Sign in</h1>
	<p class="text-sm text-neutral-400 mb-8">Use your nostr identity to manage blobs.</p>

	<div class="space-y-6">
		<div>
			<button
				onclick={signInExt}
				disabled={busy || !hasExtension}
				class="w-full rounded bg-indigo-500 hover:bg-indigo-400 disabled:opacity-40 disabled:cursor-not-allowed px-4 py-2.5 font-medium transition"
			>
				{hasExtension ? 'Login with Extension (NIP-07)' : 'No NIP-07 extension detected'}
			</button>
		</div>

		<div class="flex items-center gap-3 text-xs text-neutral-500">
			<div class="h-px bg-neutral-800 flex-1"></div>
			<span>or</span>
			<div class="h-px bg-neutral-800 flex-1"></div>
		</div>

		<form onsubmit={signInNsec} class="space-y-3">
			<label class="block text-sm">
				<span class="text-neutral-300">Login with nsec</span>
				<p class="text-xs text-amber-400/80 mt-1">
					Your private key stays in memory only. Reload = relogin.
				</p>
				<div class="mt-2 flex gap-2">
					<input
						type={reveal ? 'text' : 'password'}
						bind:value={nsec}
						placeholder="nsec1…"
						autocomplete="off"
						spellcheck="false"
						class="flex-1 rounded border border-neutral-700 bg-neutral-900 px-3 py-2 font-mono text-sm focus:outline-none focus:border-indigo-400"
					/>
					<button
						type="button"
						onclick={() => (reveal = !reveal)}
						class="rounded border border-neutral-700 px-3 text-xs text-neutral-300 hover:bg-neutral-800"
					>
						{reveal ? 'Hide' : 'Show'}
					</button>
				</div>
			</label>
			<button
				type="submit"
				disabled={busy || !nsec.trim()}
				class="w-full rounded border border-neutral-700 hover:bg-neutral-800 disabled:opacity-40 disabled:cursor-not-allowed px-4 py-2 text-sm transition"
			>
				Sign in with nsec
			</button>
		</form>

		{#if err}
			<div class="rounded border border-red-500/40 bg-red-500/10 px-3 py-2 text-sm text-red-300">
				{err}
			</div>
		{/if}
	</div>
</div>
