import type { JSX } from 'react';
import type { Player } from '../services';

interface Props {
	player: Player;
	start(id: string): void;
}

function Game({ player: { id, name }, start }: Props): JSX.Element {
	const handlePress = (): void => start(id);

	return (
		<div className='h-svh w-svw flex items-center justify-center'>
			<div className='flex flex-col'>
				<h2>GAME</h2>
				<h3>Name: {name}</h3>
				<button
					className='p-1 bg-gray-500 hover:bg-gray-400 rounded-md'
					onClick={handlePress}
				>Start</button>
			</div>
		</div>
	);
}

export default Game;
