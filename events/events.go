package events

type EventType string

const (
	PageViewEvent    EventType = "page_view"
	ButtonClickEvent EventType = "button_click"
	SearchEvent      EventType = "search_event"
	PurchaseEvent    EventType = "purchase_event"
	AdvertEvent      EventType = "advert_event"
	ScrollEvent      EventType = "scroll_event"
	UnkownEvent      EventType = "unkown_event"
)

type ButtonElement struct {
	Element string `json:"button_element,omitempty"`
}

type ButtonURL struct {
	URL string `json:"button_url,omitempty"`
}

type Event struct {
	UserID    string    `json:"user_id,omitempty"`
	EventType EventType `json:"event_type,omitempty"`
	CreatedAt string    `json:"created_at,omitempty"`
	Duration  uint32    `json:"duration_secs,omitempty"`

	PageURL             string      `json:"page_url,omitempty"`
	ClickedTarget       interface{} `json:"clicked_target,omitempty"` // button element or a url
	ProductID           string      `json:"product_id,omitempty"`
	SearchQuery         string      `json:"search_query,omitempty"`
	NumItemsPurchased   uint16      `json:"num_items,omitempty"`
	PaymentMode         string      `json:"payment_mode,omitempty"`
	AdvertID            string      `json:"advert_id,omitempty"`
	SourcePageURL       string      `json:"source_page_url,omitempty"`
	DestPageURL         string      `json:"dest_page_url,omitempty"`
	ScrolledNumProducts uint16      `json:"scrolled_products,omitempty"`
}
