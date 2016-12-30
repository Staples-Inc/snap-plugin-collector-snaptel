# Metrics
| Metrics                                        | Description                                        |
|------------------------------------------------|----------------------------------------------------|
| /staples/snaptel/task/<task name>/fail_count   | Number of failures for this task                   |
| /staples/snaptel/task/<task name>/hit_count    | Number of hits for this task                       |
| /staples/snaptel/task/<task name>/state        | Current state of this task                         |
| /staples/snaptel/task/<task name>/statecode    | Current state of this task as an integer           |
| /staples/snaptel/tasks/Disabled                | Number of Disabled tasks on the agent              |
| /staples/snaptel/tasks/Running                 | Number of Running tasks on the agent               |
| /staples/snaptel/tasks/Stopped                 | Number of Stopped tasks on the agent               |

## State Codes
| Code | Status   |
|------|----------|
|  1   | Running  |
|  2   | Stopped  |
|  3   | Disabled |
