const prompt = require('prompt-sync')({sigint: true});

const hat = '^';
const hole = 'O';
const fieldCharacter = '░';
const pathCharacter = '*';

class Field {

  constructor(inputArray) {
    this.field = inputArray;
  }

  printField() {
    for(let i = 0; i < this.field.length;i++)
      console.log(this.field[i]);
  }

  //this function updates the users location based on the direction they input
  updateLocation(direction) {
    let currentRow = 0;
    let currentColumn = 0;
    let currentCharacter = this.field[currentRow][currentColumn];
    let nextCharacter = this.field[currentRow][currentColumn];
    let nextRow = 0;
    let nextColumn = 0;

    //this switch statement updates the users location based on the direction they input
    switch(direction) {
      case 'u':
        nextRow = currentRow - 1;
        nextColumn = currentColumn;
        break;
      case 'd':
        nextRow = currentRow + 1;
        nextColumn = currentColumn;
        break;
      case 'l':
        nextRow = currentRow;
        nextColumn = currentColumn - 1;
        break;
      case 'r':
        nextRow = currentRow;
        nextColumn = currentColumn + 1;
        break;
      default:
        console.log('Please enter a valid direction');
        break;
    }
    //this if statement checks if the users next location is out of bounds
    if(nextRow < 0 || nextRow > this.field.length - 1 || nextColumn < 0 || nextColumn > this.field[0].length - 1) {
      console.log('You went out of bounds! Game Over!');
      return false;
    }
    //this if statement checks if the users next location is a hole
    else if(this.field[nextRow][nextColumn] == hole) {
      console.log('You fell in a hole! Game Over!');
      return false;
    }
    //this if statement checks if the users next location is the hat
    else if(this.field[nextRow][nextColumn] == hat) {
      console.log('You found your hat! You win!');
      return false;
    }
    //this if statement checks if the users next location is a path character
    else if(this.field[nextRow][nextColumn] == pathCharacter) {
      console.log('You already went there! Game Over!');
      return false;
    }

    //this if statement checks if the users next location is a field character
    else if(this.field[nextRow][nextColumn] == fieldCharacter) {
      this.field[currentRow][currentColumn] = pathCharacter;
      this.field[nextRow][nextColumn] = pathCharacter;
      return true;
    }
  }
}

const myField = new Field([
  ['*', '░', 'O'],
  ['░', 'O', '░'],
  ['░', '^', '░'],
  ]);

//this while loop runs until the user finds the hat or falls in a hole
while(true) {
  let direction = prompt('Which direction would you like to go? ');
  if(myField.updateLocation(direction) == false) {
    break;
  }
  myField.printField();
}
