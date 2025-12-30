package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// 4. Custom JSON marshaling with MarshalJSON
type CustomDate time.Time

func (cd CustomDate) MarshalJSON() ([]byte, error) {
	// Custom format: YYYY-MM-DD
	formatted := time.Time(cd).Format("2006-01-02")
	return json.Marshal(formatted)
}

func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}
	*cd = CustomDate(t)
	return nil
}

type Event struct {
	Title string     `json:"title"`
	Date  CustomDate `json:"date"` // Uses custom marshaling
}

func main() {
	event := Event{
		Title: "Go Conference",
		Date:  CustomDate(time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)),
	}

	// Uses custom MarshalJSON for Date field
	data, _ := json.MarshalIndent(event, "", "  ")
	fmt.Println("Event with custom date format:\n", string(data))

	// Unmarshal with custom UnmarshalJSON
	jsonStr := `{"title":"Hackathon","date":"2024-12-25"}`
	var newEvent Event
	err := json.Unmarshal([]byte(jsonStr), &newEvent)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}
	fmt.Printf("Unmarshaled Event: %+v\n", newEvent)
	fmt.Printf("Date as time.Time: %v\n", time.Time(newEvent.Date))
}
