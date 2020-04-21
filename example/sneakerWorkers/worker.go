package sneakerWorkers

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
)

type Worker struct {
	Name       string            `yaml:"name"`
	Exchange   string            `yaml:"exchange"`
	RoutingKey string            `yaml:"routing_key"`
	Queue      string            `yaml:"queue"`
	Log        string            `yaml:"log"`
	Durable    bool              `yaml:"durable"`
	Ack        bool              `yaml:"ack"`
	Options    map[string]string `yaml:"options"`
	Arguments  map[string]string `yaml:"arguments"`
	Delays     []int32           `yaml:"delays"`
	Steps      []int32           `yaml:"steps"`
	Threads    int               `yaml:"threads"`

	Logger *log.Logger
}

var AllWorkers []Worker

func InitWorkers() {
	path_str, _ := filepath.Abs("config/workers.yml")
	content, err := ioutil.ReadFile(path_str)
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(content, &AllWorkers)
	for i, _ := range AllWorkers {
		AllWorkers[i].initLogger()
	}
}

func (worker Worker) GetName() string {
	return worker.Name
}
func (worker Worker) GetExchange() string {
	return worker.Exchange
}
func (worker Worker) GetRoutingKey() string {
	return worker.RoutingKey
}
func (worker Worker) GetQueue() string {
	return worker.Queue
}
func (worker Worker) GetLog() string {
	return worker.Log
}
func (worker Worker) GetLogFolder() string {
	re := regexp.MustCompile(`\/.*\.log$`)
	return strings.TrimSuffix(worker.Log, re.FindString(worker.Log))
}
func (worker Worker) GetDurable() bool {
	return worker.Durable
}
func (worker Worker) GetAck() bool {
	return worker.Ack
}
func (worker Worker) GetOptions() map[string]string {
	return worker.Options
}
func (worker Worker) GetArguments() map[string]string {
	return worker.Arguments
}
func (worker Worker) GetDelays() []int32 {
	return worker.Delays
}
func (worker Worker) GetSteps() []int32 {
	return worker.Steps
}
func (worker Worker) GetThreads() int {
	return worker.Threads
}

func (worker *Worker) LogInfo(text ...interface{}) {
	worker.Logger.SetPrefix("INFO: " + worker.GetName() + " ")
	worker.Logger.Println(text)
}

func (worker *Worker) LogDebug(text ...interface{}) {
	worker.Logger.SetPrefix("DEBUG: " + worker.GetName() + " ")
	worker.Logger.Println(text)
}

func (worker *Worker) LogError(text ...interface{}) {
	worker.Logger.SetPrefix("ERROR: " + worker.GetName() + " ")
	worker.Logger.Println(text)
}

func (worker *Worker) initLogger() {
	err := os.MkdirAll(worker.GetLogFolder(), 0755)
	if err != nil {
		log.Fatalf("create folder error: %v", err)
	}
	file, err := os.OpenFile(worker.GetLog(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	worker.Logger = log.New(file, "", log.LstdFlags)
}
