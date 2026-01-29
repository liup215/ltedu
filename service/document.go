package service

import (
"edu/lib/converter"
"edu/lib/logger"
"edu/lib/utils"
"edu/model"
"errors"
"fmt"
"os"
"path/filepath"
"strings"
"time"

"go.uber.org/zap"
"gorm.io/gorm"
)

var DocumentSvr = &DocumentService{baseService: newBaseService()}

type DocumentService struct {
	baseService
}

func (s *DocumentService) CreateDocument(dc model.DocumentCreateEditRequest) error {
	if dc.ID != 0 {
		dc.ID = uint(0)
	}

	if dc.SyllabusId == 0 {
		return errors.New("syllabusId不能为空")
	}

	if dc.DocumentCategoryId == 0 {
		return errors.New("documentCategoryId不能为空")
	}

	if dc.AttachmentId == 0 {
		logger.Logger.Error("CreateDocument", zap.Error(errors.New("attachmentId不能为空")))
		return errors.New("attachmentId不能为空")
	}

return nil // TODO: 替换为实际文档创建逻辑
}

func (s *DocumentService) EditDocument(dc model.DocumentCreateEditRequest) error {
	if dc.ID == 0 {
		return errors.New("id不能为空")
	}

	if dc.SyllabusId == 0 {
		return errors.New("syllabusId不能为空")
	}

	if dc.DocumentCategoryId == 0 {
		return errors.New("documentCategoryId不能为空")
	}

return nil // TODO: 替换为实际文档编辑逻辑
}

func (s *DocumentService) DeleteDocument(id uint) error {
	if id == 0 {
		return errors.New("id不能为空")
	}

return nil // TODO: 替换为实际文档删除逻辑
}

func (s *DocumentService) SelectDocumentById(id uint) (*model.Document, error) {
	if id == 0 {
		return nil, errors.New("id不能为空")
	}

return nil, nil // TODO: 替换为实际查询逻辑
}

func (s *DocumentService) SelectDocumentList(q model.DocumentQueryRequest) ([]*model.DocumentQueryResponse, int64, error) {
return nil, 0, nil // TODO: 替换为实际分页查询逻辑
}

func (s *DocumentService) SelectDocumentAll(q model.DocumentQueryRequest) ([]*model.DocumentQueryResponse, error) {
return nil, nil // TODO: 替换为实际查询逻辑
}

func (s *DocumentService) parseQuery(q model.DocumentQueryRequest) any {
return nil // TODO: 替换为实际查询逻辑
}

// Set Document status
func (s *DocumentService) setDocumentStatus(id uint, status int) error {
	if id == 0 {
		return errors.New("id不能为空")
	}

return nil // TODO: 替换为实际状态更新逻辑
}

var convertDocumentRunning = false

// Set Document status
func (s *DocumentService) LoopConvertDocument() {
	if convertDocumentRunning {
		return
	}
	// 清空缓存目录
	os.RemoveAll("cache/convert")
	convertDocumentRunning = true
	sleep := 10 * time.Second

/* TODO: 替换为实际文档状态批量更新逻辑 */
	for {
		now := time.Now()
		logger.Logger.Debug("loopCovertDocument，start...")
		err := s.ConvertDocument()
		if err != nil && err != gorm.ErrRecordNotFound {
			return
		}
		logger.Logger.Debug("loopCovertDocument，end...", zap.String("cost", time.Since(now).String()))
		if err == gorm.ErrRecordNotFound {
			time.Sleep(sleep)
		}
	}
}

func (s *DocumentService) ConvertDocument() (err error) {
document := &model.Document{} // TODO: 替换为实际文档查询逻辑

// convert the file to pdf
s.setDocumentStatus(document.ID, model.DocumentStatusConverting)

attachment := &model.Attachment{} // TODO: 替换为实际附件查询逻辑

	localFile := strings.TrimLeft(attachment.Path, "./")
	baseDir := strings.TrimSuffix(localFile, filepath.Ext(localFile))
	cover := baseDir + "/cover.png"

	timeout := 30 * time.Minute

	cvt := converter.NewConverter(logger.Logger, timeout)

	defer cvt.Clean()
	dstPDF, err := cvt.ConvertToPDF(localFile)
	if err != nil {
		s.setDocumentStatus(document.ID, model.DocumentStatusFailed)
		logger.Logger.Error("ConvertDocument", zap.Error(err))
		return
	}

	document.Pages, _ = cvt.CountPDFPages(dstPDF)
	maxPreview := 20
	document.Preview = maxPreview

	if document.Preview > document.Pages {
		document.Preview = document.Pages
	}

	// PDF截取第一章图片作为封面(封面不是最重要的，期间出现错误，不影响文档转换)
	pages, err := cvt.ConvertPDFToPNG(dstPDF, 1, 1)
	if err != nil {
		logger.Logger.Error("get pdf cover", zap.Error(err))
	}

	if len(pages) > 0 {
		coverBig := baseDir + "/cover.big.png"
		utils.CopyFile(pages[0].PagePath, coverBig)
		utils.CopyFile(pages[0].PagePath, cover)
		utils.CropImage(cover, model.DocumentCoverWidth, model.DocumentCoverHeight)
		document.Width, document.Height, _ = utils.GetImageSize(coverBig) // 页面宽高
	}

	// PDF转为SVG
	toPage := document.Pages
	if maxPreview > 0 {
		toPage = maxPreview
	}
	if toPage > document.Pages {
		toPage = document.Pages
	}

	pages, err = cvt.ConvertPDFToPages(dstPDF, 1, toPage, &converter.OptionConvertPages{
		EnableSVGO: true,
		EnableGZIP: true,
		Extension:  "svg",
	})
	if err != nil {
		s.setDocumentStatus(document.ID, model.DocumentStatusFailed)
		logger.Logger.Error("ConvertDocument", zap.Error(err))
		return
	}

	ext := ".svg"
	if ext == ".svg" && true {
		ext = ".gzip.svg"
	}

	for _, page := range pages {
		dst := fmt.Sprintf(baseDir+"/%d%s", page.PageNum, ext)
		logger.Logger.Debug("ConvertDocument CopyFile", zap.String("src", page.PagePath), zap.String("dst", dst))
		errCopy := utils.CopyFile(page.PagePath, dst)
		if errCopy != nil {
			logger.Logger.Error("ConvertDocument CopyFile", zap.Error(errCopy))
		}
	}

	// 提取PDF文本以及获取文档信息
	textFile, _ := cvt.ConvertPDFToTxt(dstPDF)
	utils.CopyFile(textFile, baseDir+"/content.txt")

	// 读取文本内容，以提取关键字和摘要
	if content, errRead := os.ReadFile(textFile); errRead == nil {
		contentStr := string(content)
		replacer := strings.NewReplacer("\r", " ", "\n", " ", "\t", " ")
		document.Description = strings.TrimSpace(replacer.Replace(utils.Substr(contentStr, 255)))
	}

	document.Status = model.DocumentStatusConverted
	document.EnableGZIP = true
	document.PreviewExt = strings.TrimPrefix(ext, ".gzip")
/* TODO: 替换为实际文档属性更新逻辑 */

return
}
