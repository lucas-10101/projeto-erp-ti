# Organization API

API Destined to manage users, organizations and ther companies and plants

! Not all functionalities are implemented yet, so be pacient. Submit any bugs to issues on github or feel free to contribute on this project

#### This API handles the following resources
- Companies
- Plants
- Users
- User groups
- Application Roles
- User access between plants, companies and allowed actions by predefined roles
- Authentication of users, providing jwt tokens for api calls between multiple resources and resource servers of this ecossystem

#### Used resources and tools

- Golang 1.X (go1.24.2 linux/amd64)
- Liquibase 4.X (4.32.0)
- Java 21 (openjdk 21.0.6 2025-01-21)
- Visual Studio Code 1.X (1.98.2 x64)

#### Environment Variables

- APPLICATION_PROPERTIES_FILE: point to your application.properties location, this is required for correct operation of entire application, including tests and debug

#### How to run this application ?

###### The database

The database is managed by liquibase, so you must have it installed before proceed

- Install instructions of liquibase: https://sites.google.com/liquibase.org/liquibase/install-liquibase

After, you must run the update in your local database to create or update the database so put the required properties in liquibase/liquibase.properties file, you can follow the sample.liquibase.properties file to do that, the connection is on jdbc format for postgres

###### Prepare to launch

- If you get here after 'How to run this application ?', you are ready to make go build and run or go test ./... to test this application
- On development environment ? Just drop you database if you want to restart or do an liquibase rollback to previous version of this applications
- how ? go to ./liquibase dir containing the liquibase.properties file, run "liquibase update && liquibase rollback startdb && liquibase update" and you are ready to go