package file

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/stretchr/testify/assert"
)

type mockedFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
	sys     interface{}
}

func (fi mockedFileInfo) Name() string       { return fi.name }
func (fi mockedFileInfo) Size() int64        { return fi.size }
func (fi mockedFileInfo) Mode() os.FileMode  { return fi.mode }
func (fi mockedFileInfo) ModTime() time.Time { return fi.modTime }
func (fi mockedFileInfo) IsDir() bool        { return fi.isDir }
func (fi mockedFileInfo) Sys() interface{}   { return fi.sys }

type mockedFile struct {
	path, str string
	info      os.FileInfo
	reader    io.Reader
	err       error
}

func (f mockedFile) Path() string               { return f.path }
func (f mockedFile) String() string             { return f.str }
func (f mockedFile) Info() (os.FileInfo, error) { return f.info, f.err }
func (f mockedFile) Reader() io.Reader          { return f.reader }

func TestCheckOrError(t *testing.T) {
	assert := assert.New(t)
	expectedError := errors.New("Hi there!")

	assert.NoError(isTrueOrError(true, expectedError))

	err := isTrueOrError(false, expectedError)
	assert.Error(err)
	assert.Equal(expectedError, err)
}

func TestIsRegular(t *testing.T) {
	assert := assert.New(t)

	f, err := ioutil.TempFile("", "regular")
	assert.NoError(err)

	file := resource.NewFile(f.Name())
	assert.NoError(IsRegular(file))
	assert.NoError(os.Remove(file.Path()))
}

func TestIsDirectory(t *testing.T) {
	assert := assert.New(t)

	path, err := ioutil.TempDir("", "directory")
	assert.NoError(err)

	file := resource.NewFile(path)
	assert.NoError(IsDirectory(file))
	assert.NoError(os.Remove(file.Path()))
}

func TestIsBlockDevice(t *testing.T) {
	fi := mockedFileInfo{mode: os.ModeDevice}
	f := mockedFile{info: fi}
	assert.NoError(t, IsBlockDevice(f))
}

func TestIsCharDevice(t *testing.T) {
	fi := mockedFileInfo{mode: os.ModeCharDevice}
	f := mockedFile{info: fi}
	assert.NoError(t, IsCharDevice(f))
}

func TestIsPipe(t *testing.T) {
	fi := mockedFileInfo{mode: os.ModeNamedPipe}
	f := mockedFile{info: fi}
	assert.NoError(t, IsPipe(f))
}

func TestIsSocket(t *testing.T) {
	fi := mockedFileInfo{mode: os.ModeSocket}
	f := mockedFile{info: fi}
	assert.NoError(t, IsSocket(f))
}

func TestIsSymLink(t *testing.T) {
	fi := mockedFileInfo{mode: os.ModeSymlink}
	f := mockedFile{info: fi}
	assert.NoError(t, IsSymlink(f))
}

func TestIsModeOk(t *testing.T) {
	fi := mockedFileInfo{mode: os.ModeSticky} // non listed mode
	f := mockedFile{str: "filename", info: fi}

	assert.NoError(t, IsMode(f, os.ModeSticky))
}

func TestIsModeErrorMessages(t *testing.T) {
	assert := assert.New(t)

	fi := mockedFileInfo{mode: 0}
	f := mockedFile{str: "filename", info: fi}

	baseErrMessage := "filename file is not "
	cases := map[os.FileMode]string{
		os.ModeSymlink:   "a symlink",
		os.ModeSocket:    "a socket",
		os.ModeDir:       "a directory",
		os.ModeNamedPipe: "a named pipe",
		ModeRegular:      "a regular file",
	}

	for mode, expectedVerboseName := range cases {
		err := IsMode(f, mode)
		assert.Error(err)
		assert.Equal(baseErrMessage+expectedVerboseName, err.Error())
	}
}

func TestSize(t *testing.T) {
	fi := mockedFileInfo{size: 100}
	f := mockedFile{info: fi}

	assert.NoError(t, IsSize(f, 100))
}

func TestModTimeAfter(t *testing.T) {
	tstTime := time.Now()
	fi := mockedFileInfo{modTime: tstTime.Add(5 * time.Minute)}
	f := mockedFile{info: fi}

	assert.NoError(t, ModTimeAfter(f, tstTime))
}

func TestContains(t *testing.T) {
	assert := assert.New(t)
	data := `line1
	line2`

	f, err := ioutil.TempFile("", "regular")
	assert.NoError(err)
	_, err = f.Write([]byte(data))
	assert.NoError(err)

	file := resource.NewFile(f.Name())
	re := regexp.MustCompile(`line2$`)
	assert.NoError(Contains(file, re))
	assert.NoError(os.Remove(file.Path()))
}

func TestMD5Equal(t *testing.T) {
	assert := assert.New(t)
	data := `line1
	line2`

	f, err := ioutil.TempFile("", "regular")
	assert.NoError(err)
	_, err = f.Write([]byte(data))
	assert.NoError(err)

	file := resource.NewFile(f.Name())
	assert.NoError(MD5Equal(file, "3dfc125676228ddbac790f3b6d8d58be"))
	assert.NoError(os.Remove(file.Path()))
}

func TestSHA256Equal(t *testing.T) {
	assert := assert.New(t)
	data := `line1
	line2`

	f, err := ioutil.TempFile("", "regular")
	assert.NoError(err)
	_, err = f.Write([]byte(data))
	assert.NoError(err)

	file := resource.NewFile(f.Name())
	expSum := "2f6928c43c919915d452b6f2b90f7cf6640a7773c83412bf3d8ea1abfc699020"
	assert.NoError(SHA256Equal(file, expSum))
	assert.NoError(os.Remove(file.Path()))
}
