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

type lobbyResponse struct {
	Data  []playerLobbyResponse `json:"data"`
	Host  bool                  `json:"host"`
	You   int                   `json:"you"`
	Start bool                  `json:"start"` //flag to tell the clients to start timer
	Timer int                   `json:"timer"`
}
