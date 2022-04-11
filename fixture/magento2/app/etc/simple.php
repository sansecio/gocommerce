<?php
return array(
  'cache_types' =>
  array(
    'compiled_config' => 1,
    'config' => 1,
    'layout' => 1,
    'block_html' => 1,
    'collections' => 1,
    'reflection' => 1,
    'db_ddl' => 1,
    'eav' => 1,
    'customer_notification' => 1,
    'target_rule' => 1,
    'full_page' => 1,
    'config_integration' => 1,
    'config_integration_api' => 1,
    'translate' => 1,
    'config_webservice' => 1,
    'amasty_shopby' => 1,
  ),
  'backend' =>
  array(
    'frontName' => 'admin',
  ),
  'db' =>
  array(
    'connection' =>
    array(
      'default' =>
      array(
        'username' => 'gooduser',
        'host' => 'goodhost',
        'dbname' => 'gooddb',
        'password' => 'verylongpassword',
        'model' => 'mysql4',
        'engine' => 'innodb',
        'initStatements' => 'SET NAMES utf8;',
        'active' => '1',
      ),
      'indexer' =>
      array(
        'username' => 'indexuser',
        'host' => 'indexhost',
        'dbname' => 'indexdb',
        'password' => 'indexpass',
        'model' => 'mysql4',
        'engine' => 'innodb',
        'initStatements' => 'SET NAMES utf8;',
        'active' => '1',
        'persistent' => NULL,
      ),
    ),
    'table_prefix' => 'pref_',
  ),
  'crypt' =>
  array(
    'key' => '0de48f26a28869f3a73f7e81ab0abac7',
  ),
  'session' =>
  array(
    'save' => 'redis',
    'redis' =>
    array(
      'host' => 'redis',
      'port' => '6379',
      'database' => 0,
      'disable_locking' => 1,
    ),
  ),
  'resource' =>
  array(
    'default_setup' =>
    array(
      'connection' => 'default',
    ),
  ),
  'x-frame-options' => 'SAMEORIGIN',
  'MAGE_MODE' => 'production',
  'install' =>
  array(
    'date' => 'Tue, 08 Aug 2017 20:08:01 +0000',
  ),
  'cache' =>
  array(
    'frontend' =>
    array(
      'default' =>
      array(
        'backend' => 'Cm_Cache_Backend_Redis',
        'backend_options' =>
        array(
          'server' => 'redis',
          'port' => '6379',
          'database' => 1,
        ),
      ),
      'page_cache' =>
      array(
        'backend' => 'Cm_Cache_Backend_Redis',
        'backend_options' =>
        array(
          'server' => 'redis',
          'port' => '6379',
          'database' => 1,
        ),
      ),
    ),
  ),
  'system' =>
  array(
    'default' =>
    array(
      'dev' =>
      array(
        'js' =>
        array(
          'session_storage_key' => 'collected_errors',
        ),
        'restrict' =>
        array(
          'allow_ips' => NULL,
        ),
      ),
      'system' =>
      array(
        'full_page_cache' =>
        array(
          'varnish' =>
          array(
            'access_list' => 'localhost',
            'backend_host' => 'localhost',
            'backend_port' => '8080',
          ),
        ),
        'smtp' =>
        array(
          'host' => 'localhost',
          'port' => '25',
        ),
        'magento_scheduled_import_export_log' =>
        array(
          'error_email' => NULL,
        ),
      ),
      'web' =>
      array(
        'unsecure' =>
        array(
          'base_url' => 'http://myurl.com/',
          'base_link_url' => '{{unsecure_base_url}}',
          'base_static_url' => NULL,
          'base_media_url' => NULL,
        ),
        'secure' =>
        array(
          'base_url' => 'https://myurl.com/',
          'base_link_url' => '{{secure_base_url}}',
          'base_static_url' => NULL,
          'base_media_url' => NULL,
        ),
        'default' =>
        array(
          'front' => 'cms',
        ),
        'cookie' =>
        array(
          'cookie_path' => NULL,
          'cookie_domain' => NULL,
        ),
      ),
      'admin' =>
      array(
        'url' =>
        array(
          'custom' => NULL,
          'custom_path' => 'secretadminpath',
        ),
      ),
      'currency' =>
      array(
        'import' =>
        array(
          'error_email' => NULL,
        ),
      ),
      'customer' =>
      array(
        'create_account' =>
        array(
          'email_domain' => 'example.com',
        ),
      ),
      'catalog' =>
      array(
        'productalert_cron' =>
        array(
          'error_email' => NULL,
        ),
        'search' =>
        array(
          'elasticsearch_server_hostname' => 'elasticsearch',
          'elasticsearch_server_port' => '9200',
          'elasticsearch_index_prefix' => 'magento2',
          'elasticsearch_enable_auth' => '0',
          'elasticsearch_server_timeout' => '15',
          'solr_server_hostname' => 'localhost',
          'solr_server_port' => '8983',
          'solr_server_username' => 'admin',
          'solr_server_timeout' => '15',
          'solr_server_path' => 'solr',
        ),
        'product_video' =>
        array(
          'youtube_api_key' => 'youtubesecret api key',
        ),
      ),
      'payment' =>
      array(
        'cybersource' =>
        array(
          'sandbox_flag' => '1',
          'access_key' => NULL,
          'profile_id' => NULL,
          'secret_key' => NULL,
          'merchant_id' => NULL,
          'transaction_key' => NULL,
        ),
        'eway' =>
        array(
          'sandbox_flag' => '1',
          'live_api_key' => NULL,
          'live_api_password' => NULL,
          'live_encryption_key' => NULL,
          'sandbox_api_key' => NULL,
          'sandbox_api_password' => NULL,
          'sandbox_encryption_key' => NULL,
        ),
        'checkmo' =>
        array(
          'mailing_address' => NULL,
        ),
        'paypal_express' =>
        array(
          'debug' => '1',
          'merchant_id' => 'my merchant id',
        ),
        'paypal_express_bml' =>
        array(
          'publisher_id' => NULL,
        ),
        'payflow_express' =>
        array(
          'debug' => '0',
        ),
        'payflowpro' =>
        array(
          'user' => NULL,
          'pwd' => NULL,
          'partner' => NULL,
          'sandbox_flag' => '0',
          'debug' => '0',
        ),
        'paypal_billing_agreement' =>
        array(
          'debug' => '0',
        ),
        'payflow_link' =>
        array(
          'pwd' => NULL,
          'url_method' => 'GET',
          'sandbox_flag' => '0',
          'use_proxy' => '0',
          'debug' => '0',
        ),
        'payflow_advanced' =>
        array(
          'user' => NULL,
          'pwd' => NULL,
          'url_method' => 'GET',
          'sandbox_flag' => '0',
          'debug' => '0',
        ),
        'authorizenet_directpost' =>
        array(
          'debug' => '0',
          'email_customer' => '0',
          'login' => NULL,
          'merchant_email' => NULL,
          'test' => '1',
          'trans_key' => NULL,
          'trans_md5' => NULL,
          'cgi_url' => 'https://secure.authorize.net/gateway/transact.dll',
          'cgi_url_td' => 'https://api2.authorize.net/xml/v1/request.api',
        ),
        'braintree' =>
        array(
          'private_key' => NULL,
          'merchant_id' => NULL,
          'merchant_account_id' => NULL,
          'descriptor_phone' => NULL,
          'descriptor_url' => NULL,
        ),
        'braintree_paypal' =>
        array(
          'merchant_name_override' => NULL,
        ),
        'worldpay' =>
        array(
          'response_password' => NULL,
          'auth_password' => NULL,
          'md5_secret' => NULL,
          'sandbox_flag' => '1',
          'signature_fields' => 'instId:cartId:amount:currency',
          'installation_id' => NULL,
        ),
      ),
      'sales_email' =>
      array(
        'order' =>
        array(
          'copy_to' => 'myemail@dom.com',
        ),
        'order_comment' =>
        array(
          'copy_to' => NULL,
        ),
        'invoice' =>
        array(
          'copy_to' => NULL,
        ),
        'invoice_comment' =>
        array(
          'copy_to' => NULL,
        ),
        'shipment' =>
        array(
          'copy_to' => NULL,
        ),
        'shipment_comment' =>
        array(
          'copy_to' => NULL,
        ),
        'creditmemo' =>
        array(
          'copy_to' => NULL,
        ),
        'creditmemo_comment' =>
        array(
          'copy_to' => NULL,
        ),
        'magento_rma' =>
        array(
          'copy_to' => NULL,
        ),
        'magento_rma_auth' =>
        array(
          'copy_to' => NULL,
        ),
        'magento_rma_comment' =>
        array(
          'copy_to' => NULL,
        ),
        'magento_rma_customer_comment' =>
        array(
          'copy_to' => NULL,
        ),
      ),
      'promo' =>
      array(
        'magento_reminder' =>
        array(
          'identity' => 'general',
        ),
      ),
      'checkout' =>
      array(
        'payment_failed' =>
        array(
          'copy_to' => 'help@my.com',
        ),
      ),
      'contact' =>
      array(
        'email' =>
        array(
          'recipient_email' => 'help@my.com',
        ),
      ),
      'trans_email' =>
      array(
        'ident_custom1' =>
        array(
          'email' => 'test@test.com',
          'name' => 'Custom 1',
        ),
      ),
      'analytics' =>
      array(
        'url' =>
        array(
          'signup' => 'https://advancedreporting.rjmetrics.com/signup',
          'update' => 'https://advancedreporting.rjmetrics.com/update',
          'bi_essentials' => 'https://dashboard.rjmetrics.com/v2/magento/signup',
          'otp' => 'https://advancedreporting.rjmetrics.com/otp',
          'report' => 'https://advancedreporting.rjmetrics.com/report',
          'notify_data_changed' => 'https://advancedreporting.rjmetrics.com/report',
        ),
        'general' =>
        array(
          'token' => 'secrettoken-WiX-hoi',
        ),
      ),
      'carriers' =>
      array(
        'dhl' =>
        array(
          'account' => NULL,
          'gateway_url' => 'https://xmlpi-ea.dhl.com/XMLShippingServlet',
          'id' => NULL,
          'password' => NULL,
          'debug' => '0',
        ),
        'fedex' =>
        array(
          'account' => NULL,
          'meter_number' => NULL,
          'key' => NULL,
          'password' => NULL,
          'sandbox_mode' => '0',
          'production_webservices_url' => 'https://ws.fedex.com:443/web-services/',
          'sandbox_webservices_url' => 'https://wsbeta.fedex.com:443/web-services/',
          'smartpost_hubid' => NULL,
        ),
      ),
      'google' =>
      array(
        'analytics' =>
        array(
          'account' => 'UA-234234234-1',
        ),
      ),
      'newrelicreporting' =>
      array(
        'general' =>
        array(
          'api_url' => 'https://api.newrelic.com/deployments.xml',
          'insights_api_url' => 'https://insights-collector.newrelic.com/v1/accounts/%s/events',
        ),
      ),
      'paypal' =>
      array(
        'fetch_reports' =>
        array(
          'ftp_login' => NULL,
          'ftp_password' => NULL,
          'ftp_sandbox' => '0',
          'ftp_ip' => NULL,
          'ftp_path' => NULL,
        ),
      ),
      'fraud_protection' =>
      array(
        'signifyd' =>
        array(
          'api_url' => 'https://api.signifyd.com/v2/',
          'api_key' => NULL,
        ),
      ),
      'sitemap' =>
      array(
        'generate' =>
        array(
          'error_email' => NULL,
        ),
      ),
      'crontab' =>
      array(
        'default' =>
        array(
          'jobs' =>
          array(
            'analytics_collect_data' =>
            array(
              'schedule' =>
              array(
                'cron_expr' => '00 02 * * *',
              ),
            ),
            'analytics_subscribe' =>
            array(
              'schedule' =>
              array(
                'cron_expr' => '0 * * * *',
              ),
            ),
          ),
        ),
      ),
    ),
    'websites' =>
    array(
      'base' =>
      array(
        'trans_email' =>
        array(
          'ident_sales' =>
          array(
            'name' => 'Customer Service',
          ),
          'ident_support' =>
          array(
            'name' => 'Customer Service',
          ),
          'ident_custom1' =>
          array(
            'name' => 'Customer Service',
          ),
          'ident_custom2' =>
          array(
            'name' => 'Customer Service',
          ),
          'ident_general' =>
          array(
            'name' => 'Customer Service',
          ),
        ),
        'contact' =>
        array(
          'email' =>
          array(
            'recipient_email' => '',
          ),
        ),
      ),
    ),
  ),
  'static_content_on_demand_in_production' => 0,
  'force_html_minification' => 1,
  'cron_consumers_runner' =>
  array(
    'cron_run' => false,
    'max_messages' => 10000,
    'consumers' =>
    array(),
  ),
  'directories' =>
  array(
    'document_root_is_pub' => true,
  ),
  'http_cache_hosts' =>
  array(
    0 =>
    array(
      'host' => 'varnish',
    ),
  ),
);
