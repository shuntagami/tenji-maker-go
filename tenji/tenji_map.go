package tenji

var (
	// 母音の点字の位置を6ビットの2進数で表します
	VowelMap = map[string]byte{
		"A": byte(0b100000),
		"I": byte(0b101000),
		"U": byte(0b110000),
		"E": byte(0b111000),
		"O": byte(0b011000),
	}

	// 子音の点字の位置を6ビットの2進数で表します
	ConsonantMap = map[string]byte{
		"K": byte(0b000001),
		"S": byte(0b000101),
		"T": byte(0b000110),
		"N": byte(0b000010),
		"H": byte(0b000011),
		"M": byte(0b000111),
		"Y": byte(0b010000),
		"R": byte(0b000100),
	}
)
