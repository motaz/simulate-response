# simulate-response
Simulate API response
It a web service which is used for testing calling APIs to receive pre-defined JSON/XML response

# How to use it
You can call the simulate-response web service using any path, for example /register with content-type application/json it will read response from register.json file, and for application/xml it will read from register.xml, and
when called with plain/text or empty content-type it will read from register.txt file in the same directory of simulate-response application.
Also it can handle sub-path such as /payments/isReady, in this case put isReady.json file inside payments folder
