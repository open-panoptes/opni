// Code generated by internal/codegen. DO NOT EDIT.

// @generated by protoc-gen-es v1.3.0 with parameter "target=ts,import_extension=none,ts_nocheck=false"
// @generated from file github.com/rancher/opni/internal/cortex/config/validation/limits.proto (package validation, syntax proto3)
/* eslint-disable */

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Duration, Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message validation.Limits
 */
export class Limits extends Message<Limits> {
  /**
   * Per-user ingestion rate limit in samples per second.
   *
   * @generated from field: optional double ingestion_rate = 1;
   */
  ingestionRate?: number;

  /**
   * Whether the ingestion rate limit should be applied individually to each distributor instance (local), or evenly shared across the cluster (global).
   *
   * @generated from field: optional string ingestion_rate_strategy = 2;
   */
  ingestionRateStrategy?: string;

  /**
   * Per-user allowed ingestion burst size (in number of samples).
   *
   * @generated from field: optional int32 ingestion_burst_size = 3;
   */
  ingestionBurstSize?: number;

  /**
   * Flag to enable, for all users, handling of samples with external labels identifying replicas in an HA Prometheus setup.
   *
   * @generated from field: optional bool accept_ha_samples = 4;
   */
  acceptHaSamples?: boolean;

  /**
   * Prometheus label to look for in samples to identify a Prometheus HA cluster.
   *
   * @generated from field: optional string ha_cluster_label = 5;
   */
  haClusterLabel?: string;

  /**
   * Prometheus label to look for in samples to identify a Prometheus HA replica.
   *
   * @generated from field: optional string ha_replica_label = 6;
   */
  haReplicaLabel?: string;

  /**
   * Maximum number of clusters that HA tracker will keep track of for single user. 0 to disable the limit.
   *
   * @generated from field: optional int32 ha_max_clusters = 7;
   */
  haMaxClusters?: number;

  /**
   * This flag can be used to specify label names that to drop during sample ingestion within the distributor and can be repeated in order to drop multiple labels.
   *
   * @generated from field: repeated string drop_labels = 8;
   */
  dropLabels: string[] = [];

  /**
   * Maximum length accepted for label names
   *
   * @generated from field: optional int32 max_label_name_length = 9;
   */
  maxLabelNameLength?: number;

  /**
   * Maximum length accepted for label value. This setting also applies to the metric name
   *
   * @generated from field: optional int32 max_label_value_length = 10;
   */
  maxLabelValueLength?: number;

  /**
   * Maximum number of label names per series.
   *
   * @generated from field: optional int32 max_label_names_per_series = 11;
   */
  maxLabelNamesPerSeries?: number;

  /**
   * Maximum combined size in bytes of all labels and label values accepted for a series. 0 to disable the limit.
   *
   * @generated from field: optional int32 max_labels_size_bytes = 12;
   */
  maxLabelsSizeBytes?: number;

  /**
   * Maximum length accepted for metric metadata. Metadata refers to Metric Name, HELP and UNIT.
   *
   * @generated from field: optional int32 max_metadata_length = 13;
   */
  maxMetadataLength?: number;

  /**
   * Reject old samples.
   *
   * @generated from field: optional bool reject_old_samples = 14;
   */
  rejectOldSamples?: boolean;

  /**
   * Maximum accepted sample age before rejecting.
   *
   * @generated from field: google.protobuf.Duration reject_old_samples_max_age = 15;
   */
  rejectOldSamplesMaxAge?: Duration;

  /**
   * Duration which table will be created/deleted before/after it's needed; we won't accept sample from before this time.
   *
   * @generated from field: google.protobuf.Duration creation_grace_period = 16;
   */
  creationGracePeriod?: Duration;

  /**
   * Enforce every metadata has a metric name.
   *
   * @generated from field: optional bool enforce_metadata_metric_name = 17;
   */
  enforceMetadataMetricName?: boolean;

