package client

import (
	"github.com/gorilla/websocket"
)

var (
	Clientchan = make(chan *websocket.Conn, 1)
)

//type kubeConfigPath struct {
//	kubeConfigPath string
//}

//func ClientManager(path string) {
//
//	var kubeconfig *string
//	kube := path
//	kubeconfig = &kube
//	if _, err := os.Stat(path); err != nil {
//		log.Printf("osstat error: %v", err)
//
//	}
//
//	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
//	if err != nil {
//		log.Printf("osstate error: %v", err)
//		return
//	}
//
//
//
//}
