package middleware

import (
	"time"
)

type Message struct {
	Type           string    `json:"type"`
	ID             string    `json:"id"`
	Timestamp      time.Time `json:"timestamp"`
	LocalTimestamp string    `json:"localTimestamp"`
	ServiceURL     string    `json:"serviceUrl"`
	ChannelID      string    `json:"channelId"`
	From           struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		AadObjectID string `json:"aadObjectId"`
	} `json:"from"`
	Conversation struct {
		ConversationType string `json:"conversationType"`
		ID               string `json:"id"`
	} `json:"conversation"`
	Recipient struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"recipient"`
	TextFormat string `json:"textFormat"`
	Text       string `json:"text"`
	Entities   []struct {
		Locale   string `json:"locale"`
		Country  string `json:"country"`
		Platform string `json:"platform"`
		Timezone string `json:"timezone"`
		Type     string `json:"type"`
	} `json:"entities"`
	ChannelData struct {
		Tenant struct {
			ID string `json:"id"`
		} `json:"tenant"`
	} `json:"channelData"`
	Locale string `json:"locale"`
}

//func Request(w http.ResponseWriter, req *http.Request) {
//	var x Message
//
//	fmt.Println("request value received: ", req.Body)
//	err := json.NewDecoder(req.Body).Decode(&x)
//	if err != nil {
//		fmt.Print(err, "Error while parsing Bot request")
//	}
//	fmt.Println("value recieved: ", x.Text)
//	AccToken := Auth(w, req)
//	//fmt.Println(AccToken)
//	fmt.Println("ConversationId: ", x.Conversation.ID)
//	fmt.Println("ActivityId: ", x.ID)
//	SendResponse(x.Text, x.Conversation.ID, x.ID, AccToken)
//
//}
