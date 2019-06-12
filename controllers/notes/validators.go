package notes

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"

	"github.com/Cguilliman/post-it-note/common"
	"github.com/Cguilliman/post-it-note/models"
)

type NoteValidator struct {
	Note struct {
		Value string `form:"value" json:"value" binding:"exists,max=500"`
	} `json:"note"`
	Attachments      []*multipart.FileHeader `form:"attachments" json:"-" binding:"omitempty"`
	attachmentsFiles []string                `json:"-"`
	noteModel        models.NoteModel        `json:"-"`
}

type UpdateAttachments struct {
	ID   int                   `form:"id" json:"id" binding:"exists"`
	File *multipart.FileHeader `form:"file" json:"-" binding:"exists"`
}

func NewNoteCreationValidator() NoteValidator {
	return NoteValidator{}
}

func (self *NoteValidator) Bind(c *gin.Context) error {
	currentUser := c.MustGet("my_user_model").(models.UserModel)

	if err := common.Bind(c, self); err != nil {
		return err
	}
	self.noteModel.Note = self.Note.Value
	self.noteModel.OwnerID = currentUser.ID
	return nil
}

func (self *NoteValidator) CreateFiles(c *gin.Context) error {
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
		if err := c.SaveUploadedFile(file, "uploads/"+file.Filename); err != nil {
			return err
		}
		self.attachmentsFiles = append(self.attachmentsFiles, file.Filename)
	}
	return nil
}

func NewNoteUpdateValidator(exists models.NoteModel) NoteValidator {
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
}
