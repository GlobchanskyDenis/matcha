function b64DecodeUnicode(str) {
	// Going backwards: from bytestream, to percent-encoding, to original string.
	return decodeURIComponent(atob(str).split('').map(function (c) {
		return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
	}).join(''));
}

function b64EncodeUnicode(str) {
	// first we use encodeURIComponent to get percent-encoded UTF-8,
	// then we convert the percent encodings into raw bytes which
	// can be fed into btoa.
	return btoa(encodeURIComponent(str).replace(/%([0-9A-F]{2})/g,
		function toSolidBytes(match, p1) {
			return String.fromCharCode('0x' + p1);
		}));
}

function clearUserList() {
	var list = document.getElementById("userList");
	while (list.firstChild) {
		list.removeChild(list.firstChild);
	}
}

function clearPhotoList() {
	var list = document.getElementById("photoList");
	while (list.firstChild) {
		list.removeChild(list.firstChild);
	}
}

function setNewItem(uid, fname, lname, isLiked, isMatch) {
	var list = document.getElementById("userList");
	var newItem = document.createElement("div");
	newItem.setAttribute("class", "item");
	newItem.innerHTML = uid + " : " + fname + " " + lname + " " + isLiked + " " + isMatch;
	list.appendChild(newItem);
}

function setNewPhoto(body) {
	var list = document.getElementById("photoList");
	var newPhoto = document.createElement("img");
	newPhoto.setAttribute("class", "photo");
	newPhoto.src = body
	list.appendChild(newPhoto);
}

function handleFilters() {
	var minAgeStr = document.forms['age']['min'].value
	var maxAgeStr = document.forms['age']['max'].value
	var minRatingStr = document.forms['rating']['min'].value
	var maxRatingStr = document.forms['rating']['max'].value
	var radLatitudeStr =  document.forms['location_radius']['latitude'].value
	var radLongitudeStr = document.forms['location_radius']['longitude'].value
	var radiusStr =       document.forms['location_radius']['radius'].value
	var interestsStr = document.forms['other']['interests'].value
	var filters = new Object();
	// handle age filter
	if (minAgeStr || maxAgeStr) {
		filters.age = {}
		if (minAgeStr) {
			filters.age.min = Number(minAgeStr)
		}
		if (maxAgeStr) {
			filters.age.max = Number(maxAgeStr)
		}
	}
	// handle rating filter
	if (minRatingStr || maxRatingStr) {
		filters.rating = {}
		if (minRatingStr) {
			filters.rating.min = Number(minRatingStr)
		}
		if (maxRatingStr) {
			filters.rating.max = Number(maxRatingStr)
		}
	}
	// handle location radius filter
	if (radLatitudeStr || radLongitudeStr || radiusStr) {
		filters.radius = {}
		if (radLatitudeStr) {
			filters.radius.latitude = Number(radLatitudeStr)
		}
		if (radLongitudeStr) {
			filters.radius.longitude = Number(radLongitudeStr)
		}
		if (radiusStr) {
			filters.radius.radius = Number(radiusStr)
		}
	}
	// handle interests filter
	if (interestsStr != "") {
		// результатом выражения является массив
		filters.interests = interestsStr.split(" ")
	}
	// handle online filter
	if (document.getElementById("online").checked) {
		filters.online = {}
	}
	return filters
}

function getUsers() {
	let xhr = new XMLHttpRequest();
	// var n = document.getElementById("filter1").options.selectedIndex;
	// var filter = document.getElementById("filter1").options[n].value;
	var filters = handleFilters()
	// let tokenObj = new Map()
	// tokenObj["x-auth-token"] = document.token
	// filters.push(tokenObj)
	filters["x-auth-token"] = document.token
	console.log(JSON.stringify(filters))
	// return
	xhr.open("POST", "http://localhost:3000/search/");
	// xhr.responseType = 'json';
	// console.log("tx: get users, filter = " + filter)
	xhr.send(JSON.stringify(filters));
	xhr.onload = function () {
		console.log("rx: " + xhr.status);
		if (xhr.response) {
			// console.log("response:")
			var requestAsync = JSON.parse(xhr.response);
			if (!requestAsync) {
				console.log("no users")
				return
			}
		} else {
			console.log("Empty response")
			return
			// var requestAsync = "";
		}
		if (xhr.status != 200 && xhr.status != 201) {
			document.getElementById("errorField").innerHTML = "Что-то пошло не так: " + xhr.status + " : "
			document.getElementById("errorField").innerHTML += ((requestAsync.error) ? requestAsync.error : xhr.statusText)
			document.getElementById("responseField").innerHTML = "";
			return
		}
		// if (!requestAsync) {
		// 	console.log("Empty response")
		// 	return
		// }
		if (requestAsync.error) {
			document.getElementById("errorField").innerHTML = "Что-то пошло не так: " + xhr.status + " : " + requestAsync.error + " "
		} else {
			document.getElementById("errorField").innerHTML = ""
		}
		clearUserList();
		for (i = 0; requestAsync[i]; i++) {
			setNewItem(requestAsync[i].uid, requestAsync[i].fname, requestAsync[i].lname, requestAsync[i].isLiked, requestAsync[i].isMatch)
		}
	}
}

