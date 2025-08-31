import { useCallback, useEffect, useState } from 'react';
import { CreatePlayer, Game } from './components';
import { player, session, world, type Player } from './services';

function App() {
	const [p, setP] = useState<Player | null>(null);
	const [error, setError] = useState<Error>();

	useEffect(() => {
		const id = session.get();
		if (id === null) return;

		player.get(id)
			.then(setP)
			.catch(e => {
				setError(e);
				session.clear();
			});
	}, []);

	const createPlayer = useCallback((name: string): void => {
		player.create(name)
			.then(player => {
				setP(player);
				session.set(player.id);
			})
			.catch(setError);
	}, []);

	const createWorld = useCallback((playerId: string) => {
		world.create(playerId)
			.then(console.log)
			.catch(setError);
	}, []);

	if (p !== null) {
		return <Game player={p} start={createWorld} />
	}

  return (
		<div>
			<CreatePlayer onSubmit={createPlayer} />
			<span>{error?.message}</span>
		</div>
	);
}

export default App;
