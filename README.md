# Spark VM 0.31
SparkVM is a project made for running easily readable bytecode ina virtual environment like that of **JVM** or **Common Language Runtime**. This is the original implementation of the SparkVM written fully in **Golang**

## Synatx
The SparkVM follows the instruction operand structure where the first word of any line is an intruction and there are single or multiple operands afterwards. As of the version 0.1 there are 10 instructions and more are add continously with each update listed below:

**Note: Spark VM's Bytecode language is case sensative and all the instructions start with a lowercase letter. Here the first letter is uppercased for aesthtic reseaons**
* Set
* Show
* Add
* Sub
* Mult
* Div
* Jump
* Fun
* End
* Deref
* Setptr
* Cmp
* Eql
* Import **(Work in progress)**

## Features
Spark VM contains variety of features such as:
### Dynamic "Stack"
Stack is not a dynamic area on other languages or VM's but here it is allowing you to define dynamic variables without worrying about Memory Management as it is managed by the stack like structure of the memory
### Global Variables
The bytecode language supports global variables allowing you to make apps without strictly following the functional patterns or worrying about heap and its pointers
**NOTE: Global variables are not recommended for huge data sizes and complex structures since they don't get Garbage collected till the end of the program. and can lead to excess memory usage. Use heap and global/local pointers for that**
### Pointers
Talking about pointers. If you really want to use them for anything they are present and can be used with following instructions
* Deref: It takes two operands, first operand must be a pointer. It saves the data the first pointer is adressing to the second operand which must be variable
* Setptr: It also takes two operands. and saves the memory adress of the first variable to the second variable. **NOTE: It can only get memory adress of variables local to its function's scope and global variables**
### Functions
Spark VM and it's bytecode natively support Functions so you don't have to use weird hacks. The fuctions here are declared with *fun* instruction in the syntax *fun {functionName}({function arguements})* and ended with *end* instruction functions are treated as variables meaning they can be passed as variables or can be pointed to by pointers
## Documentation
Check our [Docs](./docs) for that
## Licensing
The project is licensed under GNU GPL 3.0 License. For more info checks the [LICENSE](./LICENSE)