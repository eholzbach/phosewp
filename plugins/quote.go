/* save and fetch quotes
   need to loop back and fix where the db lives per os/config
*/

package plugins

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/thoj/go-ircevent"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Quote(conn *irc.Connection) {
	conn.AddCallback("PRIVMSG", func(event *irc.Event) {
		if strings.HasPrefix(event.Message(), "!quote") == true {

			var replyto string
			var reply string

			if strings.HasPrefix(event.Arguments[0], "#") {
				replyto = event.Arguments[0]
			} else {
				replyto = event.Nick
			}

			query := strings.Split(event.Message(), " ")

			db, err := sql.Open("sqlite3", "./quotes.db")
			checkError(err)

			statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS quotes (quote TEXT)")
			checkError(err)
			statement.Exec()
			if len(query) > 1 {
				if query[1] == "add" && len(query) > 2 {
					a := strings.TrimPrefix(event.Message(), "!quote add ")
					b := addQuote(db, a)
					reply = fmt.Sprintf("added, id: %s", b)
				} else if query[1] != "add" && len(query) == 2 {
					i, _ := strconv.Atoi(query[1])
					reply = getQuote(db, i)
				}
			} else {
				reply = getQuote(db, -0)
			}

			conn.Privmsg(replyto, reply)
		}

	})
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func addQuote(db *sql.DB, quote string) string {
	var a string
	statement, err := db.Prepare("INSERT INTO quotes VALUES (?)")
	checkError(err)
	add, err := statement.Exec(quote)
	checkError(err)
	id, err := add.LastInsertId()
	if err != nil {
		a = "error"
	}
	a = strconv.FormatInt(id, 10)
	return a
}

func getQuote(db *sql.DB, id int) string {
	var a string
	var q string
	if id > 0 {
		a = dbQuery(db, id)
	} else {
		row, err := db.Query("SELECT Count(*) FROM quotes")
		checkError(err)
		for row.Next() {
			err := row.Scan(&q)
			checkError(err)
			a = q
		}
		rand.Seed(time.Now().Unix())
		c, _ := strconv.Atoi(a)
		if c >= 2 {
			b := rand.Intn(c-1) + 1
			a = dbQuery(db, b)
			fmt.Println("random")
		} else if c == 1 {
			fmt.Println("one")
			a = dbQuery(db, 1)
		} else {
			a = "no quotes"
		}
	}
	return a
}

func dbQuery(db *sql.DB, id int) string {
	var response string
	var q string
	a := fmt.Sprintf("SELECT quote FROM quotes WHERE ROWID = %d", id)
	row, err := db.Query(a)
	checkError(err)
	for row.Next() {
		err = row.Scan(&q)
		checkError(err)
		response = q
	}
	return response
}
