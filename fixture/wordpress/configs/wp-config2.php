<?php
define('WP_CACHE', true); // Added by WP Rocket



/**
 * The base configuration for WordPress
 *
 * The wp-config.php creation script uses this file during the
 * installation. You don't have to use the web site, you can
 * copy this file to "wp-config.php" and fill in the values.
 *
 * This file contains the following configurations:
 *
 * * MySQL settings
 * * Secret keys
 * * Database table prefix
 * * ABSPATH
 *
 * @link https://codex.wordpress.org/Editing_wp-config.php
 *
 * @package WordPress
 */

// ** MySQL settings - You can get this info from your web host ** //
/** The name of the database for WordPress */
define( 'DB_NAME', 'xx_db' );

/** MySQL database username */
define( 'DB_USER', 'xxxuser' );

/** MySQL database password */
define( 'DB_PASSWORD', 'xxpass' );

/** MySQL hostname */
define( 'DB_HOST', 'xxhost' );

/** Database Charset to use in creating database tables. */
define('DB_CHARSET', 'utf8');

/** The Database Collate type. Don't change this if in doubt. */
define('DB_COLLATE', '');

/**#@+
 * Authentication Unique Keys and Salts.
 *
 * Change these to different unique phrases!
 * You can generate these using the {@link https://api.wordpress.org/secret-key/1.1/salt/ WordPress.org secret-key service}
 * You can change these at any point in time to invalidate all existing cookies. This will force all users to have to log in again.
 *
 * @since 2.6.0
 */
define('AUTH_KEY',         'xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxhgoe135te4wxmockrrrvze4md');
define('SECURE_AUTH_KEY',  'xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxdyzixuixwvymepdzw3rbwuyzi');
define('LOGGED_IN_KEY',    'xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxo81kjvdhhxdb5caq0zyusrpyc');
define('NONCE_KEY',        'xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx4kaw9a1so4uvysg7jmhslcyl');
define('AUTH_SALT',        'xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxs4hxrwto8aswplqq3pffib');
define('SECURE_AUTH_SALT', 'xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxsigxhgfqc4mbpb2agrwmq');
define('LOGGED_IN_SALT',   'xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxfzbjhd3bkndavyfsezp1');
define('NONCE_SALT',       'xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxc4zdq2mzq2zwi0m2q4md');

/**#@-*/

/**
 * WordPress Database Table prefix.
 *
 * You can have multiple installations in one database if you give each
 * a unique prefix. Only numbers, letters, and underscores please!
 */
$table_prefix = 'wp102_';

/**
 * For developers: WordPress debugging mode.
 *
 * Change this to true to enable the display of notices during development.
 * It is strongly recommended that plugin and theme developers use WP_DEBUG
 * in their development environments.
 *
 * For information on other constants that can be used for debugging,
 * visit the Codex.
 *
 * @link https://codex.wordpress.org/Debugging_in_WordPress
 */
 // Enable WP_DEBUG mode
 define('WP_DEBUG', true);

 // Enable Debug logging to the /wp-content/debug.log file
 define('WP_DEBUG_LOG', true);

 // Disable display of errors and warnings
 define('WP_DEBUG_DISPLAY', false);
 @ini_set('display_errors',0);

 define( 'WP_MEMORY_LIMIT', '128M' );

/* That's all, stop editing! Happy blogging. */

/** Absolute path to the WordPress directory. */
if ( !defined('ABSPATH') )
        define('ABSPATH', dirname(__FILE__) . '/');

/** Sets up WordPress vars and included files. */
require_once(ABSPATH . 'wp-settings.php');

# Disables all core updates. Added by SiteGround Autoupdate:
define( 'WP_AUTO_UPDATE_CORE', false );
