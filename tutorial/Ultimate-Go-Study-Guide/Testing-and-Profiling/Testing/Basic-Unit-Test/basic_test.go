package main

import (
	"net/http"
	"testing"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

func TestBasic(t *testing.T) {
	url := "https://www.google.com/"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tTest 0:\tWhen checking %q for status code %d", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to make the Get call : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to make the Get call.", succeed)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t%s\tShould receive a %d status code.", succeed, statusCode)
			} else {
				t.Errorf("\t%s\tShould receive a %d status code : %d", failed, statusCode, resp.StatusCode)
			}
		}
	}
}
