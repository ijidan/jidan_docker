<?php

use Lib\Util\SignalDemo;
use phpspider\core\phpspider;

error_reporting(E_ALL);
defined("DS") or define("DS", DIRECTORY_SEPARATOR);

$protectedPath = dirname(__DIR__) . DS . "Protected" . DS;
require_once $protectedPath . "Jphp" . DS . "Bootstrap.php";
require_once $protectedPath . "vendor" . DS . "autoload.php";
require_once $protectedPath . "library" . DS . "phpspider" . DS . "autoloader.php";


/* Do NOT delete this comment */
/* 不要删除这段注释 */
$configs = array(
	'name'                => '糗事百科',
	'domains'             => array(
		'qiushibaike.com',
		'www.qiushibaike.com'
	),
	'scan_urls'           => array(
		'http://www.qiushibaike.com/'
	),
	'content_url_regexes' => array(
		"http://www.qiushibaike.com/article/\d+"
	),
	'list_url_regexes'    => array(
		"http://www.qiushibaike.com/8hr/page/\d+\?s=\d+"
	),
	'fields'              => array(
		array(
			// 抽取内容页的文章内容
			'name'     => "article_content",
			'selector' => "//*[@id='single-next-link']",
			'required' => true
		),
		array(
			// 抽取内容页的文章作者
			'name'     => "article_author",
			'selector' => "//div[contains(@class,'author')]//h2",
			'required' => true
		),
	),
);
$spider = new phpspider($configs);
$spider->start();
