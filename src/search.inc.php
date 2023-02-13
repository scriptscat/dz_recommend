<?php

error_reporting(E_ALL);

ini_set('display_errors', '1');


if (!defined('IN_DISCUZ')) {
    exit('Access Denied');
}

include "vendor/autoload.php";

require "lib/es.php";

$client = Es::getClient();

$resp = $client->get([
    'index' => 'dev.script',
    'id' => 1,
]);

var_dump($resp);
