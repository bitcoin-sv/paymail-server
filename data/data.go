package data

import (
	"embed"
	"path"
)

// DataDir a directory container capability data.
type DataDir struct {
	prefix string
	fs     embed.FS
}

//go:embed capabilities/*
var capabilityData embed.FS

// SpvCreateData data for creating spv envelopes.
var CapabilitiesData = DataDir{
	prefix: "capabilities",
	fs:     capabilityData,
}

// LoadAll will return each file in the folder as a slice of bytes.
func (d *DataDir) LoadAll() ([][]byte, error) {
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
