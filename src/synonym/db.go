package synonym

type DB struct {
	Reader
}

func New(filepath string, encoding string) (*DB, error) {
	var (
		e  error
		db = new(DB)
	)

	dcr, e := newDashComaReader(filepath, false)
	if e != nil {
		return nil, e
	}

	db.Reader = dcr

	return db, nil
}
