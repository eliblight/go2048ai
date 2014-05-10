package main 

import (
    "flag"
    "fmt"
    "math/rand"
    "bufio"
    "os"
    "github.com/liblight/go2048ai/client"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")
var h *bool = flag.Bool("h", false, "Help")
var help *bool = flag.Bool("help", false, "Help")
var seed *int64 = flag.Int64("seed", 0 , "Seed random function.")

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var strToDir map[byte]client.Direction = map[byte]client.Direction{
	'w' : client.Up,
	'a' : client.Left,
	's' : client.Down,
	'd' : client.Right}

func getMove() client.Direction {
	for {
		fmt.Print("move: ");
		input, _ := reader.ReadString('\n')
		if dir , ok := strToDir[input[0]] ; ok {
			return dir
		}
		fmt.Println("unkown move use wasd keys and enter")
	}
	panic("never get here")
}

func getUserMove(board client.Board, score int) client.Direction {
	fmt.Println("Score: ",score)
	board.Print()
	return getMove()
}


func main() {
    flag.Parse() // Scan the arguments list 

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION)
        return
    }
    
    if *h || *help {
    	fmt.Println("wasd: up,left, down, right and hit enter")
    	return
    }
    
   	fmt.Println("Seed:", seed)
   	rand.Seed(*seed)

	board, score := client.NewGame(getUserMove)
	fmt.Println("------ FIN -------")
	fmt.Println("Score: ",score)
	board.Print()
}
