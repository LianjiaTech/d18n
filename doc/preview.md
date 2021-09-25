# File Preview

`--preview` follow a number, how many rows to show, including the headline. It's useful for Excel or HTML files. For other plain text files, we would like to use `less`, `more` or `head` commands.

## Example

```bash
# preview xlsx file
d18n --preview 10 --file test/actor.xlsx

# preview html file
d18n --preview 10 --file test/actor.html
```
