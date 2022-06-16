package kubeutil

import (
	"bytes"
	"context"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"io"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type Client struct {
	kubernetes.Clientset
	path            string
	namespace       string
	handleMapping   map[string]func(string) (string, tea.Cmd)
	backwardMapping map[string]string
}

func InitClient() *Client {
	var client = Client{
		path: "/",
	}

	client.handleMapping = map[string]func(string) (string, tea.Cmd){
		"/":     nilCmdWrap(client.showNamespace),
		"/func": nilCmdWrap(client.getNsAndShowFunction),
		"/pod":  nilCmdWrap(client.getFuncAndShowPod),
		"/log":  nilCmdWrap(client.logPod),
		"/exec": client.execPod,
	}

	client.backwardMapping = map[string]string{
		"/log":  "/pod",
		"/exec": "/pod",
		"/pod":  "/func",
		"/func": "/",
		"/":     "/",
	}

	home := os.Getenv("HOME")
	kubeconfig := filepath.Join(home, ".kube", "config")
	config, _ := clientcmd.BuildConfigFromFlags("", kubeconfig)
	clientset, _ := kubernetes.NewForConfig(config)
	client.Clientset = *clientset

	return &client
}

func (client *Client) execute(order string) (string, tea.Cmd) {
	if order == "back" {
		client.path = client.backwardMapping[client.backwardMapping[client.path]]
	}
	f, ok := client.handleMapping[client.path]
	if !ok {
		return "i don't understand the order", nil
	}
	result, cmd := f(order)
	return client.path + "  get input : " + order + "\n" + result, cmd
}

func nilCmdWrap(f func(string) string) func(string) (string, tea.Cmd) {
	return func(s string) (string, tea.Cmd) {
		return f(s), nil
	}
}

func (client *Client) showNamespace(s string) string {
	build := strings.Builder{}
	build.WriteString("choose a namespace\n")
	for i, item := range client.listNameSpace() {
		build.WriteString(fmt.Sprintf("%d: %s \n", i, item.Name))
	}

	client.path = "/func"
	return build.String()
}

func (client *Client) getNsAndShowFunction(s string) string {
	if s != "back" {
		i, err := strconv.Atoi(s)
		if err != nil {
			return "parse index error"
		}
		nsList := client.listNameSpace()
		if i > len(nsList) {
			return "index out of range"
		}
		ns := nsList[i].Name
		client.namespace = ns
	}

	client.path = "/pod"
	return "choose function \n" +
		"1: log \n" + "2: exec"
}

func (client *Client) getFuncAndShowPod(s string) string {
	if s != "back" {
		i, err := strconv.Atoi(s)
		if err != nil {
			return "parse index error"
		}
		switch i {
		case 1:
			client.path = "/log"
		case 2:
			client.path = "/exec"
		default:
			return "no such selection"
		}
	}

	//show pod
	build := strings.Builder{}
	build.WriteString("choose a pod\n")
	for j, item := range client.listPodByNs(client.namespace) {
		build.WriteString(fmt.Sprintf("%d: %s \n", j, item.Name))
	}
	return build.String()
}

func (client *Client) logPod(s string) string {
	i, err := strconv.Atoi(s)
	if err != nil {
		return "parse index error"
	}
	nsList := client.listPodByNs(client.namespace)
	if i > len(nsList) {
		return "index out of range"
	}
	pod := nsList[i].Name

	req := client.CoreV1().Pods(client.namespace).GetLogs(pod, &v1.PodLogOptions{})
	body, err := req.Stream(context.Background())
	if err != nil {
		return "get log error" + err.Error()
	}
	var buf bytes.Buffer
	io.Copy(&buf, body)
	return buf.String()
}

func (client *Client) listNameSpace() []v1.Namespace {
	list, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	return list.Items
}

func (client *Client) listPodByNs(ns string) []v1.Pod {
	podList, _ := client.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
	return podList.Items
}

func (client *Client) execPod(s string) (string, tea.Cmd) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return "parse index error", nil
	}
	nsList := client.listPodByNs(client.namespace)
	if i > len(nsList) {
		return "index out of range", nil
	}
	pod := nsList[i].Name

	c := exec.Command("kubectl", "exec", "-it", "-n", client.namespace,
		pod, "--", "sh", "-c", "clear; (bash || sh || ash)")
	return "wait a moment", tea.Exec(tea.WrapExecCommand(c), nil)
}
