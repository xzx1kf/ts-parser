package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type booking struct {
	court		int
	days		int
	hour		int
	min		int
	timeslot	int
	playerA		string
	playerB		string
	booking_link	string
	booked		bool
}

func ExampleScrape() {
	res, err := http.Get("http://tynemouth-squash.herokuapp.com/bookings?day=0")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Available Courts")
	doc.Find(".booking div.book a.booking_link").Each(func(i int, s *goquery.Selection) {
		bl, exists := s.Attr("href")
		if exists {
			fmt.Println("Booking Link: "+bl)
		}
	})

	fmt.Println("Available to Rebook")
	doc.Find(".booking div.cancelled a").Each(func(i int, s *goquery.Selection) {
		fmt.Println("Call the club to book this court")
	})

	fmt.Println("Booked Courts")
	doc.Find(".booking div.booked a").Each(func(i int, s *goquery.Selection) {
		bl, exists := s.Attr("href")
		if exists {
			fmt.Println("Booking Link: "+bl+" "+s.Text())
		}
	})
	// TODO: add some logic around if playerA and playerB are known
}

func main() {
	ExampleScrape()
}
