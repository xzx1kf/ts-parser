package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
	"github.com/PuerkitoBio/goquery"
)

type Bookings struct {
	Date	int
	Slots	[]BookingSlot
}

type BookingSlot struct {
	Court		string	`json:"court"`
	Days		string	`json:"days"`
	Hour		string	`json;"hour"`
	Min		string	`json:"min"`
	Timeslot	string	`json:"timeslot"`
	PlayerA		string	`json:"playerA"`
	PlayerB		string	`json:"playerB"`
	Booking_link	string	`json:"bookingLink"`
	Booked		bool	`json:"booked"`
}

func parseBookingUrl(link string) BookingSlot {
	s := link[14:]
	s = strings.Replace(s, "&amp", "", -1)
	m, err := url.ParseQuery(s)
	if err != nil {
		log.Fatal(err)
	}

	bs := BookingSlot {
		Court:		m["court"][0],
		Days:		m["days"][0],
		Hour:		m["hour"][0],
		Min:		m["min"][0],
		Timeslot:	m["timeSlot"][0],
		Booked:		false,
	}

	return bs
}

func GetBookings(w http.ResponseWriter, r *http.Request) {

	slots := []BookingSlot{}

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
			bs := parseBookingUrl(bl)

			slots = append(slots, bs)

			//fmt.Println(bs)
		}
	})

	//fmt.Printf("len=%d cap=%d %v\n", len(slots), cap(slots), slots)

	bookings, err := json.Marshal(slots)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Fprintln(w, string(bookings))
	/*
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
	*/
	// TODO: add some logic around if playerA and playerB are known
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/bookings", GetBookings)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}
