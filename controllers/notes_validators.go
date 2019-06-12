package controllers

import (
	// "fmt"
	"mime/multipart"
	"github.com/gin-gonic/gin"

	"github.com/Cguilliman/post-it-note/common"
	"github.com/Cguilliman/post-it-note/models"
)

type NoteCreationValidator struct {
	Note struct {
		Value string `form:"value" json:"value" binding:"exists,max=500"`
	} `json:"note"`
	Attachments      []*multipart.FileHeader `form:"attachments" json:"-" binding:"omitempty"`
	attachmentsFiles []string                `json:"-"`
	noteModel        models.NoteModel        `json:"-"`
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

func (self *NoteCreationValidator) CreateFiles(c *gin.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files, ok := form.File["attachments"]
	if !ok {
		return nil
	}
	for _, file := range files {
		// TODO: remove file uploading in utils
		if err := c.SaveUploadedFile(file, "saved/"+file.Filename); err != nil {
			return err
		}
		self.attachmentsFiles = append(self.attachmentsFiles, file.Filename)
	}
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
