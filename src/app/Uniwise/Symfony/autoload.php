<?php
use Doctrine\Common\Annotations\AnnotationRegistry;

$loader = require( __DIR__ . '/../../ext/vendor/autoload.php' );
AnnotationRegistry::registerLoader(array($loader, 'loadClass'));
