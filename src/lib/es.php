<?php

use Elasticsearch\ClientBuilder;

class Es
{
    static $client = null;


    static function getClient()
    {
        if (self::$client == null) {
            global $_G;
            $setting = $_G['cache']['plugin']['codfrm_recommend'];
            $builder = ClientBuilder::create();
            $builder = $builder->setHosts([$setting['es_host']]);
            if ($setting["es_username"]) {
                $builder = $builder->setBasicAuthentication($setting["es_username"], $setting["es_password"]);
            }
            if ($setting["es_ssl"]) {
                $builder = $builder->setSSLVerification($setting["es_ssl"])
                    ->setSSLCert($setting["es_pem"]);
            }
            self::$client = $builder->build();
        }
        return self::$client;
    }
}
