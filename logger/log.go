package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

// Config for logging
type Config struct {
	Directory string
	Filename  string
}

type Logger struct {
	*zerolog.Logger
}

// set lab log file by lab name
func setLogFile(config Config) *os.File {
	Log, _ := os.OpenFile(path.Join(config.Directory, config.Filename), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	return Log
}

func logConfigure(config Config) *Logger {
	var writers []io.Writer

	writers = append(writers, zerolog.ConsoleWriter{
		Out:        setLogFile(config),
		TimeFormat: time.RFC1123,
	})

	mw := io.MultiWriter(writers...)

	logger := zerolog.New(mw).With().Timestamp().Logger()

	return &Logger{
		Logger: &logger,
	}
}

// MainLog main log it will be written to -> /var/log/philo/philo.log
func MainLog() *Logger {
	logMain := logConfigure(Config{
		Directory: "/var/log/philo",
		Filename:  "philo.log",
	})
	return logMain
}

// LabLog setup lab log by log name -> /var/log/philo/labs/<labname>.log
func LabLog(lab string) *Logger {
	logLab := logConfigure(Config{
		Directory: "/var/log/philo/labs",
		Filename:  lab + ".log",
	})
	return logLab
}

// Color it take function to change color of printed string
type Color struct{}

func (c *Color) Red(anymsg ...string) string {
	return "\033[31m" + strings.Join(anymsg, " ") + "\033[0m"
}

func (c *Color) Green(anymsg ...string) string {
	return "\033[32m" + strings.Join(anymsg, " ") + "\033[0m"
}

func (c *Color) Yellow(anymsg ...string) string {
	return "\033[33m" + strings.Join(anymsg, " ") + "\033[0m"
}

var color = Color{}

type Console string

func (m Console) Error() {
	fmt.Println(color.Red("[!]", string(m)))
}

func (m Console) Warn() {
	fmt.Println(color.Yellow("[!]", string(m)))
}

func (m Console) Info() {
	fmt.Println("[*]", m)
}

func (m Console) Start() {
	fmt.Println("[+]", m)
}

func (m Console) Success() {
	fmt.Println(color.Green("[✓]", string(m)))
}

func (m Console) Fail() {
	fmt.Println(color.Red("[✗]", string(m)))
}

func (m Console) Test() {
	fmt.Print("    => ", m)
}
