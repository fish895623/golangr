package filemanage_test

import (
	"testing"

	filemanage "github.com/fish895623/golangr/filemanage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_read_write_file(t *testing.T) {
	input := uuid.New().String()
	filemanage.Writefile(input)
	val := filemanage.Readfile()

	assert.Equal(t, input, val)
}
