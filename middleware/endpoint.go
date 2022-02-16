package middleware

import (
	"aksbot/client"
	"bytes"

	//"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type File struct {
	Name string `json:"name"`
	//Rname string `json:"rname"`
}

type Pods struct {
	Messages      string `json:"messages,omitempty"`
	State         string `json:"state,omitempty"`
	Reason        string `json:"reason,omitempty"`
	ContainerName string `json:"containerName,omitempty"`
	PodName       string `json:"podName,omitempty"`
}

//type PodDetails struct {
//	Name  string
//	Ip    string
//	Phase string
//	PodCond       string
//	PodStatus     string
//	ReadinessName string
//	Message string
//
//}

//func create(conn *websocket.Conn) {
//	var f File
//	var r response
//	for {
//		messageType, p, err := conn.ReadMessage()
//		if err != nil {
//			log.Println(err)
//			return
//		}
//		err = json.Unmarshal(p, &f)
//		//err = json.NewDecoder(p).Decode(&f)
//		if err != nil {
//			log.Fatalf("Unable to unmarshal the request body.  %v", err)
//		}
//		if _, err := os.Stat(f.Name); err == nil {
//			fmt.Printf("File exists\n")
//			return
//
//		}
//
//		file, err = os.Create(f.Name)
//
//		if err != nil {
//			panic(err)
//		}
//		log.Println("File Created")
//		r.Message = "File created Successfully"
//		fmt.Println(string(p))
//		p, err = json.MarshalIndent(r, "", "")
//		if err != nil {
//			log.Fatalf("Unable to marshal the response.  %v", err)
//		}
//		if err := conn.WriteMessage(messageType, p); err != nil {
//			log.Println(err)
//			return
//		}
//	}
//
//	//	log.Println(file)
//	//file.Close()
//}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Checking")
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Client connected")

	//to iterate over one client
	for {
		mt, message, err := ws.ReadMessage()

		if err != nil {
			log.Println("read1:", err)
			break
		}
		var msg = string(message)
		log.Printf("Message received:%v", msg)
		clients, err := CreateClient(msg)

		if err != nil {
			panic(err)
		}
		//fmt.Println(clients)
		//message processing (Webserver1)
		log.Printf("receive1: %s", message)

		z, err := State(clients)
		if err != nil {
			panic(err)
		}

		fmt.Printf("value of z: %v", z)
		//<----- pod details

		//if z.Status.Phase == "Pending" {
		//	CheckingPendingPod(clients, z)
		//	if msg == "Yes" || msg == "yes" {
		//		err := ws.WriteMessage(mt, []byte("Provision a bigger cluster"))
		//		if err != nil {
		//			panic(err)
		//		}
		//	}
		//} else {
		//	CheckRunningPod(clients, z)
		//}

		//data := []byte(z)
		//var data string

		//takes *v1.Pod as object and convert it into byte code []byte
		//data, _ := json.Marshal(Pods{PodName: z.Name, State: string(z.Status.Phase), ContainerName: z.Spec.Containers[0].Name})

		data, _ := json.Marshal(z)
		fmt.Printf("value of data: %v", string(data))

		data1 := bytes.SplitN(data, []byte(","), -1)
		//
		////tempdata := []byte{data1}
		//fmt.Printf("Value of data1: %v", data1)

		//data2, _ := json.Marshal(data1)
		//fmt.Printf("value of data2: %v", string(data2))
		//fmt.Printf("The value of Pods : %v", string(data))
		//data2= []byte(data1)

		//fmt.Println(z)

		data3 := bytes.Join(data1, []byte("\n\n"))
		//fmt.Printf("Value of data3: %v", string(data3))

		//err = ws.WriteMessage(mt, data)
		err = ws.WriteMessage(mt, data3)
		if err != nil {
			log.Println("write1:", err)
			//break
		}
		//fmt.Printf("The value written in err: %v", err)

		mt, message, err = ws.ReadMessage()
		if err != nil {
			log.Printf("read message")
			panic(err)
		}
		log.Printf("Received msg1:%v", string(message))

		//if msg == "Pending" {
		//	pod, s, err := CheckingPendingPod(clients, z)
		//	if err != nil {
		//		panic(err)
		//	}
		//	fmt.Println(pod, s)
		//	if msg == "Yes" || msg == "yes" {
		//		err := ws.WriteMessage(mt, []byte("Provision a bigger cluster"))
		//		if err != nil {
		//			panic(err)
		//		}
		//	}
		//} else {
		//	CheckRunningPod(clients, z)
		//}

		//if string(message) == "yes" || string(message) == "Yes" {

		if string(message) == "yes" || string(message) == "Yes" {

			err := ws.WriteMessage(mt, []byte("Fix the issue in the application"))
			if err != nil {
				log.Printf("Failed to check")
				panic(err)
			}
			//
			//	if msg == "yes" || msg == "Yes" {
			//
			//		err := ws.WriteMessage(mt, []byte("okk"))
			//		if err != nil {
			//			log.Printf("Failed to check")
			//			panic(err)
			//		}
			//
			//	}
			//} else {
			//	err := ws.WriteMessage(mt, []byte("Are the pods running?"))
			//	if err != nil {
			//		log.Printf("Failed to check")
			//		panic(err)
			//	}
			//	fmt.Printf("Message1 received is : %v", msg)
			//	if msg == "yes" || msg == "Yes" {
			//		fmt.Printf("Message2 received is : %v", msg)
			//		err := ws.WriteMessage(mt, []byte("Are the pods Ready?"))
			//		if err != nil {
			//			panic(err)
			//		}
			//	} else {
			//		fmt.Printf("Message3 received is : %v", msg)
			//		err := ws.WriteMessage(mt, []byte("Check the logs"))
			//		if err != nil {
			//			panic(err)
			//		}
			//	}
			//}

		}
		//else {
		//	err := ws.WriteMessage(mt, []byte(""))
		//}

		////message processing (Webserver2)
		//mt, message, err = ws.ReadMessage()
		//
		//if err != nil {
		//	log.Println("read:", err)
		//	break
		//}
		//
		////message processing
		//log.Printf("receive: %s", message)
		//data = []byte("Webserver2 connected")
		//
		//err = ws.WriteMessage(mt, data)
		//
		//if err != nil {
		//	log.Println("write:", err)
		//	break

	}

	//s, err := K8s(path)
	//
	//if err != nil {
	//	panic(err)
	//}
	//return s
	//return nil

}

