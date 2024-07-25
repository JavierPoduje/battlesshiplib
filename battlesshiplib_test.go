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
	coordinate := Coordinate{X: 1, Y: 2}

	// Encode to []byte
	data, err := Encode(coordinate)
	if err != nil {
		fmt.Println("Error encoding:", err)
		return
	}

	// Simulate reading `data` from Redis and decoding it
	decodedCoordinate, err := Decode[Coordinate](data)
	if err != nil {
		t.Errorf("Expected value to be a Coordinate, got: %v", decodedCoordinate)
	}

	if decodedCoordinate.X != 1 || decodedCoordinate.Y != 2 {
		t.Errorf("Expected value: meo, got: %v", decodedCoordinate)
	}

	fmt.Printf("Decoded coordinate: %+v\n", decodedCoordinate)
}
