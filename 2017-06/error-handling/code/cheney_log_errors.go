import (
	"io"
	"log"
)

func Write(w io.Writer, buf []byte) error {
	_, err := w.Write(buf)
	if err != nil {
		// annotated error goes to log file
		log.Println("unable to write:", err) // HL

		// unannotated error returned to caller
		return err // HL
	}
	return nil
}
