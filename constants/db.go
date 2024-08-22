package constants

import (
	"time"
)

const DB = "mess"
const CONNECTION_STRING = "ATLAS_URI"
const REDIS_URI = "REDIS_URI"
const REDIS_USERNAME = "REDIS_USERNAME"
const REDIS_PASSWORD = "REDIS_PASSWORD"

const COLLECTION_GROUPS = "groups"
const COLLECTION_HOSTELS = "hostels"
const COLLECTION_MENU_ITEMS = "menuItems"
const COLLECTION_MEALS = "meals"
const COLLECTION_MEAL_TYPES = "mealTypes"
const COLLECTION_MESSES = "messes"
const COLLECTION_ROOMS = "rooms"
const COLLECTION_USERS = "users"
const COLLECTION_SEMESTERS = "semesters"
const COLLECTION_ROOM_ALLOTMENTS = "room_allotments"
const COLLECTION_HOSTEL_STAFF_ALLOTMENTS = "hostel_staff_allotments"

const FIREBASE_PROJECT_ID = "FIREBASE_PROJECT_ID"

const CACHING_DURATION = 20 * time.Minute
const CACHE_CONTROL_HEADER = "Cache-Control"
const NO_CACHE = "no-cache"

// Keys of Cache
const GCP_JWKS = "GCP_JWKS"

const DB_PAGINATION = 30      // 30 results will be returned for DB pagination process
const DB_MAX_CYCLE_COUNT = 10 // The max number of cycles the DB will do to interupt the request
