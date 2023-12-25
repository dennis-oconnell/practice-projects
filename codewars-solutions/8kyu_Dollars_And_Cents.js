function formatMoney(amount) {
	const numDigits = Math.round(amount).toString().length;
	return "$" + amount.toPrecision(2 + numDigits);
}
