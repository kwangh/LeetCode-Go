package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hoanhan101/ultimate-go/go/testing/web_server/handlers"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func init() {
	handlers.Routes()
}

func TestSendJSON(t *testing.T) {
	url := "/sendjson"
	statusCode := 200

	t.Log("Given the need to test the SendJSON endpoint.")
	{
		r := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)

		t.Logf("\tTest 0:\tWhen checking %q for status code %d", url, statusCode)
		{
			if w.Code != 200 {
				t.Fatalf("\t%s\tShould receive a status code of %d for the response. Received[%d].", failed, statusCode, w.Code)
			}
			t.Logf("\t%s\tShould receive a status code of %d for the response.", succeed, statusCode)

			var u struct {
				Name  string
				Email string
			}

			if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
				t.Fatalf("\t%s\tShould be able to decode the response.", failed)
			}
			t.Logf("\t%s\tShould be able to decode the response.", succeed)

			if u.Name == "Hoanh An" {
				t.Logf("\t%s\tShould have \"Hoanh An\" for Name in the response.", succeed)
			} else {
				t.Errorf("\t%s\tShould have \"Hoanh An\" for Name in the response : %q", failed, u.Name)
			}

			if u.Email == "hoanhan101@gmail.com" {
				t.Logf("\t%s\tShould have \"hoanhan101@gmail.com\" for Email in the response.", succeed)
			} else {
				t.Errorf("\t%s\tShould have \"hoanhan101@gmail.com\" for Email in the response : %q", failed, u.Email)
			}
		}
	}
}
