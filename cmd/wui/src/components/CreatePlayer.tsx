import { useState, type FormEvent, type JSX } from 'react';
import { player, type Player } from '../services';

interface Props {
	setPlayer(player: Player): void;
}

function CreatePlayer({ setPlayer }: Props): JSX.Element {
	const [name, setName] = useState('');
	const [error, setError] = useState<Error>();

	function handleSubmit(e: FormEvent) {
		e.preventDefault();

		player.create(name)
			.then(setPlayer)
			.catch(setError);
	}

	return (
		<div className='h-svh w-svw flex items-center justify-center flex-col gap-4'>
			<h1 className='text-2xl'>Create a Player</h1>
			<form onSubmit={handleSubmit} className='flex flex-col gap-2'>
				<input
					className='p-1 border-1 border-gray-500 rounded-md'
					type='text'
					placeholder='Name'
					value={name}
					onChange={e => setName(e.target.value)}
					required
				/>
				<input className='p-1 bg-gray-500 rounded-md' type='submit' value='Submit' />
			</form>
			{error && <span>{error.toString()}</span>}
		</div>
	);
}

export default CreatePlayer;
