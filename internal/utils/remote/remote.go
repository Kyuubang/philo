package remote

import (
	"bytes"
	"fmt"
	"github.com/Kyuubang/philo/internal/utils/bash"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

type SSHClient struct {
	Config *ssh.ClientConfig
	Host   string
	Port   int
}

// PublicKeyFile read manual private key file from $HOME/.nusactl/hosts.yaml
func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Failed when read ssh-key file")
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		fmt.Println("bad private key")
	}
	return ssh.PublicKeys(key)
}

// SSHAgent read ssh agent's Host
func SSHAgent() ssh.AuthMethod {
	if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	}
	return nil
}

// create new session will be return authentication error
// can't dial -> server unreachable, auth error etc.
// create session -> timeout or something error.
func (client *SSHClient) newSession() (*ssh.Session, error) {
	connection, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", client.Host, client.Port), client.Config)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %s", err)
	}

	session, err := connection.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %s", err)
	}

	return session, nil
}

// RunRemoteCommand on VM Participant
func (client *SSHClient) RunRemoteCommand(command string) (bash.Out, error) {

	var session, err = client.newSession()
	if err != nil {

	}
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	err = session.Run(command)
	if err != nil {
		exitErr, ok := err.(*ssh.ExitError)
		if !ok {
			return bash.Out{
				StdOut:   strings.Replace(stdout.String(), "\n", "", -1),
				StdErr:   strings.Replace(stderr.String(), "\n", "", -1),
				ExitCode: -1,
			}, fmt.Errorf("Failed to run command: %v", err)
		}
		return bash.Out{
			StdOut:   strings.Replace(stdout.String(), "\n", "", -1),
			StdErr:   strings.Replace(stderr.String(), "\n", "", -1),
			ExitCode: exitErr.ExitStatus(),
		}, nil
	}
	return bash.Out{
		StdOut:   strings.Replace(stdout.String(), "\n", "", -1),
		StdErr:   strings.Replace(stderr.String(), "\n", "", -1),
		ExitCode: 0,
	}, nil
}
