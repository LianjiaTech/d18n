# Introduction

The `corpus` folder is an external folder used to store mask and detect materials:

* The `corpus/detect` folder is used by detect module
* The `corpus/mask` folder is used by mask module

## use

List of files for `corpus/detect`:

* `gse.address.ja_JP`:  Japan address information is stored,can be used as the corpus of GSE
* `address.zh_CN`:  China address information is stored,can be used as the corpus of GSE
* `regex.phone.zh_CN`: Regular expressions for some Chinese mobile phone numbers,it can be used to detect the mobile phone number ,copy from https://github.com/VincentSit/ChinaMobilePhoneNumberRegex

List of files for `corpus/mask`:

* `address.ja_JP`: can be used to generate fake Japanese addresses
* `address.zh_CN`: can be used to generate fake Chinese addresses
* `BIN.all`: Bank Identification Number，can be used to generate fake bank  card number
* `BIN.zh_CN`: China Bank Identification Number, can be used to generate fake chinese bank  card number
* `IDDevisionCode.zh_CN`: China Administrative Region Code,copy from http://www.gov.cn/test/2011-08/22/content_1930111.htm
* `LicensePlate.en_US`: Us license plate prefix,can be used to generate fake us license plate ,copy from http://www.worldlicenseplates.com/usa/US_XGOV.html
* `LicensePlate.zh_CN`: China license plate prefix,can be used to generate fake chinese license plate ,copy from https://github.com/parkingwang/go-vna/tree/master/data
* `name.ja_JP`:  Japanese name,can be used to generate fake japanese name
* `passport`:  Passport country code,can be used to generate fake passport
* `TelephoneAreaCode.zh_CN`: China Mobile area code,can be used to generate fake fixed telephone number
* `USCC.zh_CN`: China unified social credit code,can be used to generate fake fixed telephone number or code table for mask
