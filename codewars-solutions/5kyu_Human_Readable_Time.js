function humanReadable(seconds) {
	let s = seconds;

	let HH = Math.floor(s / 3600);
	s = s % 3600;

	let MM = Math.floor(s / 60);
	s = s % 60;

	let SS = s;

	function addZero(num) {
		if (Math.floor(num / 10) === 0) {
			return "0" + num;
		} else {
			return "" + num;
		}
	}

	return "" + addZero(HH) + ":" + addZero(MM) + ":" + addZero(SS) + "";
}
