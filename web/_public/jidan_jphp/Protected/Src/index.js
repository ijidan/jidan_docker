import {cube} from "./math";

function component() {

	var element = document.createElement("pre");
	element.innerHTML = [
		"hello webpack",
		"5 cubed is equal to " + cube(5)
	].join("\r\n");
	return element;
}

if(process.env.NODE_ENV !== 'production'){
	console.error('Looks like we are in Production mode!');
}

let element = component();
document.body.appendChild(element);