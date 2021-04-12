package dt

func ParseNumber(num string) (*GenericNumber, error) {
	return numberFromString(num)
}
