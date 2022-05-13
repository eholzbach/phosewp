package plugins

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/eholzbach/phosewp/config"
	irc "github.com/thoj/go-ircevent"

	"math/rand"
	"strconv"
	"strings"
	"time"

	// sqlite support
	_ "github.com/mattn/go-sqlite3"
)

// Quote saves and recalls shame
func quote(conn *irc.Connection, r string, event *irc.Event, conf config.Vars) {
	var reply string

	query := strings.Split(event.Message(), " ")

	db, err := sql.Open("sqlite3", conf.Quotes)

	if err != nil {
		log.Println(err)
		return
	}

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS quotes (quote TEXT)")

	if err != nil {
		log.Println(err)
		return
	}

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

	conn.Privmsg(r, reply)
}

// addQuote adds shame to the shame db
func addQuote(db *sql.DB, quote string) string {
	var a string

	statement, err := db.Prepare("INSERT INTO quotes VALUES (?)")

	if err != nil {
		log.Println(err)
	}

	add, err := statement.Exec(quote)

	if err != nil {
		log.Println(err)
	}

	id, err := add.LastInsertId()

	if err != nil {
		a = "error"
	}

	a = strconv.FormatInt(id, 10)

	return a
}

// getQuote retrieves shame from the shame db
func getQuote(db *sql.DB, id int) string {
	var a string
	var q string

	if id > 0 {
		a = dbQuery(db, id)
	} else {
		row, err := db.Query("SELECT Count(*) FROM quotes")

		if err != nil {
			log.Println(err)
		}

		for row.Next() {
			if err := row.Scan(&q); err != nil {
				log.Println(err)
			}

			a = q
		}

		rand.Seed(time.Now().UnixNano())
		c, _ := strconv.Atoi(a)

		if c >= 2 {
			b := rand.Intn(c-1) + 1
			a = dbQuery(db, b)
		} else if c == 1 {
			a = dbQuery(db, 1)
		} else {
			a = "no quotes"
		}
	}

	return a
}

// dbQuery queries the shame db
func dbQuery(db *sql.DB, id int) string {
	var response string
	var q string

	row, err := db.Query(fmt.Sprintf("SELECT quote FROM quotes WHERE ROWID = %d", id))

	if err != nil {
		log.Println(err)
	}

	for row.Next() {
		err = row.Scan(&q)

		if err != nil {
			log.Println(err)
		}

		response = q
	}

	return response
}
