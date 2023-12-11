function addBinary(a, b) {
	let sum = a + b;
	let ans = "";

	while (sum >= 1) {
		if (sum % 2 === 1) {
			ans = "1" + ans;
		} else {
			ans = "0" + ans;
		}

		sum = Math.floor(sum / 2);
	}

	return ans;
}
