import { useEffect } from 'react';
import { CreatePlayer, Game } from './components';
import { useApp } from './hooks';

function App() {
	const app = useApp();

	useEffect(() => app.load(), [app]);

	if (app.player !== null) {
		return <Game player={app.player} start={console.log} />
	}

  return (
		<div>
			<CreatePlayer onSubmit={app.create} />
			<span>{app.error?.message}</span>
		</div>
	);
}

export default App;
