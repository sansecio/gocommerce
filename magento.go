package gocommerce

var MagentoTables = []TableCheck{
	{"core_config_data", "path", "value", ""},
	{"cms_page", "page_id", "CONCAT_WS('',content,custom_theme)", ""},
	{"cms_block", "block_id", "CONCAT_WS('',title,content)", ""},
	{"newsletter_template", "template_code", "template_text", ""}, // newsletter_template.template_text
	{"admin_user", "'email'", "email", ""},                        // Uses 'email' as identifier, so "db:admin_user.email" can be used as pathfilter.
	{"information_schema.triggers", "event_object_table", "action_statement", ""},
	{"core_file_storage", "CONCAT_WS('/',directory,filename)", "content", ""},
	{"datafeedmanager_attributes", "attribute_name", "attribute_script", ""}, // Wyomind_Datafeedmanager, 7.14.0.2  see case rufflebutts/2021-03-25/joost.md
	{"datafeedmanager_configurations", "feed_id", "feed_product", ""},        // _ALSO_ Wyomind_Datafeedmanager, 7.14.0.2  see case rufflebutts/2022-01-26/willem.md
	{"datafeedmanager_options", "option_id", "option_script", ""},
	{"customer_eav_attribute", "attribute_id", "validate_rules", ""}, // per M1 quickview hack, see case file ayquechula.com/2022-02-03
	{"sales_order_address", "entity_id", "CONCAT_WS('', firstname, lastname, middlename, prefix, suffix, company, street, city, vat_id)", "1 ORDER BY entity_id DESC LIMIT 1000"},                        // M2 CVE 2022-24086, reduce performance load with limited query
	{"sales_order", "customer_id", "CONCAT_WS('', customer_firstname, customer_lastname, customer_middlename, customer_prefix, customer_suffix, customer_note)", "1 ORDER BY entity_id DESC LIMIT 1000"}, // M2 CVE 2022-24086, reduce performance load with limited query
	{"quote_address", "quote_id", "CONCAT_WS('', firstname, lastname, middlename, prefix, suffix, company, street, city, vat_id)", "1 ORDER BY quote_id DESC LIMIT 1000"},                                // M2 CVE 2022-24086, reduce performance load with limited query
	{"quote", "entity_id", "CONCAT_WS('', customer_firstname, customer_lastname, customer_middlename, customer_prefix, customer_suffix, customer_note)", "1 ORDER BY entity_id DESC LIMIT 1000"},         // M2 CVE 2022-24086, reduce performance load with limited query
}
