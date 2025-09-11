package dependency_container

import (
	"testing"
)

func TestNewDIContainer(t *testing.T) {
	di := NewDIContainer()
	if di == nil {
		t.Error("NewDIContainer returned nil")
	}
	if di.Services == nil {
		t.Error("NewDIContainer's Services map should not be nil")
	}
}

func TestDependencyContainer_Add(t *testing.T) {
	di := NewDIContainer()
	di.Add("a", 100)

	if val, ok := di.Services["a"]; !ok || val.(int) != 100 {
		t.Error("Service a was not added correctly")
	}
}

func TestDependencyContainer_Get(t *testing.T) {
	di := NewDIContainer()
	di.Add("a", 100)

	// Test getting an existing service
	res, err := di.Get("a")
	if err != nil {
		t.Errorf("Get returned an error for an existing service: %v", err)
	}
	expected := 100
	if res == nil {
		t.Errorf("Get did not return the expected result: got %v, want %v", res, expected)
	}

	// Test getting a non-existing service
	_, err = di.Get("b")
	if err == nil {
		t.Error("Get should have returned an error for non-existing service")
	}
}

func TestDependencyContainer_MultiGet(t *testing.T) {
	di := NewDIContainer()
	di.Add("a", 100)
	di.Add("b", 200)

	// Test getting multiple existing services
	names := []string{"a", "b"}
	res, err := di.MultiGet(names)
	if err != nil {
		t.Errorf("MultiGet returned an error for existing services: %v", err)
	}
	if val, ok := res["a"]; !ok || val.(int) != 100 {
		t.Errorf("MultiGet did not return correct value for 'a': got %v, want %v", val, 100)
	}
	if val, ok := res["b"]; !ok || val.(int) != 200 {
		t.Errorf("MultiGet did not return correct value for 'b': got %v, want %v", val, 200)
	}

	// Test getting a partially non-existing services list
	names = []string{"a", "c"}
	_, err = di.MultiGet(names)
	if err == nil {
		t.Error("MultiGet should have returned an error for partially non-existing services")
	}

	// Test getting fully non-existing services list
	names = []string{"c", "d"}
	_, err = di.MultiGet(names)
	if err == nil {
		t.Error("MultiGet should have returned an error for fully non-existing services")
	}
}

//
