-- +goose Up
-- +goose StatementBegin
insert into products( title, price, count) VALUES
('Pizza', 150, 999),
('Milk', 60, 999),
('Apple', 49, 999),
('Ice', 123, 999);

insert into statuses(id, name) VALUES
(1,'Issued'), (2,'Paid'), (3,'Sent'), (4,'Delivered')
on conflict do nothing;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