function wsConnect(xAuthToken) {
	let socket = new WebSocket("ws://localhost:3000/ws/auth/?x-auth-token=" + xAuthToken)
	console.log("attempting websocket connection")
	console.log("ws://localhost:3000/ws/auth/?x-auth-token=" + xAuthToken)
	socket.onopen = function () {
		console.log("web socket successfully connected")
		sendMessage(1, "test message")
	}
	socket.onclose = function (event) {
		console.log("Socket was closed: ", event)
	}
	socket.onerror = function (error) {
		console.log("socket error: ", error)
	}
	socket.onmessage = function (message) {
		console.log("rx: ", message.data)
	}
	function sendMessage(receiver, body) {
		var message = new Object();
		message.type = "message";
		message.uidReceiver = receiver;
		message.body = body;
		jsonMessage = JSON.stringify(message)
		console.log("tx: " + jsonMessage)
		socket.send(jsonMessage)
	}
	// setInterval(sendMessage, 7000)
}

function AuthUser() {
	var mail = document.forms['auth']['mail'].value
	var pass = document.forms['auth']['pass'].value
	var request = JSON.stringify({ "mail": mail, "pass": pass })
	let xhr = new XMLHttpRequest();
	xhr.open("POST", "http://localhost:3000/user/auth/");
	xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
	console.log("tx: " + request)
	xhr.send(request)
	xhr.onload = function () {
		if (xhr.response) {
			var requestAsync = JSON.parse(xhr.response);
			// console.log("rx: " + xhr.status + " : " + requestAsync["x-auth-token"]);
		} else {
			var requestAsync = "";
			// console.log("rx: " + xhr.status);
		}
		console.log("rx: " + xhr.status + " : " + xhr.response);
		if (xhr.status != 200) {
			document.getElementById("errorField").innerHTML = "Что-то пошло не так: " + xhr.status + " : "
			document.getElementById("errorField").innerHTML += ((requestAsync.error) ? requestAsync.error : xhr.statusText)
			document.getElementById("responseField").innerHTML = "";
			return
		}
		if (!requestAsync) {
			document.getElementById("errorField").innerHTML = "Empty request. Its not a valid case";
			document.getElementById("responseField").innerHTML = "";
			return
		}
		document.getElementById("errorField").innerHTML = ""
		if (!requestAsync.uid || !requestAsync.mail || !requestAsync["x-auth-token"]) {
			document.getElementById("errorField").innerHTML += "uid or token are empty. Its not a valid case";
			return
		}
		document.getElementById("responseField").innerHTML = "uid=" + requestAsync.uid +
			" mail=" + requestAsync.mail +
			" x-auth-token=hidden ws-auth-token=hidden"
		document.forms['auth']['mail'].value = "";
		document.forms['auth']['pass'].value = "";
		document.token = requestAsync["x-auth-token"]
		wsConnect(requestAsync["x-auth-token"])
	}
	xhr.onerror = function () {
		console.log("onError event")
	}
}

function RegUser() {
	var mail = document.forms['reg']['mail'].value
	var pass = document.forms['reg']['pass'].value
	var request = JSON.stringify({ "mail": mail, "pass": pass })
	let xhr = new XMLHttpRequest();
	xhr.open("PUT", "http://localhost:3000/user/create/");
	console.log("tx: " + request)
	xhr.send(request);
	xhr.onload = function () {
		if (xhr.response) {
			var requestAsync = JSON.parse(xhr.response);
		} else {
			var requestAsync = "";
		}
		console.log("rx: " + xhr.status + " : " + xhr.response);
		if (xhr.status != 201) {
			document.getElementById("errorField").innerHTML = "Что-то пошло не так: " + xhr.status + " : "
			document.getElementById("errorField").innerHTML += ((requestAsync.error) ? requestAsync.error : xhr.statusText)
			document.getElementById("responseField").innerHTML = "";
			return
		}
		document.getElementById("errorField").innerHTML = ""
		document.getElementById("responseField").innerHTML = "registration was done. Check your email"
		document.forms['reg']['mail'].value = "";
		document.forms['reg']['pass'].value = "";
	}
	xhr.onerror = function () {
		console.log("onError event")
	}
}

