package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/deckarep/golang-set"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type SessionStartJSONResponce struct {
	Sid string `json:"sid"`
}

type RetweetObject struct {
	UserId         string `json:"user_id_str"`
	ScreenName     string `json:"screen_name"`
	FollowersCount int    `json:"followers_count"`
	FriendsCount   int    `json:"friends_count"`
	Verified       bool   `json:"verified"`
	RtIdStr        string `json:"rt_id_str"`
}

type RetweetsResultObject struct {
	LastRetweetCount int             `json:"last_retweet_count"`
	Retweets         []RetweetObject `json:"retweets"`
}

func createSession() SessionStartJSONResponce {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	req, err := http.NewRequest("GET", "https://twren.ch/socket.io/?EIO=3&transport=polling&t=MdE7D1T", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Dnt", "1")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "ru,en;q=0.9,ru-RU;q=0.8,en-US;q=0.7")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Referer", "https://twren.ch/")
	req.Header.Set("Cookie", "connect.sid=s%3AS3kuj8PqKpgWsRdqQJW8Oge4FdwyAgIW.xbQt%2FSIkZH2S8v%2BBd6HbsXsrGcBejanbDB5oADnZBWU; io=qXnzLIBADKLoQgIFAKto")
	req.Header.Set("Connection", "keep-alive")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var sessionStart SessionStartJSONResponce
	err = json.Unmarshal(body[5:], &sessionStart)
	if err != nil {
		panic(err)
	}
	return sessionStart
}

func sendMessage(ws *websocket.Conn, message string) {
	_, err := ws.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func receiveMessage(ws *websocket.Conn) string {
	var msg = make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	return string(msg[:n])
}

func main() {
	tweetId := *flag.Uint("tweetId", 1, "The tweet id for load its retweeters")
	if tweetId == 0 {
		flag.Usage()
		os.Exit(1)
	}

	session := createSession()
	config, _ := websocket.NewConfig(fmt.Sprintf("wss://twren.ch/socket.io/?EIO=3&transport=websocket&sid=%s", session.Sid), "https://twren.ch")
	config.TlsConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "twren.ch",
	}
	ws, err := websocket.DialConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	sendMessage(ws, "2probe")
	// skip probe response
	receiveMessage(ws)
	sendMessage(ws, "5")
	sendMessage(ws, fmt.Sprintf("42[\"loadRTs\",{\"id\":\"%d\"}]", tweetId))
	var buffer bytes.Buffer
	for {
		var msg = receiveMessage(ws)
		if strings.Contains(msg, "Successfully") {
			// process final json
			arr := *new([]RetweetsResultObject)
			json.Unmarshal(buffer.Bytes(), &arr)
			result := arr[1]
			users := mapset.NewSet()
			for _, retweet := range result.Retweets {
				users.Add(retweet.ScreenName)
			}
			res, _ := json.Marshal(users)
			fmt.Printf("%s\n", res)
			buffer.Reset()
			ws.Close()
			break
		} else {
			// concat the parts of the json
			if strings.Index(msg, "42") == 0 {
				// previous json was finished, read a new one
				buffer.Reset()
				buffer.WriteString(msg[2:])
			} else {
				buffer.WriteString(msg)
			}
		}
	}
}
