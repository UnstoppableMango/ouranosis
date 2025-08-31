const key = 'session-id';

export function set(id: string): void {
	localStorage.setItem(key, id);
}

export function get(): string | null {
	return localStorage.getItem(key);
}

export function clear(): void {
	localStorage.removeItem(key);
}
