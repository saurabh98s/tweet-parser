package handlers

import (
	"encoding/json"
	"io/ioutil"

	"net/http"
	"wednesday-app/models"

	log "github.com/sirupsen/logrus"
)

func BreakWords(w http.ResponseWriter, r *http.Request) {
	log.Infoln("INFO:: inside PUT Request method call BreakWords")
	var tweetArray []models.Tweets
	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorln(err)
	}
	var temp string

	chunk := string(responseData)
	tweet := []rune(chunk)
	label := 1
	counter := 0
	for chars := 0; chars < len(tweet); chars++ {
		var text models.Tweets

		temp += string(tweet[chars])
		if chars == len(tweet)-1 {
			w.Header().Set("Content-Type", "application/json")
			text.Tweet = temp
			text.Label = label
			tweetArray = append(tweetArray, text)
		}

		if chars/50 == label {
			text.Tweet = temp
			text.Label = label
			tweetArray = append(tweetArray, text)

			temp = ""
			counter = 0
			label++
		}
		counter++
	}
	textArrayConverted, err := json.Marshal(tweetArray)
	if err != nil {
		log.Fatalf(err.Error())
	}

	_, err = w.Write(textArrayConverted)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Infoln(tweetArray)

	log.Println(len(tweet))
}
