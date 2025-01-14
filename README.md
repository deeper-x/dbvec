# VectorDB

VectorDB is a simple in-memory database for storing vectors and finding the nearest vector to a query vector using Euclidean distance.

## Installation

To install the project, clone the repository and navigate to the project directory:

```sh
git clone https://github.com/deeper-x/dbvec.git
cd dbvec
```
#### Example

```
package main

import (
    "log"
)

func main() {
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
    // OUTPUT: Nearest vector: vec_1, Distance: 1.732051
}
```

Testing
To run the tests, use the go test command:
```sh
go test -v ./...
```

# Details
```Vector``` Represents a vector with an ID, values, and distance.

```VectorDB``` Represents a database of vectors.


```Add(v Vector)``` Adds a vector to the database.

```FindNearest(query []float64) (Vector, error)``` Returns the nearest vector to the query vector and the distance between the query vector and the nearest vector.

```EuclideanDistance(v1, v2 []float64) (float64, error)``` Calculates the Euclidean distance between two vectors.

# License
This project is licensed under the MIT License.