package goExt

func CheckErr(err error) bool  {
	if err != nil {
		Log(err)
		return false
	}
	return true
}

