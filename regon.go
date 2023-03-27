package plvalidator

var weightRegonMap = map[int][]int{9: {8, 9, 2, 3, 4, 5, 6, 7}, 14: {2, 4, 8, 5, 0, 9, 7, 3, 6, 1, 2, 4, 8}}

func Regon(regon string) bool {
	regonlen := len(regon)
	weight, ok := weightRegonMap[regonlen]
	if !ok {
		return false
	}
	regonDigits, err := SplitToInt(regon)
	if err != nil {
		return false
	}
	controlDigit := CheckSum(weight, regonDigits, MultipleByWeigthAndAddToChecksum, Modulo11)
	return controlDigit == regonDigits[regonlen-1]
}
