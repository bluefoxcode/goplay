package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

type Ping struct {
	Ip      string
	Visited time.Time
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	// should be handled in its own routes world, but for now this will work as I learn.

	router := gin.Default()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/**/*")
	router.Static("/static", "static")

	router.GET("/", mainAppFunc)
	router.GET("/pings", pingsFunc)

	router.Run(":" + port)

}

func mainAppFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/index.tmpl.html", nil)
}

func pingsFunc(c *gin.Context) {
	var err error

	ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)

	if err != nil {
		fmt.Printf("Error getting your IP: %q", err)
		return
	}

	// create table if it isn't there already.
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS pings (ip text, visited timestamp)"); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error storing your IP: %q", err))
		return
	}

	// insert IP
	if _, err := db.Exec("INSERT INTO pings values($1, now())", ip); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error storing your IP: %q", err))
		return
	}

	rows, err := db.Query("SELECT ip,visited FROM pings order by visited DESC limit 10")
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprintf("Error reading pings: %q", err))
		return
	}

	defer rows.Close()

	pings := make([]*Ping, 0)

	for rows.Next() {
		ping := new(Ping)

		if err := rows.Scan(&ping.Ip, &ping.Visited); err != nil {
			c.String(http.StatusInternalServerError,
				fmt.Sprintf("Error scanning pings: %q", err))
			return
		}

		pings = append(pings, ping)
	}

	c.HTML(http.StatusOK, "pages/pings.tmpl.html", gin.H{"pings": pings})

}
