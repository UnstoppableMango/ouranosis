import { useState, type FormEvent } from 'react'

function App() {
	const [name, setName] = useState('');

	function handleSubmit(e: FormEvent) {
		e.preventDefault();
		console.log(name);
	}

  return (
		<div className='h-svh w-svw flex items-center justify-center flex-col'>
			<h1>Create Player</h1>
			<form onSubmit={handleSubmit} className='flex flex-col'>
				<input
					type='text'
					placeholder='Name'
					value={name}
					onChange={e => setName(e.target.value)}
					required
				/>
				<input type='submit' value='Submit' />
			</form>
		</div>
  );
}

export default App;
