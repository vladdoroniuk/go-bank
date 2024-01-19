<p align="center">
  <picture>
    <source media="(prefers-color-scheme: light)" srcset="./images/logo.svg">
    <img width="250" height="auto" alt="rose_logo" src="./images/logo.svg">
  </picture>
</p>

# Overview

RoSe (**Ro**bust **Se**rvices)

# Software Development Life Cycle (SDLC)

todo

<p align="center">
  <picture>
    <source media="(prefers-color-scheme: light)" srcset="./images/sdlc.png">
    <img width="auto" height="auto" alt="sdlc" src="./images/sdlc.png">
  </picture>
</p>

## Analysis

todo

## Design

### Hosting Solutions

- Self-hosted with **Docker Compose**
- Self-hosted with **Kubernetes** + **Helm**
- Cloud (provisioned with **Terraform**)

### Communication

- Sync Communication: **REST (JSON)**, **gRPC**

### Services

| Name                          | Use case                             | Cloud provider      | Can be self-hosted?         |
| ----------------------------- | ------------------------------------ | ------------------- | --------------------------- |
| **Kong**                      | API gateway                          | Kong Cloud          | +                           |
| **Auth0**                     | Identity and Access Management (IAM) | Auth0               | - (Keycloak or SuperTokens) |
| **Prometheus**                | Metrics collection and storage       | Grafana Cloud       | +                           |
| **Grafana**                   | Metrics visualization                | Grafana Cloud       | +                           |
| **Elasticsearch**             | Logs storage                         | Elastic Cloud       | +                           |
| **Logstash**                  | Logs processing pipeline             | - (runs in AWS ECS) | +                           |
| **Kibana**                    | Logs visualization                   | Elastic Cloud       | +                           |
| **Sentry**                    | Error tracking                       | Sentry              | +                           |
| **AWS WAF**                   | Firewall                             | AWS                 | ?                           |
| **AWS CloudFront**            | Content Delivery Network (CDN)       | AWS                 | ?                           |
| **AWS Route 53**              | Domain Name System (DNS)             | AWS                 | ?                           |
| **AWS ECR**                   | Container registry                   | AWS                 | +- (Docker images)          |
| **AWS ECS** + **AWS Fargate** | Container orchestration              | AWS                 | +- (Kubernetes)             |
| **AWS EFS**                   | File system for ECS volumes          | AWS                 | +- (Host file system)       |
| **AWS RDS**                   | PostgreSQL database                  | AWS                 | +- (PostgreSQL)             |

? - additional research needed

## Development

todo

## Testing

todo

## Deployment

todo

## Maintenance

todo

# Roadmap

todo
