package main

// go get github.com/sirupsen/logrus
// go get gopkg.in/yaml.v2

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// https://yaml.to-go.online/
// Notification is the http endpoint payload from civ6 notification.
type notification struct {
	Game   string `json:"value1"`
	Player string `json:"value2"`
	Turn   string `json:"value3"`
}

type notificationArray []notification

var notifications = notificationArray{}

func init() {
	civ6Logfile, err := os.OpenFile("civ6.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("[Error]: %s", err))
	}

	log.SetFormatter(&log.JSONFormatter{})
	mw := io.MultiWriter(os.Stdout, civ6Logfile)
	logrus.SetOutput(mw)

	loadConfig()
}

// UserConfig yaml config file
type UserConfig struct {
	Users []struct {
		Name         string `yaml:"name"`
		NotifyMethod string `yaml:"notify-method"`
		Sms          string `yaml:"sms"`
		Email        string `yaml:"email"`
	} `yaml:"users"`
}

var u UserConfig

func loadConfig() {
	const configFile = "config.yml"
	log.Printf("Reading config %s", configFile)
	yamlFile, err := ioutil.ReadFile(configFile)

	u = UserConfig{}
	err = yaml.Unmarshal([]byte(yamlFile), &u)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%+v\n\n", u)
	fmt.Printf("%+v\n\n", u.Users[0])

	var anyJSONMap map[string]interface{}
	yaml.Unmarshal(yamlFile, &anyJSONMap)
	log.WithFields(log.Fields{"yaml": &anyJSONMap}).Info("anyJSONMap")

	if err != nil {
		log.Printf("Error reading YAML file: %s\n", err)
		return
	}
	log.Printf(string(yamlFile))

	//var yamlConfig YamlConfig
	//err = yaml.Unmarshal(yamlFile, &yamlConfig)
	//if err != nil {
	//	fmt.Printf("Error parsing YAML file: %s\n", err)
	//}

	//log.Printf("Result: %v\n", yamlConfig)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	var newNotification notification
	reqBody, err := ioutil.ReadAll(r.Body)

	var anyJSONMap map[string]interface{}
	json.Unmarshal(reqBody, &anyJSONMap)
	//log.Println(anyJSONMap)
	//fmtanyJSONMap, _ := json.Marshal(anyJSONMap)
	//xType := fmt.Sprintf("%T", fmtanyJSONMap)
	//fmt.Println(xType)
	//log.WithFields(log.Fields(fmtanyJSONMap))
	//log.Println(string(fmtanyJSONMap))
	//log.Field{"anyJSONMap": &anyJSONMap}
	//log.Field{"fmtanyJSONMap": &fmtanyJSONMap}
	log.WithFields(log.Fields{"reqBody": &anyJSONMap}).Info("payload")
	//log.WithFields(log.Fields{"fmtanyJSONMap": &fmtanyJSONMap}).Info("payload")
	//fmt.Println("value2: %s", anyJSONMap["value2"])

	// Look for this user to send
	for _, e := range u.Users {
		if e.Name == anyJSONMap["value2"] {
			//fmt.Println("found:", e.Name)
			log.WithFields(log.Fields{"yaml": &anyJSONMap}).Info("Sending to user")
		}
	}

	if err != nil {
		fmt.Fprintf(w, "Error")
	}

	json.Unmarshal(reqBody, &newNotification)
	notifications = append(notifications, newNotification)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newNotification)

	//fmtNotification, _ := json.Marshal(newNotification)
	//log.Println(string(fmtNotification))
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/civ", homePage)
	http.HandleFunc("/new_turn", homePage)
	http.HandleFunc("/civ/new_turn", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	fmt.Println("Starting...")
	handleRequests()
}
