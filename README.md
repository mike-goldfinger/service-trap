# service-trap 
<!-- ABOUT THE PROJECT -->

## About The Project

service-trap is listening on a specified tcp port and sends the client IP to an MySQL/MariaDB database. In addition it writes random data to the client host. 

I forked this project of the great https://github.com/frelon/ssh-trap
The idea comes from https://github.com/iredmail/ fail2ban DB Store: https://docs.iredmail.org/fail2ban.sql.html

### Built With

* [Go](https://golang.org/)

## Database

Create a MySQL/MariaDB Database and Table with at least the following rows:

<!-- GETTING STARTED -->
## Getting Started

You can either use service-trap directly using Go or, coming soon also over your docker enviroment

### Prerequisites

You need to install [Go](https://golang.org/) to get it running if not using docker

### Installation

#### Directly using Go

1. Clone the repo
   ```sh
   git clone https://github.com/mike-goldfinger/service-trap .git
   ```

2. Creat a Database and a user with insert permission to it and import the service-trap.mysql to create the table
   
3. Configure it:
   
   Set the parameters in the file config.yaml according your enviroment
   
4. Run it
   ```js
   go run .
   ```

#### On your Docker enviroment using docker-compose

1. Clone the repo
   ```sh
   git https://github.com/mike-goldfinger/service-trap.git
   ```

2. Merge and modify docker-compose.yml to your needs

2. Start it:
   ```sh
	docker-compose -f <pathToFile>/docker-compose.yml up -d
   ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

1. Configure and install the trap: https://github.com/mike-goldfinger/service-trap

2. Configure and install the list creator: https://github.com/mike-goldfinger/http-ip-list-creator

3. Forward the tcp port you want to use as trap to the ssh trap docker container or start it directly on your server using go

4. Install pfBlockerNG into your pfsense Firewall and set it up to get the IP list from http-ip-list-creator
