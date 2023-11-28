let perseidDate = new Date("Aug 12, 2024 00:00:00");

function updateTimeLeft() {
	let currentDate = new Date();
	let milisecondsLeft = perseidDate - currentDate;
	let daysLeft = Math.floor(milisecondsLeft / (3600 * 1000) / 24);
	let hoursLeft = Math.floor(milisecondsLeft / (3600 * 1000) - daysLeft * 24);
	let minsLeft = Math.floor((milisecondsLeft % (3600 * 1000)) / 60000);
	let secsLeft = Math.floor((milisecondsLeft / 1000) % 60);

	document.getElementById("days").innerText = daysLeft;
	document.getElementById("hours").innerText = hoursLeft;
	document.getElementById("mins").innerText = minsLeft;
	document.getElementById("seconds").innerText = secsLeft;
}

setInterval(updateTimeLeft, 1000);

// miliseconds left / (3600 * 1000) = hours left
// 1 second for every 1000 miliseconds
//
