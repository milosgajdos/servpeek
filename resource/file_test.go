package resource

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	assert := assert.New(t)

	f, err := ioutil.TempFile("", "regular")
	assert.NoError(err)

	expectedPath := f.Name()

	file := NewFile(expectedPath)
	assert.Equal(expectedPath, file.Path())
	assert.Equal("[File] "+expectedPath, file.String())

	fi, err := file.Info()
	assert.NoError(err)
	assert.NotNil(fi)
	assert.NoError(os.Remove(file.Path()))
}
