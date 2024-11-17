# redirhunt
Hunting for Open Redirection in URL parameters

Very simple tool that reads URLs from a file, and replaces each parameter value found with a site specified, then looks for redirects in the http response.

Todo:
- More testing.
- Encoding options to support payload encoding.
- Support HTTP POST messages, so it can read parameters from the body request as well. 



```
**redirhunt -h
Usage of redirhunt:
  -list string
    	Path to the input file
  -site string
    	URL to replace the parameter value with**
```
