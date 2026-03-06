package repository

import (
	"edu/lib/cache"
	"edu/model"
	"fmt"
	"time"
)

const (
	syllabusCacheCapacity = 256
	syllabusCacheTTL      = 10 * time.Minute
)

// syllabusAllCacheKey is the key used to cache the full list returned by FindAll.
func syllabusAllCacheKey(query *model.SyllabusQuery) string {
	return fmt.Sprintf("all:%d:%s:%d", query.ID, query.Code, query.QualificationId)
}

// cachedSyllabusRepository decorates an ISyllabusRepository with an in-process
// LRU cache for read-heavy methods (FindByID, FindAll).
// Write operations (Create, Update, Delete) invalidate the relevant cache entries.
type cachedSyllabusRepository struct {
	inner    ISyllabusRepository
	byID     *cache.LRU[uint, *model.Syllabus]
	byQuery  *cache.LRU[string, []*model.Syllabus]
}

// NewCachedSyllabusRepository wraps inner with a two-level LRU cache.
func NewCachedSyllabusRepository(inner ISyllabusRepository) ISyllabusRepository {
	return &cachedSyllabusRepository{
		inner:   inner,
		byID:    cache.New[uint, *model.Syllabus](syllabusCacheCapacity, syllabusCacheTTL),
		byQuery: cache.New[string, []*model.Syllabus](64, syllabusCacheTTL),
	}
}

func (r *cachedSyllabusRepository) FindByID(id uint) (*model.Syllabus, error) {
	if v, ok := r.byID.Get(id); ok {
		return v, nil
	}
	syl, err := r.inner.FindByID(id)
	if err != nil {
		return nil, err
	}
	if syl != nil {
		r.byID.Set(id, syl)
	}
	return syl, nil
}

func (r *cachedSyllabusRepository) FindAll(query *model.SyllabusQuery) ([]*model.Syllabus, error) {
	key := syllabusAllCacheKey(query)
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

// FindPage is not cached because paginated queries are highly variable.
func (r *cachedSyllabusRepository) FindPage(query *model.SyllabusQuery, offset, limit int) ([]*model.Syllabus, int64, error) {
	return r.inner.FindPage(query, offset, limit)
}

func (r *cachedSyllabusRepository) Create(s *model.Syllabus) error {
	err := r.inner.Create(s)
	if err == nil {
		r.byQuery.Purge()
	}
	return err
}

func (r *cachedSyllabusRepository) Update(s *model.Syllabus) error {
	err := r.inner.Update(s)
	if err == nil {
		r.byID.Delete(s.ID)
		r.byQuery.Purge()
	}
	return err
}

func (r *cachedSyllabusRepository) Delete(id uint) error {
	err := r.inner.Delete(id)
	if err == nil {
		r.byID.Delete(id)
		r.byQuery.Purge()
	}
	return err
}
