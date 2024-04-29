/* Package redmine implements the Redmine API
 *
 */
package redmine

import (
	"errors"
)

var (
	ErrServerEmpty   = errors.New("server is empty")
	ErrUserEmpty     = errors.New("user is empty")
	ErrPasswordEmpty = errors.New("password is empty")
	ErrForbidden     = errors.New("403 Forbidden")
)

type Redmine struct {
	auth Authenticator
}

func NewPublic(server string) *Redmine {
	return &Redmine{auth: &Public{server}}
}

// NewRegularLogin creates a RegularLogin instance.
//
// If the parameter "become" is set to a username, then the request includes the "X-Redmine-Switch-User: user" header to impersonate the given user.
func NewRegularLogin(server, login, password, become string) *Redmine {
	return &Redmine{auth: &RegularLogin{server, login, password, become}}
}

func NewAuthKey(server, key, become string) *Redmine {
	return &Redmine{auth: &AuthKey{server, key, become}}
}

func NewHeaderKey(server, key, become string) *Redmine {
	return &Redmine{auth: &HeaderKey{server, key, become}}
}
