# GCP GCS Bucket Pulumi Module
## Key Features

### API Resource Features

- **Standardized Structure**: The `GcsBucket` API resource adheres to a consistent schema with `apiVersion`, `kind`, `metadata`, `spec`, and `status` fields. This uniformity ensures compatibility and ease of integration within Kubernetes-like environments, promoting seamless workflow incorporation and tooling interoperability.
  
- **Configurable Specifications**:
  - **GCP Credential ID**: Securely references the GCP credentials required to set up the Pulumi provider, ensuring authenticated and authorized interactions with GCP services.
  - **GCP Project ID**: Specifies the GCP project where the GCS Bucket resources will be deployed, enabling precise management and organization of cloud resources.
  - **GCP Region**: Defines the region where the GCS Bucket will be created, allowing for optimized performance and compliance with data residency requirements.
  - **Public Access Configuration**: Flags to indicate if the GCS Bucket should have external (public) access, automatically configuring access controls based on the specified setting.
  
- **Validation and Compliance**: Incorporates stringent validation rules to ensure all configurations adhere to established standards and best practices, minimizing the risk of misconfigurations and enhancing overall system reliability.

### Pulumi Module Features

- **Automated GCP Provider Setup**: Utilizes the provided GCP credentials to automatically configure the Pulumi GCP provider, enabling seamless and secure interactions with GCP Cloud Storage resources.
  
- **GCS Bucket Management**: Streamlines the creation and management of GCS Buckets based on the provided specifications. This includes setting up bucket parameters such as location, access controls, and resource labels to align with organizational requirements and performance standards.
  
- **Public Access Configuration**: Automatically configures bucket access controls based on the `is_public` flag. If set to true, the module grants public read access to the bucket and its objects, facilitating easy distribution of publicly accessible content.
  
- **Exported Stack Outputs**: Captures essential outputs such as bucket IDs and access configurations in `status.stackOutputs`. These outputs facilitate integration with other infrastructure components, enabling effective monitoring, management, and automation workflows.
  
- **Scalability and Flexibility**: Designed to accommodate a wide range of GCS Bucket configurations, the module supports varying levels of complexity and can be easily extended to meet evolving infrastructure demands, ensuring long-term adaptability.
  
- **Error Handling**: Implements robust error handling mechanisms to promptly identify and report issues encountered during deployment or configuration processes, aiding in swift troubleshooting and maintaining infrastructure integrity.

## Installation

