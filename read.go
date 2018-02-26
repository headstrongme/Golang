//Here value of n is 3 i.e. 3-gram is calculated from the code
package main

import (
	"fmt"
	"io/ioutil"
	//"log"
	//"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}

}

func dup_count(ngram map[int]string) map[string]int {

	duplicate_frequency := make(map[string]int)

	for _, item := range ngram {

		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
}

func main() {

	f, err := ioutil.ReadFile("TheRaven.txt")
	check(err)

	lines := make(map[int]string)
	ngram := make(map[int]string)
	dup_map := make(map[string]int)

	//sub := strings.ToLower(f)
	//str := fmt.Sprintf("%s", f)

	//n := bytes.IndexByte(f, 0)
	str := string(f)
	str2 := strings.ToLower(str)
	str3 := strings.Trim(str2, ".',;-!\"")

	//reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//str3 := reg.ReplaceAllString(str2, "")

	for n, line := range strings.Split(str3, " ") {
		lines[n] = line
	}

	count := 0

	for i := 0; i < len(lines); i++ {
		w1 := lines[i]
		l := len(w1)

		for j := 0; j < l-2; j++ {
			//for k := 0; k < 2; k++ {

			ngram[count] = w1[j : j+3]
			count++
			dup_map = dup_count(ngram)
			//}

		}
	}
	var test int
	var test2 string
	//fmt.Println(lines)
	//fmt.Println(ngram)
	for m, s := range dup_map {

		if s > test {
			test = s
			test2 = m
		}

	}
	fmt.Print("Most  Frequent ngram is=> ", test2, ": ", test)
	fmt.Println(" ")

	for k, v := range dup_map {
		fmt.Printf("Item : %s , Count : %d\n", k, v)
	}
}

//substring := value[4:len(value)]
