package repository

import (
	"edu/lib/cache"
	"edu/model"
	"fmt"
	"time"
)

const (
	chapterCacheCapacity = 512
	chapterCacheTTL      = 10 * time.Minute
)

// chapterQueryKey builds a deterministic cache key for FindAll queries.
func chapterQueryKey(query *model.ChapterQuery) string {
	return fmt.Sprintf("all:%d:%d", query.SyllabusId, query.ParentId)
}

// cachedChapterRepository decorates an IChapterRepository with an in-process
// LRU cache for read-heavy methods (FindByID, FindBySyllabusID, FindAll).
// Write operations (Create, Update, Delete) invalidate relevant cache entries.
type cachedChapterRepository struct {
	inner       IChapterRepository
	byID        *cache.LRU[uint, *model.Chapter]
	bySyllabus  *cache.LRU[uint, []*model.Chapter]
	byQuery     *cache.LRU[string, []*model.Chapter]
}

// NewCachedChapterRepository wraps inner with a multi-level LRU cache.
func NewCachedChapterRepository(inner IChapterRepository) IChapterRepository {
	return &cachedChapterRepository{
		inner:      inner,
		byID:       cache.New[uint, *model.Chapter](chapterCacheCapacity, chapterCacheTTL),
		bySyllabus: cache.New[uint, []*model.Chapter](128, chapterCacheTTL),
		byQuery:    cache.New[string, []*model.Chapter](128, chapterCacheTTL),
	}
}

func (r *cachedChapterRepository) FindByID(id uint) (*model.Chapter, error) {
	if v, ok := r.byID.Get(id); ok {
		return v, nil
	}
	ch, err := r.inner.FindByID(id)
	if err != nil {
		return nil, err
	}
	if ch != nil {
		r.byID.Set(id, ch)
	}
	return ch, nil
}

func (r *cachedChapterRepository) FindBySyllabusID(syllabusId uint) ([]*model.Chapter, error) {
	if v, ok := r.bySyllabus.Get(syllabusId); ok {
		return v, nil
	}
	result, err := r.inner.FindBySyllabusID(syllabusId)
	if err != nil {
		return nil, err
	}
	r.bySyllabus.Set(syllabusId, result)
	return result, nil
}

func (r *cachedChapterRepository) FindAll(query *model.ChapterQuery) ([]*model.Chapter, error) {
	key := chapterQueryKey(query)
	if v, ok := r.byQuery.Get(key); ok {
		return v, nil
	}
	result, err := r.inner.FindAll(query)
	if err != nil {
		return nil, err
	}
	r.byQuery.Set(key, result)
	return result, nil
}

// FindByParentID and FindPage are not cached as they are less frequently accessed
// in hot paths, and parent/page queries are more variable.
func (r *cachedChapterRepository) FindByParentID(parentId uint) ([]*model.Chapter, error) {
	return r.inner.FindByParentID(parentId)
}

func (r *cachedChapterRepository) FindPage(query *model.ChapterQuery, offset, limit int) ([]*model.Chapter, int64, error) {
	return r.inner.FindPage(query, offset, limit)
}

func (r *cachedChapterRepository) HasChildren(id uint) (bool, error) {
	return r.inner.HasChildren(id)
}

func (r *cachedChapterRepository) CountKnowledgePoints(id uint) (int64, error) {
	return r.inner.CountKnowledgePoints(id)
}

// invalidateBySyllabusID removes chapter entries associated with a specific syllabus.
func (r *cachedChapterRepository) invalidateBySyllabusID(syllabusId uint) {
	r.bySyllabus.Delete(syllabusId)
	r.byQuery.Purge() // query cache may include syllabus-filtered results
}

func (r *cachedChapterRepository) Create(c *model.Chapter) error {
	err := r.inner.Create(c)
	if err == nil {
		r.invalidateBySyllabusID(c.SyllabusId)
	}
	return err
}

func (r *cachedChapterRepository) Update(c *model.Chapter) error {
	err := r.inner.Update(c)
	if err == nil {
		r.byID.Delete(c.ID)
		r.invalidateBySyllabusID(c.SyllabusId)
	}
	return err
}

func (r *cachedChapterRepository) Delete(id uint) error {
	// Fetch the chapter first to know its syllabusId for cache invalidation.
	ch, _ := r.inner.FindByID(id)
	err := r.inner.Delete(id)
	if err == nil {
		r.byID.Delete(id)
		if ch != nil {
			r.invalidateBySyllabusID(ch.SyllabusId)
		} else {
			// Chapter lookup failed; we can't determine the syllabusId.
			// Purge only the query cache (bySyllabus entries will expire via TTL).
			r.byQuery.Purge()
		}
	}
	return err
}
