package src

type playerLobbyResponse struct {
	Player string `json:"name"`
	Host   bool   `json:"is_host"`
}

type response struct {
	Current int `json:"current_turn"`
	Dice    int `json:"dice_num"`
}
type payload struct {
	Dice int `json:"dice_num"`
}