function UpdateUserStatus() {
	var token = document.forms['updStatus']['x-reg-token'].value
	var request = `{"x-reg-token":"` + token + `"}`
	let xhr = new XMLHttpRequest();
	xhr.open("PATCH", "http://localhost:3000/user/update/status/");
	xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
	console.log("tx: " + request)
	xhr.send(request);
	xhr.onload = function () {
		if (xhr.response) {
			var requestAsync = JSON.parse(xhr.response);
		} else {
			var requestAsync = "";
		}
		console.log("rx: " + xhr.status + " : " + xhr.response);
		if (xhr.status != 200) {
			document.getElementById("errorField").innerHTML = "Что-то пошло не так: " + xhr.status + " : "
			document.getElementById("errorField").innerHTML += ((requestAsync.error) ? requestAsync.error : xhr.statusText)
			document.getElementById("responseField").innerHTML = "";
			return
		}
		document.getElementById("errorField").innerHTML = ""
		document.getElementById("responseField").innerHTML = "Update was done successfully"
		document.forms['updStatus']['x-reg-token'].value = "";
	}
	xhr.onerror = function () {
		console.log("onError event")
	}
}

function readURL(input) {
	if (input.files && input.files[0]) {
		var reader = new FileReader();
		image = document.getElementById('upload_photo');
		image.style.display = "block";
		// document.getElementById('video').style.opacity = 0;
		reader.onload = function (e) {
			image.setAttribute('src', e.target.result);
		};
		reader.readAsDataURL(input.files[0]);
	}
	image_statut = true;
}

function Upload() {
	var image = document.getElementById('upload_photo');
	var request = `{"x-auth-token":"` + document.token + `","src":"` + image.src + `"}` //btoa(image.src)
	document.buffer = image.src
	let xhr = new XMLHttpRequest();
	xhr.open("PUT", "http://localhost:3000/photo/upload/");
	xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
	console.log("upload tx (not display img source)")
	xhr.send(request);
	xhr.onload = function () {
		if (xhr.response) {
			var requestAsync = JSON.parse(xhr.response);
		} else {
			var requestAsync = "";
		}
		console.log("upload rx: " + xhr.status + " : " + xhr.response);
		if (xhr.status != 200) {
			document.getElementById("errorField").innerHTML = "Что-то пошло не так: " + xhr.status + " : "
			document.getElementById("errorField").innerHTML += ((requestAsync.error) ? requestAsync.error : xhr.statusText)
			document.getElementById("responseField").innerHTML = "";
			return
		}
		document.getElementById("errorField").innerHTML = ""
		document.getElementById("responseField").innerHTML = "Upload was done successfully"
	}
	xhr.onerror = function () {
		console.log("onError event")
	}
}

function Download() {
	var uid = document.forms['dload']['uid'].value
	var request = `{"x-auth-token":"` + document.token + `","uid":` + uid + `}`
	let xhr = new XMLHttpRequest();
	xhr.open("POST", "http://localhost:3000/photo/download/");
	xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
	console.log("download tx: " + request)
	xhr.send(request);
	xhr.onload = function () {
		if (xhr.response) {
			var responseAsync = JSON.parse(xhr.response);
		} else {
			var responseAsync = "";
		}
		if (xhr.status != 200) {
			document.getElementById("errorField").innerHTML = "Что-то пошло не так: " + xhr.status + " : "
			document.getElementById("errorField").innerHTML += ((responseAsync.error) ? responseAsync.error : xhr.statusText)
			document.getElementById("responseField").innerHTML = "";
			return
		}
		if (responseAsync) {//atob()
			clearPhotoList()
			for (i = 0; responseAsync[i]; i++) {
				console.log("download rx: pid=" + JSON.stringify(responseAsync[i].pid) + " uid=" + JSON.stringify(responseAsync[i].uid))
				setNewPhoto(responseAsync[i].src)
			}
		}
	}
	xhr.onerror = function () {
		console.log("onError event")
	}
}

