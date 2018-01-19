package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

func main() {
	// BEGIN MAIN OMIT
	z := New("file.zip")
	z.File("err.go")
	z.File("zip.go")

	if err := z.Close(); err != nil {
		log.Fatal(err)
	}
	// END MAIN OMIT
}

// BEGIN STRUCT OMIT
type Zip struct {
	zf *os.File
	zw *zip.Writer

	err error
}

// END STRUCT OMIT

// BEGIN NEW OMIT
func New(name string) Zip {
	f, err := os.Create(name)
	if err != nil {
		return Zip{err: err} // store error and return
	}
	return Zip{zf: f, zw: zip.NewWriter(f)}
}

// END NEW OMIT

// BEGIN FILE OMIT
func (z *Zip) File(name string) {
	if z.err != nil {
		return // skip if there was a prior error
	}

	r, err := os.Open(name)
	if err != nil {
		z.err = err
		return
	}
	defer r.Close()

	w, err := z.zw.Create(name)
	if err != nil {
		z.err = err
		return
	}
	if _, err := io.Copy(w, r); err != nil {
		z.err = err
		return
	}
}

// END FILE OMIT

// BEGIN CLOSE OMIT
func (z *Zip) Close() error {
	if z.err != nil {
		return z.err // return previous error
	}
	err := z.zw.Close()
	if err != nil {
		return err
	}
	return z.zf.Close()
}

// END CLOSE OMIT

// TODO: remove the zip if it was created
// and there was an error.
