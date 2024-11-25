package typ

import "github.com/jackc/pgx/v5/pgtype"

func Text(text string) pgtype.Text {
	return pgtype.Text{
		String: text,
		Valid:  len(text) != 0,
	}
}

func Int4(data int) pgtype.Int4 {
	return pgtype.Int4{
		Int32: int32(data),
		Valid: data != 0,
	}
}
