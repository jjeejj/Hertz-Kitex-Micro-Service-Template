package objid

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/id"
)

// FromID converts an id to objected id.
func FromID(id fmt.Stringer) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id.String())
}

// MustFromID converts an id to objected id, panics on error.
func MustFromID(id fmt.Stringer) primitive.ObjectID {
	oid, err := FromID(id)
	if err != nil {
		panic(err)
	}

	return oid
}

// ToTripID converts object id to trip id.
func ToTripID(oid primitive.ObjectID) id.TripID {
	return id.TripID(oid.Hex())
}
