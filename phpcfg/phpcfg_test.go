package phpcfg

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestParseShortArray(t *testing.T) {
	src := `
<?php
return [
    'mylist' => [
        'entry0',
        'entry1'
    ],
    'db' => [
        'prefix' => '',
        'con' => [
            'host' => 'xx',
            'driver_options' => [
                1014 => false
            ]
        ]
    ],
    'queue' => [
        'enabled' => 1
    ]
];`

	expected := map[string]string{
		"root.mylist.0":                   "entry0",
		"root.mylist.1":                   "entry1",
		"root.db.prefix":                  "",
		"root.db.con.host":                "xx",
		"root.db.con.driver_options.1014": "false",
		"root.queue.enabled":              "1",
	}

	parsed, err := Parse([]byte(src))
	if err != nil {
		t.Error("Parse failed:", err)
	} else if !cmp.Equal(parsed, expected) {
		t.Error("Parsed != expected")
	}
}

func TestParseArray(t *testing.T) {
	src := `
<?php
return array(
  'install' =>
  array(
    'date' => 'Tue, 08 Aug 2017 20:08:01 +0000',
  ),
);
`
	expected := map[string]string{
		"root.install.date": "Tue, 08 Aug 2017 20:08:01 +0000",
	}

	parsed, err := Parse([]byte(src))
	if err != nil {
		t.Error("Parse failed:", err)
	} else if !cmp.Equal(parsed, expected) {
		t.Error("Parsed != expected")
	}
}
