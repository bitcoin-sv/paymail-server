package data

import (
	"embed"
	"os"
	"path"

	"github.com/pkg/errors"
)

// Directory a directory container capability data.
type Directory struct {
	prefix string
	fs     embed.FS
}

//go:embed capabilities/*
var capabilityData embed.FS

// CapabilitiesData data for creating spv envelopes.
var CapabilitiesData = Directory{
	prefix: "capabilities",
	fs:     capabilityData,
}

// LoadAll will return each file in the folder as a slice of bytes.
func (d *Directory) LoadAll() ([][]byte, error) {
	wholeDir, err := d.fs.ReadDir(d.prefix)
	if err != nil {
		return nil, err
	}
	allFileData := make([][]byte, 0)
	for _, file := range wholeDir {
		fileData, err := d.fs.ReadFile(path.Join(d.prefix, file.Name()))
		if err != nil {
			return nil, err
		}
		allFileData = append(allFileData, fileData)
	}
	return allFileData, nil
}

// LoadStaticDocument will return the generated document if it exists.
func (d *Directory) LoadStaticDocument() ([]byte, error) {
	data, err := d.fs.ReadFile(d.prefix + "/capabilities.json")
	if err != nil {
		return nil, errors.Wrap(err, "you have to run the setup to generate the static capabilities.json")
	}
	return data, nil
}

// OverwriteStaticCapabilitiesFile is self explanitory, however the important note is that this filepath is relative.
// The capabilities_test.go and main.go must remain in the same order of nested folders otherwise this filepath will fail.
func OverwriteStaticCapabilitiesFile(data []byte) error {
	path := "../../data/capabilities/capabilities.json"
	_ = os.Remove(path)
	return os.WriteFile(path, data, 0600)
}
