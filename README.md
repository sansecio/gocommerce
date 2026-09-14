# Golang Commerce Interface

Go config parsers for several PHP eCommerce platforms:

- Magento 1
- Magento 2
- Wordpress/Woocommerce
- Prestashop 1.6 / 1.7+
- Shopware 5 / 6
- OpenCart 4
- JTL-Shop
- Sylius

## Configuration options

Discovery, platform configuration, metadata, and database connection APIs take an explicit final `gocommerce.Options` argument. Use `gocommerce.Options{}` for file-only Symfony configuration with database access enabled:

```go
store := gocommerce.FindStoreAtRoot("/path/to/store", gocommerce.Options{})
```

Shopware and Sylius preserve `.env` interpolation, nested references, defaults, and quoting. They do not read the process environment implicitly. Supply only application variables that the scanned tree is authorized to use:

```go
opts := gocommerce.Options{
    Environment: map[string]string{
        "DB_USER":     os.Getenv("DB_USER"),
        "DB_PASSWORD": os.Getenv("DB_PASSWORD"),
        "APP_ENV":     os.Getenv("APP_ENV"),
    },
    SkipDatabase: true,
}
store := gocommerce.FindStoreAtRoot("/path/to/store", opts)
urls, err := store.Platform.BaseURLs(ctx, store.DocRoot, opts)
```

Pass the same options to subsequent configuration, URL, version, and database operations. The environment map is caller-owned; do not mutate it during parsing or concurrent scans.

Explicitly supplied nonempty values retain precedence over file-local values. Supplied values are literal data: dollar characters are not recursively expanded. An explicitly supplied `DATABASE_URL` overrides the file configuration, even when the file cannot be read; `APP_ENV` controls `%kernel.environment%` substitution.

An allowed value may be sent to a destination selected by the scanned configuration. Do not pass `os.Environ()` wholesale or include unrelated scanner/CI secrets.

`SkipDatabase` makes `ConnectDB` return `ErrDatabaseDisabled` before socket discovery, DNS resolution, or dialing. Database-backed metadata obeys this policy; file-based metadata remains available. It is not a general network sandbox.
