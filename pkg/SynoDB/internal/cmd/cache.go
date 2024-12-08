package cmd

import (
	"strings"

	"github.com/khulnasoft/synodb/internal/settings"
	"github.com/khulnasoft/synodb/internal/synodb"
)

const (
	DB_CACHE_KEY         = "database_names"
	DB_CACHE_TTL_SECONDS = 30 * 60
)

func setDatabasesCache(dbNames []synodb.Database) {
	settings.SetCache(DB_CACHE_KEY, DB_CACHE_TTL_SECONDS, dbNames)
}

func getDatabasesCache() []synodb.Database {
	data, err := settings.GetCache[[]synodb.Database](DB_CACHE_KEY)
	if err != nil {
		return nil
	}
	return data
}

func invalidateDatabasesCache() {
	settings.InvalidateCache[[]synodb.Database](DB_CACHE_KEY)
}

const (
	REGIONS_CACHE_KEY         = "locations"
	REGIONS_CACHE_TTL_SECONDS = 8 * 60 * 60
)

func setLocationsCache(locations map[string]string) {
	settings.SetCache(REGIONS_CACHE_KEY, REGIONS_CACHE_TTL_SECONDS, locations)
}

func locationsCache() map[string]string {
	locations, err := settings.GetCache[map[string]string](REGIONS_CACHE_KEY)
	if err != nil {
		return nil
	}
	return locations
}

const CLOSEST_LOCATION_CACHE_KEY = "closestLocation"

func setClosestLocationCache(closest string) {
	settings.SetCache(CLOSEST_LOCATION_CACHE_KEY, REGIONS_CACHE_TTL_SECONDS, closest)
}

func closestLocationCache() string {
	defaultLocation, err := settings.GetCache[string](CLOSEST_LOCATION_CACHE_KEY)
	if err != nil {
		return ""
	}
	return defaultLocation
}

const TOKEN_VALID_CACHE_KEY_PREFIX = "token_valid."

func setTokenValidCache(token string, exp int64) {
	key := TOKEN_VALID_CACHE_KEY_PREFIX + strings.ReplaceAll(token, ".", "_")
	settings.SetCacheWithExp(key, exp, true)
}

func tokenValidCache(token string) bool {
	key := TOKEN_VALID_CACHE_KEY_PREFIX + strings.ReplaceAll(token, ".", "_")
	ok, err := settings.GetCache[bool](key)
	return err == nil && ok
}

const DATABASE_TOKEN_KEY_PREFIX = "database_token."

func setDbTokenCache(dbID, token string, exp int64) {
	key := DATABASE_TOKEN_KEY_PREFIX + dbID
	settings.SetCacheWithExp(key, exp, token)
}

func dbTokenCache(dbID string) string {
	key := DATABASE_TOKEN_KEY_PREFIX + dbID
	token, err := settings.GetCache[string](key)
	if err != nil {
		return ""
	}
	return token
}

func invalidateDbTokenCache() {
	settings.SetCacheRaw(DATABASE_TOKEN_KEY_PREFIX[:len(DATABASE_TOKEN_KEY_PREFIX)-1], struct{}{})
}

const (
	ORG_CACHE_KEY           = "organizations"
	GROUP_CACHE_KEY         = "groups"
	GROUP_CACHE_TTL_SECONDS = 30 * 60
)

func orgKey(org, suffix string) string {
	key := suffix
	if org != "" {
		key = ORG_CACHE_KEY + "." + org + "." + suffix
	}
	return key
}

func setGroupsCache(org string, groups []synodb.Group) {
	settings.SetCache(orgKey(org, GROUP_CACHE_KEY), GROUP_CACHE_TTL_SECONDS, groups)
}

func getGroupsCache(org string) []synodb.Group {
	data, err := settings.GetCache[[]synodb.Group](orgKey(org, GROUP_CACHE_KEY))
	if err != nil {
		return nil
	}
	return data
}

func invalidateGroupsCache(org string) {
	settings.InvalidateCache[[]synodb.Group](orgKey(org, GROUP_CACHE_KEY))
}
