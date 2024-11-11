package legacyconversions

import (
	"github.com/ga-mitchell-l/def-prog-exercises/safesql"
	"github.com/ga-mitchell-l/def-prog-exercises/safesql/internal/raw"
)

var trustedSQLCtor = raw.TrustedSQLCtor.(func(string) safesql.TrustedSQL)

func RiskilyAssumeTrustedSQL(trusted string) safesql.TrustedSQL {
	return trustedSQLCtor(trusted)
}
