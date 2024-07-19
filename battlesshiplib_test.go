package battlesshiplib

import (
	"fmt"
	"testing"
)

func TestBattlesshiplib_Connect(t *testing.T) {
	r := NewRedis()
	r.Ping()
	r.Set("yo", "meo")
	value := r.Get("yo")
	if value != "meo" {
		t.Errorf("Expected value: meo, got: %v", value)
	}
	defer r.rdb.Close()
}

func TestBattlesshiplib_SaveWithEncoding(t *testing.T) {
	r := NewRedis()

	coord := Coordinate{X: 1, Y: 2}
	encodedCoord, err := r.Encode(coord)
	if err != nil {
		t.Errorf("Error encoding coord: %v", err)
	}
	fmt.Println("Encoded coord:", encodedCoord)

	r.Set("yo", encodedCoord)
	value := r.Get("yo")
	fmt.Println("Value:", value)
	byteValue, ok := value.([]byte)
	fmt.Println("byteValue:", byteValue)
	if !ok {
		t.Errorf("Expected value to be a byte slice, got: %v", byteValue)
	}

	decodedCoord, err := r.Decode(byteValue)
	if err != nil {
		t.Errorf("Error decoding coord: %v", err)
	}

	correctlyTypedCoord, ok := decodedCoord.(Coordinate)
	if !ok {
		t.Errorf("Expected value to be a Coordinate, got: %v", decodedCoord)
	}

	if correctlyTypedCoord.X != 1 || correctlyTypedCoord.Y != 2 {
		t.Errorf("Expected value: meo, got: %v", value)
	}

	defer r.rdb.Close()
}
