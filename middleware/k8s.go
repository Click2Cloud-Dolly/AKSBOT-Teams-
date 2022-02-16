package middleware

import (
	//"bufio"
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	//v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	//"k8s.io/api/core/v1"
	//"k8s.io/client-go/util/retry"
)

type PodDetails struct {
	Name string
	//Ready  string
	Status string
	desPod string
	//ContainerStatus string
	//Reason          string
	//Message         string
}
type Condition struct {
	Phase string
}

//func K8s(clientset *kubernetes.Clientset) (*[]PodDetails, error) {
//
//	//path="C:/Users/raymond.mathew/Downloads/config1"
//	//var kubeconfig *string
//	//kube := path
//	//kubeconfig = &kube
//	//if _, err := os.Stat(path); err != nil {
//	//	log.Printf("osstat error: %v", err)
//	//	return nil, err
//	//}
//	//
//	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
//	//
//	//if err != nil {
//	//	log.Printf("osstate error: %v", err)
//	//	return nil, err
//	//}
//	//
//	//clientset, err := kubernetes.NewForConfig(config)
//	//if err != nil {
//	//	return nil, err
//	//}
//
//	//var poddet PodDetails
//	var poddets []PodDetails
//	clientset.EventsV1()
//
//	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
//	if err != nil {
//		return nil, err
//	}
//	//log.Printf("poddata: %v",pods)
//	outpod := fmt.Sprintf("%s", pods)
//	var b = []byte{}
//	b = []byte(outpod)
//	err = ioutil.WriteFile("pods.txt", b, 0777)
//	//// handle this error
//	if err != nil {
//		// print it out
//		fmt.Println(err)
//	}
//	//for _, pod := range pods.Items {
//	//	fmt.Println("Checking pods: ", pod.Status.Phase)
//	//
//	//	if pod.Status.Phase == "Pending" {
//	//		fmt.Println("Pending pods: ", pod.Name)
//	//
//	//		out, err := clientset.CoreV1().Pods(pod.Namespace).Get(context.TODO(), pod.Name, metav1.GetOptions{})
//	//		if err != nil {
//	//			return nil, err
//	//		}
//	//		fmt.Println("Reason for pod pending: ", out.Status.Conditions)
//	//		poddet.Name = pod.Name
//	//		poddets = append(poddets, poddet)
//	//		//var PodLogs v1.PodLogOptions
//	//		//for _,container:=range pod.Spec.Containers{
//	//		//	PodLogs.Container=container.Name
//	//		//}
//	//		//PodLog=&PodLogs
//	//		//poddets=append(poddets,poddet)
//	//		//} else if pod.Status.Phase == "Running" { //For pods in not Running state
//	//		//	fmt.Println("pods running")
//	//		//	var PodLog *v1.PodLogOptions
//	//		//	logs := clientset.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, PodLog)
//	//		//	fmt.Println("Checking logs: ", logs)
//	//		//	stream, err := logs.Stream(context.TODO())
//	//		//	if err != nil {
//	//		//		fmt.Println("Logs not found")
//	//		//		//state.Resume=1
//	//		//
//	//		//		WaitForUserResponse("Did the container died too quickly?", convid, token)
//	//		//		out, err := clientset.CoreV1().Pods(pod.Namespace).Get(context.TODO(), pod.Name, metav1.GetOptions{})
//	//		//		if err != nil {
//	//		//			return nil, err
//	//		//		}
//	//		//		fmt.Println("Pods: ", out.Name)
//	//		//		for _, poddetails := range out.Status.ContainerStatuses {
//	//		//			if !poddetails.Ready {
//	//		//				if poddetails.State.Waiting != nil {
//	//		//					if poddetails.State.Waiting.Reason == "CrashLoopBackOff" {
//	//		//						poddet.Name = pod.Name
//	//		//						poddet.Reason = poddetails.State.Waiting.Reason
//	//		//						poddet.Message = poddetails.State.Waiting.Message
//	//		//						//poddet.Status=
//	//		//						////poddet.Ip=pod.Status.PodIP
//	//		//						poddets = append(poddets, poddet)
//	//		//
//	//		//						fmt.Println("The issue is CrashLoopBackoff")
//	//		//						//goto temp
//	//		//					} else if poddetails.State.Waiting.Reason == "ImagePullBackoff" {
//	//		//						poddet.Name = pod.Name
//	//		//						poddet.Reason = poddetails.State.Waiting.Reason
//	//		//						poddet.Message = poddetails.State.Waiting.Message
//	//		//						poddets = append(poddets, poddet)
//	//		//						fmt.Println("The issue is ImagePullBackoff")
//	//		//					} else if poddetails.State.Waiting.Reason == "CrashLoopBackoff" {
//	//		//						poddet.Name = pod.Name
//	//		//						poddet.Reason = poddetails.State.Waiting.Reason
//	//		//						poddet.Message = poddetails.State.Waiting.Message
//	//		//						poddets = append(poddets, poddet)
//	//		//						fmt.Println("The issue is CrashLoopBackoff")
//	//		//					} else if poddetails.State.Waiting.Reason == "RunContainerError" {
//	//		//						poddet.Name = pod.Name
//	//		//						poddet.Reason = poddetails.State.Waiting.Reason
//	//		//						poddet.Message = poddetails.State.Waiting.Message
//	//		//						poddets = append(poddets, poddet)
//	//		//						fmt.Println("The issue is likely to be with mounting volumes.")
//	//		//					} else {
//	//		//						fmt.Println("Consult StackOverflow")
//	//		//					}
//	//		//				}
//	//		//			}
//	//		//
//	//		//		}
//	//		//		fmt.Println("pod details ", poddets)
//	//		//		//
//	//		//		//return &poddets, nil
//	//	} else {
//	//		fmt.Println("logs requested")
//	//		buf := make([]byte, 2000)
//	//		//num, _ := stream.Read(buf)
//	//		//message := string(buf[:num])
//	//		//fmt.Println("Log message: ", message)
//	//		fmt.Println("Log message: ", buf)
//	//	}
//	//
//	//	//	fmt.Println("For pods in not Running state")
//	//	//	var PodLog *v1.PodLogOptions
//	//	//	logs:=clientset.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name,PodLog)
//	//	//	stream,err:=logs.Stream(context.TODO())
//	//	//	if err != nil {
//	//	//		return nil,err
//	//	//	}
//	//	//	buf := make([]byte, 2000)
//	//	//	num,err:=stream.Read(buf)
//	//	//	if num == 0 {//if logs couldn't be fetched
//	//	//		//if{//check if the container has died too quickly
//	//	//		//
//	//	//		//}
//	//	//		//clientset.
//	//	//		//out, err :=clientset.CoreV1().Pods(pod.Namespace).Get(context.TODO(), pod.Name,metav1.GetOptions{})
//	//	//		//if err != nil {
//	//	//		//	return nil,err
//	//	//		//}
//	//	//		//fmt.Println("pod status: ",out.Status.Conditions[0].Status)
//	//	//		//fmt.Println("pod reason: ",out.Status.Conditions[0].Reason)
//	//	//		//fmt.Println("pod message: ",out.Status.Conditions[0].Message)
//	//
//	//}
//
//	//}
//	//temp:
//	//}
//
//	fmt.Println("pod details ", poddets)
//
//	return &poddets, nil
//
//}
var (
	poddet  PodDetails
	poddets []PodDetails
)

