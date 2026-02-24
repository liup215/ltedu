# Architecture Diagrams

## Current Architecture (Option 1)

```
┌─────────────────────────────────────────────────────────┐
│                    Organisation                         │
│                  (Exam Board)                           │
│              e.g., Cambridge, Edexcel                   │
└──────────────────┬──────────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────────┐
│                  Qualification                          │
│              (Qualification Type)                       │
│              e.g., IGCSE, A-Level                      │
└──────────────────┬──────────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────────┐
│                    Syllabus                             │
│                   (Curriculum)                          │
│          e.g., Mathematics 9709                         │
└──────────────────┬──────────────────────────────────────┘
                   │
                   ▼
┌─────────────────────────────────────────────────────────┐
│                    Chapter                              │
│                (Tree Structure)                         │
│         ┌──────────┬──────────┬──────────┐             │
│         │ Parent   │ Parent   │ Parent   │             │
│         └─┬──┬──┬─┘└─┬──┬──┬─┘└─┬──┬──┬─┘             │
│           │  │  │    │  │  │    │  │  │               │
│          Leaf Leaf Leaf Leaf Leaf Leaf Leaf Leaf       │
│           ▼  ▼  ▼    ▼  ▼  ▼    ▼  ▼  ▼               │
└───────────┼──┼──┼────┼──┼──┼────┼──┼──┼───────────────┘
            │  │  │    │  │  │    │  │  │
            ▼  ▼  ▼    ▼  ▼  ▼    ▼  ▼  ▼
┌──────────────────────────────────────────────────────┐
│              KnowledgePoint                          │
│         (Only on Leaf Chapters)                      │
│   ┌────────┬────────┬────────┬────────┐             │
│   │   KP   │   KP   │   KP   │   KP   │             │
│   └────┬───┴───┬────┴───┬────┴────┬───┘             │
└────────┼───────┼────────┼─────────┼─────────────────┘
         │       │        │         │
         └───────┴────────┴─────────┘
                  │
                  ▼
┌──────────────────────────────────────────────────────┐
│                  Question                            │
│            (many2many relationship)                  │
│      One question can have multiple KPs              │
│      One KP can have multiple questions              │
└──────────────────────────────────────────────────────┘

Current Status: ✅ APPROVED - Working Well
```

## Future Evolution (Option 3 - Hybrid)

```
┌─────────────────────────────────────────────────────────┐
│              Current Architecture                       │
│                  (Unchanged)                            │
│                                                         │
│    Organisation → Qualification → Syllabus             │
│         → Chapter → KnowledgePoint → Question          │
│                                                         │
└────────────────────┬────────────────────────────────────┘
                     │
                     │ Optional Association
                     ▼
┌─────────────────────────────────────────────────────────┐
│        NEW: Global Knowledge Point Library              │
│             (Independent Module)                        │
│                                                         │
│   ┌──────────────────────────────────────────┐         │
│   │   GlobalKnowledgePoint                   │         │
│   │                                           │         │
│   │  - Universal concepts                    │         │
│   │  - Cross-syllabus                        │         │
│   │  - Curated & standardized                │         │
│   │  - Category tags                         │         │
│   └──────────────────────────────────────────┘         │
│                                                         │
│   Examples:                                            │
│   • Linear Equations (Global)                          │
│   • Quadratic Functions (Global)                       │
│   • Pythagoras Theorem (Global)                        │
│                                                         │
└────────────────────┬────────────────────────────────────┘
                     │
                     │ Enables
                     ▼
┌─────────────────────────────────────────────────────────┐
│           Advanced Features                             │
│                                                         │
│  ✓ Cross-syllabus question recommendations            │
│  ✓ Knowledge graph & dependencies                      │
│  ✓ Learning path planning                              │
│  ✓ Unified analytics                                   │
│                                                         │
└─────────────────────────────────────────────────────────┘

Future Status: 📋 DOCUMENTED - Ready When Needed
```

## Comparison: Cambridge vs Edexcel Example

### Current Approach (Option 1) ✅

```
Cambridge IGCSE Mathematics
└── Algebra
    └── Quadratic Functions (Leaf Chapter)
        ├── KP: Factorization Methods (Cambridge-specific)
        ├── KP: Completing the Square (Cambridge focus)
        ├── KP: Quadratic Formula (Cambridge approach)
        └── Questions...

Edexcel IGCSE Mathematics
└── Algebra
    └── Quadratic Functions (Leaf Chapter)
        ├── KP: Standard Form (Edexcel-specific)
        ├── KP: Solving Techniques (Edexcel focus)
        ├── KP: Graphing Parabolas (Edexcel approach)
        └── Questions...

Result: Independent, precise definitions ✅
No conflicts, clear for students ✅
```

