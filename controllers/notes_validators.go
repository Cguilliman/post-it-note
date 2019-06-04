package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/Cguilliman/post-it-note/common"
	"github.com/Cguilliman/post-it-note/models"
)

type NoteCreationValidator struct {
	Note struct {
		Value string `form:"value" json:"value" binding:"exists,max=500"`
	} `json:"note"`
	noteModel models.NoteModel `json:"-"`
}

func NewNoteCreationValidator() NoteCreationValidator {
	return NoteCreationValidator{}
}

func (self *NoteCreationValidator) Bind(c *gin.Context) error {
	currentUser := c.MustGet("my_user_model").(models.UserModel)

	if err := common.Bind(c, self); err != nil {
		return err
	}
	self.noteModel.Note = self.Note.Value
	self.noteModel.OwnerID = currentUser.ID
	return nil
}

func NewNoteCreationValidatorFillWith(exists models.NoteModel) NoteCreationValidator {
	validator := NewNoteCreationValidator()
	validator.noteModel.Note = exists.Note
	validator.noteModel.CreatedAt = exists.CreatedAt
	validator.noteModel.DeletedAt = exists.DeletedAt
	validator.noteModel.OwnerID = exists.OwnerID
	return validator
}

func IsNoteOwner(c *gin.Context, note models.NoteModel) bool {
	user := c.MustGet("my_user_model").(models.UserModel)
	return note.OwnerID == user.ID
	// return false
}
