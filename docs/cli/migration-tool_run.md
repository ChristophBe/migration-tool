## migration-tool run

Runs every script that were not previously executed.

### Synopsis

Runs every script that were not previously executed.
Before running the scripts, the scripts are checked for unexecuted changes. 
In case of unexecuted changes, the script is not executed.

```
migration-tool run [flags]
```

### Options

```
  -h, --help   help for run
```

### Options inherited from parent commands

```
  -o, --execution-log-file string   File where the execution log is written to. (default "execution-log.yaml")
      --folder string               Folder where the scripts and configurations file are located.
```

### SEE ALSO

* [migration-tool](migration-tool.md)	 - migration-tool is a CLI that orchestrates the execution of scripts.

