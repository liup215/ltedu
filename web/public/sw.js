// LTEdu Service Worker - Provides offline capability and caching for educational data
const CACHE_NAME = 'ltedu-v1'
const API_CACHE_NAME = 'ltedu-api-v1'

// Static assets to pre-cache on install
const PRECACHE_URLS = [
  '/',
  '/manifest.json',
]

// API paths whose responses are safe to cache with stale-while-revalidate.
// These are read-heavy educational data endpoints that change infrequently.
const CACHEABLE_API_PATTERNS = [
  /\/api\/v1\/syllabus\/(list|byId|all)/,
  /\/api\/v1\/chapter\/(list|byId|all)/,
  /\/api\/v1\/organisation\/(list|byId|all)/,
  /\/api\/v1\/qualification\/(list|byId|all)/,
  /\/api\/v1\/exam-node\/(list|byId|all)/,
]

const API_CACHE_TTL_SECONDS = 300 // 5 minutes

// Install: cache essential static assets
self.addEventListener('install', (event) => {
  event.waitUntil(
    caches.open(CACHE_NAME).then((cache) => {
      return cache.addAll(PRECACHE_URLS).catch(() => {
        // ignore pre-cache errors (e.g. offline during install)
      })
    })
  )
  self.skipWaiting()
})

// Activate: clean up old caches
self.addEventListener('activate', (event) => {
  event.waitUntil(
    caches.keys().then((cacheNames) => {
      return Promise.all(
        cacheNames
          .filter((name) => name !== CACHE_NAME && name !== API_CACHE_NAME)
          .map((name) => caches.delete(name))
      )
    })
  )
  self.clients.claim()
})

/**
 * Returns true if the cached response is still within the TTL window.
 */
function isFresh(response) {
  if (!response) return false
  const fetchedAt = response.headers.get('x-sw-fetched-at')
  if (!fetchedAt) return false
  return (Date.now() - parseInt(fetchedAt, 10)) < API_CACHE_TTL_SECONDS * 1000
}

/**
 * Stores a network response in the API cache, annotating it with a fetch
 * timestamp so we can enforce TTL without relying on Cache-Control headers
 * that the backend may not set.
 */
async function cacheApiResponse(request, response) {
  if (!response || !response.ok) return
  const cache = await caches.open(API_CACHE_NAME)
  const cloned = response.clone()
  // Inject a custom header recording when this response was fetched.
  const headers = new Headers(cloned.headers)
  headers.set('x-sw-fetched-at', String(Date.now()))
  const annotated = new Response(await cloned.blob(), {
    status: cloned.status,
    statusText: cloned.statusText,
    headers,
  })
  await cache.put(request, annotated)
}

/**
 * Stale-while-revalidate: return the cached response immediately (if present
 * and fresh), then refresh the cache in the background.  If there is no
 * cached response at all, wait for the network.
 */
async function staleWhileRevalidate(event) {
  const cache = await caches.open(API_CACHE_NAME)
  const cached = await cache.match(event.request.clone())

  const networkPromise = fetch(event.request.clone()).then((response) => {
    cacheApiResponse(event.request.clone(), response.clone())
    return response
  })

  if (cached && isFresh(cached)) {
    // Refresh in background without blocking the response
    event.waitUntil(networkPromise.catch(() => {}))
    return cached
  }

  // No fresh cache: wait for the network (falls back to stale if network fails)
  return networkPromise.catch(() => cached || new Response(
    JSON.stringify({ code: 1, message: 'Service unavailable. Please check your connection.' }),
    { status: 503, headers: { 'Content-Type': 'application/json' } }
  ))
}

// Fetch: choose strategy based on request type
self.addEventListener('fetch', (event) => {
  const url = new URL(event.request.url)

  // Only handle same-origin requests; returning without calling
  // event.respondWith() lets the browser handle cross-origin requests natively.
  if (url.origin !== self.location.origin) return

  // Check if this is a cacheable educational API call
  if (event.request.method === 'POST') {
    const isCacheable = CACHEABLE_API_PATTERNS.some((pattern) =>
      pattern.test(url.pathname)
    )
    if (isCacheable) {
      event.respondWith(staleWhileRevalidate(event))
      return
    }
    // All other API (POST) calls go straight to network
    return
  }

  // Static assets: cache-first, network fallback
  if (event.request.method === 'GET') {
    event.respondWith(
      caches.match(event.request).then((cached) => {
        return cached || fetch(event.request)
      })
    )
  }
})
