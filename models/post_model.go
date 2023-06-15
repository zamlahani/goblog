package models

import (
	"time"
)

type Post struct {
    ID              string             `json:"id,omitempty" bson:"_id,omitempty"`
    Title     		string             `json:"title,omitempty" validate:"required"`
    Body 			string             `json:"body,omitempty" validate:"required"`
    Slug 			string             `json:"slug,omitempty"`
    CreatedAt 		time.Time          `json:"createdAt,omitempty" validate:"required"`
    LastModified 	time.Time          `json:"lastModified,omitempty" validate:"required"`
}