package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

var target string
var workDir string

func init() {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	workDir = filepath.Dir(path)
}

func main() {
	fmt.Println("启动客户端")
	fmt.Println("Ctrl C Quit")
	for {
		fmt.Println("请输入要抓取的地址...")
		_, _ = fmt.Scanln(&target)
		params := []string{
			"--url=" + target,
		}
		execCommand(workDir+"/spider", params)
	}


}

func execCommand(commandName string, params []string) bool {
	cmd := exec.Command(commandName, params...)
	//显示运行的命令
	fmt.Println(cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return false
	}
	_ = cmd.Start()
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}
	_ = cmd.Wait()
	return true
}
