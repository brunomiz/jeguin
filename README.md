# Jeguin

Simple webserver which receive Webhooks and execute list of commands

## How to use

Create a file called `jobs.txt` with a content like:

````text
[myjob]
echo 'Hello world'
[myjob2]
echo 'Hello world2'
````

Execute Jeguin providing the port to run. Ex.: jeguin 8080

This example will create two webhooks:
* http://localhost:PORT_NUMBER/myjob
* http://localhost::PORT_NUMBER/myjob2

When a request is received on the `/myjob`, it executes the command `echo 'Hello world'`  
