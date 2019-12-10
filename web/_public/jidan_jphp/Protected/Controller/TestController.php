<?php

namespace App\Controller;

use Jphp\Application\WebApplication;
use Jphp\Controller\HttpController;
use Jphp\Http\Request;
use Jphp\Http\Response;
use Jphp\Template\Template;

/**
 * Class TestController
 * @package App\Controller
 */
class TestController extends HttpController {

	/**
	 * @methods=["GET","POST"];
	 * @param Request $request
	 * @return Response
	 */
	public function indexAction(Request $request) {

		$app = WebApplication::getApp();
		$protectedPath = $app->getProtectedPath();
		$test_tpl = $protectedPath . DIRECTORY_SEPARATOR . "Dist" . DIRECTORY_SEPARATOR . "index.html";
		$data = array("a" => 1, "b" => 2);
		$template = new Template();
		$content = $template->parse($test_tpl, $data);

		return new Response($content);
	}


}
