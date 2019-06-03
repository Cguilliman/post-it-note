package controllers

import (
    "github.com/gin-gonic/gin"

    "github.com/Cguilliman/post-it-note/common"
    "github.com/Cguilliman/post-it-note/models"
)

type NewNoteValidator struct {
    Note struct {
        Note string `json:"note"`
    } `json:"note"`
    noteModel models.NoteModel `json:"-"`
}
