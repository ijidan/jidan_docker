<?php

namespace App\Controller;

use Jphp\Application\WebApplication;
use Jphp\Controller\HttpController;
use Jphp\Http\Request;
use Jphp\Http\Response;
use Jphp\Template\Template;


/**
 * 爬虫
 * Class SpiderController
 * @package App\Controller
 */
class SpiderController extends HttpController {

	/**
	 * @methods=["GET","POST"];
	 * @param Request $request
	 * @return Response
	 */
	public function indexAction(Request $request) {
		$content="我的爬虫测试...";
		return new Response($content);
	}
	

}
