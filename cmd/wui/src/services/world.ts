export async function create(playerId: string): Promise<void> {
	await fetch('world', {
		method: 'POST',
		body: JSON.stringify({ playerId }),
	});
}
