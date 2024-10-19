## ERROR PATTERN DEFINITION

To every error message in the project, it need to follow the specified documentation. At first, the messages needs to be explicity to known where the error was rised, every module must include global variables that references the module name, is the first pattern [connection] followed by the internal function [sendMessage] and a type of message as described below:
* [ERROR] - to error in operations
* [ALERT] - alert some behavior outside of the default.
* [WARNING] - to some action that can brake the running process.
* [SUCCESS] - action done to the end.