  /**
   * Enforce every sample has a metric name.
   *
   * @generated from field: optional bool enforce_metric_name = 18;
   */
  enforceMetricName?: boolean;

  /**
   * The default tenant's shard size when the shuffle-sharding strategy is used. Must be set both on ingesters and distributors. When this setting is specified in the per-tenant overrides, a value of 0 disables shuffle sharding for the tenant.
   *
   * @generated from field: optional int32 ingestion_tenant_shard_size = 19;
   */
  ingestionTenantShardSize?: number;

  /**
   * List of metric relabel configurations. Note that in most situations, it is more effective to use metrics relabeling directly in the Prometheus server, e.g. remote_write.write_relabel_configs.
   *
   * @generated from field: repeated validation.RelabelConfig metric_relabel_configs = 20;
   */
  metricRelabelConfigs: RelabelConfig[] = [];

  /**
   * Enables support for exemplars in TSDB and sets the maximum number that will be stored. less than zero means disabled. If the value is set to zero, cortex will fallback to blocks-storage.tsdb.max-exemplars value.
   *
   * @generated from field: optional int32 max_exemplars = 21;
   */
  maxExemplars?: number;

  /**
   * The maximum number of series for which a query can fetch samples from each ingester. This limit is enforced only in the ingesters (when querying samples not flushed to the storage yet) and it's a per-instance limit. This limit is ignored when running the Cortex blocks storage. When running Cortex with blocks storage use -querier.max-fetched-series-per-query limit instead.
   *
   * @generated from field: optional int32 max_series_per_query = 22;
   */
  maxSeriesPerQuery?: number;

  /**
   * The maximum number of active series per user, per ingester. 0 to disable.
   *
   * @generated from field: optional int32 max_series_per_user = 23;
   */
  maxSeriesPerUser?: number;

  /**
   * The maximum number of active series per metric name, per ingester. 0 to disable.
   *
   * @generated from field: optional int32 max_series_per_metric = 24;
   */
  maxSeriesPerMetric?: number;

  /**
   * The maximum number of active series per user, across the cluster before replication. 0 to disable. Supported only if -distributor.shard-by-all-labels is true.
   *
   * @generated from field: optional int32 max_global_series_per_user = 25;
   */
  maxGlobalSeriesPerUser?: number;

  /**
   * The maximum number of active series per metric name, across the cluster before replication. 0 to disable.
   *
   * @generated from field: optional int32 max_global_series_per_metric = 26;
   */
  maxGlobalSeriesPerMetric?: number;

  /**
   * The maximum number of active metrics with metadata per user, per ingester. 0 to disable.
   *
   * @generated from field: optional int32 max_metadata_per_user = 27;
   */
  maxMetadataPerUser?: number;

  /**
   * The maximum number of metadata per metric, per ingester. 0 to disable.
   *
   * @generated from field: optional int32 max_metadata_per_metric = 28;
   */
  maxMetadataPerMetric?: number;

  /**
   * The maximum number of active metrics with metadata per user, across the cluster. 0 to disable. Supported only if -distributor.shard-by-all-labels is true.
   *
   * @generated from field: optional int32 max_global_metadata_per_user = 29;
   */
  maxGlobalMetadataPerUser?: number;

  /**
   * The maximum number of metadata per metric, across the cluster. 0 to disable.
   *
   * @generated from field: optional int32 max_global_metadata_per_metric = 30;
   */
  maxGlobalMetadataPerMetric?: number;

  /**
   * [Experimental] Configures the allowed time window for ingestion of out-of-order samples. Disabled (0s) by default.
   *
   * @generated from field: google.protobuf.Duration out_of_order_time_window = 31;
   */
  outOfOrderTimeWindow?: Duration;

  /**
   * Maximum number of chunks that can be fetched in a single query from ingesters and long-term storage. This limit is enforced in the querier, ruler and store-gateway. 0 to disable.
   *
   * @generated from field: optional int32 max_fetched_chunks_per_query = 32;
   */
  maxFetchedChunksPerQuery?: number;

