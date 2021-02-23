package card_type

type CardType int

const (
	Visa CardType = iota
	MasterCard
	Amex
	AmericanExpress
)
