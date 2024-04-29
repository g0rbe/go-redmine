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

func NewRedmine(a Authenticator) *Redmine {
	return &Redmine{auth: a}
}
