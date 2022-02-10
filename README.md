# govmq

govmq is a type and builder helper for the usage of [VerneMQ MQTT Broker](https://vernemq.com, "VerneMQ Homepage").

- [Installation](#installation)
- [Usage](#usage)
  - [Webhooks](#webhooks)
    - [auth_on_register](#auth_on_register)
    - [auth_on_subscribe](#auth_on_subscribe)
    - [auth_on_publish](#auth_on_publish)
    - [on_register](#on_register)
    - [on_publish](#on_publish)
    - [on_subscribe](#on_subscribe)
    - [on_unsubscribe](#on_unsubscribe)
    - [on_deliver](#on_deliver)
    - [on_offline_message](#on_offline_message)
    - [on_client_wakeup](#on_client_wakeup)
    - [on_client_offline](#on_client_offline)
    - [on_client_gone](#on_client_gone)
    - [auth_on_register_m5](#auth_on_register_m5)
    - [on_auth_m5](#on_auth_m5)
    - [auth_on_subscribe_m5](#auth_on_subscribe_m5)
    - [auth_on_publish_m5](#auth_on_publish_m5)
    - [on_register_m5](#on_register_m5)
    - [on_publish_m5](#on_publish_m5)
    - [on_subscribe_m5](#on_subscribe_m5)
    - [on_unsubscribe_m5](#on_unsubscribe_m5)
    - [on_deliver_m5](#on_deliver_m5)
  - [Modifiers](#modifiers)

# Installation 

Work in progress

# Usage

## Webhooks

|                      | mountpoint | client_id | username | password | peer_addr | peer_port | clean_session | topics (array object) | topics (string) | topics (array string) | qos | payload | retain | clean_start | properties |
|----------------------|:----------:|:---------:|:--------:|:--------:|:---------:|:---------:|:-------------:|:---------------------:|:---------------:|:---------------------:|:---:|:-------:|:------:|:-----------:|:----------:|
| auth_on_register     |     X      |     X     |    X     |    X     |     X     |     X     |       X       |                       |                 |                       |     |         |        |             |            |
| auth_on_subscribe    |     X      |     X     |    X     |          |           |           |               |           X           |                 |                       |     |         |        |             |            |
| auth_on_publish      |     X      |     X     |    X     |          |           |           |               |                       |        X        |                       |  X  |    X    |   X    |             |            |
| on_register          |     X      |     X     |    X     |          |     X     |     X     |               |                       |                 |                       |     |         |        |             |            |
| on_publish           |     X      |     X     |    X     |          |           |           |               |                       |        X        |                       |  X  |    X    |   X    |             |            |
| on_subscribe         |     X      |     X     |    X     |          |           |           |               |           X           |                 |                       |     |         |        |             |            |
| on_unsubscribe       |     X      |     X     |    X     |          |           |           |               |                       |                 |           X           |     |         |        |             |            |
| on_deliver           |     X      |     X     |    X     |          |           |           |               |                       |        X        |                       |     |    X    |        |             |            |
| on_offline_message   |     X      |     X     |          |          |           |           |               |                       |        X        |                       |  X  |    X    |   X    |             |            |
| on_client_wakeup     |     X      |     X     |          |          |           |           |               |                       |                 |                       |     |         |        |             |            |
| on_client_offline    |     X      |     X     |          |          |           |           |               |                       |                 |                       |     |         |        |             |            |
| on_client_gone       |     X      |     X     |          |          |           |           |               |                       |                 |                       |     |         |        |             |            |
| auth_on_register_m5  |     X      |     X     |    X     |    X     |     X     |     X     |               |                       |                 |                       |     |         |        |      X      |     X      |
| on_auth_m5           |     X      |     X     |    X     |          |           |           |               |                       |                 |                       |     |         |        |             |     X      |
| auth_on_subscribe_m5 |     X      |     X     |    X     |          |           |           |               |           X           |                 |                       |     |         |        |             |     X      |
| auth_on_publish_m5   |     X      |     X     |    X     |          |           |           |               |                       |        X        |                       |  X  |    X    |   X    |             |     X      |
| on_register_m5       |     X      |     X     |    X     |          |     X     |     X     |               |                       |                 |                       |     |         |        |             |     X      |
| on_publish_m5        |     X      |     X     |          |          |           |           |               |           X           |                 |                       |  X  |    X    |   X    |      X      |     X      |
| on_subscribe_m5      |     X      |     X     |    X     |          |           |           |               |           X           |                 |                       |     |         |        |             |     X      |
| on_unsubscribe_m5    |     X      |     X     |    X     |          |           |           |               |                       |                 |           X           |     |         |        |             |     X      |
| on_deliver_m5        |     X      |     X     |    X     |          |           |           |               |                       |        X        |                       |     |    X    |        |             |     X      |


### auth_on_register

### auth_on_subscribe

### auth_on_publish

### on_register

### on_publish

### on_subscribe

### on_unsubscribe


### on_deliver

### on_offline_message


### on_client_wakeup

### on_client_offline

### on_client_gone

### auth_on_register_m5

### on_auth_m5

### auth_on_subscribe_m5

### auth_on_publish_m5


### on_register_m5


### on_publish_m5


### on_subscribe_m5

### on_unsubscribe_m5


### on_deliver_m5

## Modifiers

| Modifier | Description | Type |  V1  | 
|:--------:|:-----------:|:----:|:----:|
|   Test   |    test     | test | test |


### auth_on_register

|          Modifier          |  Type  | Description | 
|:--------------------------:|:------:|:-----------:|
|       allow_register       |  bool  |    test     |
|       allow_publish        |  bool  |    test     |
|      allow_subscribe       |  bool  |    test     |
|     allow_unsubscribe      |  bool  |    test     |
|      max_message_size      |  int   |    test     |
|       subscriber_id        | string |    test     |
|          username          | string |    test     |
|      max_message_rate      |  int   |    test     |
|   max_inflight_messages    |  int   |    test     |
| shared_subscription_policy |  test  |    test     |
|        upgrade_qos         |  bool  |    test     |
|  allow_multiple_sessions   |  bool  |    test     |
|    max_online_messages     |  int   |    test     |
|    max_offline_messages    |  int   |    test     |
|     queue_deliver_mode     |  test  |    test     |
|         queue_type         |  test  |    test     |
|       max_drain_time       |  int   |    test     |
|  max_msgs_per_drain_step   |  int   |    test     |






