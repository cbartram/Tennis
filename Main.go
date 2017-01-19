package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
);

/**
 * Player "Object" in Go
 */
type Player struct {
	playerScoreName string
	playerScoreValue int
	playerName string
	playerWins int
}

func main() {

	//Var declarations
	p1Name := "Player One"
	p2Name := "Player Two"

	p1Score := "Love"
	p2Score := "Love"

	p1Wins := 0
	p2Wins := 0

	mode := "game"

	playerOne := new(Player)
	playerTwo := new(Player)

	//Output the starting instructions of the game
	fmt.Println("Tennis")
	fmt.Println("-----------------")
	fmt.Println()
	fmt.Println("Enter one command line argument per line. Ex: --p1score Fifteen When you are ready to start type 'play'! ")

	//Read Console Input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Command line argument:")

	for scanner.Scan() {
		fmt.Println("Add Argument or type 'play' to start the match : ")
		text := scanner.Text()

		//if the user is ready to play
		if strings.EqualFold(text, "play") {
			break;
		}

		//If the string contains a command line parameter
		if strings.Contains(text, "--") {

			//How many command line arguments were passed
			//occurrenceCount := strings.Count(text, "--")

			index := strings.Index(text, "--")
			space := strings.Index(text, " ")


			command := text[index + 2:space]
			value   := text[space + 1:]

			//store the command/value pair in global variables
			switch command {
			case "p1name":
				p1Name = value
				fmt.Println("P1 Name set successfully")
				break

			case "p2name":
				p2Name = value
				fmt.Println("P2 Name set successfully")
				break

			case "p1score":
				//Assign the value of the command
				if strings.EqualFold(value, "Love") {
					p1Score = "Love"
				}

				if strings.EqualFold(value, "Fifteen") {
					p1Score = "Fifteen"
				}

				if strings.EqualFold(value, "Thirty") {
					p1Score = "Thirty"
				}

				if strings.EqualFold(value, "Forty") {
					p1Score = "Forty"
				}

				if strings.EqualFold(value, "Advantage") {
					p1Score = "Advantage"
				}

				//The value does not match any reasonable values
				if !strings.EqualFold(value, "Love") && !strings.EqualFold(value, "Fifteen") && !strings.EqualFold(value, "Thirty") && !strings.EqualFold(value, "Forty") && !strings.EqualFold(value, "Advantage") {
					fmt.Println("The value is incorrect for: p1Score")
				}

				break

			case "p2score":
				//Assign the value of the command
				if strings.EqualFold(value, "Love") {
					p2Score = "Love"
				}

				if strings.EqualFold(value, "Fifteen") {
					p2Score = "Fifteen"
				}

				if strings.EqualFold(value, "Thirty") {
					p2Score = "Thirty"
				}

				if strings.EqualFold(value, "Forty") {
					p2Score = "Forty"
				}

				if strings.EqualFold(value, "Advantage") {
					p2Score = "Advantage"
				}

				//The value does not match any reasonable values
				if !strings.EqualFold(value, "Love") && !strings.EqualFold(value, "Fifteen") && !strings.EqualFold(value, "Thirty") && !strings.EqualFold(value, "Forty") && !strings.EqualFold(value, "Advantage") {
					fmt.Println("The value is incorrect for: p2Score")
				}
				break;

			case "p1wins":
				intValue,_ := strconv.Atoi(value)
				//If the value is an int <= 6 and >= 0
				if intValue >= 0 && intValue <= 6 {
					p1Wins = intValue;
					fmt.Println("P1 Wins set successfully")
				} else {
					fmt.Println("The value entered for p1wins must be between 0 and 6.")
				}
				break
			case "p2wins":
 				intValue,_ := strconv.Atoi(value)
				if intValue >= 0 && intValue <= 6 {
					p2Wins = intValue;
					fmt.Println("P2 Wins set successfully")
				} else {
					fmt.Println("The value entered for p2wins must be between 0 and 6.")
				}
				break
			case "mode":
				if strings.EqualFold(value, "game") {
					mode = "game";
					fmt.Println("Game mode is currently: game")
				}
				if strings.EqualFold(value, "set") {
					mode = "set"
					fmt.Println("Game mode is currently: set")

				}
				if(!strings.EqualFold(value, "set") && !strings.EqualFold(value, "game")) {
					fmt.Println("The value you entered for mode must be either 'game' or 'set'")
				}

			default:
				fmt.Println("That command argument is not recognized.")
				break
			}

		} else {
			fmt.Println("That is not a valid command line argument. Command arguments must start with '--ARG_NAME ' ")
		}
	}

	//Set the Values
	playerOne.playerName = p1Name;
	playerOne.playerScoreName = p1Score
	playerOne.playerScoreValue = scoreNameToValue(p1Score)
	playerOne.playerWins = p1Wins

	playerTwo.playerName = p2Name
	playerTwo.playerScoreName = p2Score
	playerTwo.playerScoreValue = scoreNameToValue(p2Score)
	playerTwo.playerWins = p2Wins

	play(playerOne, playerTwo, mode)

}

