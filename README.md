# redirhunt
Hunting for Open Redirection in URL parameters

Very simple tool that reads URLs from a file, and replaces each parameter value found with a site specified, then looks for redirects in the http response.

Todo:
- Encoding options to support payload encoding.



```
redirhunt -h

Usage of redirhunt:
  -list string
    	Path to the input file
  -site string
    	URL to replace the parameter value with**

Example (Found 1 open redirect in provided url list):

redirhunt -list domains_uniq_parameters.txt -site https://google.com
Requested URL: https://public-firing-range.appspot.com/redirect/parameter?url=https%3A%2F%2Fgoogle.com
Status Code: 302
Location Header: https://google.com


```
