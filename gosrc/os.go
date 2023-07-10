package gosrc

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Cmd() {
	kexec := exec.Command("kubectl", "exec", "-it", "managed-notebook-46090-0", "--", "bash")
	kexec.Stdin = os.Stdin
	kexec.Stdout = os.Stdout
	kexec.Stderr = os.Stderr

	fmt.Print("begin")
	time.Sleep(5 * time.Second)
	if err := kexec.Run(); err != nil {
		fmt.Print(err)
	}
	fmt.Print("finished")
}