func HttpEndpoint(w http.ResponseWriter, r *http.Request) {
	var response_msg Message
	var resmsg []byte
	//fmt.Printf("request value received: %v", r.Body)
	err := json.NewDecoder(r.Body).Decode(&response_msg)

	if err != nil {
		fmt.Print(err, "Error while parsing bot request")
	}
	//log.Printf("response msg: %v", response_msg)
	if len(client.Clientchan) > 0 {
		fmt.Println("\nSame client is used")
		c := <-client.Clientchan

		err = c.WriteMessage(1, []byte(response_msg.Text))
		log.Printf("Message sent")
		if err != nil {
			panic(err)
		}

		_, msgs, err := c.ReadMessage()
		if err != nil {
			panic(err)
		}
		//fmt.Println("read message:", i, string(msgs))
		resmsg = msgs
		//
		//err = c.WriteMessage(i,msgs)
		//if err != nil {
		//	return
		//}
		client.Clientchan <- c
		//fmt.Println("value received: ", msg)
		//AccToken := Auth(w, r)
		//fmt.Println(AccToken)
		//fmt.Println("ConversationId: ", msg.Conversation.ID)
		//fmt.Println("ActivityId: ", msg.ID)
		//SendResponse(msg.Text, msg.Conversation.ID, msg.ID, AccToken)

	} else {
		fmt.Println("Client used for 1st time")
		c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8082/ws", nil)
		if err != nil {
			fmt.Println("websocket dial err")
			log.Fatal(err)
		}
		err = c.WriteMessage(1, []byte(response_msg.Text))
		if err != nil {
			panic(err)
		}

		i, msg, err := c.ReadMessage()
		if err != nil {
			panic(err)
		}
		client.Clientchan <- c
		fmt.Println("server connected", i)
		//log.Printf("Value of msg %s", msg)
		resmsg = msg
		//_, err = w.Write(msg)
		//if err != nil {
		//	return
		//}

	}
	//fmt.Printf("value received: %s", resmsg)
	AccToken := Auth(w, r)
	//fmt.Println(AccToken)
	fmt.Println("ConversationId: ", response_msg.Conversation.ID)
	fmt.Println("ActivityId: ", response_msg.ID)
	SendResponse(string(resmsg), response_msg.Conversation.ID, response_msg.ID, AccToken)

}
