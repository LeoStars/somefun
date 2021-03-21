package main

import (
	"fmt"
	"io/ioutil"
	//"os"
)

func main() {
	data, _ := ioutil.ReadFile("test.txt")
	/*dataString := ""
	for j := range data {
		//dataString += string(data[j])
		if data[j] == byte("─") {
			data[j] = byte(' ')
		}
		fmt.Println(dataString)
	}*/

	dataString := string(data)
	dataRune := []rune(dataString)
	/*for i := range dataRune {
		if dataRune[i] == '─' {
			dataRune[i] = ' '
		}
		if dataRune[i] == '╚' {
			dataRune[i] = '┗'
		}
		if dataRune[i] == '═' {
			dataRune[i] = '━'
		}
		if dataRune[i] == '╣' {
			dataRune[i] = '┫'
		}
		if dataRune[i] == '║' {
			dataRune[i] = '┃'
		}
		if dataRune[i] == '╝' {
			dataRune[i] = '┛'
		}
		if dataRune[i] == '╗' {
			dataRune[i] = '┓'
		}

		if dataRune[i] == '╔' {
			dataRune[i] = '┏'
		}

	}*/
	a := 62136
	b := 62134
	s1 := dataRune[:a]
	s2 := dataRune[a:a * 2]
	s3 := dataRune[(a * 2):(a * 2 + b)+1]

	s4 := dataRune[(a * 2 + b)+1:((a + b) * 2)+2]


	s5 := dataRune[(a + b) * 2+2:a * 3 + b * 2+3]
	s6 := dataRune[a * 3 + b * 2+3:a * 3 + b * 3]
	for i := 0; i < 62130; i ++ {
		if string(s1[i:i + 4]) == "╔══╗" {
			if string(s2[i:i + 4]) == "║╔╗║" {

				if string(s3[i:i + 4]) == "║╚╝║" {

					if string(s4[i:i + 4]) == "║╔╗║" {

						if string(s5[i:i + 4]) == "║║║║" {
							if string(s6[i:i + 4]) == "╚╝╚╝" {
								fmt.Print("а")
							}
						}
					}
				}
			}
		}
	}
	dataString = string(dataRune)
	_ = ioutil.WriteFile("test2.txt", []byte(dataString), 0644)
}


