package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"wednesday-app/models"
)

func TestBreakWords(t *testing.T) {
	var tweet []models.Tweets
	var gotPartitionWhenEmpty int
	tt := []struct {
		name       string
		payload    string
		partitions int
		errMsg     string
	}{
		{name: "long_string", payload: `After thoroughly study of humanitarian intervention over analysis of responsibility to protect the come to an end where all this argument leads to doctrine of responsibility to protect has not succeeded in establishing agreed upon conditions and methods for intervening in case of humanitarian crisis. The complete failure of United Nations peacekeeping and security council after numerous need the cries of slaughtered millions from show many years to current situation became worst slaughter in Biafra, Uganda, Burnudi, Indonesia, Bangladesh and Sudan in which current era Afghanistan, Russia Ukraine war we will finally say that legitimateness Of these mass `,
			partitions: 14,
			errMsg:     ""},

		{name: "short_string", payload: `After thoroughly study of humanitarian intervention over analysis of responsibility to protect the come to an end where all this argument`,
			partitions: 3,
			errMsg:     ""},
		{name: "no_string	", payload: ``,
			partitions: 0,
			errMsg:     "no data to parse"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			payload := strings.NewReader(tc.payload)
			req, err := http.NewRequest("PUT", "http://127.0.0.1:8000/split_tweets", payload)
			if err != nil {
				t.Fatalf("could not create request:%v", err)
			}
			rec := httptest.NewRecorder()
			BreakWords(rec, req)
			result := rec.Result()
			if result.StatusCode != http.StatusOK {
				t.Errorf("Expected Status OK! got: %v", result.StatusCode)
			}
			resBody, _ := ioutil.ReadAll(result.Body)
			resBytes := string(resBody)
			err = json.Unmarshal([]byte(resBytes), &tweet)
			if err != nil {
				t.Errorf("Unmarshalling error: %v", err)
			}
			if len(tweet) == 0 {
				gotPartitionWhenEmpty = 0
			} else {
				gotPartitionWhenEmpty = tweet[len(tweet)-1].Label
			}

			if gotPartitionWhenEmpty != tc.partitions {
				t.Errorf("Expected partitions %v  got: %v", tc.partitions, tweet[len(tweet)-1].Label)
			}
		})
	}

}
