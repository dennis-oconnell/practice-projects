var stdin = process.stdin;
stdin.setRawMode(true);
stdin.resume();
stdin.setEncoding('utf8');

stdin.on('data', function(key){
    if (key == '\u001B\u005B\u0041') {
        process.stdout.write('up'); 
    }
    if (key == '\u001B\u005B\u0043') {
        process.stdout.write('right'); 
    }
    if (key == '\u001B\u005B\u0042') {
        process.stdout.write('down'); 
    }
    if (key == '\u001B\u005B\u0044') {
        process.stdout.write('left'); 
    }

    if (key == '\u0003') { process.exit(); }    // ctrl-c
});


/*
const prompt = require('prompt-sync')({sigint: true});





// Random number from 1 - 10
const numberToGuess = Math.floor(Math.random() * 10) + 1;
// This variable is used to determine if the app should continue prompting the user for input
let foundCorrectNumber = false;

while (!foundCorrectNumber) {
  // Get user input
  let guess = prompt('Guess a number from 1 to 10: ');
  // Convert the string input to a number
  guess = Number(guess);

  // Compare the guess to the secret answer and let the user know.
  if (guess === numberToGuess) {
    console.log('Congrats, you got it!');
    foundCorrectNumber = true;
  } else {
    console.log('Sorry, guess again!');
  }
}
*/

/*
const name = prompt('What is your name?');
console.log(`Hey there ${name}`);
*/


/*
const readline = require('readline');

readline.emitKeypressEvents(process.stdin);

if (process.stdin.isTTY) {
  process.stdin.setRawMode(true);
}

process.stdin.on('keypress', (str, key) => {
  if (key.name === 'w') {
    console.log('you pressed W');
    exit();
  }
});
*/