function UpdateUser() {
	var mail = document.forms['update_1']['mail'].value
	var pass = document.forms['update_1']['pass'].value
	var fname = document.forms['update_1']['fname'].value
	var lname = document.forms['update_1']['lname'].value
	var birth = document.forms['update_2']['birth'].value
	var gender = document.forms['update_2']['gender'].value
	var orientation = document.forms['update_2']['orientation'].value
	var bio = document.forms['update_2']['bio'].value
	var avaID = document.forms['update_3']['avaID'].value
	var latitude = document.forms['update_3']['latitude'].value
	var longitude = document.forms['update_3']['longitude'].value
	var interestsString = document.forms['update_3']['interests'].value
	var interestsArr = interestsString.split(" ")
	var request = ""
	if (mail == "" && pass == "" && fname == "" && lname == "" &&
		birth == "" && gender == "" && orientation == "" && bio == "" &&
		avaID == "" && latitude == "" && longitude == "" && interestsString == "") {
		console.log("update: all fields are empty")
		return
	}
	if (mail != "") {
		request = "\"mail\":" + JSON.stringify(mail)
	}
	if (pass != "") {
		if (request != "") {
			request += ","
		}
		request += "\"pass\":" + JSON.stringify(pass)
	}
	if (fname != "") {
		if (request != "") {
			request += ","
		}
		request += "\"fname\":" + JSON.stringify(fname)
	}
	if (lname != "") {
		if (request != "") {
			request += ","
		}
		request += "\"lname\":" + JSON.stringify(lname)
	}
	if (birth != "") {
		if (request != "") {
			request += ","
		}
		request += "\"birth\":" + JSON.stringify(birth)
	}
	if (gender != "") {
		if (request != "") {
			request += ","
		}
		request += "\"gender\":" + JSON.stringify(gender)
	}
	if (orientation != "") {
		if (request != "") {
			request += ","
		}
		request += "\"orientation\":" + JSON.stringify(orientation)
	}
	if (bio != "") {
		if (request != "") {
			request += ","
		}
		request += "\"bio\":" + JSON.stringify(bio)
	}
	if (avaID != "") {
		if (request != "") {
			request += ","
		}
		var validAvaID = Number(avaID)
		request += "\"avaID\":" + JSON.stringify(validAvaID)
	}
	if (latitude != "") {
		if (request != "") {
			request += ","
		}
		request += "\"latitude\":" + latitude
	}
	if (longitude != "") {
		if (request != "") {
			request += ","
		}
		request += "\"longitude\":" + longitude
	}
	if (interestsString != "") {
		if (request != "") {
			request += ","
		}
		request += "\"interests\":" + JSON.stringify(interestsArr)
	}
	if (request != "") {
		request += ","
	}
	request += `"x-auth-token":"` + document.token + `"`
	request = "{" + request + "}"
	let xhr = new XMLHttpRequest();
	xhr.open("PATCH", "http://localhost:3000/user/update/");
	xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
	console.log("tx: " + request)
	xhr.send(request);
	xhr.onload = function () {
		if (xhr.response) {
			var requestAsync = JSON.parse(xhr.response);
		} else {
			var requestAsync = "";
		}
		console.log("rx: " + xhr.status + " : " + xhr.response);
		if (xhr.status != 200) {
			document.getElementById("errorField").innerHTML = "Что-то пошло не так: " + xhr.status + " : "
			document.getElementById("errorField").innerHTML += ((requestAsync.error) ? requestAsync.error : xhr.statusText)
			document.getElementById("responseField").innerHTML = "";
			return
		}
		document.getElementById("errorField").innerHTML = ""
		document.getElementById("responseField").innerHTML = "Update was done successfully"
		document.forms['update_1']['mail'].value = "";
		document.forms['update_1']['pass'].value = "";
		document.forms['update_1']['fname'].value = "";
		document.forms['update_1']['lname'].value = "";
		document.forms['update_2']['birth'].value = "";
		document.forms['update_2']['gender'].value = "";
		document.forms['update_2']['orientation'].value = "";
		document.forms['update_2']['bio'].value = "";
		document.forms['update_3']['avaID'].value = "";
		document.forms['update_3']['latitude'].value = "";
		document.forms['update_3']['longitude'].value = "";
		document.forms['update_3']['interests'].value = "";
	}
	xhr.onerror = function () {
		console.log("onError event")
	}
}

