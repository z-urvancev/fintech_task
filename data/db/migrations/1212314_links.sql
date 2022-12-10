-- +goose Up
-- +goose StatementBegin
CREATE TABLE links
(
    id    serial  not null primary key,
    url   varchar not null,
    short varchar not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE links;
-- +goose StatementEnd