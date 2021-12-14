package data

import (
	"embed"
	"os"
	"path"
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

// OverwriteFile is self explanitory
func (d *Directory) OverwriteFile(name string, data []byte) error {
	_ = os.Remove(d.prefix + "/" + name)
	return os.WriteFile("capabilities.json", data, 0600)
}
