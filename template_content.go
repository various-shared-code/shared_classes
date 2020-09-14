package classes

type TemplateContent struct {
	Title       string `json:"title"`
	Description string `json:"description"`

	Components []EditableComponent `json:"components"`
}
