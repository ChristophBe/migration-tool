## migration-tool

migration-tool is a CLI that orchestrates the execution of scripts.

### Synopsis

migration-tool is a CLI that orchestrates the execution of bash scripts organized as steps.

It makes sure that the scripts are executed in the correct order and that the scripts are only executed if the script have not run before.
To ensure consistency, the scripts are checked for unexecuted changes.


### Options

```
  -o, --execution-log-file string   File where the execution log is written to. (default "execution-log.yaml")
      --folder string               Folder where the scripts and configurations file are located.
  -h, --help                        help for migration-tool
```

### SEE ALSO

* [migration-tool add](migration-tool_add.md)	 - Add a file to the migration definition.
* [migration-tool regenerate-hashes](migration-tool_regenerate-hashes.md)	 - This command recalculates the hashes of the scripts.
* [migration-tool run](migration-tool_run.md)	 - Runs every script that were not previously executed.
* [migration-tool verify](migration-tool_verify.md)	 - Verify checks if the scripts have changed.

