package main

import (
	"database/sql"
	"europe/handlers"
	"fmt"
	"math/rand"
	"net"
	"time"

	//sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
	// build server
	listener, err := net.Listen("tcp", ":4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")

	//build BD
	connStr := "user=mydb1 password=123456 dbname=mydb1 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB - ok")
	defer db.Close()

	//launch a loop
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
		target := randomCountry(dict)
		conn.Write([]byte(target))

		//4 сервер получает ответ и сравнивает и отправляет ок или нет

		answer, err := handlers.Recieve(conn)
		if err != nil {
			return
		}
		caser := cases.Title(language.English)
		if caser.String(answer) == caser.String(dict[target]) {
			conn.Write([]byte("right "))
		} else {
			conn.Write([]byte(fmt.Sprint("wrong. ", dict[target], " is right ")))
		}
		time.Sleep(time.Millisecond * 5)

	}
}

func randomCountry(m map[string]string) string {
	k := rand.Intn(len(m))
	i := 0
	for x := range m {
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
