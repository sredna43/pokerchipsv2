export interface Res {
	requester: string;
	players: Players;
	message: string;
	pot: number;
	whose_turn: number;
	dealer: string;
	error: boolean;
}

export interface Req {
	name: string;
	action: string;
	amount: number;
}

export interface Player {
	name: string;
	is_host: boolean;
	is_dealer: boolean;
	chips: number;
	spot: number;
}

export interface Players {
	[key: string]: Player;
}

export interface TableCheckResponse {
	exists: boolean;
}
