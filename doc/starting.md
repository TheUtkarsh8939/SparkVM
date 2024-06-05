# First Program In Spark VM
***This tutorial expects you to have a basic knowledge of programming***

## Setting up the environment
To make and run your first program, you need to have the Spark VM Runtime installed. To do this clone this github repositry. Once you have cloned it go to the bin folder. Now you can save folder's location to the 'PATH' environment variable or you can copy paste the executable file to your project directory. For simplicity's sake we will use the latter approach.

Now we can run the bytecode using the command <br> <code>./sparkvm.exe run [your project file name]</code>

## Writing the Program
Create a file in the directory you just pasted the binary into. Name the file whatever you want. For this tutorial we will use <code>main.sp</code>. In the file. Copy and paste the following snippet of code
>fun main() <br> halt <br> end 

Between <code>fun main()</code> and <code>halt</code> you will write your code

### Explaination

Let's see what the code snippet does

<code>fun main()</code> makes a new function called main which is the starting point for our program

<code>halt</code> Stops tells the Spark VM interpreter to stop when reached this instruction. This essentialy stops the interpreter from endless loops trying to execute empty lines

<code>end</code> tells the program that the function has ended here and the lines after that are not part of the main function

### Writing a Hello World

 Let's write a hello world
>fun main()<br>set %x 'Hello World' <br> show %x<br> halt <br> end 

The given snippet of code prints hello world. Let's understand how

### Explaination

We have discussed about line no. 1 and line no. 4 and line no. 5 so let's understand line no. 2 and 3

<code>set %x 'Hello World'</code> Creates a variable <code>%x</code> with a type of string. Refer to [Variables](./variables.md) for more info

<code>show %x</code> Prints the variable <code>%x</code>

