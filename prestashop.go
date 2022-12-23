package gocommerce

func (p *prestashop) TableChecks() []TableCheck {
	return []TableCheck{
		{"configuration", "name", "value", ""},
		{"cms_lang", "id_cms", "content", ""},
		{"employee", "'email'", "email", ""},
		{"vccontentanywhere_lang", "id_vccontentanywhere", "CONCAT_WS(' ',title,content)", ""},
	}
}