  /**
   * The maximum number of unique series for which a query can fetch samples from each ingesters and blocks storage. This limit is enforced in the querier, ruler and store-gateway. 0 to disable
   *
   * @generated from field: optional int32 max_fetched_series_per_query = 33;
   */
  maxFetchedSeriesPerQuery?: number;

  /**
   * Deprecated (use max-fetched-data-bytes-per-query instead): The maximum size of all chunks in bytes that a query can fetch from each ingester and storage. This limit is enforced in the querier, ruler and store-gateway. 0 to disable.
   *
   * @generated from field: optional int32 max_fetched_chunk_bytes_per_query = 34;
   */
  maxFetchedChunkBytesPerQuery?: number;

  /**
   * The maximum combined size of all data that a query can fetch from each ingester and storage. This limit is enforced in the querier and ruler for `query`, `query_range` and `series` APIs. 0 to disable.
   *
   * @generated from field: optional int32 max_fetched_data_bytes_per_query = 35;
   */
  maxFetchedDataBytesPerQuery?: number;

  /**
   * Limit how long back data (series and metadata) can be queried, up until <lookback> duration ago. This limit is enforced in the query-frontend, querier and ruler. If the requested time range is outside the allowed range, the request will not fail but will be manipulated to only query data within the allowed time range. 0 to disable.
   *
   * @generated from field: google.protobuf.Duration max_query_lookback = 36;
   */
  maxQueryLookback?: Duration;

  /**
   * Limit the query time range (end - start time). This limit is enforced in the query-frontend (on the received query) and in the querier (on the query possibly split by the query-frontend). 0 to disable.
   *
   * @generated from field: google.protobuf.Duration max_query_length = 37;
   */
  maxQueryLength?: Duration;

  /**
   * Maximum number of split queries will be scheduled in parallel by the frontend.
   *
   * @generated from field: optional int32 max_query_parallelism = 38;
   */
  maxQueryParallelism?: number;

  /**
   * Most recent allowed cacheable result per-tenant, to prevent caching very recent results that might still be in flux.
   *
   * @generated from field: google.protobuf.Duration max_cache_freshness = 39;
   */
  maxCacheFreshness?: Duration;

  /**
   * Maximum number of queriers that can handle requests for a single tenant. If set to 0 or value higher than number of available queriers, *all* queriers will handle requests for the tenant. If the value is < 1, it will be treated as a percentage and the gets a percentage of the total queriers. Each frontend (or query-scheduler, if used) will select the same set of queriers for the same tenant (given that all queriers are connected to all frontends / query-schedulers). This option only works with queriers connecting to the query-frontend / query-scheduler, not when using downstream URL.
   *
   * @generated from field: optional double max_queriers_per_tenant = 40;
   */
  maxQueriersPerTenant?: number;

  /**
   * Maximum number of outstanding requests per tenant per request queue (either query frontend or query scheduler); requests beyond this error with HTTP 429.
   *
   * @generated from field: optional int32 max_outstanding_requests_per_tenant = 41;
   */
  maxOutstandingRequestsPerTenant?: number;

  /**
   * Duration to delay the evaluation of rules to ensure the underlying metrics have been pushed to Cortex.
   *
   * @generated from field: google.protobuf.Duration ruler_evaluation_delay_duration = 42;
   */
  rulerEvaluationDelayDuration?: Duration;

  /**
   * The default tenant's shard size when the shuffle-sharding strategy is used by ruler. When this setting is specified in the per-tenant overrides, a value of 0 disables shuffle sharding for the tenant.
   *
   * @generated from field: optional int32 ruler_tenant_shard_size = 43;
   */
  rulerTenantShardSize?: number;

