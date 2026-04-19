export function formatBytes(n: number): string {
	if (n < 1024) return `${n} B`;
	const units = ['KB', 'MB', 'GB', 'TB'];
	let v = n / 1024;
	let i = 0;
	while (v >= 1024 && i < units.length - 1) {
		v /= 1024;
		i++;
	}
	return `${v.toFixed(v >= 10 ? 0 : 1)} ${units[i]}`;
}

export function truncate(s: string, head = 8, tail = 6): string {
	if (s.length <= head + tail + 1) return s;
	return `${s.slice(0, head)}…${s.slice(-tail)}`;
}

export function timeAgo(unix: number): string {
	const diff = Date.now() / 1000 - unix;
	if (diff < 60) return 'just now';
	if (diff < 3600) return `${Math.floor(diff / 60)}m ago`;
	if (diff < 86400) return `${Math.floor(diff / 3600)}h ago`;
	if (diff < 2592000) return `${Math.floor(diff / 86400)}d ago`;
	return new Date(unix * 1000).toISOString().slice(0, 10);
}

export function isImage(mime: string): boolean {
	return mime.startsWith('image/');
}
