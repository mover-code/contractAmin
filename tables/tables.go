package tables

import (
	"strconv"
	"time"

	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

// generators is a map of table models.
//
// The key of generators is the prefix of table info url.
// The corresponding value is the Form and TableName data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "users"   => http://localhost:9033/admin/info/users
// "posts"   => http://localhost:9033/admin/info/posts
// "authors" => http://localhost:9033/admin/info/authors
//
var Generators = map[string]table.Generator{
	"contract": GetContract,
}

// ThisTime returns the current time in seconds since the epoch.
func ThisTime() int64 {
	return time.Now().Unix()
}

// Convert a time.Time to a string in the format YYYY-MM-DD HH:MM:SS.
func TimeToStr(s string) string {
	t, _ := strconv.ParseInt(s, 10, 64)
	tm := time.Unix(t, 0)
	return tm.Format("2006-01-02 15:04:05")
}
