use market;
create table USERS(
    userId       char(10) primary key not null,
    pswHash      char(128) not null
);

create table USER_INFOS(
    userId      char(10) primary key not null,
    nickName    varchar(127),
    avatar      varchar(127),
    contact     varchar(127),
    foreign key(userId) references USERS(userId)
);

create table COOKIES(
    userId      char(10) primary key not null,
    cookie      char(16)
);

create table PRODUCTS(
    productId   int primary key not null auto_increment,
    userId      char(10) not null,
    title       char(127) not null,
    price       decimal(10, 2) not null,
    status      char(31),
    description varchar(2047) not null,
    createTime  time not null,
    updateTime  time not null,
    foreign key(userId) references USERS(userId)
);

create table PRODUCT_IMAGES(
    productId   int not null,
    imagePath   varchar(127) not null,
    foreign key(productId) references PRODUCTS(productId) on delete cascade
);

create table CATEGORIES(
    categoryName char(30) primary key not null
);

create table PRODUCT_CATEGORIES(
    productId int not null,
    category  char(30),
    foreign key(productId) references PRODUCTS(productId) on delete cascade
);

create table COMMENTS(
    commentId   int primary key not null auto_increment,
    publisherId char(10) not null,
    receiverId  char(10) not null,
    content     varchar(2047) not null,
    foreign key(publisherId) references USERS(userId),
    foreign key(receiverId) references USERS(userId)
);

create table SUBSCRIBE(
    subscribeId int primary key not null auto_increment,
    userId      char(10) not null,
    productId   int      not null,
    foreign key(userId) references USERS(userId),
    foreign key(productId) references PRODUCTS(productId)
);


