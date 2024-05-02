package redmine_test

import (
	"gorbe.io/go/redmine"
)

func ExampleNewPublic() {

	_ = redmine.NewPublic("https://www.redmine.org")

}

func ExampleNewRegularLogin() {

	_ = redmine.NewRegularLogin("https://www.redmine.org", "admin1", "password")

}

func ExampleNewAuthKey() {

	_ = redmine.NewAuthKey("https://www.redmine.org", "apikey")

}

func ExampleNewHeaderKey() {

	_ = redmine.NewHeaderKey("https://www.redmine.org", "apikey")

}
