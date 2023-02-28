# Web Applications

Writing related to "web applications" as a topic.

## Ways to build Web Applications

[[Source](https://tomhummel.com/posts/four-web-apps/)]

1. Hugo Static sites + Progressive Web Apps:
    - Common Hosting Providers: [Cloudflare Pages](https://pages.cloudflare.com/), [AWS S3](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_website_configuration), Vercel, Netlify.
    - Base hosting cost: Free
    - Ideal for: high-read, low-write applications. Content sites, blogs, marketing, documentation, data-driven websites in which the data changes infrequently (20-30 updates per day is probably on the high side).
    - Not great for: high-write applications
2. Cloudflare Workers:
    - Base hosting cost: Free tier, and then $5/mo for higher volume and access to more storage options.
    - Ideal for:
        - Customizing response headers
        - Privacy-respecting usage analytics
        - Authenticating/authorizing users
        - Handling form submissions
        - Triggering transactional email
        - Dynamically building JSON payloads
        - Interacting with storage (cloudflare or third party)
        - Building A/B testing and feature flagging capabilities
        - Achieving global points of presence as close as possible to consumers
    - Not great for:
        - Complicated routing schemes or very large applications
        - Applications that require heavier frameworks and functionality
        - Applications that require a more advanced or featureful runtime
        - Multi-provider architectures (identity and networking can be cumbersome or brittle)
3. Linux Server Singleton
    - Common Hosting Providers: [AWS EC2](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance), [OCI Compute](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/core_instance), [Azure Compute](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/linux_virtual_machine)
    - Base hosting cost: Limited free tier depending on the provider, and then ~$5/mo and up. Additional costs for block and object storage and network bandwidth.
    - Ideal for:
        - Projects which need full customization of the operating system, operational agents/tools, application runtime, and frameworks.
        - Colocating your primary datastore with your application code, in process for simplified database connectivity.
        - Attempting to minimize per request costs down as far as you can.
        - Uptime objectives of 99.9% (three 9’s) availability.
        - Usecases which can be located in a single geography.
        - Taking advantage of IaaS provider identity and networking primitives while not irreversibly tying yourself to that provider.
    - Not great for:
        - Teams which don’t want the overhead of managing even a single linux server - security patching, backups, disaster recovery scenarios.
        - Uptime objectives of 99.999% (five 9’s) or higher.
        - Workloads which are likely to outpace the upper bound of vertical scaling (note: the upper bound is surprisingly high!).
        - If you need global points of presence.
4. Containers on Platform as a Service (PaaS)
    - Common Hosting Providers:
        - [ECS Fargate](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs_cluster)
        - [Azure Container Instances](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/container_group)
        - [OCI Container Instances](https://registry.terraform.io/providers/oracle/oci/latest/docs/resources/container_instances_container_instance)
        - [Google Cloud Run](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloud_run_v2_service)
        - [Fly.io](https://docs.aws.amazon.com/AmazonECS/latest/bestpracticesguide/security-shared.html)
    - Base hosting cost: Limited free tier depending on the provider, and then ~$3/mo and up. Additional costs for block/object storage and network bandwidth.
    - Ideal for:
        - Projects which need a more customized runtime than Cloudflare Workers can provide but don’t want to manage servers.
        - Projects which need to scale the application tier horizontally.
        - Projects which can afford a network hop between the stateless application tier and the persistent data tier.
        - Projects which want to try for 5+ 9’s of availability.
        - Projects which need global points of presence (note that it can be easy push your application to the edge, but it is more challenging to do so with your data)
        - Taking advantage of an IaaS provider identity and networking primitives.
    - Not great for:
        - Teams which don’t want to be in the business of building container ci/cd deployment pipelines.
        - Teams which don’t want to manage the complexity of distributed systems.
