package announcement

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gaborszakacs/who-is-on-call/pagerduty"
)

type Announcer struct {
}

func (a Announcer) DoItWithDrama() error {
	client := pagerduty.Client{}
	theOne, err := client.OnCallThisWeek()
	if err != nil {
		return errors.New("Sorry, can't do it now.")
	}

	fmt.Printf("So the one\n")
	time.Sleep(1 * time.Second)
	fmt.Printf("who is on-call this week\n")
	time.Sleep(1 * time.Second)
	fmt.Printf("is...\n")
	time.Sleep(3 * time.Second)
	fmt.Printf(a.dramatize(theOne))
	return nil
}

func (a Announcer) dramatize(str string) string {
	return strings.ToUpper(str)
}
