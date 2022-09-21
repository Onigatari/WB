-- name: create-table-deliveries
CREATE TABLE DELIVERY
(
    delivery_name VARCHAR UNIQUE,
    phone         VARCHAR(20),
    zip           VARCHAR(30),
    city          VARCHAR(30),
    address       VARCHAR(30),
    region        VARCHAR(30),
    email         VARCHAR(30)
);

-- name: create-table-payments
CREATE TABLE PAYMENT
(
    payment_transaction VARCHAR UNIQUE,
    request_id          VARCHAR(30),
    currency            VARCHAR(30),
    payment_provider    VARCHAR(30),
    amount              INT,
    payment_dt          INT,
    bank                VARCHAR(30),
    delivery_cost       INT,
    goods_total         INT,
    custom_fee          INT
);

-- name: create-table-orders
CREATE TABLE ORDERS
(
    order_uid          VARCHAR PRIMARY KEY,
    track_number       VARCHAR(25) NOT NULL,
    entry              VARCHAR(25),
    delivery           VARCHAR,
    payment            VARCHAR,
    locale             VARCHAR(3),
    internal_signature TEXT,
    customer_id        VARCHAR,
    delivery_service   VARCHAR(30),
    shardkey           VARCHAR(10),
    sm_id              INT,
    date_created       TIMESTAMPTZ,
    oof_shard          VARCHAR(5),
    CONSTRAINT delivery_fk FOREIGN KEY (delivery) REFERENCES delivery (delivery_name),
    CONSTRAINT payment_fk FOREIGN KEY (payment) REFERENCES payment (payment_transaction)
);

-- name: create-table-items
CREATE TABLE ITEM
(
    chrt_id      BIGINT,
    order_uid    VARCHAR,
    track_number VARCHAR(30),
    price        INT,
    rid          VARCHAR(30),
    item_name    VARCHAR(30),
    sale         INT,
    size         INT,
    total_price  INT,
    nm_id        BIGINT,
    brand        VARCHAR(30),
    status       INT,
    PRIMARY KEY (chrt_id),
    FOREIGN KEY (order_uid)
        REFERENCES orders (order_uid)
        ON DELETE CASCADE
);