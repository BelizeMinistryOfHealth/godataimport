# GoData Import Tool

This tool imports a csv file of covid tests to GoData.

## Installation
Download Links:

Windows: https://storage.googleapis.com/epi_moh_bz/godata-import/windows/godataimport

Linux: https://storage.googleapis.com/epi_moh_bz/godata-import/linux/godataimport

Copy the file to your execution PATH. Alternatively, copy it to the directory you
will run the program from

## Usage
1. Open your terminal.
2. Go to the directory where you have the csv file.
3. Run the import tool:

```sh
bin/godataimport -u <godata username> -p <godata password -f <csv file> -d <destination file>
```
The data in the csv file will be parsed and converted to the structure that is compatible with GoData.
