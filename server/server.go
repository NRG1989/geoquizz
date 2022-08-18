package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"

	//sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
)

var dict = map[string]string{
	"Albania":                "Tirana",
	"Andorra":                "Andorra la Vella",
	"Austria":                "Vienna",
	"Belarus":                "Minsk",
	"Belgium":                "Brussels",
	"Bosnia and Herzegovina": "Sarajevo",
	"Bulgaria":               "Sofia",
	"Croatia":                "Zagreb",
	"Czechia":                "Prague",
	"Denmark":                "Copenhagen",
	"Estonia":                "Tallinn",
	"Finland":                "Helsinki",
	"France":                 "Paris",
	"Germany":                "Berlin",
	"Greece":                 "Athens",
	"Hungary":                "Budapest",
	"Iceland":                "Reykjavik",
	"Ireland":                "Dublin",
	"Italy":                  "Rome",
	"Latvia":                 "Riga",
	"Liechtenstein":          "Vaduz",
	"Lithuania":              "Vilnius",
	"Luxembourg":             "Luxembourg",
	"Malta":                  "Valletta",
	"Moldova":                "Chisinau",
	"Monaco":                 "Monaco",
	"Montenegro":             "Podgorica",
	"Netherlands":            "Amsterdam",
	"North Macedonia":        "Skopje",
	"Norway":                 "Oslo",
	"Poland":                 "Warsaw",
	"Portugal":               "Lisbon",
	"Romania":                "Bucharest",
	"Russia":                 "Moscow",
	"San Marino":             "San Marino",
	"Serbia":                 "Belgrade",
	"Slovakia":               "Bratislava",
	"Slovenia":               "Ljubljana",
	"Spain":                  "Madrid",
	"Sweden":                 "Stockholm",
	"Switzerland":            "Bern",
	"Ukraine":                "Kiev",
	"United Kingdom":         "London",
}

func main() {
	listener, err := net.Listen("tcp", ":4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")

	connStr := "user=mydb1 password=123456 dbname=mydb1 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Errorf("the DB doesn't work")
		panic(err)
	}
	fmt.Println("DB - ok")
	defer db.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		//1 сервер загадывает страну и посылает
		target := pickCountry(dict)
		conn.Write([]byte(target))

		//4 сервер получает ответ и сравнивает и отправляет ок или нет
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		answer := string(input[0 : n-1])

		if strings.Title(answer) == strings.Title(dict[target]) {
			conn.Write([]byte("right"))
		} else {
			conn.Write([]byte(fmt.Sprint("wrong. ", dict[target], " is right")))
		}
		time.Sleep(time.Millisecond * 5)

	}
}

func pickCountry(m map[string]string) string {
	k := rand.Intn(len(m))
	i := 0
	for x, _ := range m {
		if i == k {
			return x
		}
		i++
	}
	panic("unreachable")
}

// func GuessCapital() (string, error) {
// 	qb := sq.
// 		Select(
// 			"capital",
// 		).
// 		From("europe.general").
// 		OrderByRandom().
// 		Limit(1)

// }
