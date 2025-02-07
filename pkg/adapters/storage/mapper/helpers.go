package mapper

import (
	"database/sql"
	"time"

	"github.com/hertzCodes/magnificent-bot/internal/user/domain"
)

func ToNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: !t.IsZero(),
	}
}

func RoleFromInt(roleInt uint64) domain.Role {
	switch roleInt {
	case 1:
		return domain.ValueChanger
	case 2:
		return domain.Admin
	case 3:
		return domain.Owner

	default:
		return domain.NormalUser
	}

}
