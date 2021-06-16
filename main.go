package main
import (
	d "github.com/thestupendous/snake-game/definitions"
	//"./definitions"
	"fmt"
	"time"
	"os"
	"os/exec"
	"strconv"
)

// func gameWon(s d.Snake) bool {
// 	if uint32(len(s.Tail))>= d.M*d.N-1 {
// 		return true
// 	}
// 	return false
// }


func main() {
	won,lost := false, false
	//defining dimensions of board
	d.M,d.N = 20,40
	//defining game speed
	d.TickDelay = 200
	//defining empty board of M,N dimensions
	var board d.MyBoard
	board = make([][]string,d.M)
	for  i:=0;i<int(d.M);i++  {
		board[i] = make([]string,d.N)
	}

	//initialising board
	for i:=0 ; i< int(d.M) ; i++ {
		for j:=0 ; j < int(d.N) ; j++ {
			board[i][j] = string(" ")
		}
	}

	//initialising snake
	var sn d.Snake
	// board[0][4] = "A"
	// fmt.Println(board)
	d.PlaceInitialSnake(&sn,board)
	d.PlaceFood(board)				//for the first time

	fmt.Println(board)


	// fmt.Println("len of snake before starting game ",len(sn.Tail))

	//game progression steps
	d.Score = 0
	d.Dir = d.UserDir["d"]
	d.OldDir = d.UserDir["d"]
	// ch := ""
	// for i:=0;i>0;i++ {
	// 	fmt.Scanln(&ch)
	// 	d.Dir = userDir[ch]
	// 	// fmt.Println("\033[H\033[2J")				//for clearing screen
	// 	updateBoard(board)
	// }

    ch := make(chan string)
    go func(ch chan string) {
        // disable input buffering
        exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
        // do not display entered characters on the screen
        exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
        var b []byte = make([]byte, 1)
        for {
			// time.Sleep(time.Millisecond * 1000)
            os.Stdin.Read(b)
            ch <- string(b)
        }
		if lost {
			close(ch)
		}
    }(ch)

	counter := 0
    for !lost && !won {
        select {
            case stdin, _ := <-ch:
				counter+=1
                fmt.Println("Moves : ",counter,"\nKey pressed: ", stdin)
					//updateBoard(newDir)  //update snake, placefood, lost, won
					d.Dir = d.UserDir[stdin]
					d.ChecknCorrectWrongDir()
					d.UpdateBoard(board,&sn,&won,&lost)
            default:
                fmt.Println("Moves : ",counter)
					d.Dir = d.OldDir
					d.UpdateBoard(board,&sn,&won,&lost)
        }
        time.Sleep(time.Millisecond * time.Duration(d.TickDelay))
		fmt.Println("\033[H\033[2J")			//clear the screen
		fmt.Println(board)						//printing board at every clock tick
		fmt.Println("  Score : " + strconv.Itoa(d.Score))
    }

	if lost {
		fmt.Println("GAME OVER! ",d.GameOverReason)
	}
	if won {
		fmt.Println("Hurray!!, You finished the game. You can try again with more speed of snake :)")
	}
	fmt.Println("your final Score : ",d.Score)


}
