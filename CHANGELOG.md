## v0.27.0 [2022-09-06]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.6](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v416-2022-09-02) which includes several caching and memory management improvements. ([#387](https://github.com/turbot/steampipe-plugin-gcp/pull/387))
- Recompiled plugin with Go version `1.19`. ([#383](https://github.com/turbot/steampipe-plugin-gcp/pull/383))

_Enhancements_

- Added column `self_link` to `gcp_dataproc_cluster table`. ([#386](https://github.com/turbot/steampipe-plugin-gcp/pull/386))

## v0.26.0 [2022-08-16]

_What's new?_

- New tables added
  - [gcp_dataproc_cluster](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_dataproc_cluster) ([#379](https://github.com/turbot/steampipe-plugin-gcp/pull/379))

_Enhancements_

- Added column `access_approval_settings` to `gcp_project` table. ([#381](https://github.com/turbot/steampipe-plugin-gcp/pull/381))
- Added column `essential_contacts` to `gcp_organization` table. ([#380](https://github.com/turbot/steampipe-plugin-gcp/pull/380))

## v0.25.0 [2022-07-13]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v3.3.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v332--2022-07-11) which includes several caching fixes. ([#371](https://github.com/turbot/steampipe-plugin-gcp/pull/371))

## v0.24.0 [2022-07-01]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v3.3.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v331--2022-06-30). ([#368](https://github.com/turbot/steampipe-plugin-gcp/pull/368))

## v0.23.1 [2022-05-23]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#363](https://github.com/turbot/steampipe-plugin-gcp/pull/363))

## v0.23.0 [2022-04-27]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#360](https://github.com/turbot/steampipe-plugin-gcp/pull/360))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#356](https://github.com/turbot/steampipe-plugin-gcp/pull/356))

## v0.22.1 [2022-03-25]

_Bug fixes_

- Fixed the `gcp_compute_resource_policy` table to correctly return the compute resource policies instead of an error ([#352](https://github.com/turbot/steampipe-plugin-gcp/pull/352))

## v0.22.0 [2022-03-23]

_What's new?_

- New tables added
  - [gcp_kms_key_version](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_kms_key_version) ([#345](https://github.com/turbot/steampipe-plugin-gcp/pull/345))

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v2.1.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v211--2022-03-10) ([#347](https://github.com/turbot/steampipe-plugin-gcp/pull/347))

## v0.21.0 [2022-03-17]

_Enhancements_

- Added the `autopilot_enabled` column to the `gcp_kubernetes_cluster` table ([#344](https://github.com/turbot/steampipe-plugin-gcp/pull/344))

_Bug fixes_

- The `gcp_project` table now correctly returns projects that have a project name different from their project ID

## v0.20.0 [2022-01-19]

_Enhancements_

- Added additional optional key quals, filter support, and context cancellation handling to `BigQuery`, `Cloud Functions`, `Compute Engine`, `DNS`, `KMS`, `Logging`, `Monitoring`, `Project`, `Pub/Sub` and `SQL` tables ([#337](https://github.com/turbot/steampipe-plugin-gcp/pull/337))
- Recompiled plugin with [steampipe-plugin-sdk v1.8.3](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v183--2021-12-23) ([#339](https://github.com/turbot/steampipe-plugin-gcp/pull/339))

## v0.19.0 [2021-11-23]

_Enhancements_

- Updated: Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) and Go version 1.17 ([#332](https://github.com/turbot/steampipe-plugin-gcp/pull/332)) ([#335](https://github.com/turbot/steampipe-plugin-gcp/pull/335))

## v0.18.0 [2021-10-26]

_What's new?_

- Added: The `credentials` argument can now be specified in the configuration file to pass in either the path to or the contents of a service account key file in JSON format ([#318](https://github.com/turbot/steampipe-plugin-gcp/pull/318))

_Enhancements_

- Updated: Recompiled plugin with [steampipe-plugin-sdk v1.7.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md) ([#317](https://github.com/turbot/steampipe-plugin-gcp/pull/317))
- Updated: Add additional optional key quals to `gcp_iam_role`, `gcp_compute_image` and `gcp_compute_instance` tables ([#298](https://github.com/turbot/steampipe-plugin-gcp/pull/298))
- Updated: Add `WithCache` functionality for hydrate functions across all the tables ([#298](https://github.com/turbot/steampipe-plugin-gcp/pull/298))

_Deprecated_

- The `credential_file` argument in the configuration file is now deprecated and will be removed in the next major version. We recommend using the `credentials` argument instead, which can take the same file path as the `credential_file` argument. ([#318](https://github.com/turbot/steampipe-plugin-gcp/pull/318))

## v0.17.1 [2021-10-12]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v162--2021-10-08) ([#313](https://github.com/turbot/steampipe-plugin-gcp/pull/313))
- `credential_file` parameter in the configuration file now extends tildes to the current users home directory ([#312](https://github.com/turbot/steampipe-plugin-gcp/pull/312))

## v0.17.0 [2021-09-23]

_Enhancements_

- Added the `self_link` column to `gcp_bigtable_instance`, `gcp_cloudfunctions_function`, `gcp_monitoring_notification_channel`, `gcp_pubsub_snapshot`, `gcp_pubsub_subscription`, `gcp_pubsub_topic` and `gcp_project` tables ([#304](https://github.com/turbot/steampipe-plugin-gcp/pull/304))
- Added the `labels` column to `gcp_compute_instance_template` and `gcp_kms_key` tables ([#301](https://github.com/turbot/steampipe-plugin-gcp/pull/301))
- Recompiled plugin with [steampipe-plugin-sdk v1.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v161--2021-09-21) ([#309](https://github.com/turbot/steampipe-plugin-gcp/pull/309))

_Bug fixes_

- Column name `selfLink` has been correctly renamed to `self_link` in the `gcp_kubernetes_node_pool` table ([#308](https://github.com/turbot/steampipe-plugin-gcp/pull/308))

## v0.16.1 [2021-09-13]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.5.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v151--2021-09-13) ([#303](https://github.com/turbot/steampipe-plugin-gcp/pull/303))

_Bug fixes_

- If the Google service API is disabled, then the corresponding table will now return zero records ([#287](https://github.com/turbot/steampipe-plugin-gcp/pull/287))

## v0.16.0 [2021-08-26]

_What's new?_

- New tables added
  - [gcp_kubernetes_cluster](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_kubernetes_cluster) ([#213](https://github.com/turbot/steampipe-plugin-gcp/pull/213))
  - [gcp_kubernetes_node_pool](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_kubernetes_node_pool) ([#214](https://github.com/turbot/steampipe-plugin-gcp/pull/214))

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.5.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v150--2021-08-06) ([#290](https://github.com/turbot/steampipe-plugin-gcp/pull/290))

## v0.15.0 [2021-07-08]

_What's new?_

- New tables added
  - [gcp_sql_database_instance_metric_connections](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_sql_database_instance_metric_connections) ([#267](https://github.com/turbot/steampipe-plugin-gcp/pull/267))
  - [gcp_sql_database_instance_metric_connections_hourly](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_sql_database_instance_metric_connections_hourly) ([#268](https://github.com/turbot/steampipe-plugin-gcp/pull/268))

_Enhancements_

- Updated: Add columns `location` and `project` to all metric tables ([#279](https://github.com/turbot/steampipe-plugin-gcp/pull/279))

_Bug fixes_

- Fixed: `timestamp` column type in all metric tables is now `TIMESTAMP` instead of `STRING` ([#277](https://github.com/turbot/steampipe-plugin-gcp/pull/277))
- Fixed: Typo in table name and invalid key column definition for `gcp_dns_record_set table`

## v0.14.0 [2021-07-01]

_What's new?_

- New tables added
  - [gcp_bigquery_job](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_bigquery_job) ([#247](https://github.com/turbot/steampipe-plugin-gcp/pull/247))
  - [gcp_compute_disk_metric_read_ops](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_compute_disk_metric_read_ops) ([#262](https://github.com/turbot/steampipe-plugin-gcp/pull/262))
  - [gcp_compute_disk_metric_read_ops_daily](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_compute_disk_metric_read_ops_daily) ([#264](https://github.com/turbot/steampipe-plugin-gcp/pull/264))
  - [gcp_compute_disk_metric_read_ops_hourly](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_compute_disk_metric_read_ops_hourly) ([#260](https://github.com/turbot/steampipe-plugin-gcp/pull/260))
  - [gcp_compute_disk_metric_write_ops](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_compute_disk_metric_write_ops) ([#272](https://github.com/turbot/steampipe-plugin-gcp/pull/272))
  - [gcp_compute_disk_metric_write_ops_daily](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_compute_disk_metric_write_ops_daily) ([#263](https://github.com/turbot/steampipe-plugin-gcp/pull/263))
  - [gcp_compute_disk_metric_write_ops_hourly](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_compute_disk_metric_write_ops_hourly) ([#271](https://github.com/turbot/steampipe-plugin-gcp/pull/271))
  - [gcp_compute_instance_metric_cpu_utilization](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_compute_instance_metric_cpu_utilization) ([#254](https://github.com/turbot/steampipe-plugin-gcp/pull/254))
  - [gcp_compute_instance_metric_cpu_utilization_daily](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_compute_instance_metric_cpu_utilization_daily) ([#256](https://github.com/turbot/steampipe-plugin-gcp/pull/256))
  - [gcp_compute_instance_metric_cpu_utilization_hourly](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_compute_instance_metric_cpu_utilization_hourly) ([#255](https://github.com/turbot/steampipe-plugin-gcp/pull/255))
  - [gcp_compute_machine_type](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_compute_machine_type) ([#238](https://github.com/turbot/steampipe-plugin-gcp/pull/238))
  - [gcp_compute_resource_policy](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_compute_resource_policy) ([#237](https://github.com/turbot/steampipe-plugin-gcp/pull/237))
  - [gcp_dns_record_set](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_dns_record_set) ([#212](https://github.com/turbot/steampipe-plugin-gcp/pull/212))
  - [gcp_logging_bucket](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_logging_bucket) ([#232](https://github.com/turbot/steampipe-plugin-gcp/pull/232))
  - [gcp_organization](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_organization) ([#250](https://github.com/turbot/steampipe-plugin-gcp/pull/250))
  - [gcp_project_organization_policy](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_project_organization_policy) ([#204](https://github.com/turbot/steampipe-plugin-gcp/pull/204))
  - [gcp_sql_database_instance_metric_connections_daily](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_sql_database_instance_metric_connections_daily) ([#240](https://github.com/turbot/steampipe-plugin-gcp/pull/240))
  - [gcp_sql_database_instance_metric_cpu_utilization](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_sql_database_instance_metric_cpu_utilization) ([#259](https://github.com/turbot/steampipe-plugin-gcp/pull/259))
  - [gcp_sql_database_instance_metric_cpu_utilization_daily](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_sql_database_instance_metric_cpu_utilization_daily) ([#251](https://github.com/turbot/steampipe-plugin-gcp/pull/251))
  - [gcp_sql_database_instance_metric_cpu_utilization_hourly](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_sql_database_instance_metric_cpu_utilization_hourly) ([#261](https://github.com/turbot/steampipe-plugin-gcp/pull/261))

_Enhancements_

- Updated: Add support for service account impersonation credentials ([#249](https://github.com/turbot/steampipe-plugin-gcp/pull/249))

_Bug fixes_

- Fixed: List and get calls should not error in `gcp_sql_database` table for stopped databases ([#244](https://github.com/turbot/steampipe-plugin-gcp/pull/244))

## v0.13.1 [2021-06-17]

_Bug fixes_

- Fixed: `id` column type in `gcp_compute_disk` table ([#228](https://github.com/turbot/steampipe-plugin-gcp/pull/228))
- Fixed: Example query in `gcp_compute_address` table doc ([#224](https://github.com/turbot/steampipe-plugin-gcp/pull/224))

## v0.13.0 [2021-05-27]

_What's new?_

- Updated plugin license to Apache 2.0 per [turbot/steampipe#488](https://github.com/turbot/steampipe/issues/488)
- New tables added
  - [gcp_kms_key](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_kms_key) ([#215](https://github.com/turbot/steampipe-plugin-gcp/pull/215))

## v0.12.0 [2021-05-20]

_What's new?_

- New tables added
  - [gcp_compute_project_metadata](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_compute_project_metadata) ([#192](https://github.com/turbot/steampipe-plugin-gcp/pull/192))
  - [gcp_dns_policy](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_dns_policy) ([#208](https://github.com/turbot/steampipe-plugin-gcp/pull/208))
  - [gcp_project](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_project) ([#205](https://github.com/turbot/steampipe-plugin-gcp/pull/205))

## v0.11.0 [2021-05-13]

_What's new?_

- New tables added
  - [gcp_bigquery_table](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_bigquery_table) ([#198](https://github.com/turbot/steampipe-plugin-gcp/pull/198))

_Enhancements_

- Updated: README.md and docs/index.md now contain links to our Slack community ([#202](https://github.com/turbot/steampipe-plugin-gcp/pull/202))
- Updated: Add `self_link` column to `gcp_dns_managed_zone` table ([#195](https://github.com/turbot/steampipe-plugin-gcp/pull/195))
- Updated: Add `self_link` column to `gcp_logging_sink` table ([#193](https://github.com/turbot/steampipe-plugin-gcp/pull/193))
- Updated: Bump lodash from 4.17.20 to 4.17.21 in /gcp-test ([#197](https://github.com/turbot/steampipe-plugin-gcp/pull/197))

## v0.10.0 [2021-05-06]

_What's new?_

_Enhancements_

- Updated: Add `disk_encryption_key` and `disk_encryption_key_type` columns to `gcp_compute_disk` table ([#185](https://github.com/turbot/steampipe-plugin-gcp/pull/185))
- Updated: Remove `disk_encryption_kms_key` column from `gcp_compute_disk` table ([#185](https://github.com/turbot/steampipe-plugin-gcp/pull/185))
- Updated: Add `metric_descriptor_type` column to `gcp_logging_metric` table ([#182](https://github.com/turbot/steampipe-plugin-gcp/pull/182))

## v0.9.0 [2021-04-29]

_What's new?_

- New tables added
  - [gcp_compute_target_ssl_proxy](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_compute_target_ssl_proxy) ([#156](https://github.com/turbot/steampipe-plugin-gcp/pull/156))

## v0.8.0 [2021-04-22]

_What's new?_

- New tables added
  - [gcp_kms_key_ring](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_kms_key_ring) ([#171](https://github.com/turbot/steampipe-plugin-gcp/pull/171))

## v0.7.0 [2021-04-15]

_What's new?_

- New tables added
  - [gcp_compute_ssl_policy](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_compute_ssl_policy) ([#154](https://github.com/turbot/steampipe-plugin-gcp/pull/154))
  - [gcp_monitoring_alert_policy](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_monitoring_alert_policy) ([#168](https://github.com/turbot/steampipe-plugin-gcp/pull/168))

_Bug fixes_

- Fixed: Query example in `gcp_dns_managed_zone` table docs ([#172](https://github.com/turbot/steampipe-plugin-gcp/pull/172))

## v0.6.0 [2021-04-08]

_What's new?_

- New tables added
  - [gcp_compute_target_https_proxy](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_compute_target_https_proxy) ([#151](https://github.com/turbot/steampipe-plugin-gcp/pull/151))

_Enhancements_

- Updated: `gcp_sql_backup`, `gcp_sql_database`, `gcp_sql_database_instance` tables now use the `sqladmin` package instead of the `sql` package ([#161](https://github.com/turbot/steampipe-plugin-gcp/pull/161))
- Recompiled plugin with [steampipe-plugin-sdk v0.2.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v027-2021-03-31)

_Bug fixes_

- Fixed: Backup configuration columns now return the correct data in the `gcp_sql_database_instance` table ([#166](https://github.com/turbot/steampipe-plugin-gcp/pull/166))
- Fixed: Removed unused `root_password` column in the `gcp_sql_database_instance` table ([#166](https://github.com/turbot/steampipe-plugin-gcp/pull/166))

## v0.5.1 [2021-04-02]

_Bug fixes_

- Fixed: `Table definitions & examples` link now points to the correct location ([#163](https://github.com/turbot/steampipe-plugin-gcp/pull/163))

## v0.5.0 [2021-04-01]

_What's new?_

- New tables added
  - [gcp_bigquery_dataset](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_bigquery_dataset) ([#146](https://github.com/turbot/steampipe-plugin-gcp/pull/146))
  - [gcp_dns_managed_zone](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_dns_managed_zone) ([#147](https://github.com/turbot/steampipe-plugin-gcp/pull/147))

## v0.4.0 [2021-03-25]

_What's new?_

- New tables added
  - [gcp_bigtable_instance](https://hub.steampipe.io/plugins/turbot/gcp/tables/gcp_bigtable_instance) ([#90](https://github.com/turbot/steampipe-plugin-gcp/pull/90))

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v0.2.6](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v026-2021-03-18)

## v0.3.1 [2021-03-11]

_Bug fixes_

- Removed use of deprecated `ItemFromKey` function from all tables

## v0.3.0 [2021-03-04]

_What's new?_

- New tables added
  - gcp_compute_region
  - gcp_compute_zone

## v0.2.1 [2021-03-02]

_Bug fixes_

- Recompiled plugin with latest [steampipe-plugin-sdk](https://github.com/turbot/steampipe-plugin-sdk) to resolve issue:
  - Fix tables failing with error similar to `Error: pq: rpc error: code = Internal desc = get hydrate function getIamRole failed with panic interface conversion: interface {} is nil, not *gcp.roleInfo`([#129](https://github.com/turbot/steampipe-plugin-gcp/issues/129)).

## v0.2.0 [2021-02-25]

_What's new?_

- New tables added
  - gcp_sql_backup
  - gcp_sql_database
  - gcp_sql_database_instance

_Bug fixes_

- Updated `gcp_compute_instance` table `network_tags` field to display value correctly ([#114](https://github.com/turbot/steampipe-plugin-gcp/pull/114))
- Recompiled plugin with latest [steampipe-plugin-sdk](https://github.com/turbot/steampipe-plugin-sdk) to resolve SDK issues:
  - Fix error for missing required quals [#40](https://github.com/turbot/steampipe-plugin-sdk/issues/42).
  - Queries fail with error socket: too many open files [#190](https://github.com/turbot/steampipe/issues/190)

## v0.1.1 [2021-02-22]

_Bug fixes_

- Now union query for multiple projects will work if different credential files are used for project connection in config ([#116](https://github.com/turbot/steampipe-plugin-gcp/issues/116))

- Updated `gcp_storage_bucket` table `labels` field to display value correctly ([#115](https://github.com/turbot/steampipe-plugin-gcp/issues/115))

## v0.1.0 [2021-02-18]

_What's new?_

- Added support for [connection configuration](https://github.com/turbot/steampipe-plugin-gcp/blob/main/docs/index.md#connection-configuration). You may specify gcp `project` and `credential_file` for each connection in a configuration file. You can have multiple gcp connections, each configured for a different gcp project.

- If the project id to connect to is not specified in connection configuration file or through `CLOUDSDK_CORE_PROJECT` environment variable. Now plugin will use active project, as returned by the `gcloud config get-value project` command.

_Enhancements_

- Added `location` column to `gcp_compute_image`, `gcp_compute_snapshot` and `gcp_monitoring_notification_channel`, `gcp_pubsub_snapshot`, `gcp_pubsub_subscription` and `gcp_pubsub_topic` tables.
- Added `iamPolicy` column to `gcp_compute_instance` table.
- Added `disabled` and `oauth2_client_id` columns to gcp_service_account table.

## v0.0.6 [2021-02-11]

_What's new?_

- New tables added to plugin

  - [gcp_compute_backend_bucket](https://github.com/turbot/steampipe-plugin-gcp/blob/main/docs/tables/gcp_compute_backend_bucket.md) ([#70](https://github.com/turbot/steampipe-plugin-gcp/issues/70))
  - [gcp_compute_backend_service](https://github.com/turbot/steampipe-plugin-gcp/blob/main/docs/tables/gcp_compute_backend_service.md) ([#80](https://github.com/turbot/steampipe-plugin-gcp/issues/80))
  - [gcp_compute_image](https://github.com/turbot/steampipe-plugin-gcp/blob/main/docs/tables/gcp_compute_image.md) ([#45](https://github.com/turbot/steampipe-plugin-gcp/issues/45))
  - [gcp_compute_instance_template](https://github.com/turbot/steampipe-plugin-gcp/blob/main/docs/tables/gcp_compute_instance_template.md) ([#84](https://github.com/turbot/steampipe-plugin-gcp/issues/84))
  - [gcp_compute_node_group](https://github.com/turbot/steampipe-plugin-gcp/blob/main/docs/tables/gcp_compute_node_group.md) ([#58](https://github.com/turbot/steampipe-plugin-gcp/issues/58))
  - [gcp_compute_node_template](https://github.com/turbot/steampipe-plugin-gcp/blob/main/docs/tables/gcp_compute_node_template.md) ([#87](https://github.com/turbot/steampipe-plugin-gcp/issues/87))
  - [gcp_compute_subnetwork](https://github.com/turbot/steampipe-plugin-gcp/blob/main/docs/tables/gcp_compute_subnetwork.md) ([#68](https://github.com/turbot/steampipe-plugin-gcp/issues/68))
  - [gcp_compute_target_pool](https://github.com/turbot/steampipe-plugin-gcp/blob/main/docs/tables/gcp_compute_target_pool.md) ([#81](https://github.com/turbot/steampipe-plugin-gcp/issues/81))
  - [gcp_compute_target_vpn_gateway](https://github.com/turbot/steampipe-plugin-gcp/blob/main/docs/tables/gcp_compute_target_vpn_gateway.md) ([#65](https://github.com/turbot/steampipe-plugin-gcp/issues/65))
  - [gcp_compute_url_map](https://github.com/turbot/steampipe-plugin-gcp/blob/main/docs/tables/gcp_compute_url_map.md) ([#85](https://github.com/turbot/steampipe-plugin-gcp/issues/85))
  - [gcp_compute_vpn_tunnel](https://github.com/turbot/steampipe-plugin-gcp/blob/main/docs/tables/gcp_compute_vpn_tunnel.md) ([#63](https://github.com/turbot/steampipe-plugin-gcp/issues/63))

## v0.0.5 [2021-02-04]

_What's new?_

- New tables added to plugin

  - gcp_compute_address ([#29](https://github.com/turbot/steampipe-plugin-gcp/issues/29))
  - gcp_compute_disk ([#47](https://github.com/turbot/steampipe-plugin-gcp/issues/47))
  - gcp_compute_firewall ([#42](https://github.com/turbot/steampipe-plugin-gcp/issues/42))
  - gcp_compute_forwarding_rule ([#53](https://github.com/turbot/steampipe-plugin-gcp/issues/53))
  - gcp_compute_network ([#43](https://github.com/turbot/steampipe-plugin-gcp/issues/43))
  - gcp_compute_router ([#51](https://github.com/turbot/steampipe-plugin-gcp/issues/51))
  - gcp_compute_snapshot ([#60](https://github.com/turbot/steampipe-plugin-gcp/issues/60))

_Enhancements_

- Added field `location` to resource tables that are not regional with value as `global`

## v0.0.4 [2021-01-28]

_What's new?_

- Added: `gcp_cloudfunctions_function` table
- Added: `gcp_compute_global_address` table
- Added: `gcp_compute_global_forwarding_rule` table
- Added: `gcp_compute_instance` table
- Added: `gcp_storage_bucket` table

- Updated: `gcp_iam_role` table. Added `is_gcp_managed` boolean field to distinguish between GCP managed and Customer managed roles.

_Bug fixes_

- Fixed: `gcp_iam_role` table. Updated `included_permissions` field to have details of role grants for list call.
