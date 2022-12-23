package gocommerce

func (s *shopware6) TableChecks() []TableCheck {
	return []TableCheck{
		{"cms_slot_translation", "created_at", "config", ""},
		{"mail_template_translation", "created_at", "content_html", ""},
		{"cms_page", "created_at", "config", ""},
		{"cms_block", "created_at", "custom_fields", ""},
		{"user", "'email'", "email", ""}, // TODO: figure out how we can generalize magento rogue admin pathfilter
		{"system_config", "HEX(id)", "configuration_value", ""},
	}
}
