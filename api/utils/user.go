package utils

import (
	"database/sql"
	"fmt"
	"strings"
	"unicode"

	"github.com/jmoiron/sqlx"
	"golang.org/x/text/unicode/norm"
)

func GenerateUniqueUserHandle(db *sqlx.DB, firstName, lastName sql.NullString) (string, error) {
	base := ""
	randomSuffixLen := 4

	if firstName.Valid && lastName.Valid && len(firstName.String) > 0 {
		base = fmt.Sprintf("%c%s", firstName.String[0], lastName.String)
	} else if firstName.Valid {
		base = firstName.String
	} else if lastName.Valid {
		base = lastName.String
	}
	base = NormalizeHandle(base)
	if base == "user" {
		randomSuffixLen = 6
	}

	handle := base
	var exists bool
	for {
		err := db.QueryRow(`select exists(select 1 from users where handle = $1)`, handle).Scan(&exists)
		if err != nil {
			return "", err
		}
		if !exists {
			break
		}
		randomSuffix := RandomString(randomSuffixLen)
		handle = fmt.Sprintf("%s-%s", base, randomSuffix)
	}

	return handle, nil
}

func NormalizeHandle(s string) string {
	s = norm.NFKD.String(strings.ToLower(s))

	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '.' || r == '_' || r == '-' {
			b.WriteRune(r)
		}
	}

	out := strings.Trim(b.String(), "._-")
	if out == "" {
		return "user"
	}

	return out
}

// when adding a new oauth provider and user table fields, add emails here:
const EMAIL_SQL_MATCH_CONDITION = "email = $1 or github_email = $1 or google_email = $1 or facebook_email = $1 or spotify_email = $1"

func UserEmailExistsQuery() string {
	return `exists(select 1 from users where ` + EMAIL_SQL_MATCH_CONDITION + `)`
}
func SelectUserByEmailQuery() string {
	return `select * from users where ` + EMAIL_SQL_MATCH_CONDITION
}
func SelectIdAndPasswordHashByEmailQuery() string {
	return `select id, password_hash from users where ` + EMAIL_SQL_MATCH_CONDITION
}
func SelectUserById() string {
	return "select * from users where id = $1"
}

type ThemePosibilities string

const (
	ThemeLight ThemePosibilities = "light"
	ThemeDark  ThemePosibilities = "dark"
)

// Implement ISwaggerEnum
func (e ThemePosibilities) EnumValues() []any {
	return []any{ThemeLight, ThemeDark}
}

type TranslationsPossibilities string

const (
	LangCS TranslationsPossibilities = "cs"
	LangEN TranslationsPossibilities = "en"
)

// Implement ISwaggerEnum
func (e TranslationsPossibilities) EnumValues() []any {
	return []any{LangCS, LangEN}
}
