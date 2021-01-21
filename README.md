# CodeMax Assignment
### 1.  [Usage](#usage)
### 2.  [Design Pattern](#design-pattern)
### 3.  [Middleware Logging](#middleware-logging)
### 4.  [Framework](#framework)
### 5.  [Future Integration](#future-integration)
### 6.  [Token Authentication](#token-authentication)

    The following assignment was given to create a REST API to send email to "to", "cc" and "bcc" using MailJet API and also considering the future Integration with Amazon ses or sendgrid provider

## [Usage](#usage)
### First clone this repo

git clone "https://github.com/tecsavy97/cdmx_assignment.git"


### Single Instance

    Make the necessary changes to the ./config/config.toml
    If you want to change the provide change the file ./config/emailConf.json
    {
        "hostName": "MailJet",
        "server": "in-v3.mailjet.com",
        "port": 587,
        "username": "",
        "password": "",
        "auth_keys": { //optional
            "identity":"",
            "publicKey":"",
            "secretKey":""
        },
        "isDefault": false //optional
    }

    Open a Terminal it the Cloned Directory
    Run Command
        go run main.go
### Multiple Instance
    Use the `emailhelper` and use the Function to expose a route and pass a array of the object such as 
    [
        {
        "hostName": "Provider1",
        "server": "Provider1.com",
        "port": 587,
        "username": "",
        "password": "",
        "auth_keys": { //optional
            "identity":"",
            "publicKey":"",
            "secretKey":""
            },
        "isDefault": true //optional
        },
        {
        "hostName": "Provider2",
        "server": "Provider2.com",
        "port": 587,
        "username": "",
        "password": "",
        "auth_keys": { //optional
            "identity":"",
            "publicKey":"",
            "secretKey":""
            },
        "isDefault": false //optional
        }
    ]

    Open a terminal and run command in the directory as 
        go build
    then run command after succesfull build
        ./codemax-assignment

## Routes
There are only two routes 

1. **`/o/login`** Login open route, It is a POST method where we have to pass body as
{
    "username":"Sahil",
    "password":"sahil"
}   
It returns a Token in the reponse which we have to set it in the **`Authorizarion`** Header to access the next route.
2. **`/r/send-email`** Send Email Route, it is a restricted POST method route where the token from Login route is to be set in **`Authorization`** Header and Body to be passed as 
   
        {
            "to":["xyz@gmail.com"],
            "cc":[],
            "bcc":[],
            "from":"xyz@gmail.com",
            "replyTo":"",
            "subject":"test mail",
            "body":"Hi All"
        }  
It will return success if the emails are sent or it will give error if any issue occurs.

## [Design Pattern](#design-pattern)
    1. For creating REST API for the following Project I have used Singleton and Separation of concern pattern in which where I created a single instance of EmailConfig for Single Instance or the Map of EmailConfig. 
   
    2. As per the Separation of concern for Multiple EmailConfigs I have initiated a constructor where we have to pass the necessary data and also we can call the Method that are binded, so that the flow gets executed in it.

    3. Also we can modularise this Project as per `email` and `user` where the routes, handlers and services with models can be made independent from the main service.

    4. Also for Multiple EmailConfigs, I have used Pool Object Pattern where we create a Map of configs and we can use the HostName to get the EmailConfig from the Map as required

## [Middleware Logging](#middleware-logging)
    1. For Logging I have used two different loggers 
       a. Gin Default Logger (i.e Access Loggin)
       b. Custom Sugar Logging (i.e Data Logging) 
    
    2. The reason that I have used two different loggers as I want to clearly display the routes consumed of the server and also clearly observe the necessary data if any API fails.

    3. Access Logger stores all the API consumed and Data Logging stores all the Data that is logged that can be errors or any info logging

    4. These two loggers help is resolving any issue or bugs that can occurs in the services
   
    5. As for Custom Sugar Logging I wrote a `loggerhelper` which help me to wrote "ERROR" and "INFO" logging, we can also write "DEBUG", "PANIC" and "FATAL" logs but the methods are not initiated in the loggerhelper.
   
## [Framework](#framework)
    1. Framework used while creating these Project is Gin-Gonic it is HTTP web framework as it is relatively faster than any other frameworks.

    2. The Gin framewrok has its own Recovery Method, it also has its Default Logger and it also has it Middleware and can also supports any Custom Middleware or Logger.
   
    3. Group Routing is easy in Gin. Authorized routes vs non authorized routes, different API versions can be managed. In addition, the groups can be nested unlimitedly without degrading performance.

    4. It can handle JSON, XML or HTML type data for data rendering.

## [Future Integration](#future-integration)
    1. For Future Integration with sendgrid or Amazon ses, I have provided a `emailhelper` where we can initiate Multiple provides as per the concern and can user these providers with their HostName and send Email through it.

    2. I have created a Map of Configs and can be used to Load these configs in memory and by the HostName of the Configs we can Get the ProviderData and send Email using the Method binded to it.

    3. We can call `email.NewEmail()` pass the necessary parameters to it and create the Email instance and using `SendEmail()` binded to the Email and passing the HostName of the Config we can Send Email as per the Provider Hostname initiated in the Map.

## [Token Authentication](#token-authentication)
    1. As I don't have to use JWT authentication in this Project, I have used base64 Authentication.

    2. We can generate the base64 token required for the server by using '/o/login' route by passing the data as 
   
    {
        "username":"",
        "password";""
    }

    getting the token in the response using it in the Header of "/r/send-email" request.
    
    2. I have authenticated the API by these methods