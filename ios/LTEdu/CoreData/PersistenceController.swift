import CoreData
import Foundation

// MARK: - PersistenceController

final class PersistenceController {
    static let shared = PersistenceController()

    static var preview: PersistenceController = {
        let result = PersistenceController(inMemory: true)
        let context = result.container.viewContext
        // Seed preview data
        let cachedUser = CachedUserEntity(context: context)
        cachedUser.id = 1
        cachedUser.username = "preview_user"
        cachedUser.nickname = "Preview User"
        cachedUser.email = "preview@ltedu.com"
        cachedUser.cachedAt = Date()
        try? context.save()
        return result
    }()

    let container: NSPersistentContainer

    init(inMemory: Bool = false) {
        container = NSPersistentContainer(name: "LTEdu")

        if inMemory {
            container.persistentStoreDescriptions.first?.url = URL(fileURLWithPath: "/dev/null")
        }

        container.loadPersistentStores { _, error in
            if let error = error {
                // In production, handle this gracefully rather than crashing
                print("Core Data store failed to load: \(error.localizedDescription)")
            }
        }

        container.viewContext.automaticallyMergesChangesFromParent = true
        container.viewContext.mergePolicy = NSMergeByPropertyObjectTrumpMergePolicy
    }

    var viewContext: NSManagedObjectContext {
        container.viewContext
    }

    func newBackgroundContext() -> NSManagedObjectContext {
        container.newBackgroundContext()
    }

    func save(context: NSManagedObjectContext? = nil) {
        let ctx = context ?? container.viewContext
        guard ctx.hasChanges else { return }
        do {
            try ctx.save()
        } catch {
            print("Core Data save error: \(error.localizedDescription)")
        }
    }
}

// MARK: - Cache Manager

final class CacheManager {
    static let shared = CacheManager()
    private let persistence = PersistenceController.shared

    private init() {}

    // MARK: - User Cache

    func cacheUser(_ user: User) {
        let context = persistence.newBackgroundContext()
        context.perform {
            let request: NSFetchRequest<CachedUserEntity> = CachedUserEntity.fetchRequest()
            request.predicate = NSPredicate(format: "id == %d", user.id)

            let entity = (try? context.fetch(request))?.first ?? CachedUserEntity(context: context)
            entity.id = Int64(user.id)
            entity.username = user.username
            entity.email = user.email
            entity.nickname = user.nickname
            entity.realname = user.realname
            entity.avatar = user.avatar
            entity.isAdmin = user.isAdmin ?? false
            entity.isTeacher = user.isTeacher ?? false
            entity.cachedAt = Date()

            self.persistence.save(context: context)
        }
    }

    func getCachedUser(id: Int) -> CachedUserEntity? {
        let request: NSFetchRequest<CachedUserEntity> = CachedUserEntity.fetchRequest()
        request.predicate = NSPredicate(format: "id == %d", id)
        request.fetchLimit = 1
        return try? persistence.viewContext.fetch(request).first
    }

    // MARK: - Course Cache

    func cacheCourses(_ courses: [Course]) {
        let context = persistence.newBackgroundContext()
        context.perform {
            for course in courses {
                let request: NSFetchRequest<CachedCourseEntity> = CachedCourseEntity.fetchRequest()
                request.predicate = NSPredicate(format: "id == %d", course.id)

                let entity = (try? context.fetch(request))?.first ?? CachedCourseEntity(context: context)
                entity.id = Int64(course.id)
                entity.title = course.title
                entity.courseDescription = course.shortDescription
                entity.thumb = course.thumb
                entity.isFree = course.isFreeAccess
                entity.syllabusId = Int64(course.syllabusId ?? 0)
                entity.cachedAt = Date()
            }
            self.persistence.save(context: context)
        }
    }

    func getCachedCourses(syllabusID: Int? = nil) -> [CachedCourseEntity] {
        let request: NSFetchRequest<CachedCourseEntity> = CachedCourseEntity.fetchRequest()
        if let syllabusID = syllabusID {
            request.predicate = NSPredicate(format: "syllabusId == %d", syllabusID)
        }
        request.sortDescriptors = [NSSortDescriptor(key: "cachedAt", ascending: false)]
        return (try? persistence.viewContext.fetch(request)) ?? []
    }

    // MARK: - Chat Session Cache

    func saveChatSession(_ session: ChatSession) {
        guard let data = try? JSONEncoder().encode(session) else { return }
        let context = persistence.newBackgroundContext()
        context.perform {
            let request: NSFetchRequest<CachedChatSessionEntity> = CachedChatSessionEntity.fetchRequest()
            request.predicate = NSPredicate(format: "sessionId == %@", session.id.uuidString)

            let entity = (try? context.fetch(request))?.first ?? CachedChatSessionEntity(context: context)
            entity.sessionId = session.id.uuidString
            entity.title = session.title
            entity.sessionData = data
            entity.updatedAt = session.updatedAt

            self.persistence.save(context: context)
        }
    }

    func loadChatSessions() -> [ChatSession] {
        let request: NSFetchRequest<CachedChatSessionEntity> = CachedChatSessionEntity.fetchRequest()
        request.sortDescriptors = [NSSortDescriptor(key: "updatedAt", ascending: false)]
        let entities = (try? persistence.viewContext.fetch(request)) ?? []
        return entities.compactMap { entity in
            guard let data = entity.sessionData else { return nil }
            return try? JSONDecoder().decode(ChatSession.self, from: data)
        }
    }

    func deleteChatSession(id: UUID) {
        let request: NSFetchRequest<CachedChatSessionEntity> = CachedChatSessionEntity.fetchRequest()
        request.predicate = NSPredicate(format: "sessionId == %@", id.uuidString)
        if let entity = (try? persistence.viewContext.fetch(request))?.first {
            persistence.viewContext.delete(entity)
            persistence.save()
        }
    }

    // MARK: - Cache Invalidation

    /// Removes cached data older than the given age in seconds
    func invalidateStaleCache(olderThan age: TimeInterval = 3600) {
        let cutoff = Date().addingTimeInterval(-age)
        let context = persistence.newBackgroundContext()
        context.perform {
            for entityName in ["CachedCourseEntity"] {
                let request = NSFetchRequest<NSManagedObject>(entityName: entityName)
                request.predicate = NSPredicate(format: "cachedAt < %@", cutoff as NSDate)
                if let objects = try? context.fetch(request) {
                    objects.forEach { context.delete($0) }
                }
            }
            self.persistence.save(context: context)
        }
    }
}
