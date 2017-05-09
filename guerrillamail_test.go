package guerrillamail

import (
	"fmt"
	"net/http"
	"testing"
)

var client = NewGuerrillamailClient(http.DefaultClient)

func TestGetEmailAddress(t *testing.T) {
	args := NewArgument()
	args.Set("lang", "zh")
	emailaddr, err := client.GetEmailAddress(args)
	fmt.Println(emailaddr, err)

	email, err := client.SetEmailUser(Argument{
		"email_user": "qianlnk",
		"lang":       LANGUAGE_EN,
		"sid_token":  emailaddr.SidToken,
	})

	fmt.Println(email, err)
}
