//Created by Kuttymbek Daniyar
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Global variable for DES
var (
	p10    = []int{3, 5, 2, 7, 4, 10, 1, 9, 8, 6}
	ip     = []int{2, 6, 3, 1, 4, 8, 5, 7}
	ip_r   = []int{4, 1, 3, 5, 7, 2, 8, 6}
	p8     = []int{6, 3, 7, 4, 8, 5, 10, 9}
	text   []rune
	key    = []byte{}
	keyOne []byte
	keyTwo []byte
	ep     = []int{4, 1, 2, 3, 2, 3, 4, 1}
	S0     = [4][4]int{
		{1, 0, 3, 2},
		{3, 2, 1, 0},
		{0, 2, 1, 3},
		{3, 1, 3, 1},
	}
	S1 = [4][4]int{
		{1, 1, 2, 3},
		{2, 0, 1, 3},
		{3, 0, 1, 0},
		{2, 1, 0, 3},
	}
	p4         = []int{2, 4, 3, 1}
	final_text string
	answer     int
	file       *os.File
)

func getNewKey(x []byte, y []int) []byte {
	var new []byte
	for _, i := range y {
		new = append(new, x[i-1])
	}
	return new
}
func shiftAll(x []byte, shift int) []byte {
	return append(append([]byte{}, x[shift:]...), x[:shift]...)
}
func xor(a, b []byte) []byte {
	var res []byte
	for i := range a {
		res = append(res, (a[i]^b[i])+'0')
	}
	return res
}
func pair(x, y byte) string {
	return string(x) + string(y)
}
func getValue(s []byte, x [4][4]int) int64 {
	a, _ := strconv.ParseInt(pair(s[0], s[3]), 2, 64)
	b, _ := strconv.ParseInt(pair(s[1], s[2]), 2, 64)
	return int64(S0[a][b])
}
func fx(cur_text, cur_key []byte) []byte {
	cur_text = getNewKey(cur_text, ep)
	cur_text = xor(cur_text, cur_key)
	cur_text = getNewKey([]byte(fmt.Sprintf("%02b", getValue(cur_text[:4], S0))+fmt.Sprintf("%02b", getValue(cur_text[4:], S1))), p4)
	return cur_text
}
func initilizate() {
	key = getNewKey(key, p10)
	key = append(shiftAll(key[:5], 1), shiftAll(key[5:], 1)...)
	keyOne = getNewKey(key, p8)
	key = append(shiftAll(key[:5], 2), shiftAll(key[5:], 2)...)
	keyTwo = getNewKey(key, p8)
}
func process(char, k1, k2 []byte) string {
	char = getNewKey(char, ip)
	char = append(append([]byte{}, char[4:]...), xor(char[:4], fx(char[4:], k1))...)
	char = getNewKey(append(xor(char[:4], fx(char[4:], k2)), char[4:]...), ip_r)
	return string(char)
}

func main() {
	//Scan input data on txt
	fmt.Print("Please choose :\n1 - encrypt\n2 - decrypt\n")
	fmt.Scan(&answer)
	if answer == 2 {
		file, _ = os.Open("input.txt")
	} else {
		file, _ = os.Open("output.txt")
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	key = []byte(fmt.Sprintf("%010b", num))
	scanner.Scan()
	text = []rune(scanner.Text())
	fmt.Println(string(key))
	initilizate()
	if answer == 1 {
		keyOne, keyTwo = keyTwo, keyOne
	}
	for _, i := range text {
		cur := fmt.Sprintf("%08b", i)
		ch, _ := strconv.ParseInt(process([]byte(cur), keyOne, keyTwo), 2, 64)
		final_text += string(ch)
		fmt.Printf("%c -> %c\n", i, ch)
	}
	f, _ := os.Create("output.txt")
	f.Write([]byte(strconv.Itoa(num)))
	f.Write([]byte("\n"))
	f.Write([]byte(final_text))
}
