package battlesshiplib

import (
	"fmt"
	"testing"
	"time"
)

func TestBattlesshiplib_Connect(t *testing.T) {
	r := NewRedis()
	r.Ping()
	r.Set("yo", "meo", 1)
	value, _ := r.Get("yo")
	if value != "meo" {
		t.Errorf("Expected value: meo, got: %v", value)
	}
	defer r.rdb.Close()
}

func TestBattlesshiplib_EncodeDecode(t *testing.T) {
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

func TestBattlesshiplib_SaveWithEncoding(t *testing.T) {
	r := NewRedis()

	coordinate := Coordinate{X: 5, Y: 3}

	// Encode to []byte
	data, err := Encode(coordinate)
	if err != nil {
		fmt.Println("Error encoding:", err)
		return
	}

	r.Set("coordinate", string(data), time.Second*1)

	retrievedData, err := r.Get("coordinate")

	// Cast retrievedData to []byte
	decodedData := []byte(retrievedData)

	// Simulate reading `data` from Redis and decoding it
	decodedCoordinate, err := Decode[Coordinate](decodedData)
	if err != nil {
		t.Errorf("Expected value to be a Coordinate, got: %v", decodedCoordinate)
	}

	if decodedCoordinate.X != 5 || decodedCoordinate.Y != 3 {
		t.Errorf("Expected value: meo, got: %v", decodedCoordinate)
	}

	fmt.Printf("Decoded coordinate: %+v\n", decodedCoordinate)
}
