## About the project

Hermes is a project I created to relax and spend cozy time developing software with new ideas while also learning new things. We always learn by trying to develop software, debugging, and mixing ideas, which is the main reason I chose GoLang for this project.

#### Scopes
Main: Not implemented yet.
Computers:
* Table - In the computer section, you can inspect your integrated computers, viewing details such as hostname, IP address, country, username, operating system, and a redirection link to an interaction console.
* Add - The ADD button in the computers menu allows you to add a new computer to the structure. All the data is saved in a JSON object, as the idea is to enable quick setup interaction with the machines. It would be a good approach to implement cryptography in the file and authorization for each access to the instances. At the moment, the project is just a sketch.
* Interact - The interaction feature is available after selecting the console for a specific machine. In the console, you can send commands to the target, choosing the communication protocol (currently supported: DNS, HTTP, WebSocket).

Receive: The receiver routes handle receiving data from the target and displaying it in the selected console.
## Error Pattern

Every error message in the project follows the specified documentation. First, the messages need to be explicit to identify where the error occurred. The construction is based on the module name **[connection]**, followed by the internal function **[sendMessage]**, and a type of message, as described below:
* **[ERROR]** - Error in execution.
* **[ALERT]** - Alert some behavior outside of the default.
* **[WARNING]** - Some action that can brake the running process.
* **[SUCCESS]** - Concluded action.


## Considerations of Implementations

#### Regex
Go uses [RE2](https://github.com/google/re2/wiki/Syntax) syntax,  there is a related article about automata and regular expressions in the context of regexp. Ref: https://swtch.com/~rsc/regexp/regexp1.html. 
This article focuses on describing the workings of faster regular expressions, explaining the behaviors of DFA and NFA, and comparing the default implementations in different programming languages and technologies.
The most notorious message of the article for me is: "Today, regular expressions have also become a shining example of how ignoring good theory leads to bad programs. The regular expression implementations used by today's popular tools are significantly slower than the ones used in many of those thirty-year-old Unix tools."