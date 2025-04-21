## migration-tool add

Add a file to the migration definition.

### Synopsis

Add a file to the migration definition. The file will be added to the end of the migration definition.
Only files in the same folder as the migration.yaml file or in a subfolder of this folder can be added.

```
migration-tool add [filename] [flags]
```

### Options

```
  -d, --description string   Description of the step
  -h, --help                 help for add
```

### Options inherited from parent commands

```
  -o, --execution-log-file string   File where the execution log is written to. (default "execution-log.yaml")
      --folder string               Folder where the scripts and configurations file are located.
```

### SEE ALSO

* [migration-tool](migration-tool.md)	 - migration-tool is a CLI that orchestrates the execution of scripts.

