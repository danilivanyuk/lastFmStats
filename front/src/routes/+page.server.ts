
import type { Actions } from './$types';

export const actions = {
	default: async (event) => {
		console.log('aaaa');
		// TODO log the user in
	},
	user: async (event) => {
		
		// TODO register the user
		console.log('tryt', event);
	}
} satisfies Actions;
