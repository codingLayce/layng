package object

type errUnknownVariable struct {
	Name string
}

func (e *errUnknownVariable) Error() string {
	return "unknown variable: " + e.Name
}

type errDifferentType struct {
	Base   Type
	Wanted Type
}

func (e *errDifferentType) Error() string {
	return "cannot assign the type `" + e.Wanted.String + "` to a variable of type `" + e.Base.String + "`"
}
