package serializers

import (
	"github.com/Cguilliman/post-it-note/models"
	"github.com/gin-gonic/gin"
)

type NoteSerializer struct {
	C *gin.Context
	models.NoteModel
}

type NoteResponse struct {
	ID        uint   `json:"id"`
	Note      string `json:"note"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	// Owner     UserResponse
}

func (self *NoteSerializer) Response() NoteResponse {
	response := NoteResponse{
		ID:        self.ID,
		Note:      self.Note,
		CreatedAt: self.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		UpdatedAt: self.UpdatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
	}
	return response
}

type NotesSerializer struct {
	C     *gin.Context
	Notes []models.NoteModel
}

func (self *NotesSerializer) Response() []NoteResponse {
	var response []NoteResponse
	for _, note := range self.Notes {
		noteSerializer := NoteSerializer{self.C, note}
		response = append(response, noteSerializer.Response())
	}
	return response
}
