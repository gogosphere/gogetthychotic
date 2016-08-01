package main

// https://secretserveronline.com/webservices/sswebservice.asmx API description here

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/howeyc/gopass"
	xmld "github.com/williamrhancock/xmlanswers"
)

const (
	// update these (or replace them with a Flags package which ever is better for the user)
	domain  = "YOUADDOMAIN"
	hosturl = "https://yoursecretserver.yourdomain.com/SecretServer/webservices/SSWebservice.asmx"
)

func authenticationToken(xmlpayloadsource string, contentLengthraw int) []byte {
	contentLength := strconv.Itoa(contentLengthraw)
	client := &http.Client{}
	method := "POST"

	req, err := http.NewRequest(method, hosturl, bytes.NewBuffer([]byte(xmlpayloadsource)))
	req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
	req.Header.Set("Content-Length", contentLength)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func getGoal() ([]byte, string, string, string) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("AD Username: ")
	userUncut, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	username := strings.TrimSpace(userUncut)

	fmt.Printf("Password: ")
	password, err := gopass.GetPasswdMasked()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	readerquestion := bufio.NewReader(os.Stdin)
	fmt.Printf("(S)earch or (L)ookup ")
	questionUncut, err := readerquestion.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	searchOrLookup := strings.TrimSpace(strings.ToUpper(questionUncut))

	readerterm := bufio.NewReader(os.Stdin)
	fmt.Printf("Criteria: ")
	termUncut, err := readerterm.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	term := strings.TrimSpace(termUncut)

	return password, username, searchOrLookup, term

}

func main() {

	password, username, sOrL, term := getGoal()

	uriPayLoad := xmld.WindCreds(username, string(password), domain)
	uriPostLength := len(uriPayLoad)
	tokenxml := authenticationToken(uriPayLoad, uriPostLength)
	token := xmld.UnwindToken(tokenxml)

	if sOrL == "L" {
		secretID := term
		tokenPayLoad := xmld.WindToken(token, secretID)
		tokenPostLength := len(tokenPayLoad)
		passxml := authenticationToken(tokenPayLoad, tokenPostLength)
		secrets := xmld.UnwindSecret(passxml)
		for _, v := range secrets {
			fmt.Println(string(v))
		}
		fmt.Println("END OF RESULTS: SecretID")
		os.Exit(0)

	} else if sOrL == "S" {
		searchTerm := term
		searchPayload := xmld.WindSearch(token, searchTerm)
		searchPostLength := len(searchPayload)
		searchxml := authenticationToken(searchPayload, searchPostLength)
		searchName, searchID := xmld.UnwindSearch(searchxml)
		for k, v := range searchName {
			fmt.Println(searchID[k], v)
		}
		fmt.Println("END OF RESULTS: Search")
		os.Exit(0)
	}
	log.Fatalln("END OF RESULTS: RTFM")

}
