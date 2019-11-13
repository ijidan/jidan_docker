import React, {Component} from 'react'
import config from './config.json';


module.exports = function() {
	var greet = document.createElement('div');
	greet.textContent = config.greetText;
	return greet;
};