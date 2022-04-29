package FileUtil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type fileS struct {
	path     string
	fileName string
	mu       sync.Mutex
}

func New(p, f string) *fileS {
	return &fileS{fileName: f, path: p}
}

// Read string from fileS
func (f *fileS) Read() (string, error) {
	dat, err := ioutil.ReadFile(filepath.Join(f.path, f.fileName))
	return string(dat), err
}

// Write string to fileS
func (f *fileS) Write(s string) error {
	f.mu.Lock()

	f1, err := os.Create(filepath.Join(f.path, f.fileName))
	if err != nil {
		return err
	}
	defer func() {
		f1.Close()
		f.mu.Unlock()
	}()
	_, err = f1.WriteString(s)
	//f.mu.Unlock()
	if err != nil {
		return err
	}
	f1.Sync()
	//f1.Close()
	return nil
}
