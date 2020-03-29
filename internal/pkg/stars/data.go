package stars

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

// Pixel is internal data of pixel
// Maximum grid size is 2^16 x 2^16
type Pixel struct {
	Px []uint16
	Py []uint16
}

// SimData contains simulation data
// Written by star-calc
// Read by star-sim
type SimData struct {
	Width    int
	Height   int
	Steps    int
	Model    string
	Velocity float64
	Time     time.Duration
	Summary  string
	Stars    []Pixel
}

// Write simdata to file
func Write(file string) error {
	f, err := json.MarshalIndent(Data, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(file, f, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Read simdata from file
func Read(file string, simdata *SimData) error {
	tmpdata, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmpdata, simdata)
	if err != nil {
		return err
	}
	return nil
}
