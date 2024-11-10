package server

func isNoRowsError(err error) bool {
	return err != nil && err.Error() == "no rows in dis set"
}
