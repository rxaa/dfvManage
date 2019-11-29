package utils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//调用os.MkdirAll递归创建文件夹
func MsDirs(filePath string) {
	if !Exists(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		panic(err)
	}
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

type Comd struct {
	Cmd    *exec.Cmd
	Stdout io.ReadCloser
	Stdin  io.WriteCloser
	Stderr io.ReadCloser
}

func ReadFile(filename string) string {
	json, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(json)
}

func WriteFile(filename string, data []byte, perm os.FileMode) {
	err := ioutil.WriteFile(filename, data, perm)
	panic(err)
}

func Exec(cm string) string {
	fmt.Println(cm)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", cm)
		//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	} else {
		cmd = exec.Command("bash", "-c", cm)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	defer stdout.Close() // 保证关闭输出流
	if err := cmd.Start(); err != nil { // 运行命令
		panic(err)
	}

	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil { // 读取输出结果
		panic(err)
	}
	ret := string(opBytes)
	fmt.Println(ret)
	return ret
}

func (inst *Comd) Open() string {
	if runtime.GOOS == "windows" {
		inst.Cmd = exec.Command("cmd")
		//隐藏窗口
		//inst.Cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	} else {
		inst.Cmd = exec.Command("sh")
	}

	stdout, err := inst.Cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stdin, err := inst.Cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	stderr, err := inst.Cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	inst.Stdout = stdout
	inst.Stdin = stdin
	inst.Stderr = stderr

	err = inst.Cmd.Start()
	if err != nil {
		panic(err)
	}

	return inst.ReadAll()
}

func (inst *Comd) ReadAll() string {
	var end_line = ""
	reader := bufio.NewReader(inst.Stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		end_line += line
	}
	return end_line
}

func (inst *Comd) Exec(cmd string) string {
	_, err := inst.Stdin.Write([]byte(cmd))
	if err != nil {
		panic(err)
	}

	return inst.ReadAll()
}

func (inst *Comd) Listen() {

}
