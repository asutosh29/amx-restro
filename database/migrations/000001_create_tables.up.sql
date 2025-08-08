create table users(
	id bigint auto_increment,
    email varchar(255) NOT NULL unique,
    username varchar(255) NOT NULL unique,
    first_name varchar(255) NOT NULL ,
    last_name varchar(255) NOT NULL ,
    contact varchar(14) NOT NULL,
    hashpwd varchar(100) NOT NULL,
    userRole ENUM('customer', 'chef','admin','super') default 'customer' NOT NULL,
    primary key(id)
);

create table category(
	category_id bigint auto_increment,
    category_name varchar(30),
    primary key(category_id)
);

create table items(
	item_id bigint auto_increment,
    category_id bigint not null,
    item_name varchar(30) not null,
    item_description text not null,
    img_url text not null ,
    price double not null,
    isVeg boolean not null,
    isAvailable boolean not null,
    primary key(item_id),
    foreign key(category_id) references category(category_id)
);

create table tables(
	table_id bigint,
    isAvailable boolean not null default 1,
    primary key(table_id)
);

create table orders(
	order_id bigint not null auto_increment,
    customer_id bigint not null,
    table_id bigint not null,
    extra_instructions text not null,
    order_status ENUM('placed', 'cooking', 'served', 'billed', 'paid') default 'placed' not null,
    total_amount double not null,
    order_at_time datetime not null,
    primary key(order_id),
    foreign key(table_id) references tables(table_id),
    foreign key(customer_id) references users(id)
);

create table order_item(
	item_id bigint,
    order_id bigint,
	qty bigint not null,
    primary key(item_id, order_id),
    foreign key(item_id) references items(item_id),
    foreign key(order_id) references orders(order_id)
);