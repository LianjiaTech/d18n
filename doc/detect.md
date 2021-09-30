# Detect

d18n can detect sensitive info use `--detect` flag. d18n has a built-in method, use keywords match, regexp match, and NLP words match. Users can self-define config with `--sensitive` flag.

## config example

```yaml
phone: # classify name
  key: # column name key words match, also support regexp
    - phone
    - telephone
    - phone[_]*number
  value: # value data regexp match
    - ^1(3[0-9]|4[01456879]|5[0-35-9]|6[2567]|7[0-8]|8[0-9]|9[0-35-9])\d{8}$
    - 0\d{2,3}-\d{7,8}|\(?0\d{2,3}[)-]?\d{7,8}|\(?0\d{2,3}[)-]*\d{7,8}
```

## select sensitive data detect

```bash
~ $ d18n --defaults-extra-file test/my.cnf --database sakila --query 'select * from address limit 10' -detect
{
  "address": [
    "address"
  ],
  "address2": [
    "address"
  ],
  "address_id": [
    "address"
  ],
  "city_id": [
    "address"
  ],
  "district": null,
  "last_update": null,
  "location": [
    "address"
  ],
  "phone": [
    "phone"
  ],
  "postal_code": null
}
```

## full table scan detect sensitive data

A full table scan will cost a long time.

```bash
~ $ d18n --defaults-extra-file test/my.cnf --database sakila --table address --detect
{
  "address": [
    "address"
  ],
  "address2": [
    "address"
  ],
  "address_id": [
    "address"
  ],
  "city_id": [
    "address"
  ],
  "district": null,
  "last_update": null,
  "location": [
    "address"
  ],
  "phone": [
    "phone"
  ],
  "postal_code": null
}
```

## jq filter

```bash
# show all sensitive columns and type
~ $ d18n ... | jq -r 'del(.[] | select(. == null))'
{
  "address": [
    "address"
  ],
  "address2": [
    "address"
  ],
  "address_id": [
    "address"
  ],
  "city_id": [
    "address"
  ],
  "location": [
    "address"
  ],
  "phone": [
    "phone"
  ]
}

# only show column name
~ $ d18n ... | jq -r 'del(.[] | select(. == null)) | keys | .[]'
address
address2
address_id
city_id
location
phone
```

## NLP base sensitive detection

d18n use `github.com/go-ego/gse` library to find sensitive info. Use `gse` must define corpus first. d18n has a built-in corpus about Chinese names and addresses.

```text
北京市 10 address
东城区 10 address
东华门街道 10 address
景山街道 10 address

赵 10 name
钱 10 name
孙 10 name
李 10 name
```

```text
GSE("我叫岳云鹏") => name
GSE("我住在北京市大兴区庞各庄镇") => address
```

## Corpus

Add new file into `detect/corpus` directory, file format please reference `gse.address.zh_CN`, first column is keyword, second column is score, third column is type.

Notice: file should not be too large, it will build into d18n binary file. remove no used corpus should be nice.
