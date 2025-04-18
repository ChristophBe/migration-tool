## migration-tool regenerate-hashes

This command recalculates the hashes of the scripts.

### Synopsis

Regenerate hashes of the scripts. This is useful if the scripts have been intentional changed.
Be careful, this can lead to consistent behavior while executing the scripts. It might prevent the run command to execute scripts in some cases.
It is recommended that this is only used to recalculate the hashes for scripts that were not executed in any environment before.


```
migration-tool regenerate-hashes [flags]
```

### Options

```
  -h, --help   help for regenerate-hashes
```

### Options inherited from parent commands

```
  -o, --execution-log-file string   File where the execution log is written to. (default "execution-log.yaml")
      --folder string               Folder where the scripts and configurations file are located.
```

### SEE ALSO

* [migration-tool](migration-tool.md)	 - migration-tool is a CLI that orchestrates the execution of scripts.

