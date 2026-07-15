openapi.jsonは https://manual.sakura.ad.jp/api/cloud/monitoring-suite/ からダウンロードできるJSONに以下の変更を加えたものです

```console
$ jq 'del(.paths.[].[].requestBody.content.["application/x-www-form-urlencoded", "multipart/form-data"])' openapi.json
```

以下のパッチがあたっていますが、自動生成に必要なだけで、本質的にAPIが変わっているわけではありません。

```diff
diff --git a/monitoring-suite-api.json b/openapi.json
index 7e5a3d3..6c55515 100644
--- a/monitoring-suite-api.json
+++ b/openapi.json
@@ -4885,7 +4885,21 @@
             "$ref": "#/components/schemas/MapValueNumMatcher"
           }
         ],
-        "title": "FieldMatcher"
+        "title": "FieldMatcher",
+        "discriminator": {
+          "propertyName": "type",
+          "mapping": {
+            "or": "#/components/schemas/OrMatcher",
+            "and": "#/components/schemas/AndMatcher",
+            "string": "#/components/schemas/StrMatcher",
+            "number": "#/components/schemas/NumMatcher",
+            "enum": "#/components/schemas/EnumMatcher",
+            "map-key-exists": "#/components/schemas/MapKeyExistsMatcher",
+            "map-key-value-matcher": "#/components/schemas/MapKeyValueMatcher",
+            "map-value-string": "#/components/schemas/MapValueStrMatcher",
+            "map-value-number": "#/components/schemas/MapValueNumMatcher"
+          }
+        }
       },
       "FieldModel": {
         "enum": [
```