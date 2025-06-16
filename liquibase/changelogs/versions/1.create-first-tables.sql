--liquibase formatted sql

--changeset lucas-10101:create-company_groups-table
CREATE TABLE company_groups (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    activate BOOLEAN NOT NULL,
    country_id BIGINT ,
    country_subdivision_id BIGINT 
);

--rollback DROP TABLE company_groups;

--changeset lucas-10101:create-companies-table
CREATE TABLE companies (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    activate BOOLEAN NOT NULL,
    country_id BIGINT ,
    country_subdivision_id BIGINT ,
    company_group_id BIGINT NOT NULL REFERENCES company_groups(id)
);

--rollback DROP TABLE companies;

--changeset lucas-10101:create-roles-table
CREATE TABLE roles (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

--rollback DROP TABLE roles;

--changeset lucas-10101:create-access_groups-table
CREATE TABLE access_groups (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

--rollback DROP TABLE access_groups;

--changeset lucas-10101:create-access_group_roles-table
CREATE TABLE access_group_roles (
    access_group_id BIGINT NOT NULL REFERENCES access_groups(id),
    role_id BIGINT NOT NULL REFERENCES roles(id),
    PRIMARY KEY (access_group_id, role_id)
);

--rollback DROP TABLE access_group_roles;

--changeset lucas-10101:create-users-table
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(128) NOT NULL,
    active BOOLEAN NOT NULL
);

--rollback DROP TABLE users;

--changeset lucas-10101:create-user_companies-table
CREATE TABLE user_companies (
    user_id BIGINT NOT NULL REFERENCES users(id),
    company_id BIGINT NOT NULL REFERENCES companies(id),
    PRIMARY KEY (user_id, company_id)
);

--rollback DROP TABLE user_companies;

--changeset lucas-10101:create-plants-table
CREATE TABLE plants (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    activate BOOLEAN NOT NULL
);

--rollback DROP TABLE plants;

--changeset lucas-10101:create-rouser_plantsles-table
CREATE TABLE user_plants (
    user_id BIGINT NOT NULL REFERENCES users(id),
    plant_id BIGINT NOT NULL REFERENCES plants(id),
    PRIMARY KEY (user_id, plant_id)
);

--rollback DROP TABLE user_plants;

--changeset lucas-10101:create-user_access_group-table
CREATE TABLE user_access_groups (
    user_id BIGINT NOT NULL REFERENCES users(id),
    access_group_id BIGINT NOT NULL REFERENCES access_groups(id),
    PRIMARY KEY (user_id, access_group_id)
);

--rollback DROP TABLE user_access_groups;
