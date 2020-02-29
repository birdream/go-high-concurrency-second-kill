### start the db
```cmd
sudo docker run --name first-mysql -p 3306:3306 -e MYSQL\_ROOT\_PASSWORD=123456 -d mysql:5.6

## id name is already exist then run
## get the containerId
docker ps -a

## remove container
docker rm ${containerId}

## docker run again
```

### get into the db manager create database
```sql
# root 123456
create database imooc;

use imooc

-- create product table
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `productName` varchar(255) DEFAULT NULL,
  `productNum` int(11) DEFAULT NULL,
  `productImage` varchar(255) DEFAULT NULL,
  `productUrl` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- create order table
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `userID` int(11) DEFAULT NULL,
  `productID` int(11) DEFAULT NULL,
  `orderStatus` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=274 DEFAULT CHARSET=utf8;

-- create user table
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `nickName` varchar(255) DEFAULT NULL,
  `userName` varchar(255) DEFAULT NULL,
  `HashPassword` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=274 DEFAULT CHARSET=utf8;
```

### start the project
```cmd
cd backend
go build
./backend

cd fronts
go build
./fronts
```

### visit the browser `localhost:8081`
