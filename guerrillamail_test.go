package guerrillamail

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

var client = NewGuerrillamailClient(http.DefaultClient)

func TestGetEmailAddress(t *testing.T) {
	emailaddr, err := client.GetEmailAddress(Argument{
		"lang": LANGUAGE_EN,
	})
	if err != nil {
		t.Error(err)
		return
	}

	body, err := json.MarshalIndent(emailaddr, "", "	")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(body))
}

func TestSetEmailUser(t *testing.T) {
	email, err := client.SetEmailUser(Argument{
		"email_user": "qianlnk",
		"lang":       LANGUAGE_EN,
	})
	if err != nil {
		t.Error(err)
		return
	}

	body, err := json.MarshalIndent(email, "", "	")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(string(body))
}

func TestCheckEmail(t *testing.T) {
	// emailaddr, err := client.GetEmailAddress(Argument{
	// 	"lang": LANGUAGE_EN,
	// })
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	// for {
	// 	resp, err := client.CheckEmail(Argument{
	// 		"seq": "1",
	// 	})
	// 	if err != nil {
	// 		t.Error(err)
	// 		continue
	// 	}

	// 	if len(resp.List) > 0 {

	// 	}
	// }
}
