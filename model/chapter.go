package model

const (
	ChapterLevelAS = "AS"
	ChapterLevelA2 = "A2"
)

type Chapter struct {
	Model
	Name       string      `json:"name"`
	SyllabusId uint        `json:"syllabusId"`
	Syllabus   Syllabus    `json:"syllabus"`
	ParentId   uint        `json:"parentId" gorm:"index"`
	Level      string      `json:"level"` // syllabus level: "AS", "A2", or "" for non-A-Level
	Children   []*Chapter  `gorm:"-" json:"children,omitempty"`
	IsLeaf     int         `gorm:"-" json:"isLeaf"`
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
	FilterRoot bool   `json:"filterRoot"` // If true, filter by parentId = 0 (overrides ParentId=0 check)
	Page
}
