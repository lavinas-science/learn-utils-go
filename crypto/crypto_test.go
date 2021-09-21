package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMd5(t *testing.T) {
	md5 := GetMd5("testing ofsshore")
	assert.NotNil(t, md5)
	assert.EqualValues(t, "b2f1e44edacad8e757928b8d5bbc233b", md5)
	md5 = GetMd5("")
	assert.EqualValues(t, "", md5)
}
