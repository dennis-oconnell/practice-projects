let perseidDate = new Date("Aug 12, 2024 00:20:00");

function updateTimeLeft() {
	let currentDate = new Date();
	let milisecondsLeft = perseidDate - currentDate;
	let daysLeft = Math.floor(milisecondsLeft / (3600 * 1000) / 24);
	let hoursLeft = Math.floor(milisecondsLeft / (3600 * 1000) - daysLeft * 24);
	let minsLeft = Math.floor(milisecondsLeft / (3600 * 1000) - daysLeft * 24);

	document.getElementById("days").innerText = daysLeft;
	document.getElementById("hours").innerText = hoursLeft;
}

setInterval(updateTimeLeft, 1000);
