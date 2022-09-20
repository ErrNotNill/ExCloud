package models

import (
	"net/textproto"
	"time"
)

type DocumentParams struct {
	ID     int
	File   bool
	Public bool
	Token  string
	Mime   textproto.MIMEHeader
	Grant  []User
	Name   string
	Date   time.Time
}

type Document interface {
	PostNewDocument(name string) //POST .id
	GetDocumentsList()           //GET .all
	GetDocumentById()            //GET .id
	DeleteDocument()             //DELETE .id
}