### Future Approach (Option 3) 🚀

```
Cambridge IGCSE Mathematics
└── Algebra
    └── Quadratic Functions
        ├── KP: Factorization ──┐
        ├── KP: Completing Square│
        └── KP: Quadratic Formula│
                                 │
                                 ▼
                    ┌─────────────────────────┐
                    │  Global KP Library      │
                    │  • Quadratic Functions  │
                    │  • Linear Equations     │
                    │  • Polynomials          │
                    └─────────────────────────┘
                                 ▲
                                 │
Edexcel IGCSE Mathematics       │
└── Algebra                     │
    └── Quadratic Functions     │
        ├── KP: Standard Form ──┘
        ├── KP: Solving Techniques
        └── KP: Graphing Parabolas

Result: Maintains independence ✅
        Enables cross-syllabus features ✅
        Best of both worlds ✅
```

## Decision Tree

```
                    START
                      │
                      ▼
         ┌─────────────────────────┐
         │  Current System Status  │
         └────────┬────────────────┘
                  │
                  ▼
         ┌─────────────────────────┐
         │  Is it working well?    │
         └────┬──────────────┬─────┘
              │YES           │NO
              ▼              ▼
    ┌──────────────┐   ┌──────────────┐
    │  Keep Option 1│   │ Fix issues   │
    │  (APPROVED)   │   │   first      │
    └────┬──────────┘   └──────────────┘
         │
         ▼
    ┌──────────────────────────┐
    │ Monitor Trigger Events   │
    │                          │
    │ • QB > 10K questions?    │
    │ • Cross-syllabus demand? │
    │ • Knowledge graph need?  │
    └────┬──────────┬──────────┘
         │NO        │YES
         ▼          ▼
    ┌─────────┐  ┌──────────────┐
    │Continue │  │Evolve to     │
    │Option 1 │  │Option 3      │
    └─────────┘  └──────────────┘

Current Position: ✅ Keep Option 1
```

## Implementation Phases (When Needed)

```
Phase 1: Foundation
┌─────────────────────────────────────┐
│ 1. Add GlobalKnowledgePoint table  │
│ 2. Add optional FK to KP table     │
│ 3. No functional changes           │
│                                     │
│ Duration: 1-2 weeks                │
│ Risk: Very Low                     │
└─────────────────────────────────────┘
              │
              ▼
Phase 2: Library Building
┌─────────────────────────────────────┐
│ 1. Background matching service     │
│ 2. Human review & standardization  │
│ 3. Gradual association building    │
│                                     │
│ Duration: 3-6 months               │
│ Risk: Low (background process)     │
└─────────────────────────────────────┘
              │
              ▼
Phase 3: Advanced Features
┌─────────────────────────────────────┐
│ 1. Cross-syllabus recommendations  │
│ 2. Knowledge graph                  │
│ 3. Learning path planning          │
│                                     │
│ Duration: 6-12 months              │
│ Risk: Medium (new features)         │
└─────────────────────────────────────┘
```

## Risk vs Benefit Matrix

```
                    High Benefit
                         ▲
                         │
         Option 3        │        Option 2
         (Hybrid)        │      (Independent)
                         │
         🟢              │           🟡
    Recommended          │      Complex but
    Future Path          │      powerful
                         │
─────────────────────────┼─────────────────────► High Risk
                         │
                         │
         🟢              │           🔴
    Option 1             │      Bad Choice
    (Current)            │    (Not considered)
    Perfect for Now      │
                         │
                    Low Benefit

Legend:
🟢 = Recommended
🟡 = Consider carefully
🔴 = Avoid
```

## Summary Visual

```
╔════════════════════════════════════════════════════════╗
║                  ARCHITECTURE DECISION                 ║
╠════════════════════════════════════════════════════════╣
║                                                        ║
║  Current Status:    ✅ Option 1 (APPROVED)            ║
║  Future Path:       📋 Option 3 (DOCUMENTED)          ║
║  Action Required:   ❌ None                            ║
║  Risk Level:        🟢 Low                             ║
║  Confidence:        🎯 High                            ║
║                                                        ║
║  When to Revisit:                                     ║
║    • Question Bank > 10,000                           ║
║    • Cross-syllabus demand                            ║
║    • Knowledge graph needs                            ║
║                                                        ║
╚════════════════════════════════════════════════════════╝

              👍 Decision Approved
```
