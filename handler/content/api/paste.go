package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/go-sonic/sonic/model/param"
	"github.com/go-sonic/sonic/service"
	"github.com/go-sonic/sonic/util/xerr"
)

type PasteHandler struct {
	PasteService service.PasteService
}

func NewPasteHandler(pasteService service.PasteService) *PasteHandler {
	return &PasteHandler{
		PasteService: pasteService,
	}
}

func (p *PasteHandler) Get(ctx *gin.Context) {
	key := ctx.Param("key")
	username := ctx.Query("username")
	password := ctx.Query("password")
	content, err := p.PasteService.Get(ctx, key, username, password)
	if err != nil {
		errorType := xerr.GetType(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.String(http.StatusNotFound, err.Error())
		} else if errorType == xerr.Forbidden {
			ctx.String(http.StatusForbidden, err.Error())
		}
	}
	ctx.String(http.StatusOK, content)
}

func (p *PasteHandler) Create(ctx *gin.Context) (any, error) {
	para := param.Paste{}
	err := ctx.ShouldBindJSON(&para)
	if err != nil {
		return nil, xerr.WithStatus(err, xerr.StatusBadRequest).WithMsg("parameter error")
	}
	para.ClientIP = ctx.ClientIP()
	key, err := p.PasteService.Create(ctx, para)
	if err != nil {
		return nil, xerr.WithStatus(err, xerr.StatusInternalServerError).WithMsg("create error")
	}
	return key, nil
}
