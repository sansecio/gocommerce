<?php
return [
    'downloadable_domains' => [
        'x',
        'y'
    ],
    'db' => [
        'table_prefix' => '',
        'connection' => [
            'default' => [
                'host' => 'myhost',
                'model' => 'mysql4',
                'driver_options' => [
                    1014 => false
                ],
                'active' => '1',
                'username' => 'myuser',
                'engine' => 'innodb',
                'dbname' => 'mydb',
                'password' => 'mypass',
                'initStatements' => 'SET NAMES utf8'
            ]
        ]
    ],
    'queue' => [
        'consumers_wait_for_messages' => 1
    ]
];
