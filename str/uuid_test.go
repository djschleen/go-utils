package str

import (
	"regexp"
	"testing"

	"github.com/devops-kung-fu/go-util/str"
	"github.com/stretchr/testify/assert"
)

func Test_GetUUID(t *testing.T) {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	uuid := GetUUID()
	assert.Regexp(t, r, uuid, "Not a valid UUID")
}
