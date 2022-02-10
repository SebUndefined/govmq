# govmq

govmq is a type and builder helper for the usage of [VerneMQ MQTT Broker](https://vernemq.com, "VerneMQ Homepage").

- [Installation](#installation)
- [Usage](#usage)
  - [Webhooks](#webhooks)
  - [Modifiers](#modifiers)

# Installation 

Work in progress

# Usage

## Webhooks

|                      | Struct name | mountpoint | client_id | username | password | peer_addr | peer_port | clean_session | topics (array object) | topics (string) | topics (array string) | qos | payload | retain | clean_start | properties |
|----------------------|-------------|:----------:|:---------:|:--------:|:--------:|:---------:|:---------:|:-------------:|:---------------------:|:---------------:|:---------------------:|:---:|:-------:|:------:|:-----------:|:----------:|
| auth_on_register     |             |     X      |     X     |    X     |    X     |     X     |     X     |       X       |                       |                 |                       |     |         |        |             |            |
| auth_on_subscribe    |             |     X      |     X     |    X     |          |           |           |               |           X           |                 |                       |     |         |        |             |            |
| auth_on_publish      |             |     X      |     X     |    X     |          |           |           |               |                       |        X        |                       |  X  |    X    |   X    |             |            |
| on_register          |             |     X      |     X     |    X     |          |     X     |     X     |               |                       |                 |                       |     |         |        |             |            |
| on_publish           |             |     X      |     X     |    X     |          |           |           |               |                       |        X        |                       |  X  |    X    |   X    |             |            |
| on_subscribe         |             |     X      |     X     |    X     |          |           |           |               |           X           |                 |                       |     |         |        |             |            |
| on_unsubscribe       |             |     X      |     X     |    X     |          |           |           |               |                       |                 |           X           |     |         |        |             |            |
| on_deliver           |             |     X      |     X     |    X     |          |           |           |               |                       |        X        |                       |     |    X    |        |             |            |
| on_offline_message   |             |     X      |     X     |          |          |           |           |               |                       |        X        |                       |  X  |    X    |   X    |             |            |
| on_client_wakeup     |             |     X      |     X     |          |          |           |           |               |                       |                 |                       |     |         |        |             |            |
| on_client_offline    |             |     X      |     X     |          |          |           |           |               |                       |                 |                       |     |         |        |             |            |
| on_client_gone       |             |     X      |     X     |          |          |           |           |               |                       |                 |                       |     |         |        |             |            |
| auth_on_register_m5  |             |     X      |     X     |    X     |    X     |     X     |     X     |               |                       |                 |                       |     |         |        |      X      |     X      |
| on_auth_m5           |             |     X      |     X     |    X     |          |           |           |               |                       |                 |                       |     |         |        |             |     X      |
| auth_on_subscribe_m5 |             |     X      |     X     |    X     |          |           |           |               |           X           |                 |                       |     |         |        |             |     X      |
| auth_on_publish_m5   |             |     X      |     X     |    X     |          |           |           |               |                       |        X        |                       |  X  |    X    |   X    |             |     X      |
| on_register_m5       |             |     X      |     X     |    X     |          |     X     |     X     |               |                       |                 |                       |     |         |        |             |     X      |
| on_publish_m5        |             |     X      |     X     |          |          |           |           |               |           X           |                 |                       |  X  |    X    |   X    |      X      |     X      |
| on_subscribe_m5      |             |     X      |     X     |    X     |          |           |           |               |           X           |                 |                       |     |         |        |             |     X      |
| on_unsubscribe_m5    |             |     X      |     X     |    X     |          |           |           |               |                       |                 |           X           |     |         |        |             |     X      |
| on_deliver_m5        |             |     X      |     X     |    X     |          |           |           |               |                       |        X        |                       |     |    X    |        |             |     X      |



## Modifiers

| Modifier | Description | Type |  V1  | 
|:--------:|:-----------:|:----:|:----:|
|   Test   |    test     | test | test |


### auth_on_register

Potential modifiers

|          Modifier          |         Property         |  Type  | Description | 
|:--------------------------:|:------------------------:|:------:|:-----------:|
|       allow_register       |      AllowRegister       |  bool  |    test     |
|       allow_publish        |       AllowPublish       |  bool  |    test     |
|      allow_subscribe       |      AllowSubscribe      |  bool  |    test     |
|     allow_unsubscribe      |     AllowUnsubscribe     |  bool  |    test     |
|      max_message_size      |      MaxMessageSize      |  int   |    test     |
|       subscriber_id        |       SubscriberId       | string |    test     |
|          username          |         Username         | string |    test     |
|      max_message_rate      |      MaxMessageRate      |  int   |    test     |
|   max_inflight_messages    |   MaxInflightMessages    |  int   |    test     |
| shared_subscription_policy | SharedSubscriptionPolicy | string |    test     |
|        upgrade_qos         |        UpgradeQoS        |  bool  |    test     |
|  allow_multiple_sessions   |  AllowMultipleSessions   |  bool  |    test     |
|    max_online_messages     |    maxOnlineMessages     |  int   |    test     |
|    max_offline_messages    |    MaxOfflineMessages    |  int   |    test     |
|     queue_deliver_mode     |     QueueDeliverMode     | string |    test     |
|         queue_type         |        QueueType         | string |    test     |
|       max_drain_time       |       MaxDrainTime       |  int   |    test     |
|  max_msgs_per_drain_step   | MaxMessagesPerDrainStep  |  int   |    test     |

### auth_on_register_M5

Potential modifiers

|          Modifier          |         Property         |  Type  | Description | 
|:--------------------------:|:------------------------:|:------:|:-----------:|
|       allow_register       |      AllowRegister       |  bool  |    test     |
|       allow_publish        |       AllowPublish       |  bool  |    test     |
|      allow_subscribe       |      AllowSubscribe      |  bool  |    test     |
|     allow_unsubscribe      |     AllowUnsubscribe     |  bool  |    test     |
|      max_message_size      |      MaxMessageSize      |  int   |    test     |
|       subscriber_id        |       SubscriberId       | string |    test     |
|          username          |         Username         | string |    test     |
|      max_message_rate      |      MaxMessageRate      |  int   |    test     |
|   max_inflight_messages    |   MaxInflightMessages    |  int   |    test     |
| shared_subscription_policy | SharedSubscriptionPolicy | string |    test     |
|        upgrade_qos         |        UpgradeQoS        |  bool  |    test     |
|  allow_multiple_sessions   |  AllowMultipleSessions   |  bool  |    test     |
|    max_online_messages     |    maxOnlineMessages     |  int   |    test     |
|    max_offline_messages    |    MaxOfflineMessages    |  int   |    test     |
|     queue_deliver_mode     |     QueueDeliverMode     | string |    test     |
|         queue_type         |        QueueType         | string |    test     |
|       max_drain_time       |       MaxDrainTime       |  int   |    test     |
|  max_msgs_per_drain_step   | MaxMessagesPerDrainStep  |  int   |    test     |
|         properties         |        Properties        | object |    test     |

#### Properties auth on register M5

|         Modifier          |        Property         |  Type  | Description | 
|:-------------------------:|:-----------------------:|:------:|:-----------:|
|      p_user_property      |      UserProperty       | string |    test     |
| p_session_expiry_internal |  SessionExpiryInternal  |  int   |    test     |

### auth_on_publish

Potential modifiers

|  Modifier  |  Property  |  Type  | Description | 
|:----------:|:----------:|:------:|:-----------:|
|   topic    |   Topic    | string |    test     |
|  payload   |  Payload   | string |    test     |
|    qos     |    QoS     |  int   |    test     |
|   retain   |   Retain   |  bool  |    test     |
| mountpoint | Mountpoint | string |    test     |
|  throttle  |  Throttle  |  int   |    test     |

### auth_on_publish_m5

Potential modifiers

|  Modifier  |  Property  |  Type  | Description | 
|:----------:|:----------:|:------:|:-----------:|
|   topic    |   Topic    | string |    test     |
|  payload   |  Payload   | string |    test     |
|    qos     |    QoS     |  int   |    test     |
|   retain   |   Retain   |  bool  |    test     |
| mountpoint | Mountpoint | string |    test     |
|  throttle  |  Throttle  |  int   |    test     |
| properties | Properties | object |    test     |

#### Properties auth on publish M5

|          Modifier          |        Property        |  Type  | Description | 
|:--------------------------:|:----------------------:|:------:|:-----------:|
|      p_user_property       |      UserProperty      | string |    test     |
| p_session_expiry_internal  | SessionExpiryInternal  |  int   |    test     |
|       p_content_type       |      ContentType       | string |    test     |
| p_payload_format_indicator | PayloadFormatIndicator | string |    test     |
|      p_response_topic      |     ResponseTopic      | string |    test     |
|     p_correlation_data     |    CorrelationData     | string |    test     |

