# service-trap 
<!-- ABOUT THE PROJECT -->

## About The Project

service-trap is listening on a specified tcp port and sends the client IP to a MySQL/MariaDB database. In addition, it writes random data to the client host to keep the connection up. This should keep the ugly crawling botnets occupied a bit busy

I forked this project of the great https://github.com/frelon/ssh-trap
The idea comes from https://github.com/iredmail/ fail2ban DB Store: https://docs.iredmail.org/fail2ban.sql.html

### Built With

* [Go](https://golang.org/)

## Database

Create a MySQL/MariaDB database. You can use the following mysql file to create the table: https://github.com/mike-goldfinger/service-trap/blob/master/service-trap.mysql

<!-- GETTING STARTED -->
## Getting Started

You can either use service-trap directly using Go or on your docker enviroment

### Prerequisites

You need to install [Go](https://golang.org/) to get it running if not using docker

### Installation

#### Directly using Go

1. Clone the repo
   ```sh
   git clone https://github.com/mike-goldfinger/service-trap.git
   ```

2. Create a Database and a user with insert permission to it and import the [service-trap.mysql](https://github.com/mike-goldfinger/service-trap/blob/master/service-trap.mysql) to create the table
  
3. Configure it:
   
   Set the parameters in the file config.yaml according your enviroment
   
4. Run it
   ```sh
   go run .
   ```

#### On your Docker enviroment using docker-compose

1. Clone the repo
   ```sh
   git https://github.com/mike-goldfinger/service-trap.git
   ```
2. Create a Database and a user with insert permission to it and import the [service-trap.mysql](https://github.com/mike-goldfinger/service-trap/blob/master/service-trap.mysql) to create the table

3. Merge and modify docker-compose.yml to your needs

4. Start it:
   ```sh
	docker-compose -f <pathToFile>/docker-compose.yml up -d
   ```

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->
## Usage

### How to trap IPs to fill the Database 

You can use either [service-trap](https://github.com/mike-goldfinger/service-trap) to set up an trap which listen on a specified port (Works with Go and as Docker Container) or using [fail2ban](https://github.com/fail2ban/fail2ban):

#### service.trap

Using [service-trap](https://github.com/mike-goldfinger/service-trap) you can set up a dummy service that listens on a port you have configured and sends the gathered IP addresses to your database. Just forward the port from your Firewall to this service-trap.

#### fail2ban

Using [fail2ban](https://github.com/fail2ban/fail2ban) and a part of [iRedMail](https://github.com/iredmail/iRedMail/)  you can set up a fail2ban jail that sends the blocked IP addresses to your database. Here a verry easy how to:  https://docs.iredmail.org/fail2ban.sql.html

### Block the IP addresses returned by this http list

Install [http-ip-list-creator](https://github.com/mike-goldfinger/http-ip-list-creator) let your firewall block all IP addresses returned by this http service.

#### Example using pfsense

You can install the package pfBlockerNG in your [pfsense](https://github.com/pfsense/pfsense) Firewall. Under an IPv4 Setting page you can add the IP adress or Hostname of your [http-ip-list-creator](https://github.com/mike-goldfinger/http-ip-list-creator) to the an “IPv4 List” (Format: Auto).  
