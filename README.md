# Snap collector plugin - Snaptel

This plugin queries the snapteld rest api and collects metrics on the tasks currently running.

It's used in the [Snap framework](http://github.com/intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license-and-authors)
6. [Acknowledgements](#acknowledgements)

## Getting Started

### Operating systems
* Linux/amd64
* Darwin/amd64

### Installation
#### Download the plugin binary:

You can get the pre-built binaries for your OS and architecture from the plugin's [GitHub Releases](https://github.com/Staples-Inc/snap-plugin-collector-snaptel/releases) page. Download the plugin from the latest release and load it into `snapteld` (`/opt/snap/plugins` is the default location for Snap packages).

#### To build the plugin binary:
Fork https://github.com/Staples-Inc/snap-plugin-collector-snaptel
Clone repo into `$GOPATH/src/github.com/Staples-Inc/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-snaptel.git
```

Build the Snap snaptel plugin by running make within the cloned repo:
```
$ make
```
It may take a while to pull dependencies if you haven't had them already.
This builds the plugin in `./build/`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Load the plugin and create a task, see example in [Examples](#examples).

#### Configuration parameters
| Variable   | Description                                | Default          |
|------------|--------------------------------------------|------------------|
| `address`  | Address and port of the snapteld api       | `localhost:8181` |
| `api_key`  | Rest API key for api protected snap agents | None             |
| `https`    | Boolean to request over https              | `False`          |
| `insecure` | Boolean to allow requests to unsigned certs| `False`          |


### Collected Metrics

The list of collected metrics is described in [METRICS.md](METRICS.md).

### Example
In one terminal window, start the Snap daemon (in this case with logging set to 1 and trust disabled):
```
$ snapteld -l 1 -t 0
```

In another terminal window download and load plugins:
```
$ snaptel plugin load snap-plugin-collector-snaptel
```

You can list all of available metrics:
```
$ snaptel metric list
```

Download an [example task file](examples/tasks/snaptel.json) and load it:
```
$ snaptel task create -t snaptel.json
Using task manifest to create task
Task created
ID: 96c248ba-b572-4367-809f-943f5e0cf786
Name: Task-96c248ba-b572-4367-809f-943f5e0cf786
State: Running
```

See output from snaptel task watch <task_id>

(notice, that below only the fragment of task watcher output has been presented)

```
$ snaptel task watch  96c248ba-b572-4367-809f-943f5e0cf786
Watching Task (96c248ba-b572-4367-809f-943f5e0cf786):
NAMESPACE 									                     DATA 		 TIMESTAMP
/staples/snaptel/task/example_task1/fail_count 	 0 		     2016-12-30 13:52:36.349851013 -0500 EST
/staples/snaptel/task/example_task1/hit_count    62 		   2016-12-30 13:52:36.349853713 -0500 EST
/staples/snaptel/task/example_task1/state 			 Running 	 2016-12-30 13:52:36.349841044 -0500 EST
/staples/snaptel/task/example_task1/statecode 	1 		     2016-12-30 13:52:36.349844997 -0500 EST
/staples/snaptel/tasks/Disabled 						    0 		     2016-12-30 13:52:36.349846431 -0500 EST
/staples/snaptel/tasks/Running 							    1 		     2016-12-30 13:52:36.349847677 -0500 EST
/staples/snaptel/tasks/Stopped 							    0 		     2016-12-30 13:52:36.349848711 -0500 EST
```
(Keys `ctrl+c` terminate task watcher)

These data are published to file and stored there (in this example in `/tmp/snap-docker-file.log`).

### Roadmap
Hopefully a more integrated solution.

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap.

To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support).

## License
[Snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).
