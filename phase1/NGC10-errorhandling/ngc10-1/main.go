package main

import (
	"errors"
	"fmt"
	"sort"

	"golang.org/x/exp/slices"
)

func anagram(slice1,slice2 []string) string{
	var output string
	
	if len(slice1) != len(slice2){
		output = "bukan anagram"
		return output
	}

	output = "anagram"
	for i := range slice1{
		if slice1[i] != slice2[i]{
			output = "bukan anagram"
		}
	}

	return output
}

func sortString(word string)[]string{
	slice_word := []string{}
	for i := range word{
		slice_word = append(slice_word, string(word[i]))
	}
	sort.Strings(slice_word)

	return slice_word
}


func validInput(CekSlice[]string) ([]string,error) {
	err_char := `!@#$%^&*()-+[]{};:'"<>?/., `
	var err_char_slice []string

	for i := range err_char{
		err_char_slice = append(err_char_slice, string(err_char[i]))
	}

	// check 1
	if len(CekSlice) > 10 {
		return []string{}, errors.New("Error: karakter tidak boleh lebih dari 10")
	}

	// check 2
	for _,i := range CekSlice{
		if slices.Contains(err_char_slice,i) == true{
			return []string{}, errors.New(`Error: tidak boleh ada simbol (!@#$%^&*()-+[]{};:'"<>?/.,)`)
		}
	}

	return CekSlice, nil
}


func main() {
	fmt.Printf("-PROGRAM ANAGRAM START-\n")

	for{	
		// define input word and slice
		var word1 string
		var word2 string
		var sorted_slice1 []string
		var sorted_slice2 []string

		// define valid and error
		var valid_slice1 []string
		var err1 error
		var valid_slice2 []string
		var err2 error
		// WORD 1
		// input CLI by user
		for {
			fmt.Printf("\nWrite the first word : ")
			fmt.Scanln(&word1)
	
			// cast to slice and sorting
			sorted_slice1 = sortString(word1)
	
			// valid and error handling
			valid_slice1 , err1 = validInput(sorted_slice1)
			
			if err1 != nil{
				fmt.Println(err1.Error())
				continue
			}
			break
		}
		// WORD 2
		// input CLI by user
		for{
			fmt.Printf("Write the second word : ")
			fmt.Scanln(&word2)
	
			// cast to slice and sorting
			sorted_slice2 = sortString(word2)
	
			// valid and error handling
			valid_slice2 , err2 = validInput(sorted_slice2)
	
			if err2 != nil{
				fmt.Println(err2.Error())
				continue
			}
			break
		}
		fmt.Println("")
		fmt.Println("Kata 1 dan Kata 2 : ",anagram(valid_slice1,valid_slice2))

		ask := ""
		fmt.Printf("Write 'stop' to end program : ")
		fmt.Scanln(&ask)
		if ask == "stop"{
			fmt.Printf("\n\n-PROGRAM ANAGRAM END-")
			break
		}

	}

	
}