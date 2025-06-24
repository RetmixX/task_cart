-- +goose Up
-- +goose StatementBegin

create table if not exists products
(
    id         bigserial primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    title      text,
    price      numeric not null,
    count      bigint  not null
);

create index if not exists idx_products_deleted_at
    on products (deleted_at);

create table if not exists statuses
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name       text not null
);

create index if not exists idx_statuses_deleted_at
    on statuses (deleted_at);

create table if not exists carts
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    is_ordered boolean not null
);

create index if not exists idx_carts_deleted_at
    on carts (deleted_at);

create table if not exists cart_products
(
    product_id bigint not null
        constraint fk_cart_products_product
            references products,
    cart_id    bigint not null
        constraint fk_cart_products_cart
            references carts,
    quantity   bigint not null,
    primary key (product_id, cart_id)
);

create table if not exists orders
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    cart_id    bigint,
    status_id  bigint
        constraint fk_statuses_orders
            references statuses,
    amount     numeric not null
);

create index if not exists idx_orders_status_id
    on orders (status_id);

create index if not exists idx_orders_cart_id
    on orders (cart_id);

create index if not exists idx_orders_deleted_at
    on orders (deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
