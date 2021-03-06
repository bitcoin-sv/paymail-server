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

//go:embed *
var data embed.FS

// CapabilitiesData data for creating spv envelopes.
var CapabilitiesData = Directory{
	prefix: "capabilities",
	fs:     data,
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
	data, err := os.ReadFile("/app/data/capabilities.json")
	if err != nil {
		return nil, errors.Wrap(err, "you have to run the setup to generate the static capabilities.json")
	}
	return data, nil
}

// LoadStaticDocumentV1 will return the generated document if it exists.
func (d *Directory) LoadStaticDocumentV1() ([]byte, error) {
	data, err := os.ReadFile("/app/data/capabilities_v1.json")
	if err != nil {
		return nil, errors.Wrap(err, "you have to run the setup to generate the static capabilities")
	}
	return data, nil
}

// OverwriteStaticCapabilitiesFile is self explanitory, however the important note is that this filepath is relative.
// Run server must be done from root directory of the package for this reason.
func OverwriteStaticCapabilitiesFile(data []byte) error {
	path := "/app/data/capabilities.json"
	_ = os.Remove(path)
	return os.WriteFile(path, data, 0644)
}

// OverwriteStaticCapabilitiesFile is self explanitory, however the important note is that this filepath is relative.
// Run server must be done from root directory of the package for this reason.
func OverwriteStaticCapabilitiesFileV1(data []byte) error {
	path := "/app/data/capabilities_v1.json"
	_ = os.Remove(path)
	return os.WriteFile(path, data, 0644)
}
