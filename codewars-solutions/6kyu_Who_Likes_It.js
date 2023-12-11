function likes(names) {
	// no one likes this
	if (names.length === 0) {
		let result = "no one likes this";
		return result;
	}

	//one person & two people like this
	if (names.length > 0 && names.length < 3) {
		//one person likes this
		let like = " likes this";
		let people = names[0];

		//two people like this
		for (let i = 1; i < names.length; i++) {
			like = " like this";
			if (i === names.length - 1) {
				people += " and ";
			}
			people += names[i];
		}
		let result = people + like;
		return result;
	}

	// 3 people like this
	if (names.length === 3) {
		let like = " like this";
		let people = names[0];

		for (let i = 1; i < names.length; i++) {
			if (i === names.length - 1) {
				people += " and ";
			}

			if (i === 1) {
				people += ",";
				people = people.concat(" ");
			}

			people += names[i];
		}
		let result = people + like;
		return result;
	}

	//4 people like this
	if (names.length > 3) {
		let like = " like this";
		let people = names[0] + "," + " " + names[1];
		let others = names.length - 2;

		people += " and " + others + " others";

		let result = people + like;
		return result;
	}
}
