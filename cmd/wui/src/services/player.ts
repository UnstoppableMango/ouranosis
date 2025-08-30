export interface Player {
	id: string;
	name: string;
}

export async function create(name: string): Promise<Player> {
	const res = await fetch('/player', {
		method: 'POST',
		body: JSON.stringify({ name }),
	});

	return await res.json();
}
