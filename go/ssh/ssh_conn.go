package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	// 创建 ssh 配置
	sshConfig := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("Root1234!"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	// 创建 client
	client, err := ssh.Dial("tcp", "39.108.78.55:22", sshConfig)
	checkErr(err)
	defer client.Close()

	// 获取 session
	session, err := client.NewSession()
	checkErr(err)
	defer session.Close()

	// 拿到当前终端文件描述符
	fd := int(os.Stdin.Fd())
	termWidth, termHeight, err := terminal.GetSize(fd)

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud}

	}
	// request pty
	err = session.RequestPty("xterm", termHeight, termWidth, modes)
	checkErr(err)

	// 对接 std
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	err = session.Shell()
	checkErr(err)
	err = session.Wait()
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
