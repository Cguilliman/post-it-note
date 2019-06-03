package controllers

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    
    "github.com/Cguilliman/post-it-note/models"
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
    c.JSON(http.StatusOK, gin.H{
        "nodes": serializer.Response(),
    })
}
