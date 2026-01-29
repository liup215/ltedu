**Service Layer Refactor Workflow**

1. **Define Service Interface**
   - Create an interface listing all public service methods for the entity.

2. **Implement Private Struct**
   - Implement the interface on a private struct.
   - Place all business logic and method implementations here.

3. **Export Public Variable**
   - Export a public variable of the interface type, assigned to the private struct instance.

4. **Use Unified DAO Pattern**
   - Read dao/(entity).go to learn the dao functions.
   - Replace direct model queries and legacy DAO calls with unified DAO CRUD methods and Query struct.

5. **Method Implementation**
   - Use DAO CRUD methods for data access.
   - All CRUD methods must use DAO, and all data operation should be kept.
   - Use Query struct for filtering and pagination.
   - Handle errors and wrap with context as needed.

6. **Documentation**
   - Write all code and comments in English.
   - Document method purpose, parameters, and error handling.

7. **Testing**
   - Ensure service methods are testable via the interface.
   - Mock the interface for unit tests.

**Example:**
```go
type EntityServiceInterface interface {
  CreateEntity(e model.Entity) (*model.Entity, error)
  GetEntityById(id uint) (*model.Entity, error)
  ListEntities(query model.EntityQuery) ([]*model.Entity, int64, error)
}

var EntityService EntityServiceInterface = &entityService{}

type entityService struct{}

func (svr *entityService) CreateEntity(e model.Entity) (*model.Entity, error) {
  // Use DAO CRUD methods
  return dao.EntityDao.Create(e)
}
```
Apply this workflow to each service file for consistency and maintainability.