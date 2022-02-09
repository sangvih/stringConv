package stringConv

import (
	"errors"
	"strconv"
)

func Dz(char string) (string, error) {
	var fchar, pchar string
	for i := 0; i < len(char); i++ {

		//если это символ, а не число (да это костыль))
		if char[i] > 57 {
			pchar = string(char[i])
			fchar += string(char[i])

		} else if i > 0 {

			digit, err := strconv.Atoi(string(char[i]))
			if err == nil {
				for i := 0; i < digit-1; i++ {
					fchar += pchar
				}
			} else {
				fchar += string(char[i])
			}

		} else {
			return "", errors.New("")
		}

	}
	return fchar, nil
}
