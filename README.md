bash-jenkins
============

Interact with the [Jenkins](https://jenkins.io/) API from bash.

Copy jobs, create pipelines, and trigger pipeline inputs.

Setup:
------

Set the following environment variables:

* JENKINS_ROOT_URL - The url of the Jenkins server (IE: http://localhost:8080)
* JENKINS_USER_NAME - Jenkins user with API access
* JENKINS_TOKEN - Access token obtained from Jenkins (http://localhost:8080/me/configure)

Usage:
------

```
$ ./bash-jenkins
./bash-jenkins is a tool for interacting with the Jenkins API.

Usage:

	./bash-jenkins command [arguments]

The commands are:

	abort-input     abort a pipeline input
	copy-job	    copy a job from another
	create-job      create a job
	env             print environment information
	trigger-input   trigger a pipeline input

Use "./bash-jenkins help [command]" for more information about a command.
```
