package main

import "fmt"
import "math/rand"

func main() {
	fmt.Println("begin π°roll gift !!")
	fmt.Println("--------------------")

	peoples := [13]string{
		"ε°±η Ί", "ι­ε°", "θ«ι‘»", "η©Ήθ°·", "ιζ", "θΎε«",
		"εΏθΏ", "ιΆζ‘", "θΏθ·", "δΈ­δΊ­", "δΈ­θ₯Ώ", "δΊ¦δ½",
		"θͺη"}

	counts := len(peoples)

	var (
		userInput  int
		rollNumber int
	)

	for i := 0; i < counts-1; i++ {

		for {
			fmt.Println("π θ―·θΎε₯ιζΊζ°η§ε­οΌ")
			_, err := fmt.Scanf("%d", &userInput)
			if err == nil {
				break
			}
			fmt.Println(err)
			fmt.Println("π₯Έ ιζ³ηθΎε₯οΌθ―·ιθ―")
		}

		rand.Seed(int64(userInput))
		rollNumber = rand.Intn(counts - i)

		//fmt.Println(rollNumber)
		fmt.Println()
		fmt.Println("πππππππππππππππ")
		fmt.Println("π")
		fmt.Println("π€ δ½ θ·εΎδΊ " + peoples[rollNumber] + " ηη€Όη© π")
		fmt.Println("π")
		fmt.Println("πππππππππππππππ")
		fmt.Println("-------------------------")
		fmt.Println()
		if i < counts-2 {
			fmt.Println("π θ―·" + peoples[rollNumber] + "εΌε§ζ½ε₯ ")
		}

		peoples[counts-1-i], peoples[rollNumber] = peoples[rollNumber], peoples[counts-1-i]
	}

	fmt.Println("-------------------------")
	fmt.Println()
	fmt.Println("δΈ­δΊ­ι’ε€εε€δΊδΈδ»½η€Όη©π₯³")
	fmt.Println()
	fmt.Println("π εΌε§ζ½ε₯ π")

	for {
		fmt.Println("π θ―·θΎε₯ιζΊζ°η§ε­οΌ")
		_, err := fmt.Scanf("%d", &userInput)
		if err == nil {
			break
		}
		fmt.Println(err)
		fmt.Println("π₯Έ ιζ³ηθΎε₯οΌθ―·ιθ―")
	}

	rand.Seed(int64(userInput))
	rollNumber = rand.Intn(counts - 1)

	//fmt.Println(rollNumber)
	fmt.Println()
	fmt.Println("πππππππππππππππ")
	fmt.Println("π")
	fmt.Println("π€ ζ­ε " + peoples[rollNumber] + " θ·εΎδΊηη€Όη© π")
	fmt.Println("π")
	fmt.Println("πππππππππππππππ")
	fmt.Println("-------------------------")

}
