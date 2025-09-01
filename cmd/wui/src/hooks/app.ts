import { useReducer, type ActionDispatch } from 'react';
import { player, session, type Player } from '../services';

class App {
	readonly #state: State;
	readonly #dispatch: Dispatch;

	constructor(state: State, dispatch: Dispatch) {
		this.#state = state;
		this.#dispatch = dispatch;
	}

	public get player(): Player | null {
		return this.#state.player;
	}

	public get error(): Error | null {
		return this.#state.error;
	}

	public create(name: string): void {
		player.create(name)
			.then(this.#created)
			.catch(this.#failed);
	}

	public load(): void {
		const id = session.get();
		if (id === null) return;

		player.get(id)
			.then(this.#loaded)
			.catch(this.#failed);
	}

	#created(player: Player): void {
		session.set(player.id);

		this.#dispatch({
			type: 'created',
			payload: player,
		});
	}

	#loaded(player: Player | null): void {
		this.#dispatch({
			type: 'loaded',
			payload: player,
		});
	}

	#failed(error: Error): void {
		this.#dispatch({
			type: 'failed',
			payload: error,
		});
	}
}

interface State {
	error: Error | null;
	player: Player | null;
}

const initialState: State = {
	error: null,
	player: null,
};

interface Created {
	type: 'created',
	payload: Player,
}

interface Failed {
	type: 'failed';
	payload: Error;
}

interface Loaded {
	type: 'loaded';
	payload: Player | null;
}

type Actions = Created | Failed | Loaded;

type Dispatch = ActionDispatch<[action: Actions]>;

function reducer(state: State, action: Actions): State {
	switch (action.type) {
		case 'created':
			return { ...state, player: action.payload };
		case 'failed':
			return { ...state, error: action.payload };
		case 'loaded':
			return { ...state, player: action.payload };
	}
}

export function useApp(): App {
	return new App(...useReducer(reducer, initialState));
}
