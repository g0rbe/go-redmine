package redmine_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"gorbe.io/go/redmine"
)

func TestCreateDeleteUser(t *testing.T) {

	nu := redmine.User{Login: fmt.Sprintf("Tester_%d", time.Now().Unix()), FirstName: "John", LastName: "Doe", Mail: fmt.Sprintf("%d@example.com", time.Now().Unix())}

	err := redmine.NewAuthKey(os.Getenv("REDMINE_SERVER"), os.Getenv("REDMINE_KEY"), "").CreateUser(&nu, false)
	if err != nil {
		t.Fatalf("Create: %s\n", err)
	}

	t.Logf("%v\n", nu)

	nu.MailNotification = redmine.MailNotificationNone
	err = redmine.NewAuthKey(os.Getenv("REDMINE_SERVER"), os.Getenv("REDMINE_KEY"), "").UpdateUser(&nu, false)
	if err != nil {
		t.Fatalf("Update: %s\n", err)
	}

	t.Logf("%v\n", nu)

	err = redmine.NewAuthKey(os.Getenv("REDMINE_SERVER"), os.Getenv("REDMINE_KEY"), "").DeleteUser(nu.ID)
	if err != nil {
		t.Fatalf("Delete: %s\n", err)
	}
}
