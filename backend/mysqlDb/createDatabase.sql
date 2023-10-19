use market;
create table USERS(
    userId       char(10) primary key not null,
    pswHash      char(128) not null
);

create table COOKIES(
    userId      char(10) primary key not null,
    cookie      char(16)
);

create table PRODUCTS(
    productId   int primary key not null,
    userId      char(10) not null,
    price       decimal(10, 2) not null,
    description varchar(2047) not null,
    createTime  time not null,
    updateTime  time not null,
    foreign key(userId) references USERS(userId)
);

create table PRODUCT_IMAGES(
    productId   int not null,
    imagePath   varchar(127) not null
);

create table COMMENTS(
    commentId   int primary key not null,
    publisherId char(10) not null,
    receiverId  char(10) not null,
    content     varchar(2047) not null,
    foreign key(publisherId) references USERS(userId),
    foreign key(receiverId) references USERS(userId)
);

create table SUBSCRIBE(
    subscribeId int primary key not null,
    userId      char(10) not null,
    productId   int      not null,
    foreign key(userId) references USERS(userId),
    foreign key(productId) references PRODUCTS(productId)
);


