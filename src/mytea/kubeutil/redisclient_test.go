package kubeutil

import (
	"fmt"
	"testing"
)

func Test_splitBySpace(t *testing.T) {
	fmt.Println(splitBySpace("zrange   -1     8  heelo"))
}

func Test_runCmd(t *testing.T) {
	client := NewRedisClient(manualConnector{})
	fmt.Println(client.exec("keys *"))
	fmt.Println(client.exec("get test"))
}
