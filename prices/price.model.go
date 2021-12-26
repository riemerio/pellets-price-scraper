package prices

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Price struct {
	ID        primitive.ObjectID `bson:"_id"`
	RecordId  string             `bson:"recordId"`
	Price     float64            `bson:"price"`
	Timestamp time.Time          `bson:"timestamp"`
}
