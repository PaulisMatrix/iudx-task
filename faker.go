package main

import (
	"datakaveri/events"
	"encoding/json"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var pageURLS = []string{"https://gofakeme.com", "https://faker.com", "https://faker-me.com", "https://faku-me.com", "https://imafaker.com", "https://fakirhume.com"}

var buttonElements = []string{"button-class1", "button-class2", "button-class3", "button-class4", "button-class5", "button-class6", "button-class7"}

var searchQueries = []string{"kids wear", "mens wear", "hoodies for women and men", "unisex colognes", "type C heaphones", "headphones connect", "unisex sneakers", "surprise me", "new offers on clothing"}

var paymentModes = []string{"UPI", "CASH", "CREDITCARD", "DEBITCARD", "CASHONDELIVERY", "PAYLATER", "WALLET"}

var userIDs = []string{"gwm1cH", "C8pi0O", "lWZbEZ", "SnKjjv", "ON8Wck", "mkzOnI"}

// Generate fake events depending on input type mentioned?,
// return marshalled bytes to directly produce to kafka.
func FakerProduce(eventType events.EventType) []byte {
	switch eventType {

	case events.PageViewEvent:
		data, _ := json.Marshal(fakePageViewEvent())
		return data

	case events.ButtonClickEvent:
		data, _ := json.Marshal(fakeButtonClickEvent())
		return data

	case events.SearchEvent:
		data, _ := json.Marshal(fakeSearchEvent())
		return data

	case events.PurchaseEvent:
		data, _ := json.Marshal(fakePurchaseEvent())
		return data

	case events.AdvertEvent:
		data, _ := json.Marshal(fakeAdvertEvent())
		return data

	case events.ScrollEvent:
		data, _ := json.Marshal(fakeScrollEvent())
		return data

	default:
		return nil

	}
}

func fakePageViewEvent() *events.Event {
	pageEvent := &events.Event{
		EventType: events.PageViewEvent,
		UserID:    getUserID(),
		CreatedAt: getCreatedAt(),
		Duration:  uint32(getRandomNumberRange(1000, 100)),
		PageURL:   getPageUrl(),
		ProductID: getRandomID(10),
	}
	return pageEvent
}

func fakeButtonClickEvent() *events.Event {
	buttonClickEvent := &events.Event{
		EventType: events.ButtonClickEvent,
		UserID:    getUserID(),
		CreatedAt: getCreatedAt(),
		ProductID: getRandomID(10),
		ClickedTarget: &events.ButtonElement{
			Element: getButtonElement(),
		},
	}
	return buttonClickEvent
}

func fakeSearchEvent() *events.Event {
	searchEvent := &events.Event{
		EventType:   events.SearchEvent,
		UserID:      getUserID(),
		CreatedAt:   getCreatedAt(),
		Duration:    uint32(getRandomNumberRange(1000, 100)),
		PageURL:     getPageUrl(),
		SearchQuery: getSearchQuery(),
	}
	return searchEvent
}

func fakePurchaseEvent() *events.Event {
	purchaseEvent := &events.Event{
		EventType:         events.PurchaseEvent,
		UserID:            getUserID(),
		CreatedAt:         getCreatedAt(),
		Duration:          uint32(getRandomNumberRange(1000, 100)),
		PaymentMode:       getPaymentMode(),
		NumItemsPurchased: uint16(getRandomNumberRange(10, 2)),
	}
	return purchaseEvent
}

func fakeAdvertEvent() *events.Event {
	adVertEvent := &events.Event{
		UserID:        getUserID(),
		EventType:     events.AdvertEvent,
		CreatedAt:     getCreatedAt(),
		Duration:      uint32(getRandomNumberRange(1000, 100)),
		AdvertID:      getRandomID(10),
		SourcePageURL: getPageUrl(),
		DestPageURL:   getPageUrl(),
	}
	return adVertEvent
}

func fakeScrollEvent() *events.Event {
	scrollEvent := &events.Event{
		UserID:              getUserID(),
		EventType:           events.ScrollEvent,
		CreatedAt:           getCreatedAt(),
		Duration:            uint32(getRandomNumberRange(1000, 100)),
		PageURL:             getPageUrl(),
		ScrolledNumProducts: uint16(getRandomNumberRange(10, 5)),
	}

	return scrollEvent
}

func getRandomID(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func getUserID() string {
	return userIDs[rand.Intn(len(userIDs))]
}

func getCreatedAt() string {
	now := time.Now().UnixNano()
	t := time.Unix(0, now)
	return t.Format("2006-01-02T15:04:05")
}

func getPageUrl() string {
	return pageURLS[rand.Intn(len(pageURLS))]
}

func getButtonElement() string {
	return buttonElements[rand.Intn(len(buttonElements))]
}

func getSearchQuery() string {
	return searchQueries[rand.Intn(len(searchQueries))]
}

func getRandomNumberRange(max, min int) int {
	return rand.Intn(max-min) + min
}

func getPaymentMode() string {
	return paymentModes[rand.Intn(len(paymentModes))]
}
