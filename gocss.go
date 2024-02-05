package gocss

func CreateStyle(css CSS) string {
	className := generateId()
	textContent := "." + className + "{" + css.String() + "}"
	insertRule(textContent)
	return className
}
