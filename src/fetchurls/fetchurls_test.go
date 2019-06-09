package main

import (
  "testing"
)

func TestReadFile(t *testing.T) {
    contents := ReadFile("SampleListOfUrls.txt")
    if contents == "" {
        t.Errorf("File not read")
    }
}

func TestSliceString(t *testing.T) {
    sliced := SliceString("joe bloggs\njane doe\nandrew smith\n")
    var count int = len(sliced)
    if count != 4 {
        t.Errorf("Expected %d in slice, got %d", 4, count)
    }
}
