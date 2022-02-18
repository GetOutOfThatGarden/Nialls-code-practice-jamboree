// This is some practice of how ternary operators work.
// This is the format:

// example ? then example : then example

// if this is true ? then this is returned : if not, this is returned

//Below I am declaring a constant with the value of 33.
const age = 33;

// Below I am declaring a var variable named result. I am not assigning it a value right now.
// The reason it is a var and not a const is because it's value will change later in the code.
// If it was a const, the code would throw an error when it is ran.
var result;

// Below I am changing the variable I named 'result' into another value.
// I am using the '?' and ':' ternary operators. 
// So, before the '?' I have 'age  < 33'. This asks is the 'age' variable is less than 33. (Similar to an 'if' statement)
// After the '?' and before the ':' I have 'result = "Glory Days"'. This is like the 'then' part of an 'if' statement.
// Lastly I have 'result = "All downhill from here"'. This like the 'else if' part of an 'if' statement.
result = age < 33 ? result = "Glory days" : result = "All downhill from here" ;

console.log(result)

// Heres a better explanation of a conditional (ternary) operator from mozilla.org:
// The conditional (ternary) operator is the only JavaScript operator that 
// takes three operands: a condition followed by a question mark ( ? ), then an 
// expression to execute if the condition is truthy followed by a colon ( : ), and 
// finally the expression to execute if the condition is falsy.