package store

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/avelex/passwordme/internal/store/passfile"
)

func (s *PasswordStore) firstPasswordFilename() (string, error) {
	entries := ls(s.rootPath)
	for _, v := range entries {
		if strings.HasSuffix(v, passfile.EXTENSION) {
			return v, nil
		}
	}
	return "", errors.New("password file not found")
}

func detectPassfiles(dirPath string) []*passfile.PasswordFile {
	files := ls(dirPath)
	passfiles := make([]*passfile.PasswordFile, 0, len(files))

	for _, v := range files {
		if strings.HasSuffix(v, passfile.EXTENSION) {
			path := filepath.Join(dirPath, v)
			passfiles = append(passfiles, passfile.New(path))
		}
	}

	return passfiles
}

func ls(dirPath string) []string {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return []string{}
	}

	result := make([]string, 0, len(entries))
	for _, entry := range entries {
		result = append(result, entry.Name())
	}

	return result
}
