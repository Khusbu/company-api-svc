## company-api-svc

A sample API service in Go deployed to [Heroku](https://company-api-svc.herokuapp.com/).

### Supported APIs

Base URL: https://company-api-svc.herokuapp.com/

#### POST - `/company/new`
Creates a new record with the details provided as POST data in JSON format. Note: Only JSON format input is supported.

**Arguments**

- *company_name* `required`  
  A string value containing the company name. Maximum 20 characters allowed.
- *description* `optional`  
  A string value containing a brief description about the company. Maximum 200 characters allowed.
- *logo* `optional`  
  A string value containing the image url for the logo.
- *funding_details* `optional`  
  An JSON array containing the details of each fund received by the company. Each element of the array has the given attributes:
  - *amount* `required`  
  A string value containing the funding amount value in paise (e.g., 5000 paise denotes Rs 50.00). Minimum amount is 0 paisa and maximum is 1000000000 paise.
  - *date* `optional`  
  A string value containing the date on which funding was received by the company in `dd/mm/yyyy` format.
  - *stages* `optional`  
  A string value containing the stage of the funding. It should be one of the following values: `"Series A", "Series B", "Series C", "Series D", "Series E", "Series F"`
  - *investors* `optional`  
  A string value containing the name of the investor behind this funding.
- markets `optional`  
  A string value containing the tags for the company. For example: ecommerce, edutech, adtech, healthtech
- founded_on `optional`  
  A string value containing the date on which the company was found. It should be in `dd/mm/yyyy` format.
- website `optional`
  A string value containing the company home page URL.
- linked_in `optional`
  A string value containing the LinkedIn URL of the company.
- twitter `optional`  
  A string value containing the Twitter URL of the company.
- email `optional`  
  A string value containing the contact email of the company.
- phone_number `optional`  
  A string value containing the 10 digit contact number of the company.

**Example Request**

```curl
$ curl http://localhost:5000/company/new \
-H "Content-Type: application/json" -X POST \
-d '{"name":"Example Company","description":"Startup Company","logo":"www.examplecompanylogo.com","funding_details":[\
{"amount":"1000000","date":"12/12/2012","stages":"Series A","investors":"Flipkart"}, \
{"amount":"100000000","date":"12/12/2013","stages":"Series B","investors":"Snapdeal"}],\
"markets":"ecommerce","founded_on":"12/12/2010","website":"www.examplecompanywebsite.com",\
"linked_in":"https://linkedin.com/examplecompany","twitter":"https://twitter.com/examplecompany",\
"email":"support@examplecompany.com","phone_number":"9999999999"}'
```

**Example Response**

```json
{
  "name": "Example Company",
  "description": "Startup Company",
  "logo": "www.examplecompanylogo.com",
  "funding_details": [
    {
      "ID": 1,
      "amount": "1000000",
      "date": "12\/12\/2012",
      "stages": "Series A",
      "investors": "Flipkart",
      "profile_id": "cmp_aUCpqLKTAUwAuL0h"
    },
    {
      "ID": 2,
      "amount": "100000000",
      "date": "12\/12\/2013",
      "stages": "Series B",
      "investors": "Snapdeal",
      "profile_id": "cmp_aUCpqLKTAUwAuL0h"
    }
  ],
  "markets": "ecommerce",
  "founded_on": "12\/12\/2010",
  "website": "www.examplecompanywebsite.com",
  "linked_in": "https:\/\/linkedin.com\/examplecompany",
  "twitter": "https:\/\/twitter.com\/examplecompany",
  "email": "support@examplecompany.com",
  "phone_number": "9999999999",
  "profile_id": "cmp_aUCpqLKTAUwAuL0h"
}
```

#### GET - `/company/:profile_id`
Fetches a company record with the Profile ID as `profile_id`.

**Example Request**
```curl
curl http://localhost:5000/company/cmp_aUCpqLKTAUwAuL0h
```

**Example Response**

```json
{  
   "name":"Example Company",
   "description":"Startup Company",
   "logo":"www.examplecompanylogo.com",
   "funding_details":[  
      {  
         "ID":1,
         "amount":"1000000",
         "date":"12/12/2012",
         "stages":"Series A",
         "investors":"Flipkart",
         "profile_id":"cmp_aUCpqLKTAUwAuL0h"
      },
      {  
         "ID":2,
         "amount":"100000000",
         "date":"12/12/2013",
         "stages":"Series B",
         "investors":"Snapdeal",
         "profile_id":"cmp_aUCpqLKTAUwAuL0h"
      }
   ],
   "markets":"ecommerce",
   "founded_on":"12/12/2010",
   "website":"www.examplecompanywebsite.com",
   "linked_in":"https://linkedin.com/examplecompany",
   "twitter":"https://twitter.com/examplecompany",
   "email":"support@examplecompany.com",
   "phone_number":"9999999999",
   "profile_id":"cmp_aUCpqLKTAUwAuL0h"
}
```

### Running Locally Using Heroku

Make sure you have [Go](http://golang.org/doc/install) and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

A SQL database whose URL must be set in environment as:
`export DATABASE_URL=<URL>`

```sh
$ go get -u github.com/Khusbu/company-api-svc
$ cd $GOPATH/src/github.com/Khusbu/company-api-svc
$ heroku local web
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

### Deploying to Heroku

```sh
$ heroku create <app-name>
$ git push heroku master
$ heroku open
```
