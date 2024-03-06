package file

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/celestiaorg/celestia-node/share"
)

const HeaderSize = 64

type Header struct {
	version fileVersion

	// Taken directly from EDS
	shareSize  uint16
	squareSize uint16

	// TODO(@walldiss) store all heights in the header?
	//heightы   []uint64
	datahash share.DataHash
}

type fileVersion uint8

const (
	FileV0 fileVersion = iota
)

func (h *Header) Version() fileVersion {
	return h.version
}

func (h *Header) ShareSize() int {
	return int(h.shareSize)
}

func (h *Header) SquareSize() int {
	return int(h.squareSize)
}

func (h *Header) DataHash() share.DataHash {
	return h.datahash
}

func (h *Header) WriteTo(w io.Writer) (int64, error) {
	buf := make([]byte, HeaderSize)
	buf[0] = byte(h.version)
	binary.LittleEndian.PutUint16(buf[1:3], h.shareSize)
	binary.LittleEndian.PutUint16(buf[3:5], h.squareSize)
	copy(buf[5:37], h.datahash)
	_, err := io.Copy(w, bytes.NewBuffer(buf))
	return HeaderSize, err
}

func ReadHeader(r io.Reader) (*Header, error) {
	buf := make([]byte, HeaderSize)
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}

	h := &Header{
		version:    fileVersion(buf[0]),
		shareSize:  binary.LittleEndian.Uint16(buf[1:3]),
		squareSize: binary.LittleEndian.Uint16(buf[3:5]),
		datahash:   make([]byte, 32),
	}

	copy(h.datahash, buf[5:37])
	return h, err
}
