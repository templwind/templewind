Name: {{.serviceName}}
Host: {{.host}}
Port: {{.port}}
DSN: ${DSN}
EnableWALMode: true
{{.auth}}
Site:
  Title: {{.serviceName}}
Assets:
  CSS:
    - /assets/css/styles.css
  JS:
    - /assets/js/main.js
Menus:
  main:
    - URL: /
      Title: Home
      Identifier: home
    - URL: /about
      Title: About
      Identifier: about
    - URL: /contact
      Title: Contact
      Identifier: contact
