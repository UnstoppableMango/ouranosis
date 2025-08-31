import type { JSX } from 'react';
import type { Player } from '../services';

interface Props {
	player: Player;
}

function Game({ player: { id, name } }: Props): JSX.Element {
	return (
		<div className='flex flex-col'>
			<h2>GAME</h2>
			<h3>Id: {id}</h3>
			<h3>Name: {name}</h3>
		</div>
	);
}

export default Game;