//Play the tennis match and record the results
func play(p1 *Player, p2 *Player,  mode string) {

	if(strings.EqualFold(p1.playerScoreName, "Advantage") && strings.EqualFold(p2.playerScoreName, "Advantage")) {
		fmt.Println("Only one player may be at advantage...Stopping execution")
		os.Exit(3)
	}


	//Draw the table header depending on sets or game
	if(strings.EqualFold(mode, "game")) {
		drawTable(p1.playerScoreName, p2.playerScoreName)
	} else {
		drawSetTable(p1.playerScoreName, p2.playerScoreName, p1.playerWins, p2.playerWins)
	}

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			text    := scanner.Text()

			//Player One scores
			if(strings.EqualFold(text, p1.playerName + " Scores!")) {
				if(strings.EqualFold(mode, "game")) {
					playerOneScore(p1, p2)
				} else {
					//Game mode is sets
					playerOneScoreSet(p1, p2)
				}
			}

			//Player Two Scores
			if(strings.EqualFold(text, p2.playerName + " Scores!")) {

				if(strings.EqualFold(mode, "game")) {
					playerTwoScores(p1, p2)
				} else {
					//Game mode is sets
					playerTwoScoreSet(p1, p2)
				}
			}

			//Invalid input
			if(!strings.EqualFold(text, p1.playerName + " Scores!") && !strings.EqualFold(text, p2.playerName + " Scores!")) {
				fmt.Println("Invalid Input")
			}
		}
}



func playerOneScoreSet(p1 *Player, p2 *Player) {
	//Player 1 & 2's scores before the increase
	playerOneCurrentScore := p1.playerScoreName
	playerTwoCurrentScore := p2.playerScoreName

	//Is it a deuce?
	deuce := false
	advantage := false
	advantageWin := false

	//Player 1's score & value after the increase
	playerOneNewScoreName, playerOneNewScoreValue := calculateNewScore(playerOneCurrentScore)

	//Player 1 current score is thirty and opponents score is 40 so total is going to be 40-40
	if(strings.EqualFold(playerOneCurrentScore, "Thirty") && strings.EqualFold(playerTwoCurrentScore, "Forty")) {
		deuce = true
		addSetTableRow(p1.playerWins, p2.playerWins, p1.playerScoreName, playerTwoCurrentScore, p1.playerName + " Scores!", "Deuce!")
	}

	//Advantages: p1 & 2 scores are both in deuce
	if(strings.EqualFold(playerOneCurrentScore, "Forty") && strings.EqualFold(playerTwoCurrentScore, "Forty")) {
		advantage = true
		addSetTableRow(p1.playerWins, p2.playerWins, playerOneCurrentScore, playerTwoCurrentScore, p1.playerName + " Scores!", "Advantage " + p1.playerName)

		p1.playerScoreName = "Advantage"
	}

	//if opposite player has the advantage but we won this round
	if(strings.EqualFold(playerTwoCurrentScore, "Advantage")) {
		//we go back to deuce
		deuce = true
		addSetTableRow(p1.playerWins, p2.playerWins, "", "Advantage",  p1.playerName + " Scores!", "Deuce!")

		p1.playerScoreName = "Forty"
		p2.playerScoreName = "Forty"
	}

	//Player 1 already has the advantage this time he wins
	if(strings.EqualFold(playerOneCurrentScore, "Advantage")) {
		advantageWin = true
		addSetTableRow(p1.playerWins, p2.playerWins, "Advantage", "", p1.playerName + " Scores!", p1.playerName + " Wins!")

		p1.playerWins += 1
		p1.playerScoreName = "Love"
		p1.playerScoreValue = 0

		p2.playerScoreName = "Love"
		p2.playerScoreValue = 0

	}

	//Player 1 wins
	if playerOneNewScoreValue >= scoreNameToValue(p2.playerScoreName) + 2 && strings.EqualFold(playerOneNewScoreName, "Forty") {
		//If the player has won 5 times previously (6 including this win)
		if(p1.playerWins == 5) {
			addSetTableRow(p1.playerWins + 1, p2.playerWins, playerOneNewScoreName, playerTwoCurrentScore, p1.playerName + " Scores!", p1.playerName + " Wins the game and the set " + strconv.Itoa(p1.playerWins + 1) + "-" + strconv.Itoa(p2.playerWins))
			os.Exit(3)
		} else {
			addSetTableRow(p1.playerWins + 1, p2.playerWins, playerOneNewScoreName, playerTwoCurrentScore, p1.playerName + " Scores!", p1.playerName + " Wins the game!")
		}

		p1.playerWins += 1;
		p1.playerScoreName = "Love"
		p1.playerScoreValue = 0

		p2.playerScoreName = "Love"
		p2.playerScoreValue = 0

	} else {

		if(!deuce && !advantage && !advantageWin) {
			addSetTableRow(p1.playerWins, p2.playerWins, playerOneCurrentScore, playerTwoCurrentScore, p1.playerName + " Scores!", playerOneNewScoreName + "-" + playerTwoCurrentScore)
		}

		p1.playerScoreName = playerOneNewScoreName
		p1.playerScoreValue = playerOneNewScoreValue
	}
}