  /**
   * Maximum number of rules per rule group per-tenant. 0 to disable.
   *
   * @generated from field: optional int32 ruler_max_rules_per_rule_group = 44;
   */
  rulerMaxRulesPerRuleGroup?: number;

  /**
   * Maximum number of rule groups per-tenant. 0 to disable.
   *
   * @generated from field: optional int32 ruler_max_rule_groups_per_tenant = 45;
   */
  rulerMaxRuleGroupsPerTenant?: number;

  /**
   * The default tenant's shard size when the shuffle-sharding strategy is used. Must be set when the store-gateway sharding is enabled with the shuffle-sharding strategy. When this setting is specified in the per-tenant overrides, a value of 0 disables shuffle sharding for the tenant. If the value is < 1 the shard size will be a percentage of the total store-gateways.
   *
   * @generated from field: optional double store_gateway_tenant_shard_size = 46;
   */
  storeGatewayTenantShardSize?: number;

  /**
   * The maximum number of data bytes to download per gRPC request in Store Gateway, including Series/LabelNames/LabelValues requests. 0 to disable.
   *
   * @generated from field: optional int32 max_downloaded_bytes_per_request = 47;
   */
  maxDownloadedBytesPerRequest?: number;

  /**
   * Delete blocks containing samples older than the specified retention period. 0 to disable.
   *
   * @generated from field: google.protobuf.Duration compactor_blocks_retention_period = 48;
   */
  compactorBlocksRetentionPeriod?: Duration;

  /**
   * The default tenant's shard size when the shuffle-sharding strategy is used by the compactor. When this setting is specified in the per-tenant overrides, a value of 0 disables shuffle sharding for the tenant.
   *
   * @generated from field: optional int32 compactor_tenant_shard_size = 49;
   */
  compactorTenantShardSize?: number;

  /**
   * S3 server-side encryption type. Required to enable server-side encryption overrides for a specific tenant. If not set, the default S3 client settings are used.
   *
   * @generated from field: optional string s3_sse_type = 50;
   */
  s3SseType?: string;

  /**
   * S3 server-side encryption KMS Key ID. Ignored if the SSE type override is not set.
   *
   * @generated from field: optional string s3_sse_kms_key_id = 51;
   */
  s3SseKmsKeyId?: string;

  /**
   * S3 server-side encryption KMS encryption context. If unset and the key ID override is set, the encryption context will not be provided to S3. Ignored if the SSE type override is not set.
   *
   * @generated from field: optional string s3_sse_kms_encryption_context = 52;
   */
  s3SseKmsEncryptionContext?: string;

  /**
   * Comma-separated list of network CIDRs to block in Alertmanager receiver integrations.
   *
   * @generated from field: repeated string alertmanager_receivers_firewall_block_cidr_networks = 53;
   */
  alertmanagerReceiversFirewallBlockCidrNetworks: string[] = [];

  /**
   * True to block private and local addresses in Alertmanager receiver integrations. It blocks private addresses defined by  RFC 1918 (IPv4 addresses) and RFC 4193 (IPv6 addresses), as well as loopback, local unicast and local multicast addresses.
   *
   * @generated from field: optional bool alertmanager_receivers_firewall_block_private_addresses = 54;
   */
  alertmanagerReceiversFirewallBlockPrivateAddresses?: boolean;

  /**
   * Per-user rate limit for sending notifications from Alertmanager in notifications/sec. 0 = rate limit disabled. Negative value = no notifications are allowed.
   *
   * @generated from field: optional double alertmanager_notification_rate_limit = 55;
   */
  alertmanagerNotificationRateLimit?: number;

  /**
   * Per-integration notification rate limits. Value is a map, where each key is integration name and value is a rate-limit (float). On command line, this map is given in JSON format. Rate limit has the same meaning as -alertmanager.notification-rate-limit, but only applies for specific integration. Allowed integration names: webhook, email, pagerduty, opsgenie, wechat, slack, victorops, pushover, sns, telegram, discord, webex.
   *
   * @generated from field: map<string, double> alertmanager_notification_rate_limit_per_integration = 56;
   */
  alertmanagerNotificationRateLimitPerIntegration: { [key: string]: number } = {};

