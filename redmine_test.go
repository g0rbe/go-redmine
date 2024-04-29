package redmine_test

import (
	"gorbe.io/go/redmine"
)

func ExampleNewRedmine_newPublic() {

	_ = redmine.NewRedmine(redmine.NewPublic("https://www.redmine.org"))

}

func ExampleNewRedmine_newRegularLogin() {

	// admin1 user impersonates admin2
	_ = redmine.NewRedmine(redmine.NewRegularLogin("https://www.redmine.org", "admin1", "password", "admin2"))

}

func ExampleNewRedmine_newAuthKey() {

	// User of apikey impersonates admin2 if user is admin
	_ = redmine.NewRedmine(redmine.NewAuthKey("https://www.redmine.org", "apikey", "admin2"))

}

func ExampleNewRedmine_newHeaderKey() {

	// User of apikey impersonates admin2 if user is admin
	_ = redmine.NewRedmine(redmine.NewHeaderKey("https://www.redmine.org", "apikey", "admin2"))

}