func playerTwoScoreSet(p1 *Player, p2  *Player) {
	//Player 1 & 2's scores before the increase
	playerOneCurrentScore := p1.playerScoreName
	playerTwoCurrentScore := p2.playerScoreName

	//Is it a deuce?
	deuce := false
	advantage := false
	advantageWin := false

	//Player 2's score & value after the increase
	playerTwoNewScoreName, playerTwoNewScoreValue := calculateNewScore(playerTwoCurrentScore)

	//Player 2 current score is thirty and opponents score is 40 so total is going to be 40-40
	if(strings.EqualFold(playerTwoCurrentScore, "Thirty") && strings.EqualFold(playerOneCurrentScore, "Forty")) {
		deuce = true
		addSetTableRow(p1.playerWins, p2.playerWins, p1.playerScoreName, playerTwoNewScoreName, p2.playerName + " Scores!", "Deuce!")
	}

	//Advantages: p1 & 2 scores are both in deuce
	if(strings.EqualFold(playerOneCurrentScore, "Forty") && strings.EqualFold(playerTwoCurrentScore, "Forty")) {
		advantage = true
		addSetTableRow(p1.playerWins, p2.playerWins, playerOneCurrentScore, playerTwoCurrentScore, p2.playerName + " Scores!", "Advantage " + p2.playerName)

		p2.playerScoreName = "Advantage"
	}

	//Player 2 already has the advantage this time he wins
	if(strings.EqualFold(playerTwoCurrentScore, "Advantage")) {
		advantageWin = true
		addSetTableRow(p1.playerWins, p2.playerWins, "", "Advantage", p2.playerName + " Scores!", p2.playerName + " Wins!")

		p2.playerWins += 1
		p2.playerScoreName = "Love"
		p2.playerScoreValue = 0

		p1.playerScoreName = "Love"
		p1.playerScoreValue = 0

	}

	//if opposite player has the advantage but we won this round
	if(strings.EqualFold(playerOneCurrentScore, "Advantage")) {
		//we go back to deuce
		deuce = true

		addSetTableRow(p1.playerWins, p2.playerWins, "Advantage", "",  p2.playerName + " Scores!", "Deuce!")

		p2.playerScoreName = "Forty"
		p1.playerScoreName = "Forty"
	}


	//Player 2 wins
	if playerTwoNewScoreValue >= scoreNameToValue(p1.playerScoreName) + 2 && strings.EqualFold(playerTwoNewScoreName, "Forty") {
		//If the player has won 5 times previously (6 including this win)
		if(p2.playerWins == 5) {
			addSetTableRow(p1.playerWins, p2.playerWins + 1, playerTwoNewScoreName, playerTwoCurrentScore, p2.playerName + " Scores!", p2.playerName + " Wins the game and the set " + strconv.Itoa(p1.playerWins) + "-" + strconv.Itoa(p2.playerWins + 1))
			os.Exit(3)
		} else {
			addSetTableRow(p1.playerWins, p2.playerWins + 1, playerTwoNewScoreName, playerTwoCurrentScore, p1.playerName + " Scores!", p2.playerName + " Wins the game!")
		}

		//Add points and reset for new match
		p2.playerWins += 1;
		p2.playerScoreName = "Love"
		p2.playerScoreValue = 0

		p1.playerScoreName = "Love"
		p1.playerScoreValue = 0

	} else {

		if(!deuce && !advantage && !advantageWin) {

			addSetTableRow(p1.playerWins, p2.playerWins, playerOneCurrentScore, playerTwoCurrentScore, p2.playerName + " Scores!", playerOneCurrentScore + "-" + playerTwoNewScoreName)
		}

		p2.playerScoreName = playerTwoNewScoreName
		p2.playerScoreValue = playerTwoNewScoreValue
	}


}

