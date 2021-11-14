CREATE DATABASE `receiptdb` /*!40100 DEFAULT CHARACTER SET ascii */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `receiptdb`;

CREATE TABLE `receipts` (
  `receiptID` int NOT NULL AUTO_INCREMENT,
  `storeName` varchar(255) NOT NULL,
  `dateOfPurchase` datetime NOT NULL,
  `totalCost` decimal(13,2) NOT NULL,
  `taxRate` decimal(5,2) NOT NULL,
  `totalTax` decimal(13,2) NOT NULL,
  PRIMARY KEY (`receiptID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=ascii;

CREATE TABLE `products` (
  `productID` int NOT NULL AUTO_INCREMENT,
  `receiptID` int NOT NULL,
  `productName` varchar(60) NOT NULL,
  `productCost` decimal(13,2) NOT NULL,
  `productQuantity` int NOT NULL,
  PRIMARY KEY (`productID`),
  KEY `receiptID_idx` (`receiptID`),
  CONSTRAINT `receiptID` FOREIGN KEY (`receiptID`) REFERENCES `receipts` (`receiptID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=ascii;

USE `receiptdb`;

INSERT INTO `receipts` (`storeName`,`dateOfPurchase`,`totalCost`,`taxRate`,`totalTax`) VALUES ("CostCo", curdate(), 100.99, 8.3, 1.00);
INSERT INTO `receipts` (`storeName`,`dateOfPurchase`,`totalCost`,`taxRate`,`totalTax`) VALUES ("Best Buy", curdate(), 5.21, 9.3, 2.00);
INSERT INTO `receipts` (`storeName`,`dateOfPurchase`,`totalCost`,`taxRate`,`totalTax`) VALUES ("WinCo", curdate(), 1.09, 10.4, 3.00);

INSERT INTO `products` (`receiptID`, `productName`, `productCost`, `productQuantity`) VALUES (1, "Product1", 100.00, 1);
INSERT INTO `products` (`receiptID`, `productName`, `productCost`, `productQuantity`) VALUES (1, "Product2", 200.00, 2);
INSERT INTO `products` (`receiptID`, `productName`, `productCost`, `productQuantity`) VALUES (1, "Product3", 300.00, 3);
INSERT INTO `products` (`receiptID`, `productName`, `productCost`, `productQuantity`) VALUES (2, "Product1", 100.00, 1);
INSERT INTO `products` (`receiptID`, `productName`, `productCost`, `productQuantity`) VALUES (2, "Product2", 200.00, 2);
INSERT INTO `products` (`receiptID`, `productName`, `productCost`, `productQuantity`) VALUES (2, "Product3", 300.00, 3);
INSERT INTO `products` (`receiptID`, `productName`, `productCost`, `productQuantity`) VALUES (3, "Product1", 100.00, 1);
INSERT INTO `products` (`receiptID`, `productName`, `productCost`, `productQuantity`) VALUES (3, "Product2", 200.00, 2);
INSERT INTO `products` (`receiptID`, `productName`, `productCost`, `productQuantity`) VALUES (3, "Product3", 300.00, 3);

select * from `receiptdb`.`receipts`;
select * from `receiptdb`.`products`;


#delete from `receiptdb`.`products` where receiptID <> 0;
#delete from `receiptdb`.`receipts` where receiptID <> 0;