  /**
   * Maximum size of configuration file for Alertmanager that tenant can upload via Alertmanager API. 0 = no limit.
   *
   * @generated from field: optional int32 alertmanager_max_config_size_bytes = 57;
   */
  alertmanagerMaxConfigSizeBytes?: number;

  /**
   * Maximum number of templates in tenant's Alertmanager configuration uploaded via Alertmanager API. 0 = no limit.
   *
   * @generated from field: optional int32 alertmanager_max_templates_count = 58;
   */
  alertmanagerMaxTemplatesCount?: number;

  /**
   * Maximum size of single template in tenant's Alertmanager configuration uploaded via Alertmanager API. 0 = no limit.
   *
   * @generated from field: optional int32 alertmanager_max_template_size_bytes = 59;
   */
  alertmanagerMaxTemplateSizeBytes?: number;

  /**
   * Maximum number of aggregation groups in Alertmanager's dispatcher that a tenant can have. Each active aggregation group uses single goroutine. When the limit is reached, dispatcher will not dispatch alerts that belong to additional aggregation groups, but existing groups will keep working properly. 0 = no limit.
   *
   * @generated from field: optional int32 alertmanager_max_dispatcher_aggregation_groups = 60;
   */
  alertmanagerMaxDispatcherAggregationGroups?: number;

  /**
   * Maximum number of alerts that a single user can have. Inserting more alerts will fail with a log message and metric increment. 0 = no limit.
   *
   * @generated from field: optional int32 alertmanager_max_alerts_count = 61;
   */
  alertmanagerMaxAlertsCount?: number;

  /**
   * Maximum total size of alerts that a single user can have, alert size is the sum of the bytes of its labels, annotations and generatorURL. Inserting more alerts will fail with a log message and metric increment. 0 = no limit.
   *
   * @generated from field: optional int32 alertmanager_max_alerts_size_bytes = 62;
   */
  alertmanagerMaxAlertsSizeBytes?: number;

