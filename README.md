<h1 align="center">philo</h1>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/Kyuubang/philo">
    <img src="https://goreportcard.com/badge/github.com/Kyuubang/philo" alt="Go Report Card" />
  </a>
  <img alt="GitHub go.mod Go version (subdirectory of monorepo)" src="https://img.shields.io/github/go-mod/go-version/Kyuubang/philo">
  <a href="https://github.com/Kyuubang/philo/blob/master/LICENSE">
    <img alt="GitHub" src="https://img.shields.io/github/license/Kyuubang/philo">
  </a>
</p>

Philo is a simple CLI tools that help you to automatically check your assignment based on test case files. It allows students 
to test their work on a local machine or virtual machine and automatically submit their results to a central server for 
grading. Philo is designed to work with a local or virtual machine infrastructure, which can be managed using Vagrant as 
an Infrastructure as Code (IaC) tool. Philo also supports cross-platform use, which enables students to use it regardless 
of their operating system.

## Installation

### Download binary

see release page here match with your architecture. note: **tested on linux only**

### From source

#### Requirements

- [Go](https://golang.org/dl/) 1.16 or higher
- [Vagrant](https://www.vagrantup.com/)
- [VirtualBox](https://www.virtualbox.org/)

Clone repository with following command

```bash
git clone https://github.com/Kyuubang/philo.git
cd philo
```

Build Philo with following command

```bash
go build -o philo
```

for admin user (to use `philo admin` command)

```bash
go build -o philo -tags admin
```

**Note:** If you want to build Philo for a different operating system, you can use the `GOOS` and `GOARCH` environment
variables. For example, to build Philo for Windows on 64-bit architecture, you can use the following command:

```bash
GOOS=windows GOARCH=amd64 go build -o philo.exe
```


## Contributing

This project is open to contributions. and I realize that many things can be improved. If you want to contribute to Philo, 
follow these steps:

1. Fork the Philo repository. 
2. Create a new branch with a descriptive name. 
3. Make your changes and commit them. 
4. Push your changes to your forked repository. 
5. Submit a pull request to the Philo repository.

**Note:** If you want to contribute to Philo, you can fix available open issue or if you add new feature, please open new 
issue first.

## License

Philo is licensed under the MIT License. See the [LICENSE](https://github.com/Kyuubang/philo/blob/master/LICENSE) file for details.
