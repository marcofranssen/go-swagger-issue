# Swagger issue

This example showcases the unreliable behavior of the following [go-swagger issue](https://github.com/go-swagger/go-swagger/issues/2066).

Run `make install` to install the dependencies.

Then run `make` a bunch of times to reproduce the issue.

Without make you can run the commands as defined in `make reproduce` and `make generate` manually from the terminal.

You will figure that at random moments there will be a `git diff` on the `docs/swagger.json` sometimes every other run, sometimes it takes 10 or 20 runs to have the issue occur.

See below for the reproduced issue....

```bash
marco@DESKTOP-B249DU4 MINGW64 /c/code/go-swagger-issue (master)
$ make
Generate code

marco@DESKTOP-B249DU4 MINGW64 /c/code/go-swagger-issue (master)
$ make
Generate code

marco@DESKTOP-B249DU4 MINGW64 /c/code/go-swagger-issue (master)
$ make
Generate code
diff --git a/docs/swagger.json b/docs/swagger.json
index 42ce3b3..991b945 100644
--- a/docs/swagger.json
+++ b/docs/swagger.json
@@ -47,6 +47,26 @@
     }
   },
   "definitions": {
+    "Example": {
+      "description": "Example the response of an example resource",
+      "type": "object",
+      "properties": {
+        "description": {
+          "type": "string",
+          "x-go-name": "Description"
+        },
+        "priority": {
+          "type": "integer",
+          "format": "int64",
+          "x-go-name": "Priority"
+        },
+        "title": {
+          "type": "string",
+          "x-go-name": "Title"
+        }
+      },
+      "x-go-package": "github.com/marcofranssen/go-swagger-issue/domain"
+    },
     "ExampleListResponse": {
       "description": "ExampleListResponse the response of a list of example resources",
       "type": "object",

marco@DESKTOP-B249DU4 MINGW64 /c/code/go-swagger-issue (master)
$ make
Generate code

marco@DESKTOP-B249DU4 MINGW64 /c/code/go-swagger-issue (master)
$ make
Generate code

marco@DESKTOP-B249DU4 MINGW64 /c/code/go-swagger-issue (master)
$ make
Generate code
diff --git a/docs/swagger.json b/docs/swagger.json
index 42ce3b3..991b945 100644
--- a/docs/swagger.json
+++ b/docs/swagger.json
@@ -47,6 +47,26 @@
     }
   },
   "definitions": {
+    "Example": {
+      "description": "Example the response of an example resource",
+      "type": "object",
+      "properties": {
+        "description": {
+          "type": "string",
+          "x-go-name": "Description"
+        },
+        "priority": {
+          "type": "integer",
+          "format": "int64",
+          "x-go-name": "Priority"
+        },
+        "title": {
+          "type": "string",
+          "x-go-name": "Title"
+        }
+      },
+      "x-go-package": "github.com/marcofranssen/go-swagger-issue/domain"
+    },
     "ExampleListResponse": {
       "description": "ExampleListResponse the response of a list of example resources",
       "type": "object",
```
