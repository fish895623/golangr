package filemanage_test

import (
	"testing"

	filemanage "github.com/fish895623/golangr/filemanage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_read_write_file(t *testing.T) {
	t.Run("Exists", func(t *testing.T) {
		val := filemanage.CheckFileExists()
		assert.Equal(t, true, val)
	})
	t.Run("Create", func(t *testing.T) {
		filemanage.CreateFile()
		assert.True(t, true, filemanage.CheckFileExists())
	})
	t.Run("Write", func(t *testing.T) {
		input := uuid.New().String()
		filemanage.Writefile(input)
		val := filemanage.Readfile()

		assert.Equal(t, input, val)
	})
	t.Run("Delete", func(t *testing.T) {
		assert.True(t, filemanage.DeleteFile())
	})
}