func CreateClient(path string) (*kubernetes.Clientset, error) {
	var kubeconfig *string
	kube := path
	kubeconfig = &kube
	if _, err := os.Stat(path); err != nil {
		log.Printf("osstat error: %v", err)
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		log.Printf("osstate error: %v", err)
	}
	//log.Printf("Checking config:%v", config)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func State(clientset *kubernetes.Clientset) (*[]PodDetails, error) {

	pods, err := GetPodsList(clientset, "default")
	//log.Printf("PodList is: %v", pods.Items)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Length of pods: %v\n", len(pods.Items))

	for _, pod := range (*pods).Items {

		poddet.Name = pod.Name
		//poddet.Ready = string(pod.Status.Conditions[0].Type)
		poddet.Status = string(pod.Status.Phase)

		fmt.Println("Checking pods: ", pod.Status.Phase)

		if pod.Status.Phase == "Pending" {
			fmt.Println("Pending pods: ", pod.Name)

			//pods, _ := CheckingPendingPod(clientset)
			//fmt.Printf("Describe Pending Pods: %v", pods)

			//return &pod, nil

			//poddet.desPod = string(pods)

			//break
			//out, err := GetPod(clientset, pod.Namespace, pod.Name)
			//if err != nil {
			// return "", err
			//}
		} else if pod.Status.Phase == "Running" {
			fmt.Println("Running pods: ", pod.Name)
			//return &pod, nil

			//var podlogs *v1.PodLogOptions
			//podlg, err := GetPodLogs(clientset, "default", pod.Name, podlogs)
			//if err != nil {
			//	return podlg, nil
			//}
			//fmt.Printf("Logs of pod: %v", podlg)
			//
			//if pod.Status.Conditions[0].Type == "Ready" {
			//	return "'Run the command kubectl port-forward <pod-name> 8080:<pod-port>'", nil
			//} else {
			//	return "'Describe pod'", nil
			//}

		}
		poddets = append(poddets, poddet)
		////else if pod.Status.Phase != "Running" {
		////
		////	for {
		//		var podlogs *v1.PodLogOptions
		//		podlg, err := GetPodLogs(clientset, "default", pod.Name, podlogs)
		//		if err != nil {
		//			panic(err)
		//		}
		//		fmt.Printf("Logs of pod: %v", podlg)
		//		return podlg, nil
		//		break
		//	}
		//	return "Can you see the app logs?", nil
		//return "Can you see the log for the app?", nil
		//fmt.Printf("The pods not in a running state are: %v", pod.Name)
		//return "Check the logs", nil
		//}
		//res := bytes.Split(poddets, []byte(","))

	}
	return &poddets, nil
}

