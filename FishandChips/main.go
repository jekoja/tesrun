package main

func FishandChips(s sting, n int) {
	if n < 0 {
		print("error")
	}
	if n%2 == 0 && n%3 == 0 {
		print("fish and chips")
	}
	if n%2 == 0 {
		print("fish")
	}
	if n%3 == 0 {
		print("chips")
	}
}
