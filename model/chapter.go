package model

type Chapter struct {
	Model
	Name       string      `json:"name"`
	SyllabusId uint        `json:"syllabusId"`
	Syllabus   Syllabus    `json:"syllabus"`
	ParentId   uint        `json:"parentId" gorm:"index"`
	Children   []*Chapter  `gorm:"-" json:"children,omitempty"`
	IsLeaf     int         `gorm:"-" json:"isLeaf"`
	Questions  []*Question `gorm:"many2many:question_chapters;"`
}

func BuildChapterString(chapters []*Chapter, currentID uint) string {
	var chapterString string
	for _, chapter := range chapters {
		if chapter.ID == currentID {
			chapterString = chapter.Name
			if chapter.ParentId != 0 {
				chapterString = BuildChapterString(chapters, chapter.ParentId) + "/" + chapterString
			}
			break
		}
	}
	return chapterString
}

type ChapterQuery struct {
	Model
	Name       string `json:"name"`
	SyllabusId uint   `json:"syllabusId"`
	ParentId   uint   `json:"parentId"`
	Page
}
