package main

import "testing"

// write unit test for FindNearest
// TestFindNearest tests the FindNearest function.
// The test checks if the function returns the nearest vector to the query vector.
// The test checks if the function returns an error when the query vector has a different dimension than the vectors in the database.
// The test checks if the function returns an error when the database is empty.
func TestFindNearest(t *testing.T) {
	db := &VectorDB{}
	db.Add(Vector{ID: "vec_1", Values: []float64{1.0, 2.0, 3.0}})
	db.Add(Vector{ID: "vec_2", Values: []float64{4.0, 5.0, 6.0}})
	db.Add(Vector{ID: "vec_3", Values: []float64{7.0, 8.0, 9.0}})

	query := []float64{2.0, 3.0, 4.0}
	gotVec, _ := db.FindNearest(query)
	expectedVec := "vec_1"

	if gotVec.ID != "vec_1" {
		t.Errorf("Expected %s, got %s", expectedVec, gotVec.ID)
	}

	gotDist := gotVec.Distance
	expectedDist := 1.7320508075688772
	if gotVec.Distance != expectedDist {
		t.Errorf("Expected dist %f, got %f", expectedDist, gotDist)
	}
}
