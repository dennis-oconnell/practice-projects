package main

import "fmt"

//TASK STRUCT
type Task struct {
	Title string
	Description string
	Status string
}	

//TASKMANAGER STRUCT TO HOLD A LIST OF TASKS
type TaskManager struct {
	Tasks []*Task
}

//METHOD TO INITIALIZE NEW TASK
//This method returns a pointer so that we can work with the same instance of the struct across different parts of the program
//This method is handy because if we didn't have it, we'd have to initialize a task in a much more verbose way:
	/*
	myFirstTask := Task{
		Title: "myFirstTitle",
		Description: "myFirstDescription",
		Status: "incomplete",
	 }
	*/
func newTask(title string, description string) *Task {
	
	newt := &Task{
		Title: title,
		Description: description,
		Status: "incomplete",
	}

	return newt
}

//METHOD TO MARK TASK AS COMPLETED
//This is a method with a receiver
//We could do the same thing with a function that receives a pointer and modifies the task like so:
	/*
	func markComplete(task *Task){
		task.Status = "Complete"
	}
	*/
//The only thing that would change would be how we would call the function
	//For the Function with Pointer Parameter version we would do something like: 	//markComplete(myTask)
	//For the Method with a Receiver version we do this instead: 					//myTask.CompleteTask()
func (task *Task) markComplete(){
	task.Status = "Complete"
}

//METHOD TO DISPLAY DETAILS OF TASK
//We also want this method to receive a pointer rather than a value so that we can be consistent with our other methods
func displayTask(task *Task){
	fmt.Printf(" Task Title: %s\n Task Description: %s\n Task Status: %s\n", task.Title, task.Description, task.Status)
}

// Methods for TaskManager to add tasks, mark as completed, and display all tasks

// Implement methods here...

func main() {
	 mySecondTask := newTask("mySecondTask", "This is the second task made by this Go program!")

	 mySecondTask.Title = "Actually this is the title"
	 mySecondTask.markComplete()

	 displayTask(mySecondTask)
}