<h1 align="center">
Ad Registration Service
</h1>

Cloud computing first project/homework. In this project we are going to create a **Ad Registration Service**. We are building two Backend services for getting user information and processing the next steps.

## First API

In first API we are going to get all of our clients information about an Ad. This information consists of:

- User Email
- Ad title (description)
- Ad image

Then we need to store User email and Ad Title in a [MongoDB cluster](https://www.mongodb.com/) and Ad Image in 
Amazon [S3 database](https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&cad=rja&uact=8&ved=2ahUKEwiA_d3T5un6AhUFzqQKHXySDwQQFnoECBcQAQ&url=https%3A%2F%2Faws.amazon.com%2Fs3%2F&usg=AOvVaw3NS_rqXKJpiZug3wHxUGKs). After that we need to publish this information on a [RabbitMQ Cluster](https://www.cloudamqp.com/).

In this API we need to return the information of each Ad. The schema is like this:

```
+----------+----------------------+----------------+----------------+-------------------+
| id (int) | description (string) | email (string) | state (string) | category (string) |
+----------+----------------------+----------------+----------------+-------------------+
```

## Second API

In this API we need to process the Ad. We need to subscribe over RabbitMQ Cluster to 
receive events from first API. After that we are going to get Image from S3 cluster and 
then it will send that image to [Image tagging service](). It will get a response for
the suggested tags and stores them inside our MongoDB. If the Ad was valid, it will send an email to client via [Mailgun]().

## Requirements 

- Cloud based service, which we are going to use [Arvan Cloud](https://www.arvancloud.com/fa).
- DBaaS (Database as a service), which we are going to use [MongoDB Atlas](https://www.mongodb.com/cloud/atlas/lp/try4?utm_content=controldbaasterms&utm_source=google&utm_campaign=search_gs_pl_evergreen_atlas_core_prosp-brand_gic-null_emea-nl_ps-all_desktop_eng_lead&utm_term=mongodb%20dbaas&utm_medium=cpc_paid_search&utm_ad=e&utm_ad_campaign_id=12212624536&adgroup=115749708663).
- Amazon S3 database for Object storage, which we are going to get from [Arvan Cloud](https://www.arvancloud.com/en/products/cloud-storage).
- RabbitMQ service, which are going to use [Cloud AMQP](https://www.cloudamqp.com/).
- Image tagging service, which we are going to use [Imagga](https://imagga.com/).
- Mail service, which we are going to use [Mailgun](https://www.mailgun.com/).

## API routes

Creating two Golang services to provide the APIs.

### First API

#### Routes

```json
[
  {
    "method": "GET",
    "uri": "api/{id}",
    "response": {
      "id": "1",
      "description": "about ad",
      "email": "sender@gmail.com",
      "status": "progress/failed/successful",
      "category": "something"
    }
  },
  {
    "method": "POST",
    "uri": "api/",
    "request body": {
      "description": "about ad",
      "email": "sender@gmail.com",
      "file": "file object"
    },
    "response": {
      "id": "1"
    }
  }
]
```

#### Configs

```yaml
port: 5000
db:
  mong:
    url: "address to mongo cluster"
  s3:
    url: "address to S3"
    token: "S3 auth token"
rabbit:
  url: "rabbit url"
```

### Second API

This API job is to listen on RabbitMQ and process the Ads.

#### Configs

```yaml
port: 10251
mail:
  address: "mailgun service address"
  token: "api token"
db:
  mong:
    url: "address to mongo cluster"
  s3:
    url: "address to S3"
    token: "S3 auth token"
rabbit:
  url: "rabbit url"
imagga:
  url: "immage url"
```
