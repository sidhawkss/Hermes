## ERROR PATTERN DEFINITION

To every error message in the project, it need to follow the specified documentation. At first, the messages needs to be explicity to known where the error was rised, every module must include global variables that references the module name, is the first pattern [connection] followed by the internal function [sendMessage] and a type of message as described below:
* [ERROR] - to error in operations
* [ALERT] - alert some behavior outside of the default.
* [WARNING] - to some action that can brake the running process.
* [SUCCESS] - action done to the end.



# REGEX

Go uses RE2 syntax: https://github.com/google/re2/wiki/Syntax. 
There is a related article about automata and regular expressions: https://swtch.com/~rsc/regexp/regexp1.html. 
This article focuses on describing the workings of faster regular expressions, explaining the behaviors of DFA and NFA, and comparing the default implementations in different programming languages and technologies.
The most notorious message of the article for me is: "Today, regular expressions have also become a shining example of how ignoring good theory leads to bad programs. The regular expression implementations used by today's popular tools are significantly slower than the ones used in many of those thirty-year-old Unix tools."