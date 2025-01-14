package main

import (
	"errors"
	"log"
	"math"
)

// Vector represents a vector with an ID, values, and distance.
// The ID is a unique identifier for the vector.
// The values are the components of the vector.
// The distance is the distance between the vector and another vector.
// The distance is used to find the nearest vector to a query vector.
// The distance is set when finding the nearest vector to a query vector.
type Vector struct {
	ID       string
	Values   []float64
	Distance float64
}

// VectorDB represents a database of vectors.
// The database stores vectors and provides a method to find the nearest vector to a query vector.
// The database uses Euclidean distance to calculate the distance between vectors.
type VectorDB struct {
	vects []Vector
}

// Add adds a vector to the database.
// The time complexity of this function is O(1).
// The space complexity of this function is O(1).
// The space complexity of the VectorDB is O(n).
func (vdb *VectorDB) Add(v Vector) {
	vdb.vects = append(vdb.vects, v)
}

// FindNearest returns the nearest vector to the query vector
// and the distance between the query vector and the nearest vector.
// It uses Euclidean distance to calculate the distance between vectors.
// The time complexity of this function is O(n) where n is the number of vectors in the database.
// The space complexity of this function is O(1).
// The space complexity of the VectorDB is O(n).
// The space complexity of the Vector is O(1).
// The space complexity of the query vector is O(1).
// The space complexity of the distance is O(1).
func (vdb *VectorDB) FindNearest(query []float64) (Vector, error) {
	nearest := Vector{
		Distance: math.MaxFloat64,
	}

	for _, v := range vdb.vects {
		dist, err := EuclideanDistance(query, v.Values)
		if err != nil {
			return nearest, err
		}

		if dist < nearest.Distance {
			nearest = v
			nearest.Distance = dist
		}
	}

	return nearest, nil
}

// EuclideanDistance calculates the Euclidean distance between two vectors.
// The time complexity of this function is O(n) where n is the length of the vectors.
// The space complexity of this function is O(1).
// The space complexity of the first vector is O(n).
// The space complexity of the second vector is O(n).
func EuclideanDistance(v1, v2 []float64) (float64, error) {
	if len(v1) != len(v2) {
		return 0.0, errors.New("vector must be of equal length")
	}

	var sum float64
	for i := 0; i < len(v1); i++ {
		diff := v1[i] - v2[i]

		sum += diff * diff
	}

	return math.Sqrt(sum), nil
}

func main() {
	// Example usage
	// Create a VectorDB and add some vectors to it
	// Find the nearest vector to a query vector
	// and the distance between the query vector and the nearest vector
	db := &VectorDB{}
	db.Add(Vector{ID: "vec_1", Values: []float64{1.0, 2.0, 3.0}})
	db.Add(Vector{ID: "vec_2", Values: []float64{4.0, 5.0, 6.0}})
	db.Add(Vector{ID: "vec_3", Values: []float64{7.0, 8.0, 9.0}})

	query := []float64{2.0, 3.0, 4.0}
	nearest, err := db.FindNearest(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Nearest vector: %s, Distance: %f", nearest.ID, nearest.Distance)
}
