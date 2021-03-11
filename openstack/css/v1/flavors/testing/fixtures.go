package testing

const listResponseBody = `
{
  "versions": [
    {
      "flavors": [
        {
          "cpu": 1,
          "ram": 8,
          "name": "css.medium.8",
          "region": "eu-de",
          "diskrange": "40,640",
          "flavor_id": "ced8d1a7-eff8-4e30-a3de-cd9578fd518f"
        },
        {
          "cpu": 2,
          "ram": 8,
          "name": "css.large.4",
          "region": "eu-de",
          "diskrange": "40,800",
          "flavor_id": "b58dfcfb-5f73-4f05-8c06-5171b12fd618"
        },
        {
          "cpu": 2,
          "ram": 16,
          "name": "css.large.8",
          "region": "eu-de",
          "diskrange": "40,1280",
          "flavor_id": "d3952f68-86d5-4a41-8be3-73f99b6be7f8"
        },
        {
          "cpu": 4,
          "ram": 8,
          "name": "css.xlarge.2",
          "region": "eu-de",
          "diskrange": "40,800",
          "flavor_id": "d9dc06ae-b9c4-4ef4-acd8-953ef4205e27"
        },
        {
          "cpu": 4,
          "ram": 16,
          "name": "css.xlarge.4",
          "region": "eu-de",
          "diskrange": "40,1600",
          "flavor_id": "35b060a4-f152-48ce-8773-36559ceb81f2"
        },
        {
          "cpu": 4,
          "ram": 32,
          "name": "css.xlarge.8",
          "region": "eu-de",
          "diskrange": "40,2560",
          "flavor_id": "cb1a3408-33a7-4ced-b834-c2ad419876d8"
        },
        {
          "cpu": 8,
          "ram": 16,
          "name": "css.2xlarge.2",
          "region": "eu-de",
          "diskrange": "80,1600",
          "flavor_id": "b9fe19f1-835b-40db-879c-cddcd64a9648"
        },
        {
          "cpu": 8,
          "ram": 32,
          "name": "css.2xlarge.4",
          "region": "eu-de",
          "diskrange": "80,3200",
          "flavor_id": "c3ed3633-e62c-4903-af4e-dd7aec13816d"
        },
        {
          "cpu": 8,
          "ram": 64,
          "name": "css.2xlarge.8",
          "region": "eu-de",
          "diskrange": "80,5120",
          "flavor_id": "c2f6e2f5-a8de-49d7-864b-19559f127954"
        },
        {
          "cpu": 16,
          "ram": 32,
          "name": "css.4xlarge.2",
          "region": "eu-de",
          "diskrange": "100,3200",
          "flavor_id": "70c5a33e-dfcb-41ee-8992-129deb069b26"
        },
        {
          "cpu": 16,
          "ram": 64,
          "name": "css.4xlarge.4",
          "region": "eu-de",
          "diskrange": "100,6400",
          "flavor_id": "d4e2a36b-3d93-41df-97e7-498816b7bd2e"
        },
        {
          "cpu": 16,
          "ram": 128,
          "name": "css.4xlarge.8",
          "region": "eu-de",
          "diskrange": "160,10240",
          "flavor_id": "3da413a1-2ad0-4f09-905c-28fbde50c1d4"
        },
        {
          "cpu": 32,
          "ram": 64,
          "name": "css.8xlarge.2",
          "region": "eu-de",
          "diskrange": "320,10240",
          "flavor_id": "d62976d4-6019-95a2-7e30-1db3fbfaa87d"
        },
        {
          "cpu": 32,
          "ram": 128,
          "name": "css.8xlarge.4",
          "region": "eu-de",
          "diskrange": "160,10240",
          "flavor_id": "b3512820-9033-42a5-9412-69a3af653a3d"
        }
      ],
      "type": "ess",
      "version": "7.6.2"
    },
    {
      "flavors": [
        {
          "cpu": 1,
          "ram": 8,
          "name": "css.medium.8",
          "region": "eu-de",
          "diskrange": "40,640",
          "flavor_id": "ced8d1a7-eff8-4e30-a3de-cd9578fd518f"
        },
        {
          "cpu": 2,
          "ram": 8,
          "name": "css.large.4",
          "region": "eu-de",
          "diskrange": "40,800",
          "flavor_id": "b58dfcfb-5f73-4f05-8c06-5171b12fd618"
        },
        {
          "cpu": 2,
          "ram": 16,
          "name": "css.large.8",
          "region": "eu-de",
          "diskrange": "40,1280",
          "flavor_id": "d3952f68-86d5-4a41-8be3-73f99b6be7f8"
        },
        {
          "cpu": 4,
          "ram": 8,
          "name": "css.xlarge.2",
          "region": "eu-de",
          "diskrange": "40,800",
          "flavor_id": "d9dc06ae-b9c4-4ef4-acd8-953ef4205e27"
        },
        {
          "cpu": 4,
          "ram": 16,
          "name": "css.xlarge.4",
          "region": "eu-de",
          "diskrange": "40,1600",
          "flavor_id": "35b060a4-f152-48ce-8773-36559ceb81f2"
        },
        {
          "cpu": 4,
          "ram": 32,
          "name": "css.xlarge.8",
          "region": "eu-de",
          "diskrange": "40,2560",
          "flavor_id": "cb1a3408-33a7-4ced-b834-c2ad419876d8"
        },
        {
          "cpu": 8,
          "ram": 16,
          "name": "css.2xlarge.2",
          "region": "eu-de",
          "diskrange": "80,1600",
          "flavor_id": "b9fe19f1-835b-40db-879c-cddcd64a9648"
        },
        {
          "cpu": 8,
          "ram": 32,
          "name": "css.2xlarge.4",
          "region": "eu-de",
          "diskrange": "80,3200",
          "flavor_id": "c3ed3633-e62c-4903-af4e-dd7aec13816d"
        },
        {
          "cpu": 8,
          "ram": 64,
          "name": "css.2xlarge.8",
          "region": "eu-de",
          "diskrange": "80,5120",
          "flavor_id": "c2f6e2f5-a8de-49d7-864b-19559f127954"
        },
        {
          "cpu": 16,
          "ram": 32,
          "name": "css.4xlarge.2",
          "region": "eu-de",
          "diskrange": "100,3200",
          "flavor_id": "70c5a33e-dfcb-41ee-8992-129deb069b26"
        },
        {
          "cpu": 16,
          "ram": 64,
          "name": "css.4xlarge.4",
          "region": "eu-de",
          "diskrange": "100,6400",
          "flavor_id": "d4e2a36b-3d93-41df-97e7-498816b7bd2e"
        },
        {
          "cpu": 16,
          "ram": 128,
          "name": "css.4xlarge.8",
          "region": "eu-de",
          "diskrange": "160,10240",
          "flavor_id": "3da413a1-2ad0-4f09-905c-28fbde50c1d4"
        },
        {
          "cpu": 32,
          "ram": 64,
          "name": "css.8xlarge.2",
          "region": "eu-de",
          "diskrange": "320,10240",
          "flavor_id": "d62976d4-6019-95a2-7e30-1db3fbfaa87d"
        },
        {
          "cpu": 32,
          "ram": 128,
          "name": "css.8xlarge.4",
          "region": "eu-de",
          "diskrange": "160,10240",
          "flavor_id": "b3512820-9033-42a5-9412-69a3af653a3d"
        }
      ],
      "type": "ess-master",
      "version": "7.6.2"
    },
    {
      "flavors": [
        {
          "cpu": 1,
          "ram": 8,
          "name": "css.medium.8",
          "region": "eu-de",
          "diskrange": "40,640",
          "flavor_id": "ced8d1a7-eff8-4e30-a3de-cd9578fd518f"
        },
        {
          "cpu": 2,
          "ram": 8,
          "name": "css.large.4",
          "region": "eu-de",
          "diskrange": "40,800",
          "flavor_id": "b58dfcfb-5f73-4f05-8c06-5171b12fd618"
        },
        {
          "cpu": 2,
          "ram": 16,
          "name": "css.large.8",
          "region": "eu-de",
          "diskrange": "40,1280",
          "flavor_id": "d3952f68-86d5-4a41-8be3-73f99b6be7f8"
        },
        {
          "cpu": 4,
          "ram": 8,
          "name": "css.xlarge.2",
          "region": "eu-de",
          "diskrange": "40,800",
          "flavor_id": "d9dc06ae-b9c4-4ef4-acd8-953ef4205e27"
        },
        {
          "cpu": 4,
          "ram": 16,
          "name": "css.xlarge.4",
          "region": "eu-de",
          "diskrange": "40,1600",
          "flavor_id": "35b060a4-f152-48ce-8773-36559ceb81f2"
        },
        {
          "cpu": 4,
          "ram": 32,
          "name": "css.xlarge.8",
          "region": "eu-de",
          "diskrange": "40,2560",
          "flavor_id": "cb1a3408-33a7-4ced-b834-c2ad419876d8"
        },
        {
          "cpu": 8,
          "ram": 16,
          "name": "css.2xlarge.2",
          "region": "eu-de",
          "diskrange": "80,1600",
          "flavor_id": "b9fe19f1-835b-40db-879c-cddcd64a9648"
        },
        {
          "cpu": 8,
          "ram": 32,
          "name": "css.2xlarge.4",
          "region": "eu-de",
          "diskrange": "80,3200",
          "flavor_id": "c3ed3633-e62c-4903-af4e-dd7aec13816d"
        },
        {
          "cpu": 8,
          "ram": 64,
          "name": "css.2xlarge.8",
          "region": "eu-de",
          "diskrange": "80,5120",
          "flavor_id": "c2f6e2f5-a8de-49d7-864b-19559f127954"
        },
        {
          "cpu": 16,
          "ram": 32,
          "name": "css.4xlarge.2",
          "region": "eu-de",
          "diskrange": "100,3200",
          "flavor_id": "70c5a33e-dfcb-41ee-8992-129deb069b26"
        },
        {
          "cpu": 16,
          "ram": 64,
          "name": "css.4xlarge.4",
          "region": "eu-de",
          "diskrange": "100,6400",
          "flavor_id": "d4e2a36b-3d93-41df-97e7-498816b7bd2e"
        },
        {
          "cpu": 16,
          "ram": 128,
          "name": "css.4xlarge.8",
          "region": "eu-de",
          "diskrange": "160,10240",
          "flavor_id": "3da413a1-2ad0-4f09-905c-28fbde50c1d4"
        },
        {
          "cpu": 32,
          "ram": 64,
          "name": "css.8xlarge.2",
          "region": "eu-de",
          "diskrange": "320,10240",
          "flavor_id": "d62976d4-6019-95a2-7e30-1db3fbfaa87d"
        },
        {
          "cpu": 32,
          "ram": 128,
          "name": "css.8xlarge.4",
          "region": "eu-de",
          "diskrange": "160,10240",
          "flavor_id": "b3512820-9033-42a5-9412-69a3af653a3d"
        }
      ],
      "type": "ess-client",
      "version": "7.6.2"
    },
    {
      "flavors": [
        {
          "cpu": 1,
          "ram": 8,
          "name": "css.medium.8",
          "region": "eu-de",
          "diskrange": "40,640",
          "flavor_id": "ced8d1a7-eff8-4e30-a3de-cd9578fd518f"
        },
        {
          "cpu": 2,
          "ram": 8,
          "name": "css.large.4",
          "region": "eu-de",
          "diskrange": "40,800",
          "flavor_id": "b58dfcfb-5f73-4f05-8c06-5171b12fd618"
        },
        {
          "cpu": 2,
          "ram": 16,
          "name": "css.large.8",
          "region": "eu-de",
          "diskrange": "40,1280",
          "flavor_id": "d3952f68-86d5-4a41-8be3-73f99b6be7f8"
        },
        {
          "cpu": 4,
          "ram": 8,
          "name": "css.xlarge.2",
          "region": "eu-de",
          "diskrange": "40,800",
          "flavor_id": "d9dc06ae-b9c4-4ef4-acd8-953ef4205e27"
        },
        {
          "cpu": 4,
          "ram": 16,
          "name": "css.xlarge.4",
          "region": "eu-de",
          "diskrange": "40,1600",
          "flavor_id": "35b060a4-f152-48ce-8773-36559ceb81f2"
        },
        {
          "cpu": 4,
          "ram": 32,
          "name": "css.xlarge.8",
          "region": "eu-de",
          "diskrange": "40,2560",
          "flavor_id": "cb1a3408-33a7-4ced-b834-c2ad419876d8"
        },
        {
          "cpu": 8,
          "ram": 16,
          "name": "css.2xlarge.2",
          "region": "eu-de",
          "diskrange": "80,1600",
          "flavor_id": "b9fe19f1-835b-40db-879c-cddcd64a9648"
        },
        {
          "cpu": 8,
          "ram": 32,
          "name": "css.2xlarge.4",
          "region": "eu-de",
          "diskrange": "80,3200",
          "flavor_id": "c3ed3633-e62c-4903-af4e-dd7aec13816d"
        },
        {
          "cpu": 8,
          "ram": 64,
          "name": "css.2xlarge.8",
          "region": "eu-de",
          "diskrange": "80,5120",
          "flavor_id": "c2f6e2f5-a8de-49d7-864b-19559f127954"
        },
        {
          "cpu": 16,
          "ram": 32,
          "name": "css.4xlarge.2",
          "region": "eu-de",
          "diskrange": "100,3200",
          "flavor_id": "70c5a33e-dfcb-41ee-8992-129deb069b26"
        },
        {
          "cpu": 16,
          "ram": 64,
          "name": "css.4xlarge.4",
          "region": "eu-de",
          "diskrange": "100,6400",
          "flavor_id": "d4e2a36b-3d93-41df-97e7-498816b7bd2e"
        },
        {
          "cpu": 16,
          "ram": 128,
          "name": "css.4xlarge.8",
          "region": "eu-de",
          "diskrange": "160,10240",
          "flavor_id": "3da413a1-2ad0-4f09-905c-28fbde50c1d4"
        },
        {
          "cpu": 32,
          "ram": 64,
          "name": "css.8xlarge.2",
          "region": "eu-de",
          "diskrange": "320,10240",
          "flavor_id": "d62976d4-6019-95a2-7e30-1db3fbfaa87d"
        },
        {
          "cpu": 32,
          "ram": 128,
          "name": "css.8xlarge.4",
          "region": "eu-de",
          "diskrange": "160,10240",
          "flavor_id": "b3512820-9033-42a5-9412-69a3af653a3d"
        }
      ],
      "type": "ess-cold",
      "version": "7.6.2"
    }
  ]
}
`
