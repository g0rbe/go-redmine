/* Package redmine implements the Redmine API
 *
 */
package redmine

import (
	"errors"
)

var (
	ErrServerEmpty         = errors.New("server is empty")
	ErrUserEmpty           = errors.New("user is empty")
	ErrPasswordEmpty       = errors.New("password is empty")
	ErrLoginEmpty          = errors.New("login is empty")
	ErrFirstnameEmpty      = errors.New("firstname is empty")
	ErrLastnameEmpty       = errors.New("lastname is empty")
	ErrMailEmpty           = errors.New("mail is empty")
	ErrIDEmpty             = errors.New("id is empty")
	ErrForbidden           = errors.New("403 Forbidden")
	ErrUnprocessableEntity = errors.New("422 Unprocessable Entity")
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
func NewRegularLogin(server, login, password string) *Redmine {
	return &Redmine{auth: &RegularLogin{server: server, login: login, password: password}}
}

func NewAuthKey(server, key string) *Redmine {
	return &Redmine{auth: &AuthKey{server: server, key: key}}
}

func NewHeaderKey(server, key string) *Redmine {
	return &Redmine{auth: &HeaderKey{server: server, key: key}}
}

func (r *Redmine) Impersonate(name string) {
	r.auth.Impersonate(name)
}
