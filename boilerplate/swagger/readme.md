# swagger


```yaml
Swagger UI Version	Release Date	OpenAPI Spec compatibility	Notes
3.18.3	2018-08-03	2.0, 3.0	tag v3.18.3

https://github.com/swagger-api/swagger-ui/tree/v3.18.3
```

## you may have the need to customize the `swagger.yml` file url
open `index.html` edit LINE 40 `url: "/swagger/doc.yml",`
```html
  window.onload = function () {

        // Build a system
        const ui = SwaggerUIBundle({
            url: "/swagger/doc.yml",
            dom_id: '#swagger-ui',
            deepLinking: true,
            presets: [
                SwaggerUIBundle.presets.apis,
                SwaggerUIStandalonePreset
            ],
            plugins: [
                SwaggerUIBundle.plugins.DownloadUrl
            ],
            layout: "StandaloneLayout"
        })
```