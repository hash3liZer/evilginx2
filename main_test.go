package main

import (
	"os"
	"io"
	"log"
	"testing"
	"time"
	"os/exec"
)

func TestStart(t *testing.T) {
	/*log.Println("Starting evilginx2")
	path, _ := os.Getwd()
	terminal := Start(true, path + "/phishlets", true, true, "")
	if terminal == nil {
		t.Error("Could not be started")
	}
	time.Sleep(time.Second)
	
	log.Println("Started, running commands")
	rlc := terminal.GetConfig()
	
	r := ioutil.NopCloser(strings.NewReader("help\n"))
	rlc.Stdin = r
	time.Sleep(time.Second)
	
	log.Println("Finished")
	*/
	
	path, _ := os.Getwd()
	
	/*buildProcess := exec.Command("go", "build", path)
	buildProcess.Dir = path
	buildOutput, err := buildProcess.CombinedOutput()
	log.Println("build:", buildOutput, err)*/
	
	subProcess := exec.Command(path+"/evilginx2.exe", "-debug", "-developer")
	stdin, _ := subProcess.StdinPipe()
	//defer stdin.Close()
	
	subProcess.Stdout = os.Stdout
    subProcess.Stderr = os.Stderr
	
	log.Println("start")
	if err := subProcess.Start(); err != nil {
        t.Error("Could not be started")
    }
	time.Sleep(3*time.Second)
	
	log.Println("write now")
	io.WriteString(stdin, "help\n")
	time.Sleep(3*time.Second)
    subProcess.Wait()
	
	log.Println("end")
}