/*
 * Handles player one scoring in game mode
 */
func playerOneScore(p1 *Player, p2 *Player)  {
	//Player 1 & 2's scores before the increase
	playerOneCurrentScore := p1.playerScoreName
	playerTwoCurrentScore := p2.playerScoreName

	//Is it a deuce?
	deuce := false
	advantage := false
	advantageWin := false

	//Player 1's score & value after the increase
	playerOneNewScoreName, playerOneNewScoreValue := calculateNewScore(playerOneCurrentScore)


	//Player 1 current score is thirty and opponents score is 40 so total is going to be 40-40
	if(strings.EqualFold(playerOneCurrentScore, "Thirty") && strings.EqualFold(playerTwoCurrentScore, "Forty")) {
		deuce = true
		addTableRow(p1.playerScoreName, playerTwoCurrentScore, p1.playerName + " Scores!", "Deuce!")
	}

	//Advantages: p1 & 2 scores are both in deuce
	if(strings.EqualFold(playerOneCurrentScore, "Forty") && strings.EqualFold(playerTwoCurrentScore, "Forty")) {
		advantage = true
		addTableRow(playerOneCurrentScore, playerTwoCurrentScore, p1.playerName + " Scores!", "Advantage " + p1.playerName)

		p1.playerScoreName = "Advantage"
	}

	//if opposite player has the advantage but we won this round
	if(strings.EqualFold(playerTwoCurrentScore, "Advantage")) {
		//we go back to deuce
		deuce = true
		addTableRow("", "Advantage",  p1.playerName + " Scores!", "Deuce!")

		p1.playerScoreName = "Forty"
		p2.playerScoreName = "Forty"
	}

	//Player 1 already has the advantage this time he wins
	if(strings.EqualFold(playerOneCurrentScore, "Advantage")) {
		advantageWin = true
		addTableRow("Advantage", "", p1.playerName + " Scores!", p1.playerName + " Wins!")

		p1.playerWins += 1
		p1.playerScoreName = "Love"
		p1.playerScoreValue = 0

		p2.playerScoreName = "Love"
		p2.playerScoreValue = 0

	}

	//Player 1 wins
	if playerOneNewScoreValue >= scoreNameToValue(p2.playerScoreName) + 2 && strings.EqualFold(playerOneNewScoreName, "Forty") {
		addTableRow(playerOneNewScoreName, playerTwoCurrentScore, p1.playerName + " Scores!", p1.playerName + " Wins!")

		p1.playerWins += 1;
		p1.playerScoreName = "Love"
		p1.playerScoreValue = 0

		p2.playerScoreName = "Love"
		p2.playerScoreValue = 0

	} else {

		if(!deuce && !advantage && !advantageWin) {
			addTableRow(playerOneCurrentScore, playerTwoCurrentScore, p1.playerName + " Scores!", playerOneNewScoreName + "-" + playerTwoCurrentScore)
		}

		p1.playerScoreName = playerOneNewScoreName
		p1.playerScoreValue = playerOneNewScoreValue
	}

}

/**
 * Handles player 2 scoring
 */
