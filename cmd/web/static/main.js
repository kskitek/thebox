const digits = document.getElementsByClassName("digit");
for(var i = 0; i < digits.length; i++) {
	digits[i].onkeypress = validateNumber;
}

function validateNumber({target}) {
	setValid(target);
	return target.value < 1;
}

const button = document.getElementById("send");
button.onclick = validate;

function validate() {
	var code = "";
  var isInvalid = false;
	for(var i = 0; i < digits.length; i++) {
		const digit = digits[i].value;
		if (digit.length !== 1) {
			setInvalid(digits[i]);
      isInvalid = true;
		} else {
      code += digit;
    }
	}

  if (isInvalid) {
    return;
  }
  sendCode(code);
	console.log(code);
}

function setInvalid(digit) {
	digit.classList.add("invalid");
}

function setValid(digit) {
	digit.classList.remove("invalid");
}

function sendCode(code) {
  showError(false);
  const resp = fetch("https://thebox.skitek.dev/test", {
    method: "POST",
    mode: "cors",
    cache: "no-cache",
    headers: {
      "Content-Type": "text/plain",
      "magic-cookie": getCookie(),
    },
    body: code
  });

  resp
    .then(r => r.json())
    .then(d => console.log(d))
    .catch(err => {
      console.error(err);
      showError(true);
    }
    );
}

function getCookie() {
  return window.location.search.split("=")[1];
}

function showError(shouldShow) {
  var err = document.getElementById("error");
  console.log(shouldShow);
  if (shouldShow) {
    err.style.display = "block";
  } else {
    err.style.display = "none";
  }
}
