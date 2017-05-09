# Tennis
Tennis Command Line Program Written in Go for Keeping track of scores

### Steps to Use this Start the game

Run the file "Main.go" to start the console program

You can input the following settings with one setting per line before the program starts

- "--p1Name <PLAYER_NAME>"
- "--p1Score <PLAYER_SCORE>"
- "--p2Name <PLAYER_NAME>"
- "--p2Score <PLAYER_SCORE>"
- "--p1Wins <PLAYER_WINS>"
- "--p2Wins <PLAYER_WINS>"
- "--mode <MODE_NAME>"

<PLAYER_NAME> Can be any name of your choosing
<PLAYER_SCORE> Can be either "Love", "Fifteen", "Thirty", "Forty" or "Advantage" All other text is invalid
<PLAYER_WINS> Must be an integer between 0 and 6. This can include 0 and 6.
<MODE_NAME> Must be either "set" or "game". Set mode shows both player's wins and "game" mode shows their scores only.

 When you are ready to start the game type ***"play"*** if no settings are specified default values will be used. Default player
 names are "Player One" and "Player Two" while default scores are "Love"

 ### How to Play

 - After you have typed "play" you are ready to start the game
 - Type "<PLAYER_NAME> scores!" to determine which player scores where <PLAYER_NAME> is either the specified or default player name
 - (Note) If an invalid player name is entered the program will say "Invalid Input" ensure all spaces and spelling is correct!
 - After a player scores output will be sent to the console showing the up to date results of the match.

 - Have fun!




