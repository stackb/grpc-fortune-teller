package fileutil

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

// Exists - return true if a file entry exists
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// ListFiles - convenience debugging function to log the files under a given dir
func ListFiles(dir string) error {
	log.Println("Listing files under " + dir)
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("%v\n", err)
			return err
		}
		log.Println(path)
		return nil
	})
}

// MustMkdirAll - make all dirs and panic if fail
func MustMkdirAll(dirs ...string) {
	for _, dir := range dirs {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("Failed mkdir %s: %v", dir, err))
		}
	}
}

// CopyFile - copy bytes from one file to another
func CopyFile(src, dst string) error {
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return fmt.Errorf("CopyFile: srcFile not found: %s", src)
	}

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return out.Close()
}

// MustRestore - Restore assets.
func MustRestore(baseDir string, assets map[string][]byte, mappings map[string]string) {
	// unpack variable is provided by the go_embed data and is a
	// map[string][]byte such as {"/usr/share/games/fortune/literature.dat":
	// bytes... }
	for basename, bytes := range assets {
		if mappings != nil {
			replacement := mappings[basename]
			if replacement != "" {
				basename = replacement
			}
		}
		filename := path.Join(baseDir, basename)
		dirname := path.Dir(filename)
		//log.Printf("file %s, dir %s, rel %d, abs %s, absdir: %s", file, dir, rel, abs, absdir)
		if err := os.MkdirAll(dirname, os.ModePerm); err != nil {
			log.Fatalf("Failed to create asset dir %s: %v", dirname, err)
		}

		if err := ioutil.WriteFile(filename, bytes, os.ModePerm); err != nil {
			log.Fatalf("Failed to write asset %s: %v", filename, err)
		}
		log.Printf("Restored %s", filename)
	}

	log.Printf("Assets restored to %s", baseDir)
}
