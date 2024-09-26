Here are a few example scenarios for the **GCP GCS Bucket Pulumi Module** that demonstrate different configurations for managing Google Cloud Storage (GCS) buckets. These examples cover a range of use cases, from basic public bucket creation to more advanced setups involving versioning, access control, and lifecycle management.

---

### Example 1: Basic GCS Bucket with Public Access

This example creates a basic GCS bucket with public access enabled. The bucket is configured to allow anyone to access its contents, which is useful for hosting public assets like images or files.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcsBucket
metadata:
  name: public-gcs-bucket
spec:
  gcp_credential_id: gcp-public-cred
  project_id: my-gcp-project
  bucket_config:
    bucket_name: public-bucket
    location: US
    public_access: true
    versioning_enabled: false
```

---

### Example 2: GCS Bucket with Private Access and Versioning

In this example, a GCS bucket is created with private access, meaning only authenticated users can access the bucket. Versioning is also enabled to retain multiple versions of objects for recovery purposes.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcsBucket
metadata:
  name: private-versioned-bucket
spec:
  gcp_credential_id: gcp-private-cred
  project_id: my-private-gcp-project
  bucket_config:
    bucket_name: versioned-private-bucket
    location: US
    public_access: false
    versioning_enabled: true
```

---

### Example 3: GCS Bucket with CORS Configuration

This example sets up a GCS bucket with Cross-Origin Resource Sharing (CORS) enabled, allowing external applications from specific domains to access the contents of the bucket.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcsBucket
metadata:
  name: cors-enabled-bucket
spec:
  gcp_credential_id: gcp-cors-cred
  project_id: my-gcp-project
  bucket_config:
    bucket_name: cors-bucket
    location: US
    public_access: false
    versioning_enabled: false
    cors_settings:
      allowed_origins:
        - https://www.example.com
      allowed_methods:
        - GET
        - POST
      max_age_seconds: 3600
```

---

### Example 4: GCS Bucket with Lifecycle Rules for Object Deletion

This example creates a GCS bucket with lifecycle rules that automatically delete objects older than 30 days. This is useful for managing log files or temporary data.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcsBucket
metadata:
  name: lifecycle-bucket
spec:
  gcp_credential_id: gcp-lifecycle-cred
  project_id: lifecycle-project
  bucket_config:
    bucket_name: lifecycle-bucket
    location: US
    public_access: false
    versioning_enabled: true
    lifecycle_rules:
      - action: Delete
        age: 30
```

---

### Example 5: Multi-Region GCS Bucket for High Availability

This example provisions a GCS bucket with a multi-region configuration to ensure high availability and redundancy. This is ideal for applications requiring data storage across multiple geographic locations.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcsBucket
metadata:
  name: multi-region-bucket
spec:
  gcp_credential_id: gcp-multi-region-cred
  project_id: high-availability-project
  bucket_config:
    bucket_name: multi-region-bucket
    location: multi-region
    public_access: true
    versioning_enabled: true
```

---

### Example 6: GCS Bucket with Logging Enabled

This example creates a GCS bucket with access logging enabled, which tracks requests made to the bucket and stores logs in a separate log bucket. This is useful for monitoring and auditing access to the bucket's contents.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcsBucket
metadata:
  name: logging-enabled-bucket
spec:
  gcp_credential_id: gcp-logging-cred
  project_id: logging-project
  bucket_config:
    bucket_name: logging-bucket
    location: US
    public_access: true
    versioning_enabled: false
    logging:
      log_bucket_name: gcs-access-logs
      log_object_prefix: log-
```

---

### Example 7: GCS Bucket with Custom Object Retention Policy

This example provisions a GCS bucket with a custom object retention policy that prevents objects from being deleted or overwritten for 365 days. This setup is ideal for regulatory compliance or archival purposes.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcsBucket
metadata:
  name: retention-policy-bucket
spec:
  gcp_credential_id: gcp-retention-cred
  project_id: retention-policy-project
  bucket_config:
    bucket_name: retention-policy-bucket
    location: US
    public_access: false
    versioning_enabled: true
    retention_policy:
      retention_period_days: 365
```

---

### Example 8: GCS Bucket with Object Encryption Using Customer-Supplied Keys

This example creates a GCS bucket with object-level encryption using customer-supplied encryption keys (CSEK). This configuration ensures that the bucket's contents are encrypted with a key provided by the user.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcsBucket
metadata:
  name: encrypted-bucket
spec:
  gcp_credential_id: gcp-encryption-cred
  project_id: encryption-project
  bucket_config:
    bucket_name: encrypted-bucket
    location: US
    public_access: false
    versioning_enabled: true
    encryption_settings:
      encryption_type: CSEK
      encryption_key: my-customer-supplied-key
```

---

### Example 9: GCS Bucket with Uniform Access Control

This example provisions a GCS bucket with uniform access control enabled. This enforces the same permissions across all objects within the bucket, making it easier to manage access policies at the bucket level.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GcsBucket
metadata:
  name: uniform-access-bucket
spec:
  gcp_credential_id: gcp-uniform-cred
  project_id: uniform-access-project
  bucket_config:
    bucket_name: uniform-access-bucket
    location: US
    public_access: false
    versioning_enabled: false
    uniform_access_control: true
```

---

### Applying the Configurations

To deploy any of these GCP GCS Bucket configurations, create a YAML file with the desired example content and use the following command to apply the configuration:

```shell
planton apply -f <yaml-path>
```
