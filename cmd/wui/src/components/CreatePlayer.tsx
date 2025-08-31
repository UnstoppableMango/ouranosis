import { useState, type FormEvent, type JSX } from 'react';

interface Props {
	onSubmit(name: string): void;
}

function CreatePlayer({ onSubmit }: Props): JSX.Element {
	const [name, setName] = useState('');

	const handleSubmit = (e: FormEvent): void => {
		e.preventDefault();
		onSubmit(name);
	};

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
		</div>
	);
}

export default CreatePlayer;
