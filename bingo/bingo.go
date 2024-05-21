package bingo

import (
	"encoding/json"
	"math/rand/v2"
)

/*
0 - 60 | 24 => 60

row 1 (B) | 5 | range(01,12)
row 2 (I) | 5 | range(13,25)
row 3 (N) | 4 | range(26,38)
row 4 (G) | 5 | range(39,51)
row 5 (O) | 5 | range(52,60)

json -> {

	[ id_table - auto_increment : {
	    B : [1,2,3,4,5]
	    I : [13,14,15,16,17]
	    N : [26,27,null,28,29]
	    G : [39,40,41,42,43]
	    O : [52,53,54,55,56]
	}, ...]

}

Logic => {
	lobby : {

	},
	User Create Game Session -> Auto Generate Token(CODE) Room
	GameSession{
		config: => Create Room {
			id: int(id), autogenerate || not necessary in Front req
			roomName: string, => if empty -> autogenerate
			mode: normal | fast | insane -> string,
			waitTime: 10 | 15 | 20 | personalized - ALL in seconds -> int,
			autoStart: boolean,
			playersQuantity:  min(2) - max(100),
			privateSession: boolean,
			privatePassword: string - length.min(8) | length.min(80),
			hasAwards: boolean, |===| merely illustrative
			awardItems : [
				{
					name : string,
					place : int
				}
			],
		}
		config - API needs to do -> {
			create room with cfg in DB,
			return {
				roomID : int(id),
				ownerNick : string,
				ownerToken : string(jwt),
				config : {
					cfg | the same config info
				}
			}
			return roomID and in Front End router to page/roomID
		},
		roomMoment's : Wait | Game | End
		room session -> {
			User's Control: => | IN SITE {
				startGame  : boolean = false, | Button | only owner | Wait
				banPlayer | private | WAIT && GAME(VOTE) : {
					roomID : int(id),
					NickName : string,
					ownerToken: string(jwt),
				}, => in a public rooms this func will not be possible to use
				chat : WAIT && GAME && END {
					NickName : string, -> maybe Trade for a Token
					message: string,
					roomID : int(id)
				},
				quit: WAIT && GAME && END {
					roomID : int(id),
					NickName : string | Token
				},
				changePasswordRoom : WAIT {
					ownerToken : string(jwt),
					roomID: int(id),
					newPassword : string - length.min(8) | length.min(80),
				}
			},
			In-Game : {

			}
		}
	}
}
*/

type BingoData struct {
	B []int  `json:"B"`
	I []int  `json:"I"`
	N []*int `json:"N"`
	G []int  `json:"G"`
	O []int  `json:"O"`
}

type BingoGame struct {
	ID    int       `json:"id"`
	Table BingoData `json:"table"`
}

type BingoJSON struct {
	ID int    `json:"id"`
	B  []int  `json:"B"`
	I  []int  `json:"I"`
	N  []*int `json:"N"`
	G  []int  `json:"G"`
	O  []int  `json:"O"`
}

func GenerateColumn(rangeStart, rangeEnd, count int) []int {
	numbers := rand.Perm(rangeEnd - rangeStart + 1)
	numbers = numbers[:count]
	for i := range numbers {
		numbers[i] += rangeStart
	}
	return numbers
}

func GenerateColumnWithNull(rangeStart, rangeEnd, count int) []*int {
	numbers := rand.Perm(rangeEnd - rangeStart + 1)
	numbers = numbers[:count]
	column := make([]*int, count)
	nullIndex := count / 2

	for i := range numbers {
		if i == nullIndex {
			column[i] = nil
		} else {
			value := numbers[i] + rangeStart
			column[i] = &value
		}
	}
	return column
}

func Bingo() BingoData {
	return BingoData{
		B: GenerateColumn(1, 12, 5),
		I: GenerateColumn(13, 25, 5),
		N: GenerateColumnWithNull(26, 38, 5),
		G: GenerateColumn(39, 51, 5),
		O: GenerateColumn(52, 60, 5),
	}
}

func ConvertToJSON(games []BingoGame) ([]byte, error) {
	var bingoJSON []BingoJSON

	for _, game := range games {
		jsonData := BingoJSON{
			ID: game.ID,
			B:  game.Table.B,
			I:  game.Table.I,
			N:  game.Table.N,
			G:  game.Table.G,
			O:  game.Table.O,
		}
		bingoJSON = append(bingoJSON, jsonData)
	}

	data, err := json.Marshal(bingoJSON)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func CreateGame() []BingoGame {
	var games []BingoGame

	for i := 0; i < 5; i++ {
		bingoTable := Bingo()

		game := BingoGame{
			ID:    i + 1,
			Table: bingoTable,
		}

		games = append(games, game)
	}

	return games
}
