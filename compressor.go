package d1util

var zipKey = []string{"_a", "_b", "_c", "_d", "_e", "_f", "_g", "_h", "_i", "_j", "_k", "_l", "_m", "_n", "_o", "_p",
	"_q", "_r", "_s", "_t", "_u", "_v", "_w", "_x", "_y", "_z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K",
	"L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7",
	"8", "9", "-", "_"}
var zkArray = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
	"t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
	"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-", "_"}

var hashCodes map[string]int

func init() {
	hashCodes = make(map[string]int)
	for j := len(zkArray) - 1; j >= 0; j-- {
		hashCodes[zkArray[j]] = j
	}
}

func decode64(codedValue string) int {
	return hashCodes[codedValue]
}

func encode64(value int) string {
	return zkArray[value]
}
