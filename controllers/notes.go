package controllers

import (
    "fmt"
    "strconv"
    "net/http"
    "github.com/gin-gonic/gin"
    
    "github.com/Cguilliman/post-it-note/models"
    "github.com/Cguilliman/post-it-note/common"
    "github.com/Cguilliman/post-it-note/serializers"
)

func UserNotesList(c *gin.Context) {
    myUserModel := c.MustGet("my_user_model").(models.UserModel)
    notes, err := models.GetNotes(&models.NoteModel{
        OwnerID: myUserModel.ID,
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
        c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
        return
    }
    if err := models.NodeSaveOne(&validator.noteModel); err != nil {
        c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
        return
    }
    serializer := serializers.NoteSerializer{c, validator.noteModel}
    c.JSON(http.StatusCreated, gin.H{"note": serializer.Response()})
}

func NoteUpdate(c *gin.Context) {
    var (
        err         error
        notePK      uint64
        currentNote models.NoteModel
    )
    if notePK, err = strconv.ParseUint(c.Param("pk"), 10, 32); err != nil {
        fmt.Println("1111111")
        c.JSON(http.StatusNotFound, nil)
        return
    }
    if currentNote, err = models.GetNote(&models.NoteModel{ID: uint(notePK)}); err != nil {
        fmt.Println("2222222")
        c.JSON(http.StatusUnprocessableEntity, common.NewError("pk", err))
        return
    }

    validator := NewNoteCreationValidatorFillWith(currentNote)
    if err = validator.Bind(c); err != nil {
        fmt.Println("33333333")
        c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
        return
    }
    if currentNote, err = currentNote.Update(validator.noteModel); err != nil {
        fmt.Println("44444444")
        c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
        return
    }

    serializer := serializers.NoteSerializer{c, currentNote}
    c.JSON(http.StatusOK, gin.H{"note": serializer.Response()})
}
