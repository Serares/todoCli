### cli todo app

### Usage

 After the compilation

You can run the program with -h flag to get the available arguments

- run the programm with one of the following arguments:

- -list --> to list the current todos
- -complete (number) --> to complete a task
- -task "Name of the task" --> to create a new todo
- -filename --> Point to a different filename from the prompt
by default the program will look for TODO_FILENAME env variable and use that value to declare the tasks filename. The fallback case is to use '.todo.json' as a filename