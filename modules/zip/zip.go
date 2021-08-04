package zip

import (
	"archive/zip"
	"github.com/dop251/goja"
	"os"
)

type Module struct {
	runtime *goja.Runtime
}

func New(verbose bool) *Module {
	return &Module{}
}

func (m *Module) WithWriter(output string, callback func(writer *Writer) error) error {
	writer, err := m.Writer(output)
	if err != nil {
		return err
	}

	err = callback(writer)
	if err != nil {
		writer.Close()
		return err
	}

	return writer.Close()
}

func (m *Module) Writer(output string) (*Writer, error) {
	f, err := os.Create(output)
	if err != nil {
		return nil, err
	}

	w := zip.NewWriter(f)

	return &Writer{
		writer: w,
		stream: f,
	}, err
}
