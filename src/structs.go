package src

type response struct {
	Current int `json:"current_turn"`
	Dice    int `json:"dice_num"`
}
type payload struct {
	Dice int `json:"dice_num"`
}
type responseSkeleton struct {
	Data interface{} `json:"data"`
}
