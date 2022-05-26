package command

import (
	"encoding/base64"
	"strings"
)

// command separator variable
const separator string = "[/]"

// Builder is used to convert a command of type
//vstring of slice into a string to send over a REST API
type Builder struct {
	// base64 encoded command
	cmd string
}

// Marshal takes as a slice of type string and
// returns a pointer of tu
func Marshal(command []string) *Builder {
	// join the slice into a string using separator
	cmd := []byte(strings.Join(command, separator))

	// base64 encode the string
	return &Builder{
		cmd: base64.StdEncoding.EncodeToString(cmd),
	}
}

// String returns a base64 encoded command
func (r *Builder) String() string {
	return r.cmd
}

// Unmarshal reverts the command to its initial form
func (r *Builder) Unmarshal() ([]string, error) {
	// basee64 decode the command
	s, err := base64.StdEncoding.DecodeString(r.cmd)
	if err != nil {
		return nil, err
	}

	// split the string using separator
	return strings.Split(string(s), separator), nil
}
