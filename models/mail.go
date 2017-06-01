package models

// Mail notifies a user about an action in the system e.g:
// Register, reset password, reports and so on.
type Mail struct{}

// SendTo ...
func (m *Mail) SendTo(u *User) error {
	return nil
}
