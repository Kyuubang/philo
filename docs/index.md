---
layout: default
title: Home
nav_order: 1
description: "Philo is a simple CLI tools that help you to automatically check your assignment based on test case files"
permalink: /
---

# Task-centric Assessment Approach
{: .fs-9 }

Self Grading System enable you to focus on the task and not grading process.
{: .fs-6 .fw-300 }

[Get started now](#getting-started){: .btn .btn-primary .fs-5 .mb-4 .mb-md-0 .mr-2 }
[View it on GitHub][philo repo]{: .btn .fs-5 .mb-4 .mb-md-0 }

---

# philo
philo is a simple CLI tools that help you to automatically check your assignment based on test case files. It allows 
students to test their work on a local machine or virtual machine and automatically submit their results to a central 
server for grading. Philo is designed to supports cross-platform use, which enables students to use it 
regardless of their operating system.

## Getting started

### Installation

to install philo, you can download binary from release page or build from source. for other platform, we recommend you to 
build from source.

#### Download binary

see release page here match with your architecture. 

{: .warning }
This project **tested on linux** only, if you want to contribute to this project, please test it on other platform and 
open issue if you found any problem.

#### From source

##### Requirements

if you want to enable IaC feature, you must install [Vagrant](https://www.vagrantup.com/) and VirtualBox

- [Go](https://golang.org/dl/) 1.16 or higher
- [Vagrant](https://www.vagrantup.com/)(optional)
- [VirtualBox](https://www.virtualbox.org/)(optional)

Clone repository with following command

```bash
git clone https://github.com/Kyuubang/philo.git
cd philo
```

Build philo with following command

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

[philo repo]: https://github.com/Kyuubang/philo
