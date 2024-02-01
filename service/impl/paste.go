package impl

import (
	"context"

	"github.com/go-sonic/sonic/dal"
	"github.com/go-sonic/sonic/model/entity"
	"github.com/go-sonic/sonic/model/param"
	"github.com/go-sonic/sonic/service"
	"github.com/go-sonic/sonic/util"
	"github.com/go-sonic/sonic/util/xerr"
)

type pasteServiceImpl struct {
}

func NewPasteService() service.PasteService {
	return &pasteServiceImpl{}
}

func (p *pasteServiceImpl) Get(ctx context.Context, key, username, password string) (string, error) {
	permanentPaste := dal.GetQueryByCtx(ctx).Permanent
	paste, err := permanentPaste.WithContext(ctx).
		Where(permanentPaste.Key.Eq(key)).Take()
	if paste != nil && paste.Password != nil {
		if *paste.Password != util.Md5Hex(password) {
			return "", xerr.Forbidden.New("failed to validate password")
		}
	}
	if err != nil {
		return "", err
	}
	return paste.Content, nil
}

func (p *pasteServiceImpl) Create(ctx context.Context, paste param.Paste) (any, error) {
	key := paste.Key
	if key == "" {
		key = util.Generator(8, false, &paste)
	}
	permanent := &entity.Permanent{
		Key:      key,
		Lang:     paste.Lang,
		Content:  paste.Content,
		Username: &paste.Username,
		ClientIP: paste.ClientIP,
	}
	if paste.Password != "" {
		password := util.Md5Hex(paste.Password)
		permanent.Password = &password
	}
	permanentPaste := dal.GetQueryByCtx(ctx).Permanent
	if err := permanentPaste.WithContext(ctx).Create(permanent); err != nil {
		return "", err
	}
	return permanent, nil
}

func (p *pasteServiceImpl) Update(ctx context.Context, paste param.Paste) error {
	key := paste.Key
	password := ""
	if paste.Password != "" {
		password = util.Md5Hex(paste.Password)
	}
	permanentPaste := dal.GetQueryByCtx(ctx).Permanent
	updateResult, err := permanentPaste.WithContext(ctx).
		Where(permanentPaste.Key.Eq(key), permanentPaste.Password.Eq(password)).
		UpdateSimple(permanentPaste.Content.Value(paste.Content))

	if err != nil {
		return WrapDBErr(err)
	}
	if updateResult.RowsAffected != 1 {
		return xerr.NoType.New("update permanent failed").WithMsg("update permanent failed").WithStatus(xerr.StatusInternalServerError)
	}
	return nil
}
