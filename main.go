package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	currentTime := time.Now()
	url := os.Getenv("GITLAB_URL")
	fmt.Println("=== ACTION DATE ===")
	fmt.Println(currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("===================")
	fmt.Println("URL:", url)

	data := strings.NewReader(os.Getenv("MODE") + `=.*&keep_n=` + os.Getenv("KEEP"))
	req, err := http.NewRequest(os.Getenv("REQUEST"), url, data)
	req.Header.Set("PRIVATE-TOKEN", os.Getenv("API_KEY"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println("DATA:", data)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	src, err := ioutil.ReadAll(resp.Body)
	dst := &bytes.Buffer{}
	if err := json.Indent(dst, src, "", "  "); err != nil {
		panic(err)
	}

	fmt.Println(dst.String())

}
