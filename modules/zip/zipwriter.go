package zip

import (
	"archive/zip"
	"io"
	"os"
)

type Writer struct {
	writer *zip.Writer
	stream io.WriteCloser
}

func (w *Writer) Close() error {
	err := w.writer.Close()
	if err != nil {
		return err
	}
	return w.stream.Close()
}

func (w *Writer) AddFile(path, dest string) error {
	fileToZip, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
	header.Name = dest

	// Change to deflate to gain better compression
	// see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	writer, err := w.writer.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}