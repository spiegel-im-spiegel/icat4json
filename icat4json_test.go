package icat4json

import "testing"

func TestGetW(t *testing.T) {
	json, err := Get(ToolICATW)
	if err != nil {
		t.Errorf("Get() error = \"%v\", want nil.", err)
	} else if json == "" {
		t.Errorf("Get() = \"\", not want \"\".")
	}
}

func TestGetH(t *testing.T) {
	json, err := Get(ToolICATH)
	if err != nil {
		t.Errorf("Get() error = \"%v\", want nil.", err)
	} else if json == "" {
		t.Errorf("Get() = \"\", not want \"\".")
	}
}
