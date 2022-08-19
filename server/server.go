package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"

	"europe/handlers"
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
		//1 Server takes random country and sends to Client
		target := randomCountry(dict)
		conn.Write([]byte(target))

		//4 Server recieves the answer, compares it and sends the message if it is right or not

		answer, err := handlers.Recieve(conn)
		if err != nil {
			return
		}
		if strings.EqualFold(answer, dict[target]) {
			conn.Write([]byte("right "))
		} else {
			conn.Write([]byte(fmt.Sprint("wrong. ", dict[target], " is right ")))
		}
		time.Sleep(time.Millisecond)

	}
}

//take a random country
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
