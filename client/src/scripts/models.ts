export interface Res {
    requester: string;
    players: Map<string, Player>;
    message: string;
    pot: number;
    whose_turn: string;
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
    chips: number;
}

export interface TableCheckResponse {
    exists: boolean;
}