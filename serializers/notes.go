package serializers

import (
	// "fmt"
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
	Attachments []AttachmentResponse `json:"attachments"`
}

type AttachmentResponse struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Image     string `json:"image"`
}

func (self *NoteSerializer) Response() NoteResponse {
	objs, _ := models.GetAttachments(
		models.AttachmentModel{NoteID: self.ID},
	)
	var attachments []AttachmentResponse
	for _, attachment := range objs {
		attachments = append(
			attachments, 
			AttachmentResponse{
				ID:        attachment.ID, 
				Image:     *attachment.Image,
				CreatedAt: attachment.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
				UpdatedAt: attachment.UpdatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
			},
		)
	}
	response := NoteResponse{
		ID:          self.ID,
		Note:        self.Note,
		CreatedAt:   self.CreatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		UpdatedAt:   self.UpdatedAt.UTC().Format("2006-01-02T15:04:05.999Z"),
		Attachments: attachments,
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
