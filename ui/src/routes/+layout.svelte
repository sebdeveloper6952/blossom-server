<script lang="ts">
	import '../app.css';
	import { session, logout } from '$lib/auth';
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import { truncate } from '$lib/format';

	let { children } = $props();

	function signOut() {
		logout();
		goto(`${base}/`);
	}
</script>

<div class="min-h-screen flex flex-col">
	<header class="border-b border-neutral-800 bg-neutral-900/50 backdrop-blur">
		<div class="mx-auto max-w-7xl px-4 py-3 flex items-center justify-between">
			<a href="{base}/" class="font-semibold tracking-tight">🌸 Blossom Admin</a>
			{#if $session}
				<div class="flex items-center gap-3 text-sm">
					{#if $session.isAdmin}
						<span class="px-2 py-0.5 rounded bg-indigo-500/20 text-indigo-300 text-xs">admin</span>
					{/if}
					<span class="font-mono text-neutral-400">{truncate($session.pubkey, 8, 6)}</span>
					<button
						onclick={signOut}
						class="rounded border border-neutral-700 px-2 py-1 text-neutral-300 hover:bg-neutral-800"
					>
						Sign out
					</button>
				</div>
			{/if}
		</div>
	</header>
	<main class="flex-1 mx-auto w-full max-w-7xl px-4 py-6">
		{@render children()}
	</main>
</div>
