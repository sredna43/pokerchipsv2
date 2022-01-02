export interface Request {
	name: string;
	action: string;
	amount: number;
}

export interface Player {
    name: string;
    is_host: boolean;
}

export interface Response {
    players: [Player];
    message: string;
    pot: number;
    whose_turn: string;
    dealer: string;
}