To integrate the GCP GCS Bucket Pulumi Module into your project, retrieve it from the [GitHub repository](https://github.com/your-repo/gcs-bucket-pulumi-module). Ensure that you have both Pulumi and Go installed and properly configured in your development environment.

```shell
git clone https://github.com/your-repo/gcs-bucket-pulumi-module.git
cd gcs-bucket-pulumi-module
```

## Usage

Refer to the [example section](#examples) for usage instructions.

## Module Details

### Input Configuration

The module expects a `GcsBucketStackInput` which includes:

- **Pulumi Input**: Configuration details required by Pulumi for managing the stack, such as stack names, project settings, and any necessary Pulumi configurations.
- **Target API Resource**: The `GcsBucket` resource defining the desired GCS Bucket configuration, including project ID, region, and public access settings.
- **GCP Credential**: Specifications for the GCP credentials used to authenticate and authorize Pulumi operations, ensuring secure interactions with GCP services.

### Exported Outputs

Upon successful execution, the module exports the following outputs to `status.stackOutputs`:

- **Bucket ID**: The unique identifier assigned to the created GCS Bucket, which can be used for management and monitoring purposes.
- **Access Configuration**: Information regarding the access controls applied to the GCS Bucket, including public access settings if enabled.

These outputs facilitate integration with other infrastructure components, provide essential information for monitoring and management, and enable automation workflows to utilize the deployed GCS Bucket resources effectively.

## Contributing

We welcome contributions to enhance the GCP GCS Bucket Pulumi Module. Please refer to our [contribution guidelines](CONTRIBUTING.md) for more information on how to get involved, including coding standards, submission processes, and best practices.

## License

This project is licensed under the [MIT License](LICENSE), granting you the freedom to use, modify, and distribute the software with minimal restrictions. Please review the LICENSE file for more details.

## Support

For support, please contact our [support team](mailto:support@planton.cloud). We are here to help you with any issues, questions, or feedback you may have regarding the GCP GCS Bucket Pulumi Module.

## Acknowledgements

Special thanks to all contributors and the Planton Cloud community for their ongoing support and feedback. Your efforts and dedication are instrumental in making this module robust and effective.

## Changelog

A detailed changelog is available in the [CHANGELOG.md](CHANGELOG.md) file. It documents all significant changes, enhancements, bug fixes, and updates made to the GCP GCS Bucket Pulumi Module over time.

## Roadmap

We are continuously working to enhance the GCP GCS Bucket Pulumi Module. Upcoming features include:

- **Advanced IAM Configurations**: Implementing more granular permission controls for GCP resources to enhance security and compliance.
- **Enhanced Monitoring Integrations**: Integrating with monitoring and logging tools such as Prometheus, Grafana, and the ELK stack for better observability and performance tracking.
- **Support for Additional Cloud Providers**: Extending support to more cloud platforms to increase flexibility and accommodate diverse infrastructure requirements.
- **Automated Scaling and Optimization**: Introducing automated scaling capabilities based on workload demands and performance metrics to optimize resource utilization.
- **Comprehensive Documentation and Tutorials**: Expanding documentation and providing step-by-step tutorials to assist users in effectively leveraging the module's capabilities.

Stay tuned for more updates as we continue to develop and refine the module to meet your infrastructure management needs.

## Contact

For any inquiries or feedback, please reach out to us at [contact@planton.cloud](mailto:contact@planton.cloud). We value your input and are committed to providing the support you need to effectively manage your GCS Bucket resources.

## Disclaimer

*This project is maintained by Planton Cloud and is not affiliated with any third-party services unless explicitly stated. While we strive to ensure the accuracy and reliability of this module, users are encouraged to review and test configurations in their environments.*

## Security

If you discover any security vulnerabilities, please report them responsibly by contacting our security team at [security@planton.cloud](mailto:security@planton.cloud). We take security seriously and are committed to addressing any issues promptly to protect our users and their infrastructure.

## Code of Conduct

Please adhere to our [Code of Conduct](CODE_OF_CONDUCT.md) when interacting with the project. We are committed to fostering an inclusive and respectful community where all contributors feel welcome and valued.

## References

- [Pulumi Documentation](https://www.pulumi.com/docs/)
- [GCP Cloud Storage Documentation](https://cloud.google.com/storage/docs)
- [Planton Cloud APIs](https://buf.build/plantoncloud/planton-cloud-apis/docs)

## Getting Started

To get started with the GCP GCS Bucket Pulumi Module, follow the installation instructions above and refer to the examples section for detailed usage guidelines. Our comprehensive documentation will guide you through configuring your API resources, setting up Pulumi stacks, and deploying your GCS Bucket configurations with ease.

---

*Thank you for choosing Planton Cloud's GCP GCS Bucket Pulumi Module. We are dedicated to supporting your infrastructure management needs and look forward to helping you achieve seamless and efficient Cloud Storage deployments.*

## Examples

### Basic Example

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: todo-list-api
spec:
  environmentInfo:
    envId: my-org-prod
  version: main
  container:
    app:
      image:
        repo: nginx
        tag: latest
      ports:
        - appProtocol: http
          containerPort: 8080
          isIngressPort: true
          servicePort: 80
      resources:
        requests:
          cpu: 100m
          memory: 100Mi
        limits:
          cpu: 2000m
          memory: 2Gi
```

### Example with Environment Variables

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: todo-list-api
spec:
  environmentInfo:
    envId: my-org-prod
  version: main
  container:
    app:
      env:
        variables:
          DATABASE_NAME: todo
      image:
        repo: nginx
        tag: latest
      ports:
        - appProtocol: http
          containerPort: 8080
          isIngressPort: true
          name: rest-api
          networkProtocol: TCP
          servicePort: 80
      resources:
        requests:
          cpu: 100m
          memory: 100Mi
        limits:
          cpu: 2000m
          memory: 2Gi
```

### Example with Environment Secrets

The below example assumes that the secrets are managed by Planton Cloud's [GCP Secrets Manager](https://buf.build/plantoncloud/planton-cloud-apis/docs/main:cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanager) deployment module.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: todo-list-api
spec:
  environmentInfo:
    envId: my-org-prod
  version: main
  container:
    app:
      env:
        secrets:
          # value before dot 'gcpsm-my-org-prod-gcp-secrets' is the id of the gcp-secret-manager resource on planton-cloud
          # value after dot 'database-password' is one of the secrets list in 'gcpsm-my-org-prod-gcp-secrets' is the id of the gcp-secret-manager resource on planton-cloud
          DATABASE_PASSWORD: ${gcpsm-my-org-prod-gcp-secrets.database-password}
        variables:
          DATABASE_NAME: todo
      image:
        repo: nginx
        tag: latest
      ports:
        - appProtocol: http
          containerPort: 8080
          isIngressPort: true
          name: rest-api
          networkProtocol: TCP
          servicePort: 80
      resources:
        requests:
          cpu: 100m
          memory: 100Mi
        limits:
          cpu: 2000m
          memory: 2Gi
```

### Example with Multiple Containers

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: multi-container-app
spec:
  environmentInfo:
    envId: dev-env
  version: v1.0.0
  container:
    app:
      image:
        repo: myapp/frontend
        tag: v1.0.0
      ports:
        - appProtocol: http
          containerPort: 80
          isIngressPort: true
          servicePort: 8080
      resources:
        requests:
          cpu: 200m
          memory: 256Mi
        limits:
          cpu: 1000m
          memory: 1Gi
    sidecar:
      image:
        repo: myapp/logging
        tag: v1.0.0
      ports:
        - appProtocol: tcp
          containerPort: 5000
          isIngressPort: false
          servicePort: 5000
      resources:
        requests:
          cpu: 100m
          memory: 128Mi
        limits:
          cpu: 500m
          memory: 512Mi
```

### Example with Different Resource Limits

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: high-memory-service
spec:
  environmentInfo:
    envId: staging-env
  version: beta
  container:
    app:
      image:
        repo: highmemapp/backend
        tag: latest
      ports:
        - appProtocol: http
          containerPort: 9090
          isIngressPort: true
          servicePort: 9090
      resources:
        requests:
          cpu: 500m
          memory: 512Mi
        limits:
          cpu: 4000m
          memory: 8Gi
```

### Example with Annotations and Labels

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: annotated-service
  labels:
    app: annotated-app
    tier: backend
  annotations:
    description: "This service handles user authentication."
spec:
  environmentInfo:
    envId: production-env
  version: release
  container:
    app:
      image:
        repo: auth-service/image
        tag: release
      ports:
        - appProtocol: https
          containerPort: 8443
          isIngressPort: true
          servicePort: 443
      resources:
        requests:
          cpu: 250m
          memory: 256Mi
        limits:
          cpu: 1500m
          memory: 2Gi
```

### Example with Health Checks

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: healthcheck-service
spec:
  environmentInfo:
    envId: test-env
  version: test
  container:
    app:
      image:
        repo: healthapp/service
        tag: test
      ports:
        - appProtocol: http
          containerPort: 8000
          isIngressPort: true
          servicePort: 8000
      resources:
        requests:
          cpu: 100m
          memory: 128Mi
        limits:
          cpu: 500m
          memory: 1Gi
      livenessProbe:
        httpGet:
          path: /healthz
          port: 8000
        initialDelaySeconds: 30
        periodSeconds: 10
      readinessProbe:
        httpGet:
          path: /ready
          port: 8000
        initialDelaySeconds: 10
        periodSeconds: 5
```

### Example with Empty Spec

*Note: This module is not completely implemented. Certain features may be missing or not fully functional. Future updates will address these limitations.*

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: incomplete-service
spec: {}
```

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: another-incomplete-service
spec: {}
```
