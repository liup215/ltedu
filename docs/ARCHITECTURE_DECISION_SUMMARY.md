# Architecture Decision Summary

## Knowledge Point Structure Design Decision

**Date**: 2026-02-18  
**Status**: ✅ DECIDED  
**Decision**: Maintain current chapter-bound architecture (Option 1) with clear evolution path to hybrid approach (Option 3)

---

## Quick Summary

### Question
Should knowledge points be:
1. Bound to chapters (current)
2. Independent global entities
3. Hybrid approach (chapter-bound + optional global reference)

### Decision
**Keep current architecture (Option 1)**, with documented evolution path to **Option 3** when needs arise.

### Reasoning
✅ Current scale doesn't justify complex architecture  
✅ Educational reality: different exam boards define concepts differently  
✅ Syllabus independence > data deduplication  
✅ Simple architecture = stable system  
✅ Clear evolution path when needed  

---

## When to Revisit

Trigger conditions for evolution to Option 3:
- [ ] Question bank exceeds 10,000 questions
- [ ] Clear cross-syllabus learning demand from users
- [ ] Need for knowledge graph features
- [ ] Advanced analytics requirements
- [ ] Business expansion to multiple exam boards at scale

---

## Implementation Path (When Needed)

### Phase 1: Foundation
```sql
-- Add optional field to existing table
ALTER TABLE knowledge_points 
ADD COLUMN global_knowledge_point_id INT NULL;

-- Create global knowledge point library
CREATE TABLE global_knowledge_points (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    category VARCHAR(100),
    tags JSON
);
```

### Phase 2: Matching Service
- Background service to identify similar knowledge points
- Human review and standardization
- Gradual association building

### Phase 3: Advanced Features
- Cross-syllabus question recommendations
- Knowledge graph
- Learning path planning

---

## Key Principles

1. **Start Simple**: Don't build for imaginary future needs
2. **Syllabus Independence**: Respect educational differences
3. **Incremental Evolution**: Change based on data, not speculation
4. **Backward Compatibility**: Never break existing functionality

---

## References

- Full Analysis (Chinese): `KNOWLEDGE_POINT_ARCHITECTURE_ANALYSIS.md`
- Full Analysis (English): `KNOWLEDGE_POINT_ARCHITECTURE_ANALYSIS_EN.md`

---

## Sign-off

**Recommended by**: Architecture Team  
**Status**: Ready for stakeholder review  
**Next Action**: Monitor usage patterns, revisit when triggers met
