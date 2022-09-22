package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
	"wednesday-app/models"
	"wednesday-app/utils"

	log "github.com/sirupsen/logrus"
)

func BreakWords(w http.ResponseWriter, r *http.Request) {
	log.Infoln("INFO:: inside PUT Request method call BreakWords")
	var tweetArray []models.Tweets
	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorln(err)
	}
	chunk := string(responseData)
	if chunk == "" {
		log.Infoln("no data recieved in body")
		w.Write([]byte("no data recieved in body"))
		return
	}
	resultArray, _ := utils.CountAndPartition(chunk)
	for i, v := range resultArray {
		
		temp:= fmt.Sprintf("%d/%d:%s", i+1, len(resultArray), v)
		tweetArray = append(tweetArray, models.Tweets{Tweet: temp})
	}
	textArrayConverted, err := json.Marshal(tweetArray)
	if err != nil {
		log.Fatalf(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(textArrayConverted)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Infoln(tweetArray)
}
