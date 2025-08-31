import z from 'zod/v4';

const Player = z.object({
	id: z.string(),
	name: z.string(),
});

const Response = z.object({
	player: Player,
});

export type Player = z.infer<typeof Player>;

export async function create(name: string): Promise<Player> {
	const res = await fetch('/player', {
		method: 'POST',
		body: JSON.stringify({ name }),
	});

	return Response.parse(await res.json()).player;
}

export async function get(id: string): Promise<Player | null> {
	const res = await fetch(`/player/${id}`);

	if (res.status === 404) return null;

	return Response.parse(await res.json()).player;
}
