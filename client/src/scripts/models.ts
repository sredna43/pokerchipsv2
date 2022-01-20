export interface Res {
	requester: string;
	message: string;
	error: boolean;
    table: Table;
}

export interface Req {
	name: string;
	action: string;
	amount: number;
}

export interface Table {
    has_host: boolean;
    players: Players;
    whose_turn: number;
    dealer: string;
	dealer_spot: string;
    pot: number;
	playing: boolean;
    initial_chips: number;
    can_check: boolean;
    betting_round: number;
    current_bet: number;
	hand_won: boolean;
}

export interface Player {
	name: string;
	is_host: boolean;
	is_dealer: boolean;
	chips: number;
	spot: number;
    folded: boolean;
	called: boolean;
}

export interface Players {
	[key: string]: Player;
}

export interface TableCheckResponse {
	exists: boolean;
}

export interface NewTableResponse {
    id: string;
}