  constructor(data?: PartialMessage<Limits>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "validation.Limits";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "ingestion_rate", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
    { no: 2, name: "ingestion_rate_strategy", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "ingestion_burst_size", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 4, name: "accept_ha_samples", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 5, name: "ha_cluster_label", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 6, name: "ha_replica_label", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 7, name: "ha_max_clusters", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 8, name: "drop_labels", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 9, name: "max_label_name_length", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 10, name: "max_label_value_length", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 11, name: "max_label_names_per_series", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 12, name: "max_labels_size_bytes", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 13, name: "max_metadata_length", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 14, name: "reject_old_samples", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 15, name: "reject_old_samples_max_age", kind: "message", T: Duration },
    { no: 16, name: "creation_grace_period", kind: "message", T: Duration },
    { no: 17, name: "enforce_metadata_metric_name", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 18, name: "enforce_metric_name", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 19, name: "ingestion_tenant_shard_size", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 20, name: "metric_relabel_configs", kind: "message", T: RelabelConfig, repeated: true },
    { no: 21, name: "max_exemplars", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 22, name: "max_series_per_query", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 23, name: "max_series_per_user", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 24, name: "max_series_per_metric", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 25, name: "max_global_series_per_user", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 26, name: "max_global_series_per_metric", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 27, name: "max_metadata_per_user", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 28, name: "max_metadata_per_metric", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 29, name: "max_global_metadata_per_user", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 30, name: "max_global_metadata_per_metric", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 31, name: "out_of_order_time_window", kind: "message", T: Duration },
    { no: 32, name: "max_fetched_chunks_per_query", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 33, name: "max_fetched_series_per_query", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 34, name: "max_fetched_chunk_bytes_per_query", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 35, name: "max_fetched_data_bytes_per_query", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 36, name: "max_query_lookback", kind: "message", T: Duration },
    { no: 37, name: "max_query_length", kind: "message", T: Duration },
    { no: 38, name: "max_query_parallelism", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 39, name: "max_cache_freshness", kind: "message", T: Duration },
    { no: 40, name: "max_queriers_per_tenant", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
    { no: 41, name: "max_outstanding_requests_per_tenant", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 42, name: "ruler_evaluation_delay_duration", kind: "message", T: Duration },
    { no: 43, name: "ruler_tenant_shard_size", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 44, name: "ruler_max_rules_per_rule_group", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 45, name: "ruler_max_rule_groups_per_tenant", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 46, name: "store_gateway_tenant_shard_size", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
    { no: 47, name: "max_downloaded_bytes_per_request", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 48, name: "compactor_blocks_retention_period", kind: "message", T: Duration },
    { no: 49, name: "compactor_tenant_shard_size", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 50, name: "s3_sse_type", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 51, name: "s3_sse_kms_key_id", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 52, name: "s3_sse_kms_encryption_context", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 53, name: "alertmanager_receivers_firewall_block_cidr_networks", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 54, name: "alertmanager_receivers_firewall_block_private_addresses", kind: "scalar", T: 8 /* ScalarType.BOOL */, opt: true },
    { no: 55, name: "alertmanager_notification_rate_limit", kind: "scalar", T: 1 /* ScalarType.DOUBLE */, opt: true },
    { no: 56, name: "alertmanager_notification_rate_limit_per_integration", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 1 /* ScalarType.DOUBLE */} },
    { no: 57, name: "alertmanager_max_config_size_bytes", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 58, name: "alertmanager_max_templates_count", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 59, name: "alertmanager_max_template_size_bytes", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 60, name: "alertmanager_max_dispatcher_aggregation_groups", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 61, name: "alertmanager_max_alerts_count", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
    { no: 62, name: "alertmanager_max_alerts_size_bytes", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Limits {
    return new Limits().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Limits {
    return new Limits().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Limits {
    return new Limits().fromJsonString(jsonString, options);
  }

  static equals(a: Limits | PlainMessage<Limits> | undefined, b: Limits | PlainMessage<Limits> | undefined): boolean {
    return proto3.util.equals(Limits, a, b);
  }
}

/**
 * @generated from message validation.RelabelConfig
 */
export class RelabelConfig extends Message<RelabelConfig> {
  /**
   * @generated from field: repeated string source_labels = 1;
   */
  sourceLabels: string[] = [];

  /**
   * @generated from field: optional string separator = 2;
   */
  separator?: string;

  /**
   * @generated from field: string regex = 3;
   */
  regex = "";

  /**
   * @generated from field: optional uint64 modulus = 4;
   */
  modulus?: bigint;

  /**
   * @generated from field: optional string target_label = 5;
   */
  targetLabel?: string;

  /**
   * @generated from field: optional string replacement = 6;
   */
  replacement?: string;

  /**
   * @generated from field: optional string action = 7;
   */
  action?: string;

  constructor(data?: PartialMessage<RelabelConfig>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "validation.RelabelConfig";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "source_labels", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "separator", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "regex", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "modulus", kind: "scalar", T: 4 /* ScalarType.UINT64 */, opt: true },
    { no: 5, name: "target_label", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 6, name: "replacement", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 7, name: "action", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RelabelConfig {
    return new RelabelConfig().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RelabelConfig {
    return new RelabelConfig().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RelabelConfig {
    return new RelabelConfig().fromJsonString(jsonString, options);
  }

  static equals(a: RelabelConfig | PlainMessage<RelabelConfig> | undefined, b: RelabelConfig | PlainMessage<RelabelConfig> | undefined): boolean {
    return proto3.util.equals(RelabelConfig, a, b);
  }
}

