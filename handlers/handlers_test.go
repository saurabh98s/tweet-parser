package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBreakWords(t *testing.T) {
	// var tweet []models.Tweets
	tt := []struct {
		name     string
		payload  string
		response string
	}{
		{name: "long_string", payload: `After thoroughly study of humanitarian intervention over analysis of responsibility to protect the come to an end where all this argument leads to doctrine of responsibility to protect has not succeeded in establishing agreed upon conditions and methods for intervening in case of humanitarian crisis. The complete failure of United Nations peacekeeping and security council after numerous need the cries of slaughtered millions from show many years to current situation became worst slaughter in Biafra, Uganda, Burnudi, Indonesia, Bangladesh and Sudan in which current era Afghanistan, Russia Ukraine war we will finally say that legitimateness Of these mass `,
			response: `[{"tweet":"1/19:After thoroughly study of"},{"tweet":"2/19: humanitarian intervention over analysis of"},{"tweet":"3/19: responsibility to protect the come to"},{"tweet":"4/19: an end where all this argument leads"},{"tweet":"5/19: to doctrine of responsibility to"},{"tweet":"6/19: protect has not succeeded in"},{"tweet":"7/19: establishing agreed upon conditions and"},{"tweet":"8/19: methods for intervening in case of"},{"tweet":"9/19: humanitarian crisis. The complete failure"},{"tweet":"10/19: of United Nations peacekeeping and"},{"tweet":"11/19: security council after numerous need"},{"tweet":"12/19: the cries of slaughtered millions"},{"tweet":"13/19: from show many years to current"},{"tweet":"14/19: situation became worst slaughter in"},{"tweet":"15/19: Biafra, Uganda, Burnudi, Indonesia,"},{"tweet":"16/19: Bangladesh and Sudan in which"},{"tweet":"17/19: current era Afghanistan, Russia Ukraine"},{"tweet":"18/19: war we will finally say that"},{"tweet":"19/19: legitimateness Of these mass "}]`,

			},

		{name: "short_string", payload: `After thoroughly study of humanitarian intervention over analysis of responsibility to protect the come to an end where all this argument`,
			response: `[{"tweet":"1/3:After thoroughly study of humanitarian"},{"tweet":"2/3: intervention over analysis of responsibility to"},{"tweet":"3/3: protect the come to an end where all this argument"}]`,
			},
		{name: "no_string	", payload: "",
			response: `no data recieved in body`,
			},
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
			println( resBytes)
			if resBytes != tc.response {
				t.Errorf("Expected Response as: %v got: %v", tc.response, resBytes)
			}
			// err = json.Unmarshal([]byte(resBytes), &tweet)
			if err != nil {
				t.Errorf("Unmarshalling error: %v", err)
			}

		})
	}

}
