package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/garyburd/redigo/redis"
)

func handler(w http.ResponseWriter, r *http.Request) {
	host := os.Getenv("HOSTNAME")
	fmt.Fprintf(w, "<p>Hi there, from <b>%s</b>!", host)
	c, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()
	c.Do("INCR", host)
	keys, _ := redis.Strings(c.Do("KEYS", "*"))
	fmt.Fprintf(w, "<hr/>")
	fmt.Fprintf(w, "<table style='width: 10em; border-collapse: collapse;'><tr><th style='border: 2px dotted green;'>Container</th><th style='padding: 5px; border: 2px dotted green;'>#</th></tr>")
	for _, key := range keys {
		value, _ := redis.Int(c.Do("GET", key))
		fmt.Fprintf(w, "<tr><td style='border: 1px solid green;'>%s</td>", key)
		fmt.Fprintf(w, "<td style='border: 1px solid green; text-align: center;'>%d</td></tr>", value)
	}
	fmt.Fprintf(w, "</table>")
}

func main() {
	http.HandleFunc("/demo", handler)
	log.Println("counter app started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