//func CheckRunningPod(clientset *kubernetes.Clientset ) (*[]PodDetails, string, error) {
//
//	for _, container := range poddet.ContainerStatus {
//
//		if container.Ready {
//			status := "Ready"
//			log.Printf(status)
//		} else {
//			// pod is not in Ready state
//			desPod, err := GetPod(clientset, "default", pod.Name)
//			if err != nil {
//				panic(err)
//			}
//			if desPod.Spec.Containers[0].ReadinessProbe != nil {
//				return "Fix the Readiness probe", nil
//			} else {
//				return "Unknown state", nil
//			}
//		}
//	}
//	return "", nil
//}

//func CheckingPendingPod(clientset *kubernetes.Clientset) (*[]PodDetails, error) {
//
//	pods, err := GetPodsList(clientset, "default")
//
//	if err != nil {
//		panic(err)
//	}
//	for _, pod := range (*pods).Items {
//		//for pods := range pod.Status.Phase {
//		if pod.Status.Phase == "Pending" {
//			desPods, err := GetPod(clientset, "default", pod.Name)
//			if err != nil {
//				panic(err)
//			}
//			//poddet = poddet.Name
//			//return desPods, "Is the cluster full?", nil
//
//			//return desPods, nil
//			fmt.Printf("despods: %v \n ", desPods)
//		}
//	}
//	poddets = append(poddets, poddet)
//	return &poddets, nil
//
//}

func GetPodsList(clientset *kubernetes.Clientset, namespace string) (*v1.PodList, error) {
	podlist, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return podlist, nil
}

// Describe Pod

//func GetPod(clientset *kubernetes.Clientset, namespace, podname string) (*v1.Pod, error) {
//	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podname, metav1.GetOptions{})
//	if err != nil {
//		return nil, err
//	}
//	return pod, nil
//}

//func PortForward(clientset *kubernetes.Clientset, namespace string, podname string, podport int) {
//
//}

//func GetPodLogs(clientset *kubernetes.Clientset, namespace, podname string, podlog *v1.PodLogOptions) (string, error) {
//	logs := clientset.CoreV1().Pods(namespace).GetLogs(podname, podlog)
//	fmt.Println("Checking logs: ", logs)
//	stream, err := logs.Stream(context.TODO())
//	if err != nil {
//		return "", err
//	}
//	fmt.Println("logs requested")
//	buf := make([]byte, 2000)
//	num, _ := stream.Read(buf)
//	message := string(buf[:num])
//	fmt.Println("Log message: ", message)
//	return message, nil
//}
