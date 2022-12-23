package gocommerce

func (w *woocommerce) TableChecks() []TableCheck {
	return []TableCheck{
		{"posts", "ID", "post_content", "post_type != 'shop_order' AND length(post_content) > 64"}, // somewhat arbitrary number to kill small files
		{"options", "option_id", "option_value", "length(option_value) > 64"},
		// Temp disabled, silver has 280K users which would take multiple hours to scan
		// There seems no simple way to filter for admin accounts
		// {"users", "user_login", "user_email", ""},
	}
}
