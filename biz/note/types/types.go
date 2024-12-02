package types

import (
	"github.com/dgdts/UniversalServer/biz/biz_context"
	"github.com/dgdts/UniversalServer/biz/model/note"
	"github.com/dgdts/UniversalServer/biz/note/model"
)

const (
	NoteTypeMarkdown = "markdown"
)

type NoteHandler interface {
	CreateNote(ctx *biz_context.BizContext, req *model.Node) (*note.CreateNoteResponse, error)
	GetNote(ctx *biz_context.BizContext, req *note.GetNoteRequest) (*model.Node, error)
	UpdateNote(ctx *biz_context.BizContext, req *model.UpdateNode) (*note.UpdateNoteResponse, error)
	DeleteNote(ctx *biz_context.BizContext, req *note.DeleteNoteRequest) (*note.DeleteNoteResponse, error)
}