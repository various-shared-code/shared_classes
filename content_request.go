package classes

type ContentRequest struct {
	Domain     string `json:"domain"`
	Path       string `json:"path"`
	TemplateID int64  `json:"template_id"`
}
