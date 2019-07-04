package messages

import uuid "github.com/satori/go.uuid"

type DetailRequest struct {
	ID uuid.UUID `json:"id" binding:"required"`
}

