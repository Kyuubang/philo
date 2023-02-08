package remote

// TODO: Create a remote command executor that can be used to execute commands on a remote host.

import (
	"fmt"
	"github.com/Kyuubang/philo/internal/utils/bash"
	"io/ioutil"
	"net"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
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

// RunCommand with output stdout and stderr
// currently remote command just support output
// with string type data
func (client *SSHClient) RunCommand(cmd string, lab string) bash.Out {
	//mainConfig, _ := config.InitConfig()

	session, err := client.newSession()
	defer func() bash.Out {
		msgErr := recover()
		if msgErr != nil {
			fmt.Println("unsuccessful run remote command")
		}
		return bash.Out{}
	}()

	if err != nil {
		fmt.Println("Cant Create new session for remote command!")
		panic("Cant Create new session for remote command!")
	}

	stdout, stderr := session.CombinedOutput(cmd)

	if stderr != nil {
		return bash.Out{
			StdOut:   "",
			StdErr:   strings.Replace(string(stdout), "\n", "", -1),
			ExitCode: 1,
		}
	} else {
		return bash.Out{
			StdOut:   strings.Replace(string(stdout), "\n", "", -1),
			StdErr:   "",
			ExitCode: 0,
		}
	}
}