function SetClaim() {
	var uid = document.forms['claim']['uid'].value
	var request = {
		"x-auth-token": document.token,
		"otherUid":	Number(uid),
	}
	console.log(JSON.stringify(request))

	let xhr = new XMLHttpRequest();
	xhr.open("PUT", "http://localhost:3000/claim/set/");
	xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
	xhr.send(JSON.stringify(request));

	xhr.onload = function () {
		if (xhr.response) {
			var responseAsync = JSON.parse(xhr.response);
		} else {
			var responseAsync = "";
		}
		if (xhr.status != 200) {
			document.getElementById("errorField").innerHTML = "Что-то пошло не так: " + xhr.status + " : "
			document.getElementById("errorField").innerHTML += ((responseAsync.error) ? responseAsync.error : xhr.statusText)
			document.getElementById("responseField").innerHTML = "";
			console.log(responseAsync)
			return
		} else {
			document.getElementById("errorField").innerHTML = ""
			document.getElementById("responseField").innerHTML = "Success";
		}
	}
	xhr.onerror = function () {
		console.log("onError event")
	}
}

function UnsetClaim() {
	var uid = document.forms['claim']['uid'].value
	console.log(uid)
	var request = {
		"x-auth-token": document.token,
		"otherUid":	Number(uid),
	}
	console.log(JSON.stringify(request))

	let xhr = new XMLHttpRequest();
	xhr.open("DELETE", "http://localhost:3000/claim/unset/");
	xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
	xhr.send(JSON.stringify(request));

	xhr.onload = function () {
		if (xhr.response) {
			var responseAsync = JSON.parse(xhr.response);
		} else {
			var responseAsync = "";
		}
		if (xhr.status != 200) {
			document.getElementById("errorField").innerHTML = "Что-то пошло не так: " + xhr.status + " : "
			document.getElementById("errorField").innerHTML += ((responseAsync.error) ? responseAsync.error : xhr.statusText)
			document.getElementById("responseField").innerHTML = "";
			console.log(responseAsync)
			return
		} else {
			document.getElementById("errorField").innerHTML = ""
			document.getElementById("responseField").innerHTML = "Success";
		}
	}
	xhr.onerror = function () {
		console.log("onError event")
	}
}

function SetIgnore() {
	var uid = document.forms['ignore']['uid'].value
	var request = {
		"x-auth-token": document.token,
		"otherUid":	Number(uid),
	}
	console.log(JSON.stringify(request))

	let xhr = new XMLHttpRequest();
	xhr.open("PUT", "http://localhost:3000/ignore/set/");
	xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
	xhr.send(JSON.stringify(request));

	xhr.onload = function () {
		if (xhr.response) {
			var responseAsync = JSON.parse(xhr.response);
		} else {
			var responseAsync = "";
		}
		if (xhr.status != 200) {
			document.getElementById("errorField").innerHTML = "Что-то пошло не так: " + xhr.status + " : "
			document.getElementById("errorField").innerHTML += ((responseAsync.error) ? responseAsync.error : xhr.statusText)
			document.getElementById("responseField").innerHTML = "";
			console.log(responseAsync)
			return
		} else {
			document.getElementById("errorField").innerHTML = ""
			document.getElementById("responseField").innerHTML = "Success";
		}
	}
	xhr.onerror = function () {
		console.log("onError event")
	}
}

function UnsetIgnore() {
	var uid = document.forms['ignore']['uid'].value
	var request = {
		"x-auth-token": document.token,
		"otherUid":	Number(uid),
	}
	console.log(JSON.stringify(request))

	let xhr = new XMLHttpRequest();
	xhr.open("DELETE", "http://localhost:3000/ignore/unset/");
	xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
	xhr.send(JSON.stringify(request));

	xhr.onload = function () {
		if (xhr.response) {
			var responseAsync = JSON.parse(xhr.response);
		} else {
			var responseAsync = "";
		}
		if (xhr.status != 200) {
			document.getElementById("errorField").innerHTML = "Что-то пошло не так: " + xhr.status + " : "
			document.getElementById("errorField").innerHTML += ((responseAsync.error) ? responseAsync.error : xhr.statusText)
			document.getElementById("responseField").innerHTML = "";
			console.log(responseAsync)
			return
		} else {
			document.getElementById("errorField").innerHTML = ""
			document.getElementById("responseField").innerHTML = "Success";
		}
	}
	xhr.onerror = function () {
		console.log("onError event")
	}
}

document.token = "rDRrPM2M3WTtXGLdWaozGw8EOPvdX_PHr0u2qa8="
var buffer = ""

window.onload = function(){
	console.log("try to reconnect socket")
	wsConnect(document.token)
}