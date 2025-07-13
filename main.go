package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

var hash = []rune{
	'X', 'r', 'q', 'a', 'H', 'N', 'p', 'd', 'S', 'Y', 'w', '8', '6', '2', '1', '5',
}

type NetData struct {
	BSSID string
}

type DlinkKeyCalculator struct{}

func (d *DlinkKeyCalculator) GetKey(network NetData) ([]string, error) {

	trimmedBSSID := strings.ReplaceAll(network.BSSID, ":", "")

	if len(trimmedBSSID) != 12 {
		return nil, fmt.Errorf("invalid BSSID format: expected 12 characters, got %d", len(trimmedBSSID))
	}

	key := make([]rune, 20)
	newKey := make([]rune, 20)

	key[0] = rune(trimmedBSSID[11])
	key[1] = rune(trimmedBSSID[0])
	key[2] = rune(trimmedBSSID[10])
	key[3] = rune(trimmedBSSID[1])
	key[4] = rune(trimmedBSSID[9])
	key[5] = rune(trimmedBSSID[2])
	key[6] = rune(trimmedBSSID[8])
	key[7] = rune(trimmedBSSID[3])
	key[8] = rune(trimmedBSSID[7])
	key[9] = rune(trimmedBSSID[4])
	key[10] = rune(trimmedBSSID[6])
	key[11] = rune(trimmedBSSID[5])
	key[12] = rune(trimmedBSSID[1])
	key[13] = rune(trimmedBSSID[6])
	key[14] = rune(trimmedBSSID[8])
	key[15] = rune(trimmedBSSID[9])
	key[16] = rune(trimmedBSSID[11])
	key[17] = rune(trimmedBSSID[2])
	key[18] = rune(trimmedBSSID[4])
	key[19] = rune(trimmedBSSID[10])

	for i := 0; i < 20; i++ {
		t := key[i]
		var index int

		if t >= '0' && t <= '9' {
			index = int(t - '0')
		} else {
			t = unicode.ToUpper(t)
			if t >= 'A' && t <= 'F' {
				index = int(t-'A') + 10
			} else {
				return nil, fmt.Errorf("invalid character in BSSID: %c", t)
			}
		}

		newKey[i] = hash[index]
	}

	result := []string{string(newKey)}
	return result, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <BSSID>")
		fmt.Println("Example: go run main.go 00:1B:2F:12:34:56")
		os.Exit(1)
	}

	bssid := os.Args[1]

	network := NetData{BSSID: bssid}
	calculator := &DlinkKeyCalculator{}

	keys, err := calculator.GetKey(network)
	if err != nil {
		fmt.Printf("Error generating key: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("BSSID: %s\n", bssid)
	fmt.Printf("Generated Key: %s\n", keys[0])
}
