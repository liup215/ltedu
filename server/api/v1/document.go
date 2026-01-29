package v1

import (
	"edu/conf"
	"edu/lib/logger"
	"edu/lib/net/http"
	"edu/model"
	"edu/service"
	"fmt"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

var DocumentCtrl = &DocumentController{
	documentSvr: service.DocumentSvr,
}

type DocumentController struct {
	documentSvr *service.DocumentService
}

func (ctrl *DocumentController) CreateDocument(c *gin.Context) {
	dc := model.DocumentCreateEditRequest{}
	if err := c.BindJSON(&dc); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	if err := ctrl.documentSvr.CreateDocument(dc); err != nil {
		logger.Logger.Error("创建失败!", zap.Error(err))
		http.ErrorData(c, "创建失败", nil)
		return
	}

	http.SuccessData(c, "创建成功!", nil)
}

func (ctrl *DocumentController) EditDocument(c *gin.Context) {
	dc := model.DocumentCreateEditRequest{}
	if err := c.BindJSON(&dc); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	if err := ctrl.documentSvr.EditDocument(dc); err != nil {
		http.ErrorData(c, "编辑失败"+err.Error(), nil)
		return
	}

	http.SuccessData(c, "编辑成功!", nil)
}

func (ctrl *DocumentController) DeleteDocument(c *gin.Context) {
	q := model.DocumentQueryRequest{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	if err := ctrl.documentSvr.DeleteDocument(q.ID); err != nil {
		http.ErrorData(c, "删除失败", nil)
		return
	}

	http.SuccessData(c, "删除成功!", nil)
}

func (ctrl *DocumentController) SelectDocumentById(c *gin.Context) {
	q := model.DocumentQueryRequest{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	o, err := ctrl.documentSvr.SelectDocumentById(q.ID)
	if err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", o)
}

func (ctrl *DocumentController) SelectDocumentList(c *gin.Context) {
	q := model.DocumentQueryRequest{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	list, total, err := ctrl.documentSvr.SelectDocumentList(q)
	if err != nil {
		http.ErrorData(c, "参数解析失败,"+err.Error(), nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": total,
	})
}

func (ctrl *DocumentController) SelectDocumentAll(c *gin.Context) {
	oq := model.DocumentQueryRequest{}
	if err := c.BindJSON(&oq); err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}
	list, err := ctrl.documentSvr.SelectDocumentAll(oq)
	if err != nil {
		http.ErrorData(c, "数据获取失败!", nil)
		return
	}

	http.SuccessData(c, "数据获取成功!", gin.H{
		"list":  list,
		"total": len(list),
	})
}

func (ctrl *DocumentController) DownloadDocument(c *gin.Context) {

	q := model.DocumentQueryRequest{}
	if err := c.BindJSON(&q); err != nil {
		http.ErrorData(c, "参数解析失败", nil)
		return
	}

	document, err := ctrl.documentSvr.SelectDocumentById(q.ID)
	if err != nil {
		http.ErrorData(c, "文档查询失败！", nil)
		return
	}

	if document.AttachmentId == 0 {
		http.ErrorData(c, "文档附件不存在！", nil)
		return
	}

	attachment, err := service.AttachmentSvr.SelectAttachmentById(document.AttachmentId)
	if err != nil {
		http.ErrorData(c, "文档附件查询失败！", nil)
		return
	}

	if attachment.Hash == "" {
		http.ErrorData(c, "文档附件不存在！", nil)
		return
	}

	link, err := ctrl.generateDownloadUrl(document, attachment.Hash)

	if err != nil {
		http.ErrorData(c, "生成下载地址失败: "+err.Error(), nil)
		return
	}

	http.SuccessData(c, "获取下载地址成功！", gin.H{
		"link": link,
	})
}

func (ctrl *DocumentController) generateDownloadUrl(doc *model.Document, hash string) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(5))),
		ID:        hash,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(conf.Conf.SecretKey))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("/download/%s?filename=%s", tokenString, url.QueryEscape(doc.Name)), nil
}