func playerTwoScores(p1 *Player, p2 *Player) {
	//Player 1 & 2's scores before the increase
	playerOneCurrentScore := p1.playerScoreName
	playerTwoCurrentScore := p2.playerScoreName

	//Is it a deuce?
	deuce := false
	advantage := false
	advantageWin := false

	//Player 2's score & value after the increase
	playerTwoNewScoreName, playerTwoNewScoreValue := calculateNewScore(playerTwoCurrentScore)


	//Player 2 current score is thirty and opponents score is 40 so total is going to be 40-40
	if(strings.EqualFold(playerTwoCurrentScore, "Thirty") && strings.EqualFold(playerOneCurrentScore, "Forty")) {
		deuce = true
		addTableRow(p1.playerScoreName, playerTwoNewScoreName, p2.playerName + " Scores!", "Deuce!")
	}

	//Advantages: p1 & 2 scores are both in deuce
	if(strings.EqualFold(playerOneCurrentScore, "Forty") && strings.EqualFold(playerTwoCurrentScore, "Forty")) {
		advantage = true
		addTableRow(playerOneCurrentScore, playerTwoCurrentScore, p2.playerName + " Scores!", "Advantage " + p2.playerName)

		p2.playerScoreName = "Advantage"
	}

	//Player 2 already has the advantage this time he wins
	if(strings.EqualFold(playerTwoCurrentScore, "Advantage")) {
		advantageWin = true
		addTableRow("", "Advantage", p2.playerName + " Scores!", p2.playerName + " Wins!")

		p2.playerWins += 1
		p2.playerScoreName = "Love"
		p2.playerScoreValue = 0

		p1.playerScoreName = "Love"
		p1.playerScoreValue = 0

	}

	//if opposite player has the advantage but we won this round
	if(strings.EqualFold(playerOneCurrentScore, "Advantage")) {
		//we go back to deuce
		deuce = true

		addTableRow("Advantage", "",  p2.playerName + " Scores!", "Deuce!")

		p2.playerScoreName = "Forty"
		p1.playerScoreName = "Forty"
	}

	//Player 2 wins
	if playerTwoNewScoreValue >= scoreNameToValue(p1.playerScoreName) + 2 && strings.EqualFold(playerTwoNewScoreName, "Forty") {
		addTableRow(playerOneCurrentScore, playerTwoCurrentScore, p2.playerName + " Scores!", p2.playerName + " Wins!")

		p2.playerWins += 1;
		p2.playerScoreName = "Love"
		p2.playerScoreValue = 0

		p1.playerScoreName = "Love"
		p1.playerScoreValue = 0

	} else {

		if(!deuce && !advantage && !advantageWin) {

			addTableRow(playerOneCurrentScore, playerTwoCurrentScore, p2.playerName + " Scores!", playerOneCurrentScore + "-" + playerTwoNewScoreName)
		}

		p2.playerScoreName = playerTwoNewScoreName
		p2.playerScoreValue = playerTwoNewScoreValue
	}

}

	/**
	 * Calculates increasing a score after a round has been won
	 */
	func calculateNewScore(oldScore string) (string, int) {
		if(strings.EqualFold(oldScore, "Love")) {
			return "Fifteen", 1
		}
		if(strings.EqualFold(oldScore, "Fifteen")) {
			return "Thirty", 2
		}
		if(strings.EqualFold(oldScore, "Thirty")) {
			return "Forty", 3
		}
		if(strings.EqualFold(oldScore, "Forty")) {
			return "Advantage", 4
		}

		return "Love", 0
	}

	/**
	 * Converts a scores name to its respective value
	 */
	func scoreNameToValue(name string) (int) {
		if(strings.EqualFold(name, "Love")) {
			return 0;
		}

		if(strings.EqualFold(name, "Fifteen")) {
			return 1;
		}

		if(strings.EqualFold(name, "Thirty")) {
			return 2;
		}

		if(strings.EqualFold(name, "Forty")) {
			return 3
		}

		//Else return 4 for advantage
		return 4
	}

	func drawTable(player1Score string, player2Score string) {
		fmt.Println("-----------------------------------------------------------------")
		fmt.Printf("|%12s|%12s|%18s|%15s\n", "player1score", "player2score", "scoreEvent", "result")
		fmt.Println("-----------------------------------------------------------------")
		fmt.Printf("|%12s|%12s|%18s|%19s|\n", player1Score, player2Score, "Game Start", player1Score + "-" + player2Score)
	}

	func addTableRow(p1Score string, p2Score string, scoreEvent string, result string) {
		fmt.Printf("|%12s|%12s|%18s|%19s|\n", p1Score, p2Score, scoreEvent, result)
	}

	func drawSetTable(player1Score string, player2Score string, player1Wins int, player2wins int) {
		fmt.Println("----------------------------------------------------------------------------------------------------------")
		fmt.Printf("|%12s|%12s|%12s|%12s|%18s|%15s\n","player1wins", "player2wins", "player1score", "player2score", "scoreEvent", "result")
		fmt.Println("----------------------------------------------------------------------------------------------------------")
		fmt.Printf("|%12d|%12d|%12s|%12s|%18s|%15s\n", player1Wins, player2wins, player1Score, player2Score, "Game Start", "")
	}

	func addSetTableRow(p1Wins int, p2Wins int, p1Score string, p2Score string, scoreEvent string, result string) {
		fmt.Printf("|%12d|%12d|%12s|%12s|%18s|%15s\n", p1Wins, p2Wins, p1Score, p2Score, scoreEvent, result)

	}
