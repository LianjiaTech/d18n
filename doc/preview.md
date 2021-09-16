# File Preview

`--preview` follow a number, how many rows to show, including the headline. It's useful for Excel or HTML files. For other plain text files, we would like to use `less`, `more` or `head` commands.

## Example

```bash
# preview xlsx file
d18n --preview 10 --file test.xlsx
```

```bash
# --preview should not be used with --query, --lint, --import, --detect flag
# If test.csv not exists, d18n will dump data into test.csv, won't preview file.
# If test.csv exists, d18n just preview file, won't dump data.
d18n --defaults-extra-file my.cnf --query "select 1" --file test.csv --preview 1
```
