package main

import (
	"bytes"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "mombestpicture"
	pathKey := CASPathTransformFunc(key)
	expectedFileName := "cf5d4b01c4d9438c22c56c832f83bd3e8c6304f9"
	expectedPathName := "cf5d4/b01c4/d9438/c22c5/6c832/f83bd/3e8c6/304f9"
	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathName)
	}
	if pathKey.FileName != expectedFileName {
		t.Errorf("have %s want %s", pathKey.FileName, expectedFileName)
	}

}

func TestStore(t *testing.T) {
	
	s := newStore()
	defer teardown(t, s)

	key := "foobar"
	data := []byte("some jpg bytes")

	if _, err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	_, r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}

	if ok := s.Has(key); !ok {
		t.Errorf("expected to have key %s", key)
	}

	b, err := io.ReadAll(r)
	if err != nil {
		t.Error(err)
	}
	if string(b) != string(data) {
		t.Errorf("want %s have %s", data, b)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}

	if ok := s.Has(key); ok {
		t.Errorf("expected to not have the key %s", key)
	}
}


func TestStoreDeleteKey(t *testing.T) {
	s := newStore()
	defer teardown(t, s)

	data := []byte("some jpg bytes")
	key := "mom_specials"

	if _, err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	return NewStore(opts)
}

func teardown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}