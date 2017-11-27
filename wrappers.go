package wayland

import (
	"io"
)

func sizeOf_new_id(ObjectId) uint16 { return 4 }
func sizeOf_int(int32) uint16       { return 4 }
func sizeOf_uint(uint32) uint16     { return 4 }
func sizeOf_Fixed(Fixed) uint16     { return 4 }
func sizeOf_object(Object) uint16   { return 4 }
func sizeOf_fd(int) uint16          { return 0 }

func sizeOf_string(s string) uint16 {
	// XXX: we need to make sure this doesn't overflow somehow.
	return uint16(4 + ceil32(len(s)+1))
}

func writeU32(w io.Writer, val uint32) (int64, error) {
	var buf [4]byte
	hostEndian.PutUint32(buf[:], val)
	n, err := w.Write(buf[:])
	return int64(n), err
}

func write_new_id(w io.Writer, val ObjectId) (int64, error) { return writeU32(w, uint32(val)) }
func write_int(w io.Writer, val int32) (int64, error)       { return writeU32(w, uint32(val)) }
func write_uint(w io.Writer, val uint32) (int64, error)     { return writeU32(w, uint32(val)) }
func write_Fixed(w io.Writer, val Fixed) (int64, error)     { return writeU32(w, val.value) }
func write_object(w io.Writer, val Object) (int64, error)   { return writeU32(w, uint32(val.Id())) }
func write_fd(w io.Writer, val ObjectId) (int64, error)     { return 0, nil }

func write_string(w io.Writer, s string) (n int64, err error) {
	n, err = writeU32(w, uint32(len(s)))
	if err != nil {
		return
	}
	n_, err := w.Write([]byte(s))
	n += int64(n_)
	if err != nil {
		return
	}
	var padding [4]byte
	padding_size := ceil32(len(s)+1) - len(s)
	n_, err = w.Write(padding[:padding_size])
	n += int64(n_)
	return
}
