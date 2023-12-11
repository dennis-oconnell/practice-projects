function getCount(str) {
	let vCount = 0;

	for (let i = 0; i < str.length; i++) {
		switch (str.charAt(i)) {
			case "a":
				vCount++;
				break;
			case "e":
				vCount++;
				break;
			case "i":
				vCount++;
				break;
			case "o":
				vCount++;
				break;
			case "u":
				vCount++;
				break;
		}
	}
	return vCount;
}
