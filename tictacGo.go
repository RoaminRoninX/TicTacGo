package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func errorChecker(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func clearScreen() {
	os.Stdout.Write([]byte{0x1B, 0x5B, 0x33, 0x3B, 0x4A, 0x1B, 0x5B, 0x48, 0x1B, 0x5B, 0x32, 0x4A})
}

func printBoard(gameBoard [8][3]string, playerTurn int64) {
	clearScreen()
	fmt.Printf("\nRoaminRoninX's TicTacGo\n")
	fmt.Printf("Player %v's Turn\n\n", playerTurn)
	fmt.Println(" ___ ___ ___")
	for i := 1; i < 8; i++ {
		for j := 0; j < 3; j++ {
			if i == 0 {
				fmt.Printf("%v", gameBoard[i][j])
			} else if j%2 == 0 && j != 0 {
				if i == 7 {
					fmt.Printf("|%v|\n", gameBoard[i][j])
				} else {
					fmt.Printf("|%v|\n|   |   |   |\n", gameBoard[i][j])
				}
			} else {
				fmt.Printf("|%v", gameBoard[i][j])
			}
		}
	}
}

func findChoice() int64 {
	fmt.Println("Please Choose a number between 1 to 9 to mark one of the boxes")
	answerOption := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	reader := bufio.NewReader(os.Stdin)
	i, err := reader.ReadString('\n')
	errorChecker(err)
	input := strings.Trim(i, "\n")
	count := 0
	for _, num := range answerOption {
		if num == input {
			count++
			break
		}
	}
	if count == 0 {
		fmt.Printf("%v is not an option\n", input)
		return findChoice()
	}
	choice, err := strconv.ParseInt(input, 10, 64)
	errorChecker(err)
	return choice
}

func playAgain() bool {
	var choice bool
	fmt.Printf("Do you want to play again?[y/n]: ")
	answerOption := [4]string{"y", "Y", "n", "N"}
	reader := bufio.NewReader(os.Stdin)
	i, err := reader.ReadString('\n')
	errorChecker(err)
	input := strings.Trim(i, "\n")
	count := 0
	for _, answer := range answerOption {
		if answer == input {
			count++
			break
		}
	}
	if count == 0 {
		fmt.Printf("%v is not an option\nY for yes and N for no\n", input)
		return playAgain()
	}
	switch input {
	case "Y":
		choice = true
	case "y":
		choice = true
	case "N":
		choice = false
	case "n":
		choice = false
	}
	return choice
}

func findDuplicates(chooseWinner map[int64]string) []int64 {
	choices := []int64{}
	for i := 1; i <= 9; i++ {
		_, ok := chooseWinner[int64(i)]
		if ok {
			choices = append(choices, int64(i))
		}
	}
	return choices
}

func checkForWinner(chooseWinner map[int64]string) int64 {
	var winner int64 = 0
	winningLetter := ""
	var count int64 = 0
	var countSquares int64 = 0
	for true {
		// _________________123_____________________________//

		for i := 1; i <= 2; i++ {
			_, exist1 := chooseWinner[int64(i)]
			_, exist2 := chooseWinner[int64(i+1)]
			if exist1 && exist2 {
				if chooseWinner[int64(i)] == chooseWinner[int64(i+1)] {
					count++
				}
			}
		}

		if count == 2 {
			winningLetter = chooseWinner[int64(1)]
			break
		} else {
			count = 0
		}

		// ___________________456___________________________//

		for i := 4; i <= 5; i++ {
			_, exist1 := chooseWinner[int64(i)]
			_, exist2 := chooseWinner[int64(i+1)]
			if exist1 && exist2 {
				if chooseWinner[int64(i)] == chooseWinner[int64(i+1)] {
					count++
				}
			}
		}
		if count == 2 {
			winningLetter = chooseWinner[int64(4)]
			break
		} else {
			count = 0
		}
		// _____________________789_________________________//

		for i := 7; i <= 8; i++ {
			_, exist1 := chooseWinner[int64(i)]
			_, exist2 := chooseWinner[int64(i+1)]
			if exist1 && exist2 {
				if chooseWinner[int64(i)] == chooseWinner[int64(i+1)] {
					count++
				}
			}
		}

		if count == 2 {
			winningLetter = chooseWinner[int64(7)]
			break
		} else {
			count = 0
		}

		// __________________147____________________________//

		for i := 1; i <= 4; i += 3 {
			_, exist1 := chooseWinner[int64(i)]
			_, exist2 := chooseWinner[int64(i+3)]
			if exist1 && exist2 {
				if chooseWinner[int64(i)] == chooseWinner[int64(i+3)] {
					count++
				}
			}
		}
		if count == 2 {
			winningLetter = chooseWinner[int64(1)]
			break
		} else {
			count = 0
		}
		// ___________________258___________________________//

		for i := 2; i <= 5; i += 3 {
			_, exist1 := chooseWinner[int64(i)]
			_, exist2 := chooseWinner[int64(i+3)]
			if exist1 && exist2 {
				if chooseWinner[int64(i)] == chooseWinner[int64(i+3)] {
					count++
				}
			}
		}
		if count == 2 {
			winningLetter = chooseWinner[int64(2)]
			break
		} else {
			count = 0
		}
		// _____________________369_________________________//

		for i := 3; i <= 6; i += 3 {
			_, exist1 := chooseWinner[int64(i)]
			_, exist2 := chooseWinner[int64(i+3)]
			if exist1 && exist2 {
				if chooseWinner[int64(i)] == chooseWinner[int64(i+3)] {
					count++
				}
			}
		}
		if count == 2 {
			winningLetter = chooseWinner[int64(3)]
			break
		} else {
			count = 0
		}

		// _________________357_____________________________//

		for i := 3; i <= 5; i += 2 {
			_, exist1 := chooseWinner[int64(i)]
			_, exist2 := chooseWinner[int64(i+2)]
			if exist1 && exist2 {
				if chooseWinner[int64(i)] == chooseWinner[int64(i+2)] {
					count++
				}
			}
		}
		if count == 2 {
			winningLetter = chooseWinner[int64(3)]
			break
		} else {
			count = 0
		}

		// _______________________159_______________________//

		for i := 1; i <= 5; i += 4 {
			_, exist1 := chooseWinner[int64(i)]
			_, exist2 := chooseWinner[int64(i+4)]
			if exist1 && exist2 {
				if chooseWinner[int64(i)] == chooseWinner[int64(i+4)] {
					count++
				}
			}
		}
		if count == 2 {
			winningLetter = chooseWinner[int64(1)]
			break
		} else {
			count = 0
		}
		break
	}

	switch winningLetter {
	case "X":
		winner = 1
	case "Y":
		winner = 2
	case "":
		for i := 1; i <= 9; i++ {
			_, ok := chooseWinner[int64(i)]
			if ok {
				countSquares++
			}
		}

		switch countSquares {
		case 9:
			winner = 3
		default:
			winner = 0
		}
	}

	return winner
}

func gameLogic() {
	var playerTurn int64 = 1
	var draws int64 = 0
	var alreadyExist int64 = 0
	gameBoard := [8][3]string{
		{"___", "___", "___"},
		{"   ", "   ", "   "},
		{" 1 ", " 2 ", " 3 "},
		{"___", "___", "___"},
		{" 4 ", " 5 ", " 6 "},
		{"___", "___", "___"},
		{" 7 ", " 8 ", " 9 "},
		{"___", "___", "___"},
	}
	boardSquares := map[int64][]int64{
		1: {2, 0},
		2: {2, 1},
		3: {2, 2},
		4: {4, 0},
		5: {4, 1},
		6: {4, 2},
		7: {6, 0},
		8: {6, 1},
		9: {6, 2},
	}
	finalScore := map[int64]int64{
		1: 0,
		2: 0,
	}
	chooseWinner := make(map[int64]string)
	printBoard(gameBoard, playerTurn)

	for true {
		choice := findChoice()
		takenChoices := findDuplicates(chooseWinner)
		for i := 0; i < len(takenChoices); i++ {
			if choice == takenChoices[i] {
				fmt.Println("This Square is already Taken. Choose Again!")
				alreadyExist++
				break
			}
		}
		if alreadyExist > 0 {
			alreadyExist = 0
			continue
		}
		switch playerTurn {
		case 1:
			gameBoard[boardSquares[choice][0]][boardSquares[choice][1]] = " X "
			chooseWinner[choice] = "X"
			playerTurn = 2
		case 2:
			gameBoard[boardSquares[choice][0]][boardSquares[choice][1]] = " Y "
			chooseWinner[choice] = "Y"
			playerTurn = 1
		}
		// ___________________________________________________
		printBoard(gameBoard, playerTurn)
		winner := checkForWinner(chooseWinner)
		switch winner {
		case 1:
			finalScore[1]++
			fmt.Println("Player 1 wins!")
			break

		case 2:
			finalScore[2]++
			fmt.Println("Player 2 wins!")
			break
		case 3:
			draws++
			fmt.Println("Game ended in a Draw!")
			break
		}

		if winner == 1 || winner == 2 || winner == 3 {
			play := playAgain()
			if play {
				gameBoard = [8][3]string{
					{"___", "___", "___"},
					{"   ", "   ", "   "},
					{" 1 ", " 2 ", " 3 "},
					{"___", "___", "___"},
					{" 4 ", " 5 ", " 6 "},
					{"___", "___", "___"},
					{" 7 ", " 8 ", " 9 "},
					{"___", "___", "___"},
				}
				playerTurn = 1
				for k := range chooseWinner {
					delete(chooseWinner, k)
				}
				printBoard(gameBoard, playerTurn)
				continue

			} else {
				break
			}
		}
	}
	clearScreen()
	fmt.Println("Final Score")
	for key, score := range finalScore {
		fmt.Printf("Player %v won %v times\n", key, score)
	}
	fmt.Printf("There were %v draws\n", draws)
	fmt.Printf("GAME OVER!\nThank you for playing RoaminRoninX's TicTacGo\n")
}

func main() {
	gameLogic()
}
