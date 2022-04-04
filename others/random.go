package main

import "fmt"
import "math/rand"

func main() {
	fmt.Println("begin 💰roll gift !!")
	fmt.Println("--------------------")

	peoples := [13]string{
		"就砺","铭少","莫须","穹谷","锐晟","辛八",
		"心远","银桑","远跖","中亭","中西","亦余",
		"自疏"}

	counts := len(peoples)

	var (
		userInput int
		rollNumber int
	)

	for i := 0; i < counts-1; i++ {

		for {
			fmt.Println("😄 请输入随机数种子：")
			_, err := fmt.Scanf("%d",&userInput)
			if(err == nil) {
				break;
			}
			fmt.Println(err)
			fmt.Println("🥸 非法的输入，请重试")
		}

		rand.Seed(int64(userInput))
		rollNumber = rand.Intn(counts-i)

		//fmt.Println(rollNumber)
		fmt.Println()
		fmt.Println("🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉")
		fmt.Println("🎉")
		fmt.Println("🤗 你获得了 "+peoples[rollNumber]+" 的礼物 🎁")
		fmt.Println("🎉")
		fmt.Println("🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉")
		fmt.Println("-------------------------") 
		fmt.Println()
		if(i<counts-2) {
			fmt.Println("😎 请"+peoples[rollNumber]+"开始抽奖 ")
		}

		peoples[counts-1-i],peoples[rollNumber] = peoples[rollNumber],peoples[counts-1-i]
	}

	fmt.Println("-------------------------") 
	fmt.Println()
	fmt.Println("中亭额外准备了一份礼物🥳")
	fmt.Println()
	fmt.Println("😎 开始抽奖 😎")

	for {
		fmt.Println("😄 请输入随机数种子：")
		_, err := fmt.Scanf("%d",&userInput)
		if(err == nil) {
			break;
		}
		fmt.Println(err)
		fmt.Println("🥸 非法的输入，请重试")
	}

	rand.Seed(int64(userInput))
	rollNumber = rand.Intn(counts-1)

	//fmt.Println(rollNumber)
	fmt.Println()
	fmt.Println("🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉")
	fmt.Println("🎉")
	fmt.Println("🤗 恭喜 "+peoples[rollNumber]+" 获得了的礼物 🎁")
	fmt.Println("🎉")
	fmt.Println("🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉🎉")
	fmt.Println("-------------------------") 

}