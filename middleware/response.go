package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendResponse(Text, ConversationId, ActivityId, Token string) {
	var resp Response
	//var response string
	//var Name, Message, ContainerStatus, Reason, Status string
	//var details string
	//var end ="--------"
	//x,err:=K8s(Text,ConversationId,Token)
	//fmt.Printf("ConversationId: %s", ConversationId)
	//s, err := CreateClient(Text)
	//fmt.Printf("\nValue os s: %v", s)
	//var x *[]PodDetails
	//fmt.Println("checking poddetails: ", x)
	method := "POST"
	resp.Type = "message"
	//for _, x := range *x {
	//	Name = x.Name
	//	Message = x.Message
	//	ContainerStatus = x.ContainerStatus
	//	Reason = x.Reason
	//	Status = x.Status
	//	det := fmt.Sprintf("Name : %s\nMessage: %s\nContainerStatus : %s\nReason: %s\nStatus: %s\n", Name, Message, ContainerStatus, Reason, Status)
	//det := fmt.Sprintf("text: %s", Text)
	//details = det + "\n" + details
	//}
	//fmt.Print("testing details: \n", det)
	//	payload := strings.NewReader(`{`+"
	//+`
	//	"type": "message",`+"
	//"+`
	//	"text": "This message has been updated"`+"
	//"+`
	//}`)
	//`
	//	payload:=strings.NewReader("{" +
	//		"\"type\" : \"message\"," +
	//		" \"text\": "+ details+
	//		"}")

	resp.Text = Text

	//log.Printf("response text %v ", resp)

	//if err != nil {
	//	resp.Text = fmt.Sprint("Error: ", err)
	//} else {
	//	resp.Text = det
	//}
	//str:=fmt.Sprintf("{\n %q: %q, \n %q: \"%v\"\n}","type","message","text",details)
	//
	//payload:=strings.NewReader(str)
	//fmt.Print("testing str: \n",payload)
	//marshal, err := json.Marshal(resp)
	//if err != nil {
	//	return
	//}
	//fmt.Println("checking marshal: ",marshal)
	//fmt.Printf("ConversationIdR:%q",ConversationId)
	//fmt.Printf("ActivityIdR:%q",ActivityId)
	url := fmt.Sprintf("https://smba.trafficmanager.net/amer/v3/conversations/%s/activities/%s", ConversationId, ActivityId)
	//fmt.Println("Checkurl: ", url)

	//var resppp []byte
	//resppp=resp
	//finalresponse:=fmt.Sprintf("%s",resp)
	//for _,x:= range strings.Split(finalresponse,"{"){
	//	for _,y:=range strings.Split(x,"}"){
	//		response=y+response
	//	}
	//}
	//fmt.Printf("25: %s\n",response)
	//payload:=strings.NewReader(response)
	//payload:=strings.NewReader(fmt.Sprintf("%s",resp))
	client := &http.Client{}
	//fmt.Println("checking payload: ",payload)
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(resp)
	//fmt.Print("testing str: \n", payloadBuf)
	req, err := http.NewRequest(method, url, payloadBuf)

	if err != nil {
		fmt.Println(err)
		return
	}
	bearer := fmt.Sprintf("Bearer %s", Token)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

type Response struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func SendResponseWOActivity(ConversationId, Token string) {
	var resp Response
	//var response string
	//var Name, Message,ContainerStatus,Reason,Status string
	//var details string
	//var end ="--------"
	//x,err:=K8s(Text,ConversationId,Token)
	//fmt.Println("checking poddetails: ",x)
	method := "POST"
	resp.Type = "message"
	//for _,x:=range *x{
	//	Name=x.Name
	//	Message=x.Message
	//	ContainerStatus=x.ContainerStatus
	//	Reason =x.Reason
	//	Status=x.Status
	//	det:=fmt.Sprintf("Name : %s\nMessage: %s\nContainerStatus : %s\nReason: %s\nStatus: %s\n",Name,Message,ContainerStatus,Reason,Status)
	//
	//	details=det+"\n"+details
	//}
	//fmt.Print("testing details: \n",details)
	//	payload := strings.NewReader(`{`+"
	//+`
	//	"type": "message",`+"
	//"+`
	//	"text": "This message has been updated"`+"
	//"+`
	//}`)
	//`
	//	payload:=strings.NewReader("{" +
	//		"\"type\" : \"message\"," +
	//		" \"text\": "+ details+
	//		"}")

	//if err!=nil{
	//	resp.Text=fmt.Sprint("Error: ",err)
	//}else{
	//resp.Text = Text
	//}

	url := fmt.Sprintf("https://smba.trafficmanager.net/amer/v3/conversations/%s/activities", ConversationId)
	fmt.Println("Checkurl: ", url)
	//var resppp []byte
	//resppp=resp
	//finalresponse:=fmt.Sprintf("%s",resp)
	//for _,x:= range strings.Split(finalresponse,"{"){
	//	for _,y:=range strings.Split(x,"}"){
	//		response=y+response
	//	}
	//}
	//fmt.Printf("25: %s\n",response)
	//payload:=strings.NewReader(response)
	//payload:=strings.NewReader(fmt.Sprintf("%s",resp))
	client := &http.Client{}
	marshal, err := json.Marshal(resp)
	if err != nil {
		return
	}
	payload := strings.NewReader(string(marshal))
	//payloadBuf := new(bytes.Buffer)
	//json.NewEncoder(payloadBuf).Encode(resp)
	//fmt.Print("testing str: \n",payloadBuf)
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	bearer := fmt.Sprintf("Bearer %s", Token)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
