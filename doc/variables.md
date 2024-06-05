# Variables in Spark VM Bytecode
Varibles in Spark VM Bytecode language is defined using <code>set</code> keyword. There are 2 types of variables in the Spark VM: 
* Global
* Local

Local variables are limited to thier function while global variables can be read and modified from anywhere in the program

## Syntax
Variables in Spark VM Bytecode are defined using the syntax <br><code>set [Variable Name] [Variable Value]</code>

For local variables the name must be prefixed with an '%' and for global variables it must be prefixed with '$' 

The variable value can be an Immediate, string or <br> a boolean **(Coming Soon)**

## Usage
Variables can be used for instructions that use them like <code>add</code> or <code>show</code>
but they must be prefied with thier respective prefixes so '%' for local and '$' for global

