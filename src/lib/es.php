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
                $builder = $builder->setSSLVerification(false)
                    ->setSSLCert($setting["es_pem"]);
            }
            self::$client = $builder->build();
        }
        return self::$client;
    }

    static function insertThread($tid, $fid, $title, $content, $authorid, $createtime)
    {
        global $_G;
        $data = [
            "tid" => $tid,
            "fid" => $fid,
            "title" => $title,
            "content" => $content,
            "authorid" => $authorid,
            "createtime" => $createtime,
        ];
        $resp = self::getClient()->index([
            "index" => "dz.forum_thread",
            "id" => $tid,
            "body" => $data,
        ]);
        return $resp;
    }

    static function updateThread($tid, $fid, $title, $content)
    {
        global $_G;
        $data = [
            "fid" => $fid,
            "title" => $title,
            "content" => $content,
        ];
        $resp = self::getClient()->update([
            "index" => "dz.forum_thread",
            "id" => $tid,
            "body" => [
                "doc" => $data
            ]
        ]);
        return $resp;
    }
}
