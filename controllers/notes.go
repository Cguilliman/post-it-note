package controllers

import (
	"errors"
	"fmt"
	"reflect"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/Cguilliman/post-it-note/common"
	"github.com/Cguilliman/post-it-note/models"
	"github.com/Cguilliman/post-it-note/serializers"
)

func getNote(c *gin.Context) (models.NoteModel, error) {
	var (
		err         error
		notePK      uint64
		currentNote models.NoteModel
	)
	if notePK, err = strconv.ParseUint(c.Param("pk"), 10, 32); err != nil {
		return currentNote, err
	}
	if currentNote, err = models.GetNote(&models.NoteModel{ID: uint(notePK)}); err != nil {
		return currentNote, err
	}
	return currentNote, err
}

func getOwnerNote(c *gin.Context) (models.NoteModel, bool) {
	currentNote, err := getNote(c)
	if err != nil || !IsNoteOwner(c, currentNote) {
		c.JSON(http.StatusNotFound, errors.New("Permission denied"))
		return currentNote, true
	}
	return currentNote, false
}

func NoteRetrieve(c *gin.Context) {
	note, isExit := getOwnerNote(c)
	if isExit {
		return
	}
	serializer := serializers.NoteSerializer{c, note}
	c.JSON(http.StatusOK, gin.H{"note": serializer.Response()})
}

func NotesList(c *gin.Context) {
	// myUserModel := c.MustGet("my_user_model").(models.UserModel)
	notes, err := models.GetNotes(&models.NoteModel{
		// OwnerID: myUserModel.ID,
		// OwnerID: myUserModel.ID,
	})
	if err != nil {
		fmt.Println(err)
	}

	serializer := serializers.NotesSerializer{c, notes}
	c.JSON(http.StatusOK, gin.H{"nodes": serializer.Response()})
}

func NoteCreate(c *gin.Context) {
	validator := NewNoteCreationValidator()
	if err := validator.Bind(c); err != nil {
		// TODO: re-factor error displaying
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	if err := models.NoteSaveOne(&validator.noteModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	if err := validator.CreateFiles(c); err != nil {
		// TODO: re-factor error displaying
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	if err := validator.noteModel.AddAttachments(validator.attachmentsFiles); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := serializers.NoteSerializer{c, validator.noteModel}
	c.JSON(http.StatusCreated, gin.H{"note": serializer.Response()})
}

func NoteUpdate(c *gin.Context) {
	var err error
	currentNote, isExit := getOwnerNote(c)
	if isExit {
		return
	}

	validator := NewNoteCreationValidatorFillWith(currentNote)
	if err = validator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	if currentNote, err = currentNote.Update(validator.noteModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	serializer := serializers.NoteSerializer{c, currentNote}
	c.JSON(http.StatusOK, gin.H{"note": serializer.Response()})
}

func NoteDelete(c *gin.Context) {
	currentNote, isExit := getOwnerNote(c)
	if isExit {
		return
	}

	if err := models.NoteDelete(&models.NoteModel{ID: currentNote.ID}); err != nil {
		c.JSON(http.StatusNotFound, common.NewError("note", errors.New("Invalid id")))
		return
	}
	c.JSON(http.StatusOK, gin.H{"note": "Deleted success"})
}

func TestImageUploading(c *gin.Context) {
	form, err := c.MultipartForm()
	fmt.Println(form, err)
	files, ok := form.File["image"]
	fmt.Println(files, ok)

	for _, file := range files {
		fmt.Println(reflect.TypeOf(file))
		err := c.SaveUploadedFile(file, "saved/"+file.Filename)
		if err != nil {
			fmt.Println(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
