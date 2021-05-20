package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
)

type Marshaling struct {
	path string
}

var PersistenceService = Marshaling{}

func (m *Marshaling) SetPath() {
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("%v", err)
	}

	m.path = homePath + "/.rdv-sante-data"
}

func (m *Marshaling) Save(v interface{}) error {
	f, err := os.Create(m.path)
	if err != nil {
		return err
	}
	defer f.Close()

	var Marshal = func(v interface{}) (io.Reader, error) {
		b, err := json.MarshalIndent(v, "", "\t")
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(b), nil
	}

	r, err := Marshal(v)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, r)
	return err
}

func (m *Marshaling) Retrieve(ref interface{}) error {
	var Unmarshal = func(r io.Reader, v interface{}) error {
		return json.NewDecoder(r).Decode(v)
	}

	f, err := os.Open(m.path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	defer f.Close()
	return Unmarshal(f, &ref)
}
